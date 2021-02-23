package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudRamGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_group",
		Description: "Resource Access Management groups who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listRamGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getRamGroup,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("GroupName"), Description: "The name of the RAM user group."},
			// TODO: Not available - {Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("GroupId"), Description: "The ID of the RAM user group."},
			// TODO: Not avialable - {Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the RAM group."},
			// Other columns
			{Name: "comments", Type: proto.ColumnType_STRING, Description: "The description of the RAM user group."},
			{Name: "create_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the RAM user group was created."},
			{Name: "update_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the RAM user group was modified."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(groupToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("GroupName"), Description: ColumnDescriptionTitle},
		},
	}
}

func listRamGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.listRamGroup", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListGroupsRequest()
	request.Scheme = "https"

	for {
		response, err := client.ListGroups(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ram_group.listRamGroup", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Groups.Group {
			plugin.Logger(ctx).Warn("listRamGroup", "item", i)
			d.StreamListItem(ctx, i)
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

func getRamGroup(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRamGroup", "connection_error", err)
		return nil, err
	}

	var name string

	if h.Item != nil {
		i := h.Item.(ram.GroupInListGroups)
		name = i.GroupName
	} else {
		quals := d.KeyColumnQuals
		name = quals["name"].GetStringValue()
	}

	request := ram.CreateGetGroupRequest()
	request.Scheme = "https"
	request.GroupName = name

	response, err := client.GetGroup(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "EntityNotExist.Group" {
			plugin.Logger(ctx).Warn("alicloud_ram_group.getRamGroup", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_ram_group.getRamGroup", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.Group, nil
}

func groupToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch d.Value.(type) {
	case ram.GroupInListGroups:
		i := d.Value.(ram.GroupInListGroups)
		return "acs:ram::" + "ACCOUNT_ID" + ":group/" + i.GroupName, nil
	case ram.GroupInGetGroup:
		i := d.Value.(ram.GroupInGetGroup)
		return "acs:ram::" + "ACCOUNT_ID" + ":group/" + i.GroupName, nil
	}
	return nil, nil
}
