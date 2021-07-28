package alicloud

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/sethvargo/go-retry"

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
			KeyColumns:        plugin.AllColumns([]string{"db_instance_id", "region"}),
			ShouldIgnoreError: isNotFoundError([]string{"InvalidDBInstanceId.NotFound"}),
			Hydrate:           getRdsInstance,
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
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the RDS instance.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRdsInstanceARN,
				Transform:   transform.FromValue(),
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
				Hydrate:     getRdsInstance,
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
				Name:        "sql_collector_retention",
				Type:        proto.ColumnType_INT,
				Hydrate:     getSqlCollectorRetention,
				Transform:   transform.FromField("ConfigValue"),
				Description: "The log backup retention duration that is allowed by the SQL explorer feature on the instance.",
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
				Hydrate:     getRdsInstance,
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
				Hydrate:     getRdsInstance,
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
				Name:        "ssl_status",
				Type:        proto.ColumnType_STRING,
				Description: "The SSL encryption status of the Instance",
				Hydrate:     getSSLDetails,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tde_status",
				Type:        proto.ColumnType_STRING,
				Description: "The TDE status at the instance level. Valid values: Enable | Disable.",
				Hydrate:     getTDEDetails,
				Transform:   transform.FromField("TDEStatus"),
			},
			{
				Name:        "security_ips",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsInstanceIPArrayList,
				Transform:   transform.FromField("Items.DBInstanceIPArray").Transform(getSecurityIps),
				Description: "An array that consists of IP addresses in the IP address whitelist.",
			},
			{
				Name:        "security_ips_src",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsInstanceIPArrayList,
				Transform:   transform.FromField("Items.DBInstanceIPArray"),
				Description: "An array that consists of IP details.",
			},
			{
				Name:        "parameters",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsInstanceParameters,
				Transform:   transform.FromValue(),
				Description: "The list of running parameters for the instance.",
			},
			{
				Name:        "readonly_db_instance_ids",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRdsInstance,
				Transform:   transform.FromField("ReadOnlyDBInstanceIds"),
				Description: "An array that consists of the IDs of the read-only instances attached to the primary instance.",
			},
			{
				Name:        "sql_collector_policy",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSqlCollectorPolicy,
				Transform:   transform.FromValue(),
				Description: "The status of the SQL Explorer (SQL Audit) feature.",
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
				Hydrate:     getRdsInstanceARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
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
	region := d.KeyColumnQualString(matrixKeyRegion)

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
	regionMatrix := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, regionMatrix)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstance", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		id = databaseID(h.Item)
	} else {
		id = d.KeyColumnQuals["db_instance_id"].GetStringValue()
		region := d.KeyColumnQuals["region"].GetStringValue()
		if region != regionMatrix {
			return nil, nil
		}
	}

	request := rds.CreateDescribeDBInstanceAttributeRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	var response *rds.DescribeDBInstanceAttributeResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeDBInstanceAttribute(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if response.Items.DBInstanceAttribute != nil && len(response.Items.DBInstanceAttribute) > 0 {
		return response.Items.DBInstanceAttribute[0], nil
	}

	return nil, nil
}

func getRdsInstanceIPArrayList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstanceIPArrayList", "connection_error", err)
		return nil, err
	}

	var id = databaseID(h.Item)

	request := rds.CreateDescribeDBInstanceIPArrayListRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceIPArrayList(request)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstanceIPArrayList", "query_error", err, "request", request)
		return nil, err
	}

	if response.Items.DBInstanceIPArray != nil && len(response.Items.DBInstanceIPArray) > 0 {
		return response, nil
	}

	return nil, nil
}

func getTDEDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getTDEDetails", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		id = databaseID(h.Item)
	} else {
		id = d.KeyColumnQuals["db_instance_id"].GetStringValue()
	}

	request := rds.CreateDescribeDBInstanceTDERequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceTDE(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" || serverErr.ErrorCode() == "InstanceEngineType.NotSupport" || serverErr.ErrorCode() == "InvaildEngineInRegion.ValueNotSupported" {
			plugin.Logger(ctx).Warn("alicloud_rds_instance.getTDEDetails", "error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getTDEDetails", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func getSSLDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getSSLDetails", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		id = databaseID(h.Item)
	} else {
		id = d.KeyColumnQuals["db_instance_id"].GetStringValue()
	}

	request := rds.CreateDescribeDBInstanceSSLRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeDBInstanceSSL(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidDBInstanceId.NotFound" {
			plugin.Logger(ctx).Warn("alicloud_rds_instance.getSSLDetails", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getSSLDetails", "query_error", err, "request", request)
		return nil, err
	}
	if len(response.SSLExpireTime) > 0 {
		return "Enabled", nil
	}
	return "Disabled", nil
}

func getRdsInstanceParameters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstanceParameters", "connection_error", err)
		return nil, err
	}

	var id = databaseID(h.Item)

	request := rds.CreateDescribeParametersRequest()
	request.Scheme = "https"
	request.DBInstanceId = id
	response, err := client.DescribeParameters(request)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsInstanceParameters", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func getRdsTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getRdsTags", "connection_error", err)
		return nil, err
	}

	request := rds.CreateDescribeTagsRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.DBInstanceId = databaseID(h.Item)
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

func getRdsInstanceARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var region, instanceID string
	switch item := h.Item.(type) {
	case rds.DBInstance:
		region = item.RegionId
		instanceID = item.DBInstanceId
	case rds.DBInstanceAttribute:
		region = item.RegionId
		instanceID = item.DBInstanceId
	}
	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID
	return "arn:acs:rds:" + region + ":" + accountID + ":instance/" + instanceID, nil
}

func getSqlCollectorPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getSqlCollectorPolicy", "connection_error", err)
		return nil, err
	}

	request := rds.CreateDescribeSQLCollectorPolicyRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.DBInstanceId = databaseID(h.Item)
	response, err := client.DescribeSQLCollectorPolicy(request)
	if err != nil {
		plugin.Logger(ctx).Error("getSqlCollectorPolicy", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func getSqlCollectorRetention(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getSqlCollectorRetention", "connection_error", err)
		return nil, err
	}

	request := rds.CreateDescribeSQLCollectorRetentionRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.DBInstanceId = databaseID(h.Item)
	response, err := client.DescribeSQLCollectorRetention(request)
	if err != nil {
		plugin.Logger(ctx).Error("getSqlCollectorRetention", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

//// TRANSFORM FUNCTIONS

func getSecurityIps(_ context.Context, d *transform.TransformData) (interface{}, error) {
	IpArray := d.Value.([]rds.DBInstanceIPArray)

	if len(IpArray) == 0 {
		return nil, nil
	}
	var IpList []string
	for _, i := range IpArray {
		IpList = append(IpList, i.SecurityIPList)
	}
	return IpList, nil
}

func rdsInstanceTagsSrc(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.(*rds.DescribeTagsResponse)
	if tags == nil {
		return nil, nil
	}

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

func databaseID(item interface{}) string {
	switch item := item.(type) {
	case rds.DBInstance:
		return item.DBInstanceId
	case rds.DBInstanceAttribute:
		return item.DBInstanceId
	}
	return ""
}
