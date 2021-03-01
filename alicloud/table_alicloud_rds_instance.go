package alicloud

import (
	"context"

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
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
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
				Description: "A list of user CIDRs.",
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpcId"),
				Description: "The unique ID of the VPC.",
			},

			// Other columns
			{
				Name:        "Category",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the VPC.",
			},
			{
				Name:        "lock_reason",
				Type:        proto.ColumnType_STRING,
				Description: "The IPv4 CIDR block of the VPC.",
			},
			{
				Name:        "ins_id",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("InsId"),
				Description: "The IPv6 CIDR block of the VPC.",
			},
			{
				Name:        "guard_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GuardDBInstanceId"),
				Description: "The ID of the VRouter.",
			},
			{
				Name:        "db_instance_description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceDescription"),
				Description: "The description of the VPC.",
			},
			{
				Name:        "engine",
				Type:        proto.ColumnType_STRING,
				Description: "True if the VPC is the default VPC in the region.",
			},
			{
				Name:        "vpc_name",
				Type:        proto.ColumnType_STRING,
				Description: "",
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
				Description: "Indicates whether the VPC is attached to any Cloud Enterprise Network (CEN) instance.",
			},
			{
				Name:        "vpc_cloud_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the owner of the VPC.",
			},
			{
				Name:        "destroy_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "",
			},
			{
				Name:        "dedicated_host_id_for_master",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "dedicated_host_name_for_log",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the DHCP options set associated to vpc.",
			},
			{
				Name:        "region_id",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the VPC network that is associated with the DHCP options set. Valid values: InUse and Pending",
			},
			{
				Name:        "instance_network_type",
				Type:        proto.ColumnType_STRING,
				Description: "The list of Cloud Enterprise Network (CEN) instances to which the VPC is attached. No value is returned if the VPC is not attached to any CEN instance.",
			},
			{
				Name:        "resource_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "True if the ClassicLink function is enabled.",
			},
			{
				Name:        "db_instance_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceType"),
				Description: "The list of resources in the VPC.",
			},
			{
				Name:        "expire_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "A list of VSwitches in the VPC.",
			},

			{
				Name:        "db_instance_storage_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceStorageType"),
				Description: "A list of IDs of NAT Gateways.",
			},
			{
				Name:        "Mutri_or_signle",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("MutriORsignle"),
				Description: "A list of IDs of route tables.",
			},
			{
				Name:        "dedicated_host_zone_id_for_master",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},

			{
				Name:        "dedicated_host_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_id_for_log",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostIdForLog"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_group_name",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "engine_version",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "pay_type",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "vswitch_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VSwitchId"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},

			{
				Name:        "master_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_zone_id_for_slave",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostZoneIdForSlave"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "temp_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TempDBInstanceId"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "db_instance_status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceStatus"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "zone_id",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "replicate_id",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_name_for_slave",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_zone_id_for_log",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "connection_mode",
				Type:        proto.ColumnType_STRING,
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_name_for_master",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostNameForMaster"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "auto_upgrade_minor_version",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AutoUpgradeMinorVersion"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "lock_mode",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LockMode"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "dedicated_host_id_for_slave",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostIdForSlave"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "readonly_db_instance_ids",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ReadOnlyDBInstanceIds"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
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
	request.VpcId = id
	response, err := client.DescribeDBInstances(request)
	if err != nil {
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
	return []string{"acs:rds:" + i.RegionId + ":" + accountID + ":instance/" + i.VpcId}, nil
}
