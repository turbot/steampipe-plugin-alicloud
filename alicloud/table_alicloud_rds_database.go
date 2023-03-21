package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_database",
		Description: "Alibaba Cloud ApsaraDB for RDS (Relational Database Service) is a stable and reliable online database service that scales elastically.",
		List: &plugin.ListConfig{
			ParentHydrate: listRdsInstances,
			Hydrate:       listRdsdatabases,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "db_instance_id",
					Require: plugin.Optional,
				},
				{
					Name:    "db_name",
					Require: plugin.Optional,
				},
				{
					Name:    "db_status",
					Require: plugin.Optional,
				},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "db_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBName"),
				Description: "The name of the database.",
			},
			{
				Name:        "db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceId"),
				Description: "The unique ID of the instance.",
			},
			{
				Name:        "db_status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBStatus"),
				Description: "The status of the database. Valid values: Creating, Running and Deleting.",
			},
			{
				Name:        "db_description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBDescription"),
				Description: "The description of the database.",
			},
			{
				Name:        "character_set_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the character set.",
			},
			{
				Name:        "engine",
				Type:        proto.ColumnType_STRING,
				Description: "The database engine of the instance to which the database belongs.",
			},
			{
				Name:        "tde_status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TDEStatus"),
				Description: "The TDE status of the database.",
			},
			{
				Name:        "accounts",
				Type:        proto.ColumnType_JSON,
				Description: "An array that consists of the details of the accounts. Each account has specific permissions on the database.",
			},

			// Steampipe standard Columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBName"),
				Description: ColumnDescriptionTitle,
			},

			// alicloud common columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
			},
			{
				Name:        "account_id",
				Description: ColumnDescriptionAccount,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID"),
			},
		},
	}
}

//// LIST FUNCTION

func listRdsdatabases(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	dbInstance := h.Item.(rds.DBInstance)
	if d.EqualsQualString("db_instance_id") != "" {
		if d.EqualsQualString("db_instance_id") != dbInstance.DBInstanceId {
			return nil, nil
		}
	}

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_rds_database.listRdsdatabases", "connection_error", err)
		return nil, err
	}
	request := rds.CreateDescribeDatabasesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(100)
	request.PageNumber = requests.NewInteger(1)
	request.DBInstanceId = dbInstance.DBInstanceId

	if d.EqualsQualString("db_name") != "" {
		request.DBName = d.EqualsQualString("db_name")
	}
	if d.EqualsQualString("db_status") != "" {
		request.DBStatus = d.EqualsQualString("db_status")
	}

	for {
		response, err := client.DescribeDatabases(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_rds_database.listRdsdatabases", "query_error", err, "request", request)
			return nil, err
		}

		// Response body doesn't contain record count and page number details, so we need to handle it manually
		if len(response.Databases.Database) > 0 {
			request.PageNumber = request.PageNumber + requests.NewInteger(1)
		} else {
			break
		}

		for _, i := range response.Databases.Database {
			d.StreamListItem(ctx, i)
		}

	}

	return nil, nil
}

