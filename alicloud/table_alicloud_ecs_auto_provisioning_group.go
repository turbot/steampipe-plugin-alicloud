package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsAutoProvisioningGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_auto_provisioning_group",
		Description: "Alicloud ECS Auto Provisioning Group",
		List: &plugin.ListConfig{
			Hydrate: listEcsAutosProvisioningGroups,
			Tags:    map[string]string{"service": "ecs", "action": "DescribeAutoProvisioningGroups"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("auto_provisioning_group_id"),
			Hydrate:    getEcsAutosProvisioningGroup,
			Tags:       map[string]string{"service": "ecs", "action": "DescribeAutoProvisioningGroups"},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getEcsAutosProvisioningGroupInstances,
				Tags: map[string]string{"service": "ecs", "action": "DescribeAutoProvisioningGroupInstances"},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AutoProvisioningGroupName"),
			},
			{
				Name:        "auto_provisioning_group_id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auto_provisioning_group_type",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "allocation_strategy",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PayAsYouGoOptions.AllocationStrategy"),
			},
			{
				Name:        "excess_capacity_termination_policy",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "launch_template_id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "launch_template_version",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_spot_price",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "terminate_instances",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "terminate_instances_with_expiration",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "valid_from",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "valid_until",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "instances",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutosProvisioningGroupInstances,
				Transform:   transform.FromField("Instances.Instance"),
			},
			{
				Name:        "launch_template_configs",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "spot_options",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "target_capacity_specification",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsAutosProvisioningGroupAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ecsAutosProvisioningGroupTitle),
			},

			// Alicloud standard columns
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

func listEcsAutosProvisioningGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_auto_provisioning_group.listEcsAutosProvisioningGroups", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeAutoProvisioningGroupsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeAutoProvisioningGroups(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_auto_provisioning_group.listEcsAutosProvisioningGroups", "query_error", err, "request", request)
			return nil, err
		}
		for _, group := range response.AutoProvisioningGroups.AutoProvisioningGroup {
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

func getEcsAutosProvisioningGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsAutosProvisioningGroup")

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_auto_provisioning_group.getEcsAutosProvisioningGroup", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["auto_provisioning_group_id"].GetStringValue()

	request := ecs.CreateDescribeAutoProvisioningGroupsRequest()
	request.Scheme = "https"
	request.AutoProvisioningGroupId = &[]string{id}

	response, err := client.DescribeAutoProvisioningGroups(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_auto_provisioning_group.getEcsAutosProvisioningGroup", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.AutoProvisioningGroups.AutoProvisioningGroup) > 0 {
		return response.AutoProvisioningGroups.AutoProvisioningGroup[0], nil
	}

	return nil, nil
}

func getEcsAutosProvisioningGroupInstances(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsAutosProvisioningGroupInstances")
	data := h.Item.(ecs.AutoProvisioningGroup)

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_auto_provisioning_group.getEcsAutosProvisioningGroupInstances", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeAutoProvisioningGroupInstancesRequest()
	request.Scheme = "https"
	request.AutoProvisioningGroupId = data.AutoProvisioningGroupId

	response, err := client.DescribeAutoProvisioningGroupInstances(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_auto_provisioning_group.getEcsAutosProvisioningGroupInstances", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}

func getEcsAutosProvisioningGroupAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsAutosProvisioningGroupAka")
	data := h.Item.(ecs.AutoProvisioningGroup)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ecs:" + data.RegionId + ":" + accountID + ":auto-provisioning-group/" + data.AutoProvisioningGroupId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func ecsAutosProvisioningGroupTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(ecs.AutoProvisioningGroup)

	// Build resource title
	title := data.AutoProvisioningGroupId

	if len(data.AutoProvisioningGroupName) > 0 {
		title = data.AutoProvisioningGroupName
	}

	return title, nil
}
