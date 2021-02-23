package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudEcsInstance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_instance",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listEcsInstance,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the instance.",
				Transform: transform.FromField("InstanceName")},
			{Name: "host_name", Type: proto.ColumnType_STRING, Description: "The hostname of the instance."},
			{Name: "image_id", Type: proto.ColumnType_STRING, Description: "The ID of the image that the instance is running."},
			{Name: "instance_type", Type: proto.ColumnType_STRING, Description: "The instance type."},
			{Name: "auto_release_time", Type: proto.ColumnType_TIMESTAMP, Description: "The automatic release time of the pay-as-you-go instance."},
			{Name: "last_invoked_time", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "os_type", Type: proto.ColumnType_STRING, Description: "The operating system type of the instance, consisting of Windows Server and Linux."},
			{Name: "device_available", Type: proto.ColumnType_BOOL, Description: "Indicates whether data disks can be attached to the instance."},
			{Name: "instance_network_type", Type: proto.ColumnType_STRING, Description: "The network type of the instance. Valid values:classic,vpc"},
			{Name: "registration_time", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "local_storage_amount", Type: proto.ColumnType_INT, Description: "The number of local disks attached to the instance."},
			{Name: "network_type", Type: proto.ColumnType_STRING, Description: "The network type of the instance. Valid values:classic,vpc"},
			{Name: "intranet_ip", Type: proto.ColumnType_STRING, Description: "The EIP"},
			{Name: "is_spot", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "instance_charge_type", Type: proto.ColumnType_STRING, Description: "The billing method of the instance. Valid values:PostPaid: pay-as-you-go,PrePaid: subscription"},
			{Name: "machine_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "private_pool_options_id", Type: proto.ColumnType_STRING, Description: "The hostname of the instance."},
			{Name: "cluster_id", Type: proto.ColumnType_STRING, Description: "The ID of the cluster."},
			{Name: "private_pool_options_match_criteria", Type: proto.ColumnType_STRING, Description: "The name of the instance."},
			{Name: "deployment_set_group_no", Type: proto.ColumnType_INT, Description: "The group No. of the instance in a deployment set when the deployment set is used to distribute instances across multiple physical machines."},
			{Name: "credit_specification", Type: proto.ColumnType_STRING, Description: "The performance mode of the burstable instance."},
			{Name: "gpu_amount", Type: proto.ColumnType_INT, Description: "The number of GPUs for the instance type."},
			{Name: "connected", Type: proto.ColumnType_BOOL, Description: ""},

			{Name: "invocation_count", Type: proto.ColumnType_INT, Description: ""},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "The start time of the bidding mode for the preemptible instance."},

			{Name: "zone_id", Type: proto.ColumnType_STRING, Description: "The ID of the zone."},
			{Name: "internet_charge_type", Type: proto.ColumnType_STRING, Description: "The billing method for network usage. Valid values:PayByBandwidth,PayByTraffic"},
			{Name: "internet_max_bandwidth_in", Type: proto.ColumnType_INT, Description: "The maximum inbound bandwidth from the Internet. Unit: Mbit/s."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the instance. Valid values:Pending,Running,Starting,Stopping,Stopped"},
			{Name: "cpu", Type: proto.ColumnType_INT, Description: "The number of vCPUs."},
			{Name: "isp", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "os_version", Type: proto.ColumnType_STRING, Description: "The name of the operating system for the instance."},
			{Name: "spot_price_limit", Type: proto.ColumnType_DOUBLE, Description: "The maximum hourly price for the instance. It can be accurate to three decimal places. This parameter takes effect when the SpotStrategy parameter is set to SpotWithPriceLimit."},
			{Name: "os_name", Type: proto.ColumnType_STRING, Description: "The name of the operating system for the instance."},
			{Name: "os_name_en", Type: proto.ColumnType_STRING, Description: "The English name of the operating system for the instance."},
			{Name: "serial_number", Type: proto.ColumnType_STRING, Description: "The serial number of the instance."},
			{Name: "region_id", Type: proto.ColumnType_STRING, Description: "The region ID of the instance."},

			{Name: "io_optimized", Type: proto.ColumnType_BOOL, Description: "Specifies whether the instance is I/O optimized."},
			{Name: "internet_max_bandwidth_out", Type: proto.ColumnType_INT, Description: "The maximum outbound bandwidth to the Internet. Unit: Mbit/s."},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the instance belongs. If this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response."},
			{Name: "activation_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "instance_type_family", Type: proto.ColumnType_STRING, Description: "The instance family."},
			{Name: "instance_id", Type: proto.ColumnType_STRING, Description: "The ID of the instance."},
			{Name: "deployment_set_id", Type: proto.ColumnType_STRING, Description: "The ID of the deployment set."},
			{Name: "gpu_spec", Type: proto.ColumnType_STRING, Description: "The category of GPUs for the instance type."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the instance."},
			{Name: "recyclable", Type: proto.ColumnType_BOOL, Description: "Indicates whether the instance can be recycled."},
			{Name: "sale_cycle", Type: proto.ColumnType_STRING, Description: "The billing cycle of the instance."},
			{Name: "expired_time", Type: proto.ColumnType_TIMESTAMP, Description: "The expiration time of the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format."},

			{Name: "internet_ip", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "The memory size of the instance. Unit: MiB."},
			{Name: "creation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the instance was created."},
			{Name: "agent_version", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "key_pair_name", Type: proto.ColumnType_STRING, Description: "The name of the SSH key pair for the instance."},
			{Name: "hpc_cluster_id", Type: proto.ColumnType_STRING, Description: "The ID of the HPC cluster to which the instance belongs."},
			{Name: "local_storage_capacity", Type: proto.ColumnType_INT, Description: "The capacity of local disks attached to the instance."},
			{Name: "vlan_id", Type: proto.ColumnType_STRING, Description: "The VLAN ID of the instance."},
			{Name: "stopped_mode", Type: proto.ColumnType_STRING, Description: "Indicates whether the instance continues to be billed after it is stopped."},
			{Name: "spot_strategy", Type: proto.ColumnType_STRING, Description: "The preemption policy for the pay-as-you-go instance."},

			{Name: "spot_duration", Type: proto.ColumnType_INT, Description: "The protection period of the preemptible instance. Unit: hours."},
			{Name: "deletion_protection", Type: proto.ColumnType_BOOL, Description: "The VLAN ID of the instance."},
			{Name: "security_group_ids", Type: proto.ColumnType_JSON, Description: "The IDs of security groups to which the instance belongs."},
			{Name: "inner_ip_address", Type: proto.ColumnType_JSON, Description: "The internal IP addresses of classic network-type instances. This parameter takes effect when InstanceNetworkType is set to classic. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,)."},
			{Name: "public_ip_address", Type: proto.ColumnType_JSON, Description: "The public IP addresses of instances."},
			{Name: "metadata_options", Type: proto.ColumnType_JSON, Description: "The collection of metadata options."},
			{Name: "dedicated_host_attribute", Type: proto.ColumnType_JSON, Description: "The host attribute array that consists of the DedicatedHostClusterId, DedicatedHostId, and DedicatedHostName values."},
			{Name: "eip_address", Type: proto.ColumnType_JSON, Description: "The information of the EIP associated with the instance."},
			{Name: "cpu_options", Type: proto.ColumnType_JSON, Description: "The configuration details of CPU."},

			{Name: "ecs_capacity_reservation_attr", Type: proto.ColumnType_JSON, Description: "The capacity reservation attribute of the ECS instance."},
			{Name: "dedicated_instance_attribute", Type: proto.ColumnType_JSON, Description: "The attribute of the instance on a dedicated host."},
			{Name: "vpc_attributes", Type: proto.ColumnType_JSON, Description: "The VPC attributes of the instance."},
			{Name: "operation_locks", Type: proto.ColumnType_JSON, Description: "Details about the reasons why the instance was locked."},
			{Name: "network_interfaces", Type: proto.ColumnType_JSON, Description: "Details about the ENIs bound to the instance."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: "A list of tags attached with the resource."},

			// steampipe standard columns
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag").Transform(ecsTagsToMap), Description: ColumnDescriptionTags},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("InstanceName"), Description: ColumnDescriptionTitle},
			{Name: "akas", Type: proto.ColumnType_JSON, Hydrate: getEcsInstanceAka, Transform: transform.FromValue(), Description: ColumnDescriptionAkas},

			// alicloud standard columns
			{Name: "region", Description: ColumnDescriptionRegion, Type: proto.ColumnType_STRING, Transform: transform.FromField("RegionId")},
			{Name: "account_id", Description: ColumnDescriptionAccount, Type: proto.ColumnType_STRING, Hydrate: getCommonColumns, Transform: transform.FromField("AccountID")},
		},
	}
}

func listEcsInstance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_bucket.listEcsInstance", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	// if quals["is_default"] != nil {
	// 	request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	// }
	if quals["id"] != nil {
		request.InstanceIds = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeInstances(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs.listEcsInstance", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Instances.Instance {
			plugin.Logger(ctx).Warn("alicloud_ecs.listEcsInstance", "tags", i.Tags, "item", i)
			d.StreamListItem(ctx, i)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}

func getEcsInstanceAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsInstanceAka")
	instance := h.Item.(ecs.Instance)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + instance.RegionId + ":" + accountID + ":instance/" + instance.InstanceId}

	return akas, nil
}
