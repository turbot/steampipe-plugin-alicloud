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
				Name:        "category",
				Type:        proto.ColumnType_STRING,
				Description: "The RDS edition of the instance.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getRdsInstance,
				Description: "The creation time of the Instance.",
			},
			{
				Name:        "lock_reason",
				Type:        proto.ColumnType_STRING,
				Description: "The reason why the instance is locked.",
			},
			{
				Name:      "ins_id",
				Type:      proto.ColumnType_INT,
				Transform: transform.FromField("InsId"),
			},
			{
				Name:        "guard_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GuardDBInstanceId"),
				Description: "The ID of the disaster recovery instance that is attached to the instance if a disaster recovery instance is deployed.",
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
				Description: "The ID of the cloud instance on which the specified VPC is deployed.",
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
				Description: "The type of storage media that is used by the instance.",
			},
			{
				Name:        "dedicated_host_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the dedicated cluster to which the instances belong if the instances are created in a dedicated cluster.",
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
				Description: "The ID of the primary instance to which the instance is attached. If this parameter is not returned, the instance is a primary instance.",
			},
			{
				Name:        "temp_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TempDBInstanceId"),
				Description: "The ID of the temporary instance that is attached to the instance if a temporary instance is deployed.",
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
				Name:        "connection_mode",
				Type:        proto.ColumnType_STRING,
				Description: "The connection mode of the instances.",
			},

			{
				Name:        "auto_upgrade_minor_version",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AutoUpgradeMinorVersion"),
				Description: "The method that is used to update the minor engine version of the instance.",
			},
			{
				Name:        "lock_mode",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LockMode"),
				Description: "The lock mode of the instance.",
			},
			{
				Name:        "time_zone",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("TimeZone"),
				Description: "The time zone of the instance.",
			},
			{
				Name:      "temp_upgrade_time_start",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeTimeStart"),
			},
			{
				Name:      "temp_upgrade_recovery_time",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeRecoveryTime"),
			},
			{
				Name:      "temp_upgrade_recovery_max_iops",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeRecoveryMaxIOPS"),
			},
			{
				Name:      "db_instance_disk_used",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("DBInstanceDiskUsed"),
			},
			{
				Name:        "advanced_features",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("AdvancedFeatures"),
				Description: "An array that consists of advanced features. The advanced features are separated by commas (,). This parameter is supported only for instances that run SQL Server.",
			},
			{
				Name:        "db_max_quantity",
				Type:        proto.ColumnType_INT,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("DBMaxQuantity"),
				Description: "The maximum number of databases that can be created on the instance.",
			},
			{
				Name:        "db_instance_cpu",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("DBInstanceCPU"),
				Description: "The number of CPUs that are configured for the instance.",
			},
			{
				Name:        "max_connections",
				Type:        proto.ColumnType_INT,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("MaxConnections"),
				Description: "The maximum number of concurrent connections that are allowed by the instance.",
			},
			{
				Name:        "increment_source_db_instance_id",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("IncrementSourceDBInstanceId"),
				Description: "The ID of the instance from which incremental data comes. The incremental data of a disaster recovery or read-only instance comes from its primary instance. If this parameter is not returned, the instance is a primary instance.",
			},
			{
				Name:      "multiple_temp_upgrade",
				Type:      proto.ColumnType_BOOL,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("MultipleTempUpgrade"),
			},
			{
				Name:      "temp_upgrade_recovery_class",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeRecoveryClass"),
			},
			{
				Name:        "db_instance_memory",
				Type:        proto.ColumnType_DOUBLE,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("DBInstanceMemory"),
				Description: "The memory capacity of the instance. Unit: MB.",
			},
			{
				Name:        "security_ip_list",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("SecurityIPList"),
				Description: "An array that consists of IP addresses in the IP address whitelist.",
			},
			{
				Name:      "latest_kernel_version",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("LatestKernelVersion"),
			},
			{
				Name:      "support_upgrade_account_type",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("SupportUpgradeAccountType"),
			},
			{
				Name:        "max_iops",
				Type:        proto.ColumnType_INT,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("MaxIOPS"),
				Description: "The maximum number of I/O requests that the instance can process per second.",
			},
			{
				Name:        "maintain_time",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("MaintainTime"),
				Description: "The maintenance window of the instance. The maintenance window is displayed in UTC+8 in the ApsaraDB RDS console.",
			},
			{
				Name:        "db_instance_storage",
				Type:        proto.ColumnType_INT,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("DBInstanceStorage"),
				Description: "The type of storage media that is used by the instance.",
			},
			{
				Name:      "support_create_super_account",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("SupportCreateSuperAccount"),
			},
			{
				Name:      "ip_type",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("IPType"),
			},
			{
				Name:        "collation",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("Collation"),
				Description: "The character set collation of the instance.",
			},
			{
				Name:      "account_type",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("AccountType"),
			},
			{
				Name:        "super_permission_mode",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("SuperPermissionMode"),
				Description: "Indicates whether the instance supports superuser accounts, such as the system administrator (SA) account, Active Directory (AD) account, and host account.",
			},
			{
				Name:        "console_version",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("ConsoleVersion"),
				Description: "The type of proxy that is enabled on the instance.",
			},
			{
				Name:      "temp_upgrade_time_end",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeTimeEnd"),
			},

			{
				Name:      "temp_upgrade_recovery_memory",
				Type:      proto.ColumnType_INT,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeRecoveryMemory"),
			},
			{
				Name:      "dispense_mode",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("DispenseMode"),
			},
			{
				Name:      "origin_configuration",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("OriginConfiguration"),
			},
			{
				Name:        "proxy_type",
				Type:        proto.ColumnType_INT,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("ProxyType"),
				Description: "The type of proxy that is enabled on the instance.",
			},
			{
				Name:        "account_max_quantity",
				Type:        proto.ColumnType_INT,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("AccountMaxQuantity"),
				Description: "The maximum number of accounts that can be created on the instance.",
			},
			{
				Name:      "temp_upgrade_recovery_max_connections",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeRecoveryMaxConnections"),
			},
			{
				Name:        "port",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("Port"),
				Description: "The internal port of the instance.",
			},
			{
				Name:        "security_ip_mode",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("SecurityIPMode"),
				Description: "The network isolation mode of the instance.",
			},
			{
				Name:      "temp_upgrade_recovery_cpu",
				Type:      proto.ColumnType_INT,
				Hydrate:   getRdsInstance,
				Transform: transform.FromField("TempUpgradeRecoveryCpu"),
			},
			{
				Name:        "connection_string",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("ConnectionString"),
				Description: "The internal endpoint of the instance.",
			},
			{
				Name:        "availability_value",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("AvailabilityValue"),
				Description: "The availability status of the instance. Unit: %.",
			},
			{
				Name:        "db_instance_ip_array_name",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstanceIPArrayList,
				Transform:   transform.FromField("DBInstanceIPArrayName"),
				Description: "The name of the IP address whitelist.",
			},
			{
				Name:        "db_instance_ip_array_attribute",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstanceIPArrayList,
				Transform:   transform.FromField("DBInstanceIPArrayAttribute"),
				Description: "The attribute of the IP address whitelist.",
			},
			{
				Name:        "security_ip_type",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstanceIPArrayList,
				Transform:   transform.FromField("SecurityIPType"),
				Description: "The type of the IP address.",
			},
			{
				Name:      "whitelist_network_type",
				Type:      proto.ColumnType_STRING,
				Hydrate:   getRdsInstanceIPArrayList,
				Transform: transform.FromField("WhitelistNetworkType"),
			},

			{
				Name:        "readonly_db_instance_ids",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("ReadOnlyDBInstanceIds"),
				Description: "An array that consists of the IDs of the read-only instances attached to the primary instance.",
			},

			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsTags,
				Transform:   transform.FromValue().Transform(rdsInstanceTagsSrc),
				Description: ColumnDescriptionTags,
			},

			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsTags,
				Transform:   transform.FromValue().Transform(rdsInstanceTags),
				Description: ColumnDescriptionTags,
			},
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
			d.StreamListItem(ctx, rds.DBInstanceAttribute{
				DBInstanceId:            i.DBInstanceId,
				Engine:                  i.Engine,
				DBInstanceDescription:   i.DBInstanceDescription,
				PayType:                 i.PayType,
				DBInstanceType:          i.DBInstanceType,
				InstanceNetworkType:     i.InstanceNetworkType,
				ConnectionMode:          i.ConnectionMode,
				RegionId:                i.RegionId,
				ExpireTime:              i.ExpireTime,
				DBInstanceStatus:        i.DBInstanceStatus,
				DBInstanceNetType:       i.DBInstanceNetType,
				LockMode:                i.LockMode,
				LockReason:              i.LockReason,
				MasterInstanceId:        i.MasterInstanceId,
				GuardDBInstanceId:       i.GuardDBInstanceId,
				TempDBInstanceId:        i.TempDBInstanceId,
				AutoUpgradeMinorVersion: i.AutoUpgradeMinorVersion,
				Category:                i.Category,
				DBInstanceClass:         i.DBInstanceClass,
				DBInstanceStorageType:   i.DBInstanceStorageType,
				DedicatedHostGroupId:    i.DedicatedHostGroupId,
				EngineVersion:           i.EngineVersion,
				ResourceGroupId:         i.ResourceGroupId,
				VSwitchId:               i.VSwitchId,
				VpcCloudInstanceId:      i.VpcCloudInstanceId,
				VpcId:                   i.VpcId,
				ZoneId:                  i.ZoneId,
				InsId:                   i.InsId,
			})
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
		rds := h.Item.(rds.DBInstanceAttribute)
		id = rds.DBInstanceId
	} else {
		id = d.KeyColumnQuals["db_instance_id"].GetStringValue()
	}

	request := rds.CreateDescribeDBInstanceAttributeRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceAttribute(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {
			plugin.Logger(ctx).Warn("alicloud_rds_instance.getRdsInstance", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getRdsInstance", "query_error", err, "request", request)
		return nil, err
	}

	if response.Items.DBInstanceAttribute != nil && len(response.Items.DBInstanceAttribute) > 0 {
		return response.Items.DBInstanceAttribute[0], nil
	}

	return nil, nil
}

func getRdsInstanceIPArrayList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstanceIPArrayList", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		rds := h.Item.(rds.DBInstanceAttribute)
		id = rds.DBInstanceId
	} else {
		id = d.KeyColumnQuals["db_instance_id"].GetStringValue()
	}

	request := rds.CreateDescribeDBInstanceIPArrayListRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceIPArrayList(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {
			plugin.Logger(ctx).Warn("alicloud_rds_instance.getRdsInstanceIPArrayList", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getRdsInstanceIPArrayList", "query_error", err, "request", request)
		return nil, err
	}

	if response.Items.DBInstanceIPArray != nil && len(response.Items.DBInstanceIPArray) > 0 {
		return response.Items.DBInstanceIPArray[0], nil
	}

	return nil, nil
}

func getRdsTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsTags", "connection_error", err)
		return nil, err
	}

	var id string
	rdsInstance := h.Item.(rds.DBInstanceAttribute)
	id = rdsInstance.RegionId

	request := rds.CreateDescribeTagsRequest()
	request.Scheme = "https"
	request.RegionId = id
	response, err := client.DescribeTags(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {
			plugin.Logger(ctx).Warn("alicloud_rds_instance.getRdsTags", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getRdsTags", "query_error", err, "request", request)
		return nil, err
	}

	if response != nil {
		return response, nil
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func getRdsInstanceAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	i := h.Item.(rds.DBInstanceAttribute)
	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID
	return []string{"acs:rds:" + i.RegionId + ":" + accountID + ":instance/" + i.DBInstanceId}, nil
}
func rdsInstanceTagsSrc(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.(*rds.DescribeTagsResponse)
	var turbotTagsMap []map[string]string

	if tags.Items.TagInfos != nil {
		for _, i := range tags.Items.TagInfos {
			turbotTagsMap = append(turbotTagsMap, map[string]string{"Key": i.TagKey, "Value": i.TagValue})
		}
	}

	return turbotTagsMap, nil
}

func rdsInstanceTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.(*rds.DescribeTagsResponse)
	var turbotTagsMap map[string]string

	if tags.Items.TagInfos != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags.Items.TagInfos {
			turbotTagsMap[i.TagKey] = i.TagValue
		}
	}

	return turbotTagsMap, nil
}
