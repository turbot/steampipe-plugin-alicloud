package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type groupInfo = struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}

//// TABLE DEFINITION

func tableAlicloudRAMGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_group",
		Description: "Resource Access Management groups who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listRAMGroup,
			Tags:    map[string]string{"service": "ram", "action": "ListGroups"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getRAMGroup,
			Tags:       map[string]string{"service": "ram", "action": "GetGroup"},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getRAMGroupUsers,
				Tags: map[string]string{"service": "ram", "action": "ListUsersForGroup"},
			},
			{
				Func: getRAMGroupPolicies,
				Tags: map[string]string{"service": "ram", "action": "ListPoliciesForGroup"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupName"),
				Description: "The name of the RAM user group.",
			},
			// TODO: Not available - {Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("GroupId"), Description: "The ID of the RAM user group."},
			// TODO: Not avialable - {Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the RAM group."},
			// Other columns
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the RAM user group.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGroupArn,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "comments",
				Description: "The description of the RAM user group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_date",
				Description: "The time when the RAM user group was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "update_date",
				Description: "The time when the RAM user group was modified.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "attached_policy",
				Description: "A list of policies attached to a RAM user group.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMGroupPolicies,
				Transform:   transform.FromField("Policies.Policy"),
			},
			{
				Name:        "users",
				Description: "A list of users in the group.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMGroupUsers,
				Transform:   transform.FromValue(),
			},

			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGroupArn,
				Transform:   transform.FromValue().Transform(ensureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupName"),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromConstant("global"),
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

func listRAMGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.listRAMGroup", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListGroupsRequest()
	request.Scheme = "https"

	for {
		response, err := client.ListGroups(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ram_group.listRAMGroup", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Groups.Group {
			plugin.Logger(ctx).Warn("listRAMGroup", "item", i)
			d.StreamListItem(ctx, groupInfo{i.GroupName, i.Comments, i.CreateDate, i.UpdateDate})
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getRAMGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMGroup")

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMGroup", "connection_error", err)
		return nil, err
	}

	var name string

	if h.Item != nil {
		i := h.Item.(ram.Group)
		name = i.GroupName
	} else {
		quals := d.EqualsQuals
		name = quals["name"].GetStringValue()
	}

	request := ram.CreateGetGroupRequest()
	request.Scheme = "https"
	request.GroupName = name

	response, err := client.GetGroup(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMGroup", "query_error", err, "request", request)
		return nil, err
	}

	data := response.Group
	return groupInfo{data.GroupName, data.Comments, data.CreateDate, data.UpdateDate}, nil
}

func getRAMGroupUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMGroupUsers")
	data := h.Item.(groupInfo)

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMGroupUsers", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListUsersForGroupRequest()
	request.Scheme = "https"
	request.GroupName = data.GroupName

	response, err := client.ListUsersForGroup(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMGroupUsers", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.Users.User, nil
}

func getRAMGroupPolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMGroupPolicies")
	data := h.Item.(groupInfo)

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMGroupPolicies", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListPoliciesForGroupRequest()
	request.Scheme = "https"
	request.GroupName = data.GroupName

	response, err := client.ListPoliciesForGroup(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMGroupPolicies", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}

func getGroupArn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getGroupAkas")
	data := h.Item.(groupInfo)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return "acs:ram::" + accountID + ":group/" + data.GroupName, nil
}
