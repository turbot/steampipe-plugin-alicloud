package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsAutoscalingGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_autoscaling_group",
		Description: "Elastic Compute Autoscaling Group",
		List: &plugin.ListConfig{
			Hydrate: listEcsAutoscalingGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("scaling_group_id"),
			Hydrate:    getEcsAutoscalingGroup,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ScalingGroupName"),
			},
			{
				Name:        "scaling_group_id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "scaling_policy",
				Description: "Specifies the reclaim policy of the scaling group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "group_deletion_protection",
				Description: "Indicates whether scaling group deletion protection is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "vswitch_id",
				Description: "The ID of the VSwitch that is associated with the scaling group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VSwitchId"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC to which the scaling group belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "active_capacity",
				Description: "The number of ECS instances that have been added to the scaling group and are running properly.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "active_scaling_configuration_id",
				Description: "The ID of the active scaling configuration in the scaling group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "compensate_with_on_demand",
				Description: "Specifies whether to automatically create pay-as-you-go instances to meet the requirement for the number of ECS instances in the scaling group when the number of preemptible instances cannot be reached due to reasons such as cost or insufficient resources.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "creation_time",
				Description: "The time when the scaling group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "default_cooldown",
				Description: "The default cooldown period of the scaling group (in seconds).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "desired_capacity",
				Description: "The expected number of ECS instances in the scaling group. Auto Scaling automatically keeps the ECS instances at this number.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "health_check_type",
				Description: "The health check mode of the scaling group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "launch_template_id",
				Description: "The ID of the launch template used by the scaling group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "launch_template_version",
				Description: "The version of the launch template used by the scaling group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "life_cycle_state",
				Description: "The lifecycle status of the scaling group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_size",
				Description: "The maximum number of ECS instances in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "min_size",
				Description: "The minimum number of ECS instances in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "modification_time",
				Description: "The time when the scaling group was modified.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "multi_az_policy",
				Description: "The ECS instance scaling policy for a multi-zone scaling group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MultiAZPolicy"),
			},
			{
				Name:        "on_demand_base_capacity",
				Description: "The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "on_demand_percentage_above_base_capacity",
				Description: "The percentage of pay-as-you-go instances to be created when instances are added to the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "pending_capacity",
				Description: "The number of ECS instances that are being added to the scaling group, but are still being configured.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "pending_wait_capacity",
				Description: "The number of ECS instances that are in the pending state to be added in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "protected_capacity",
				Description: "The number of ECS instances that are in the protected state in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "removing_capacity",
				Description: "The number of ECS instances that are being removed from the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "removing_wait_capacity",
				Description: "The number of ECS instances that are in the pending state to be removed from the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "spot_instance_pools",
				Description: "The number of available instance types. Auto Scaling will create preemptible instances of multiple instance types available at the lowest cost. Valid values: 0 to 10.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "spot_instance_remedy",
				Description: "Specifies whether to supplement preemptible instances when the target capacity of preemptible instances is not fulfilled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "standby_capacity",
				Description: "The number of instances that are in the standby state in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "stopped_capacity",
				Description: "The number of instances that are in the stopped state in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "total_capacity",
				Description: "The total number of ECS instances in the scaling group.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "db_instance_ids",
				Description: "The IDs of the ApsaraDB RDS instances that are associated with the scaling group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DBInstanceIds.DBInstanceId"),
			},
			{
				Name:        "load_balancer_ids",
				Description: "The IDs of the SLB instances that are associated with the scaling group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("LoadBalancerIds.LoadBalancerId"),
			},
			{
				Name:        "removal_policies",
				Description: "Details about policies for removing ECS instances from the scaling group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RemovalPolicies.RemovalPolicy"),
			},
			{
				Name:        "suspended_processes",
				Description: "The scaling activity that is suspended. If no scaling activity is suspended, the returned value is null.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SuspendedProcesses.SuspendedProcess"),
			},
			{
				Name:        "vserver_groups",
				Description: "Details about backend server groups.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VServerGroups.VServerGroup"),
			},
			{
				Name:        "vswitch_ids",
				Description: "A collection of IDs of the VSwitches that are associated with the scaling group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VSwitchIds.VSwitchId"),
			},
			{
				Name:        "scaling_configurations",
				Description: "A list of scaling configurations.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutoscalingGroupConfigurations,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "scaling_instances",
				Description: "A list of ECS instances in a scaling group.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutoscalingGroupScalingInstances,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutoscalingGroupTags,
				Transform:   transform.FromField("TagResources.TagResource").Transform(modifyEssSourceTags),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutoscalingGroupTags,
				Transform:   transform.FromField("TagResources.TagResource").Transform(essTurbotTags),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutoscalingGroupAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ScalingGroupName"),
			},

			// alicloud standard columns
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

func listEcsAutoscalingGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := AutoscalingService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.listEcsAutoscalingGroup", "connection_error", err)
		return nil, err
	}
	request := ess.CreateDescribeScalingGroupsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeScalingGroups(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.listEcsAutoscalingGroup", "query_error", err, "request", request)
			return nil, err
		}
		for _, group := range response.ScalingGroups.ScalingGroup {
			plugin.Logger(ctx).Warn("listEcsAutoscalingGroup", "item", group)
			d.StreamListItem(ctx, group)
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

func getEcsAutoscalingGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsAutoscalingGroup")

	// Create service connection
	client, err := AutoscalingService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroup", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(ess.ScalingGroup)
		id = data.ScalingGroupId
	} else {
		id = d.KeyColumnQuals["scaling_group_id"].GetStringValue()
	}

	request := ess.CreateDescribeScalingGroupsRequest()
	request.Scheme = "https"
	request.ScalingGroupId1 = id
	response, err := client.DescribeScalingGroups(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroup", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.ScalingGroups.ScalingGroup != nil && len(response.ScalingGroups.ScalingGroup) > 0 {
		return response.ScalingGroups.ScalingGroup[0], nil
	}

	return nil, nil
}

func getEcsAutoscalingGroupConfigurations(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsAutoscalingGroupConfigurations")
	data := h.Item.(ess.ScalingGroup)

	// Create service connection
	client, err := AutoscalingService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroupConfigurations", "connection_error", err)
		return nil, err
	}

	request := ess.CreateDescribeScalingConfigurationsRequest()
	request.Scheme = "https"
	request.ScalingGroupId = data.ScalingGroupId

	response, err := client.DescribeScalingConfigurations(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroupConfigurations", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.ScalingConfigurations.ScalingConfiguration != nil {
		return response.ScalingConfigurations.ScalingConfiguration, nil
	}

	return nil, nil
}

func getEcsAutoscalingGroupScalingInstances(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsAutoscalingGroupScalingInstances")
	data := h.Item.(ess.ScalingGroup)

	// Create service connection
	client, err := AutoscalingService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroupScalingInstances", "connection_error", err)
		return nil, err
	}

	request := ess.CreateDescribeScalingInstancesRequest()
	request.Scheme = "https"
	request.ScalingGroupId = data.ScalingGroupId

	response, err := client.DescribeScalingInstances(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroupScalingInstances", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.ScalingInstances.ScalingInstance != nil {
		return response.ScalingInstances.ScalingInstance, nil
	}

	return nil, nil
}

func getEcsAutoscalingGroupTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsAutoscalingGroupTags")
	data := h.Item.(ess.ScalingGroup)

	// Create service connection
	client, err := AutoscalingService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroupTags", "connection_error", err)
		return nil, err
	}

	request := ess.CreateListTagResourcesRequest()
	request.Scheme = "https"
	request.ResourceType = "scalingGroup"
	request.ResourceId = &[]string{data.ScalingGroupId}

	response, err := client.ListTagResources(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_autoscaling_group.getEcsAutoscalingGroupTags", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}

func getEcsAutoscalingGroupAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsAutoscalingGroupAka")
	data := h.Item.(ess.ScalingGroup)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ess:" + data.RegionId + ":" + accountID + ":scalinggroup/" + data.ScalingGroupId}

	return akas, nil
}
