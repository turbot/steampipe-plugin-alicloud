package alicloud

import (
	"context"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func tableAlicloudEcsInstance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_instance",
		Description: "Alicloud Elastic Compute Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("instance_id"),
			Hydrate:    getEcsInstance,
		},
		List: &plugin.ListConfig{
			Hydrate: listEcsInstance,
			KeyColumns: plugin.KeyColumnSlice{
				// String columns
				{Name: "image_id", Require: plugin.Optional},
				{Name: "resource_group_id", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
				{Name: "zone", Require: plugin.Optional},
				{Name: "billing_method", Require: plugin.Optional},
				{Name: "family", Require: plugin.Optional},
				{Name: "instance_network_type", Require: plugin.Optional},
				{Name: "instance_type", Require: plugin.Optional},
				{Name: "internet_charge_type", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},

				// Boolean columns
				{Name: "device_available", Require: plugin.Optional, Operators: []string{"<>", "="}},
				{Name: "io_optimized", Require: plugin.Optional, Operators: []string{"<>", "="}},
			},
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceName"),
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the ECS instance.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsInstanceARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "instance_type",
				Description: "The type of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vpc_id",
				Description: "The type of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpcAttributes.VpcId"),
			},
			{
				Name:        "status",
				Description: "The status of the instance. Possible values are: Pending, Running, Starting, Stopping, and Stopped",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "billing_method",
				Description: "The billing method for network usage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceChargeType"),
			},
			{
				Name:        "creation_time",
				Description: "The time when the instance was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "deletion_protection",
				Description: "Indicates whether you can use the ECS console or call the DeleteInstance operation to release the instance.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "instance_network_type",
				Description: "The network type of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "family",
				Description: "The instance family of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceTypeFamily"),
			},
			{
				Name:        "activation_id",
				Description: "The activation Id if the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "agent_version",
				Description: "The agent version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auto_release_time",
				Description: "The automatic release time of the pay-as-you-go instance.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "connected",
				Description: "Indicates whether the instance is connected..",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cpu",
				Description: "The number of vCPUs.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "cpu_options_core_count",
				Description: "The number of CPU cores.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("CpuOptions.CoreCount"),
			},
			{
				Name:        "cpu_options_numa",
				Description: "The number of threads allocated.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CpuOptions.Numa").NullIfZero(),
			},
			{
				Name:        "cpu_options_threads_per_core",
				Description: "The number of threads per core.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("CpuOptions.ThreadsPerCore"),
			},
			{
				Name:        "credit_specification",
				Description: "The performance mode of the burstable instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dedicated_host_cluster_id",
				Description: "The cluster ID of the dedicated host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostAttribute.DedicatedHostClusterId"),
			},
			{
				Name:        "dedicated_host_id",
				Description: "The ID of the dedicated host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostAttribute.DedicatedHostId"),
			},
			{
				Name:        "dedicated_host_name",
				Description: "The name of the dedicated host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostAttribute.DedicatedHostName"),
			},
			{
				Name:        "dedicated_instance_affinity",
				Description: "Indicates whether the instance on a dedicated host is associated with the dedicated host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedInstanceAttribute.Affinity"),
			},
			{
				Name:        "dedicated_instance_tenancy",
				Description: "Indicates whether the instance is hosted on a dedicated host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedInstanceAttribute.Tenancy"),
			},
			{
				Name:        "deployment_set_group_no",
				Description: "The group No. of the instance in a deployment set when the deployment set is used to distribute instances across multiple physical machines.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "deployment_set_id",
				Description: "The ID of the deployment set.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "device_available",
				Description: "Indicates whether data disks can be attached to the instance.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ecs_capacity_reservation_id",
				Description: "The ID of the capacity reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EcsCapacityReservationAttr.CapacityReservationId"),
			},
			{
				Name:        "ecs_capacity_reservation_preference",
				Description: "The preference of the ECS capacity reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EcsCapacityReservationAttr.CapacityReservationPreference"),
			},
			{
				Name:        "expired_time",
				Description: "The expiration time of the instance.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "gpu_amount",
				Description: "The number of GPUs for the instance type.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("GPUAmount"),
			},
			{
				Name:        "gpu_spec",
				Description: "The category of GPUs for the instance type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GPUSpec"),
			},
			{
				Name:        "host_name",
				Description: "The hostname of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hpc_cluster_id",
				Description: "The ID of the HPC cluster to which the instance belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "image_id",
				Description: "The ID of the image that the instance is running.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "internet_charge_type",
				Description: "The billing method for network usage. Valid values:PayByBandwidth,PayByTraffic",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "internet_max_bandwidth_in",
				Description: "The maximum inbound bandwidth from the Internet (in Mbit/s).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "internet_max_bandwidth_out",
				Description: "The maximum outbound bandwidth to the Internet (in Mbit/s).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "invocation_count",
				Description: "The count of instance invocation",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "io_optimized",
				Description: "Specifies whether the instance is I/O optimized.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_spot",
				Description: "Indicates whether the instance is a spot instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "key_pair_name",
				Description: "The name of the SSH key pair for the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_invoked_time",
				Description: "The time when the instance is last invoked.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "local_storage_amount",
				Description: "The number of local disks attached to the instance.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "local_storage_capacity",
				Description: "The capacity of local disks attached to the instance.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "memory",
				Description: "The memory size of the instance (in MiB).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "network_type",
				Description: "The type of the network.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "os_name",
				Description: "The name of the operating system for the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OSName"),
			},
			{
				Name:        "os_name_en",
				Description: "The English name of the operating system for the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OSNameEn"),
			},
			{
				Name:        "os_type",
				Description: "The type of the operating system. Possible values are: windows and linux.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OSType"),
			},
			{
				Name:        "os_version",
				Description: "The version of the operating system.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "recyclable",
				Description: "Indicates whether the instance can be recycled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "registration_time",
				Description: "The time when the instance is registered.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the instance belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sale_cycle",
				Description: "The billing cycle of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "serial_number",
				Description: "The serial number of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "spot_duration",
				Description: "The protection period of the preemptible instance (in hours).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "spot_price_limit",
				Description: "The maximum hourly price for the instance.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "spot_strategy",
				Description: "The preemption policy for the pay-as-you-go instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start_time",
				Description: "The start time of the bidding mode for the preemptible instance.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "stopped_mode",
				Description: "Indicates whether the instance continues to be billed after it is stopped.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vlan_id",
				Description: "The VLAN ID of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "eip_address",
				Description: "The information of the EIP associated with the instance.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "inner_ip_address",
				Description: "The internal IP addresses of classic network-type instances. This parameter takes effect when InstanceNetworkType is set to classic. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("InnerIpAddress.IpAddress"),
			},
			{
				Name:        "metadata_options",
				Description: "The collection of metadata options.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "network_interfaces",
				Description: "Details about the ENIs bound to the instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("NetworkInterfaces.NetworkInterface"),
			},
			{
				Name:        "operation_locks",
				Description: "Details about the reasons why the instance was locked.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OperationLocks.LockReason"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP addresses of instances.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VpcAttributes.PrivateIpAddress.IpAddress"),
			},
			{
				Name:        "public_ip_address",
				Description: "The public IP addresses of instances.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("PublicIpAddress.IpAddress"),
			},
			{
				Name:        "rdma_ip_address",
				Description: "The RDMA IP address of HPC instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RdmaIpAddress.IpAddress"),
			},
			{
				Name:        "security_group_ids",
				Description: "The IDs of security groups to which the instance belongs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SecurityGroupIds.SecurityGroupId"),
			},
			{
				Name:        "vpc_attributes",
				Description: "The VPC attributes of the instance.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(ecsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsInstanceARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceName"),
			},

			// Alicloud standard columns
			{
				Name:        "zone",
				Description: "The zone in which the instance resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ZoneId"),
			},
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

func listEcsInstance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_instance.listEcsInstance", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(100)
	request.PageNumber = requests.NewInteger(1)
	request.RegionId = d.KeyColumnQualString(matrixKeyRegion)
	quals := d.Quals

	if value, ok := GetBoolQualValue(quals, "device_available"); ok {
		request.DeviceAvailable = requests.NewBoolean(*value)
	}
	if value, ok := GetBoolQualValue(quals, "io_optimized"); ok {
		request.IoOptimized = requests.NewBoolean(*value)
	}
	if value, ok := GetStringQualValue(quals, "image_id"); ok {
		request.ImageId = *value
	}
	if value, ok := GetStringQualValue(quals, "resource_group_id"); ok {
		request.ResourceGroupId = *value
	}
	if value, ok := GetStringQualValue(quals, "vpc_id"); ok {
		request.VpcId = *value
	}
	if value, ok := GetStringQualValue(quals, "zone"); ok {
		request.ZoneId = *value
	}
	if value, ok := GetStringQualValue(quals, "billing_method"); ok {
		request.InstanceChargeType = *value
	}
	if value, ok := GetStringQualValue(quals, "family"); ok {
		request.InstanceTypeFamily = *value
	}
	if value, ok := GetStringQualValue(quals, "instance_network_type"); ok {
		request.InstanceNetworkType = *value
	}
	if value, ok := GetStringQualValue(quals, "instance_type"); ok {
		request.InstanceType = *value
	}
	if value, ok := GetStringQualValue(quals, "internet_charge_type"); ok {
		request.InternetChargeType = *value
	}
	if value, ok := GetStringQualValue(quals, "name"); ok {
		request.InstanceName = *value
	}
	if value, ok := GetStringQualValue(quals, "status"); ok {
		request.Status = *value
	}

	// If the request no of items is less than the paging max limit
	// update limit to requested no of results.
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		pageSize, err := request.PageSize.GetValue64()
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_instance.listEcsInstance", "page_size_error", err)
			return nil, err
		}
		if *limit < pageSize {
			request.PageSize = requests.NewInteger(int(*limit))
		}
	}

	count := 0
	for {
		// https://partners-intl.aliyun.com/help/doc-detail/25506.htm?spm=a2c63.p38356.a3.13.24665a4cJb014m#t9865.html
		response, err := client.DescribeInstances(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_instance.listEcsInstance", "query_error", err, "request", request)
			return nil, err
		}
		for _, instance := range response.Instances.Instance {
			d.StreamListItem(ctx, instance)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if plugin.IsCancelled(ctx) {
				return nil, nil
			}
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getEcsInstance(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsInstance")

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_instance.getEcsInstance", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		instance := h.Item.(ecs.Instance)
		id = instance.InstanceId
	} else {
		id = d.KeyColumnQuals["instance_id"].GetStringValue()
	}

	// In SDK, the Datatype of InstanceIds is string, though the value should be passed as
	// ["i-bp67acfmxazb4p****", "i-bp67acfmxazb4p****", ... "i-bp67acfmxazb4p****"]
	input, err := json.Marshal([]string{id})
	if err != nil {
		return nil, err
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.InstanceIds = string(input)

	response, err := client.DescribeInstances(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_instance.getEcsInstance", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.Instances.Instance != nil && len(response.Instances.Instance) > 0 {
		return response.Instances.Instance[0], nil
	}

	return nil, nil
}

func getEcsInstanceARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsInstanceARN")
	instance := h.Item.(ecs.Instance)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + instance.RegionId + ":" + accountID + ":instance/" + instance.InstanceId

	return arn, nil
}
