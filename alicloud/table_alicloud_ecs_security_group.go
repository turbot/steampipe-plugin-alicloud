package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type securityGroupInfo = struct {
	SecurityGroup ecs.SecurityGroup
	Region        string
}

func tableAlicloudEcsSecurityGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_security_group",
		Description: "AliCloud ECS Security Group",
		List: &plugin.ListConfig{
			Hydrate: listEcsSecurityGroups,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEcsSecurityGroup,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroup.SecurityGroupName"),
			},
			{
				Name:        "id",
				Description: "The ID of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroup.SecurityGroupId"),
			},
			{
				Name:        "type",
				Description: "The type of the security group. Possible values are: normal, and enterprise.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroup.SecurityGroupType"),
			},
			{
				Name:        "vpc_id",
				Description: "he ID of the VPC to which the security group belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroup.VpcId"),
			},
			{
				Name:        "creation_time",
				Description: "The time when the security group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("SecurityGroup.CreationTime"),
			},
			{
				Name:        "description",
				Description: "The description of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroup.Description"),
			},
			{
				Name:        "inner_access_policy",
				Description: "The description of the security group.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getSecurityGroupAttribute,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the security group belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroup.ResourceGroupId"),
			},
			{
				Name:        "service_id",
				Description: "The ID of the distributor to which the security group belongs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("SecurityGroup.ServiceID"),
			},
			{
				Name:        "service_managed",
				Description: "Indicates whether the user is an Alibaba Cloud service or a distributor.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("SecurityGroup.ServiceManaged"),
			},
			{
				Name:        "permissions",
				Description: "Details about the security group rules.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSecurityGroupAttribute,
				Transform:   transform.FromField("Permissions.Permission"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the security group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SecurityGroup.Tags.Tag"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(ecsSecurityGroupTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(ecsSecurityGroupTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsSecurityGroupAka,
				Transform:   transform.FromValue(),
			},

			// alicloud standard columns
			{
				Name:        "region_id",
				Description: "The name of the region where the resource belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Region"),
			},
			{
				Name:        "account_id",
				Description: "The alicloud Account ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID"),
			},
		},
	}
}

//// LIST FUNCTION

func listEcsSecurityGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.listEcsSecurityGroups", "connection_error", err)
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
			plugin.Logger(ctx).Error("alicloud_ecs_security_group.listEcsSecurityGroups", "query_error", err, "request", request)
			return nil, err
		}
		for _, securityGroup := range response.SecurityGroups.SecurityGroup {
			plugin.Logger(ctx).Warn("alicloud_ecs_security_group.listEcsSecurityGroups", "query_error", err, "item", securityGroup)
			d.StreamListItem(ctx, securityGroupInfo{securityGroup, response.RegionId})
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

func getEcsSecurityGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSecurityGroup")

	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getEcsSecurityGroup", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(ecs.SecurityGroup)
		id = data.SecurityGroupId
	} else {
		id = d.KeyColumnQuals["id"].GetStringValue()
	}

	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.Scheme = "https"
	request.SecurityGroupId = id

	response, err := client.DescribeSecurityGroups(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getEcsSecurityGroup", "query_error", err, "request", request)
		return nil, err
	}

	if response.SecurityGroups.SecurityGroup != nil && len(response.SecurityGroups.SecurityGroup) > 0 {
		return securityGroupInfo{response.SecurityGroups.SecurityGroup[0], response.RegionId}, nil
	}

	return nil, nil
}

func getSecurityGroupAttribute(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSecurityGroupAttribute")
	data := h.Item.(securityGroupInfo)

	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getVSecurityGroupAttribute", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeSecurityGroupAttributeRequest()
	request.Scheme = "https"
	request.SecurityGroupId = data.SecurityGroup.SecurityGroupId

	response, err := client.DescribeSecurityGroupAttribute(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getVSecurityGroupAttribute", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func getEcsSecurityGroupAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSecurityGroupAka")
	data := h.Item.(securityGroupInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + data.Region + ":" + accountID + ":securitygroup/" + data.SecurityGroup.SecurityGroupId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func ecsSecurityGroupTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(securityGroupInfo)
	param := d.Param.(string)

	// Build resource title
	title := data.SecurityGroup.SecurityGroupId

	if len(data.SecurityGroup.SecurityGroupName) > 0 {
		title = data.SecurityGroup.SecurityGroupName
	}

	if param == "Title" {
		return title, nil
	}

	return ecsTagsToMap(data.SecurityGroup.Tags.Tag)
}
