package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudEcsSecurityGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_security_group",
		Description: "VSwitches to divide the VPC network into one or more subnets.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listSecurityGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSecurityGroupAttribute,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecurityGroupName"), Description: "The name of the security group."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecurityGroupId"), Description: "The unique ID of the destination security group.."},
			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the security group."},
			{Name: "creation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The creation time of the security group."},
			{Name: "vpc_id", Type: proto.ColumnType_STRING, Description: "The ID of the VPC. If a VPC ID is returned, the network type of the security group is VPC. Otherwise, the network type of the security group is classic network."},
			{Name: "security_group_type", Type: proto.ColumnType_STRING, Description: "The type of the security group. Valid values: normal and enterprise"},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the security group belongs."},
			{Name: "service_id", Type: proto.ColumnType_INT, Description: "The ID of the distributor to which the security group belongs."},
			{Name: "ServiceManaged", Type: proto.ColumnType_BOOL, Description: "Indicates whether the user is an Alibaba Cloud service or a distributor."},
			{Name: "region_id", Type: proto.ColumnType_STRING, Description: "The region ID of the security group.", Hydrate: getSecurityGroupAttribute},
			{Name: "permissions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Permissions.Permission"), Description: "Details about the security group rules.", Hydrate: getSecurityGroupAttribute},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "The tags of the security groups."},
			// Resource interface
			// {Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(vswitchToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			// TODO - It appears that Tags are not returned by the go SDK?
			// {Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("NaSecurityGroupNameme"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listSecurityGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.listSecurityGroups", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeSecurityGroups(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_security_group.listSecurityGroups", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.SecurityGroups.SecurityGroup {
			plugin.Logger(ctx).Warn("alicloud_ecs_security_group.listSecurityGroups", "query_error", err, "item", i)
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

func getSecurityGroupAttribute(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getVSecurityGroupAttribute", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeSecurityGroupAttributeRequest()
	request.Scheme = "https"
	// i := h.Item.(ecs.SecurityGroup)
	// request.SecurityGroupId = i.SecurityGroupId

	quals := d.KeyColumnQuals
	if quals["id"] != nil {
		request.SecurityGroupId = quals["id"].GetStringValue()
	}

	if len(request.SecurityGroupId) < 1 {
		return nil, nil
	}
	response, err := client.DescribeSecurityGroupAttribute(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getVSecurityGroupAttribute", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

// func vswitchToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := d.Value.(vpc.VSwitch)
// 	return "acs:vswitch:" + i.ZoneId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vswitch/" + i.VSwitchId, nil
// }
