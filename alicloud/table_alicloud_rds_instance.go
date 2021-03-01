package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsInstance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_instance",
		Description: "Provides an RDS instance resource. A DB instance is an isolated database environment in the cloud. A DB instance can contain multiple user-created databases.",
		List: &plugin.ListConfig{
			Hydrate: listRdsInstances,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("db_instance_id"),
			Hydrate:    getRdsInstance,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceId"),
				Description: "The ID of the single instance to query.",
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpcId"),
				Description: "The ID of the VPC to which the instances belong.",
			},

			// Other columns
			{
				Name:        "Category",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the Instance.",
			},
			{
				Name:        "lock_reason",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "ins_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("InsId"),
				Description: "",
			},
			{
				Name:        "guard_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GuardDBInstanceId"),
				Description: "",
			},
			{
				Name:        "db_instance_description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceDescription"),
				Description: "The description of the DB Instance.",
			},
			{
				Name:        "engine",
				Type:        proto.ColumnType_STRING,
				Description: "The database engine that the instances run.",
			},
			{
				Name:        "vpc_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the VPC to which the instances belong.",
			},
			{
				Name:        "db_instance_net_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceNetType"),
				Description: "The ID of the resource group to which the VPC belongs.",
			},
			{
				Name:        "db_instance_class",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceClass"),
				Description: "The instance type of the instances.",
			},
			{
				Name:        "vpc_cloud_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "destroy_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "dedicated_host_id_for_master",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the host to which the instances belong if the instances are created in a dedicated cluster.",
			},
			{
				Name:        "dedicated_host_name_for_log",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the DHCP options set associated to vpc.",
			},
			{
				Name:        "region_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the region to which the instances belong.",
			},
			{
				Name:        "instance_network_type",
				Type:        proto.ColumnType_STRING,
				Description: "The network type of the instances.",
			},
			{
				Name:        "resource_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the resource group to which the instances belong.",
			},
			{
				Name:        "db_instance_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceType"),
				Description: "The role of the instances.",
			},
			{
				Name:        "expire_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Instance expire time",
			},

			{
				Name:        "db_instance_storage_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceStorageType"),
				Description: "",
			},
			{
				Name:        "Mutri_or_signle",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("MutriORsignle"),
				Description: "",
			},
			{
				Name:        "dedicated_host_zone_id_for_master",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the host to which the instances belong if the instances are created in a dedicated cluster",
			},

			{
				Name:        "dedicated_host_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the dedicated cluster to which the instances belong if the instances are created in a dedicated cluster.",
			},
			{
				Name:        "dedicated_host_id_for_log",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostIdForLog"),
				Description: "",
			},
			{
				Name:        "dedicated_host_group_name",
				Type:        proto.ColumnType_STRING,
				Description: "The Name of the dedicated cluster to which the instances belong if the instances are created in a dedicated cluster.",
			},
			{
				Name:        "engine_version",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the database engine that the instances run.",
			},
			{
				Name:        "pay_type",
				Type:        proto.ColumnType_STRING,
				Description: "The billing method of the instances.",
			},
			{
				Name:        "vswitch_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VSwitchId"),
				Description: "The ID of the vSwitch associated with the specified VPC.",
			},

			{
				Name:        "master_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "dedicated_host_zone_id_for_slave",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostZoneIdForSlave"),
				Description: "",
			},
			{
				Name:        "temp_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TempDBInstanceId"),
				Description: "",
			},
			{
				Name:        "db_instance_status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceStatus"),
				Description: "The status of the instances",
			},
			{
				Name:        "zone_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the zone to which the instances belong.",
			},
			{
				Name:        "replicate_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "dedicated_host_name_for_slave",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "dedicated_host_zone_id_for_log",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "connection_mode",
				Type:        proto.ColumnType_STRING,
				Description: "The connection mode of the instances.",
			},
			{
				Name:        "dedicated_host_name_for_master",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostNameForMaster"),
				Description: "The name of the host to which the instances belong if the instances are created in a dedicated cluster.",
			},
			{
				Name:        "auto_upgrade_minor_version",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AutoUpgradeMinorVersion"),
				Description: "",
			},
			{
				Name:        "lock_mode",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LockMode"),
				Description: "",
			},
			{
				Name:        "dedicated_host_id_for_slave",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostIdForSlave"),
				Description: "",
			},
			{
				Name:        "readonly_db_instance_ids",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ReadOnlyDBInstanceIds"),
				Description: "",
			},

			// {
			// 	Name:        "tags_src",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   transform.FromField("Tags.Tag"),
			// 	Description: ColumnDescriptionTags,
			// },

			// Resource interface
			// {
			// 	Name:        "tags",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   transform.FromField("Tags.Tag").Transform(vpcTurbotTags),
			// 	Description: ColumnDescriptionTags,
			// },
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceId"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsInstanceAkas,
				Transform:   transform.FromValue(),
				Description: ColumnDescriptionAkas,
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

func listRdsInstances(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_rds.listRdsInstances", "connection_error", err)
		return nil, err
	}
	request := rds.CreateDescribeDBInstancesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeDBInstances(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_rds.DescribeDBInstances", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Items.DBInstance {
			plugin.Logger(ctx).Warn("alicloud_rds.DescribeDBInstances", "item", i)
			d.StreamListItem(ctx, i)
			count++
		}
		if count >= response.TotalRecordCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getRdsInstance(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstance", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		rds := h.Item.(rds.DBInstance)
		id = rds.DBInstanceId
	} else {
		id = d.KeyColumnQuals["db_instance_id"].GetStringValue()
	}

	request := rds.CreateDescribeDBInstancesRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstances(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {
			plugin.Logger(ctx).Warn("alicloud_rds_instance.getRdsInstance", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getRdsInstance", "query_error", err, "request", request)
		return nil, err
	}

	if response.Items.DBInstance != nil && len(response.Items.DBInstance) > 0 {
		return response.Items.DBInstance[0], nil
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func getRdsInstanceAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	i := h.Item.(rds.DBInstance)
	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID
	return []string{"acs:rds:" + i.RegionId + ":" + accountID + ":instance/" + i.DBInstanceId}, nil
}
