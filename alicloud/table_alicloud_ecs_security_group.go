package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlicloudEcsSecurityGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_security_group",
		Description: "ECS Security Group",
		List: &plugin.ListConfig{
			Hydrate: listEcsSecurityGroups,
			Tags:    map[string]string{"service": "ecs", "action": "DescribeSecurityGroups"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("security_group_id"),
			Hydrate:    getEcsSecurityGroup,
			Tags:       map[string]string{"service": "ecs", "action": "DescribeSecurityGroups"},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getSecurityGroupAttribute,
				Tags: map[string]string{"service": "ecs", "action": "DescribeSecurityGroupAttribute"},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroupName"),
			},
			{
				Name:        "security_group_id",
				Description: "The ID of the security group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the ECS security group.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsSecurityGroupARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "type",
				Description: "The type of the security group. Possible values are: normal, and enterprise.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecurityGroupType"),
			},
			{
				Name:        "vpc_id",
				Description: "he ID of the VPC to which the security group belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the security group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description of the security group.",
				Type:        proto.ColumnType_STRING,
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
			},
			{
				Name:        "service_id",
				Description: "The ID of the distributor to which the security group belongs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ServiceID"),
			},
			{
				Name:        "service_managed",
				Description: "Indicates whether the user is an Alibaba Cloud service or a distributor.",
				Type:        proto.ColumnType_BOOL,
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
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ecsSecurityGroupTitle),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsSecurityGroupARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: "The name of the region where the resource belongs.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getSecurityGroupRegion,
				Transform:   transform.FromValue(),
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
	// Create service connection
	client, err := ECSService(ctx, d)
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
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeSecurityGroups(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_security_group.listEcsSecurityGroups", "query_error", err, "request", request)
			return nil, err
		}
		for _, securityGroup := range response.SecurityGroups.SecurityGroup {
			plugin.Logger(ctx).Warn("alicloud_ecs_security_group.listEcsSecurityGroups", "query_error", err, "item", securityGroup)
			d.StreamListItem(ctx, securityGroup)
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
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getEcsSecurityGroup", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(ecs.SecurityGroup)
		id = data.SecurityGroupId
	} else {
		id = d.EqualsQuals["security_group_id"].GetStringValue()
	}

	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.Scheme = "https"
	request.SecurityGroupId = id

	response, err := client.DescribeSecurityGroups(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getEcsSecurityGroup", "query_error", err, "request", request)
		return nil, err
	}

	if len(response.SecurityGroups.SecurityGroup) > 0 {
		return response.SecurityGroups.SecurityGroup[0], nil
	}

	return nil, nil
}

func getSecurityGroupAttribute(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSecurityGroupAttribute")
	data := h.Item.(ecs.SecurityGroup)

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getVSecurityGroupAttribute", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeSecurityGroupAttributeRequest()
	request.Scheme = "https"
	request.SecurityGroupId = data.SecurityGroupId

	response, err := client.DescribeSecurityGroupAttribute(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_security_group.getVSecurityGroupAttribute", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func getEcsSecurityGroupARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSecurityGroupARN")
	data := h.Item.(ecs.SecurityGroup)
	region := d.EqualsQualString(matrixKeyRegion)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + region + ":" + accountID + ":securitygroup/" + data.SecurityGroupId

	return arn, nil
}

func getSecurityGroupRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	return region, nil
}

//// TRANSFORM FUNCTIONS

func ecsSecurityGroupTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(ecs.SecurityGroup)

	// Build resource title
	title := data.SecurityGroupId

	if len(data.SecurityGroupName) > 0 {
		title = data.SecurityGroupName
	}

	return title, nil
}
