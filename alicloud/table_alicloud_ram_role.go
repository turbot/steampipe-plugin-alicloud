package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	//"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudRamRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_role",
		Description: "Resource Access Management roles who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listRamRole,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getRamRole,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("RoleName"), Description: "The name of the RAM role."},
			{Name: "arn", Type: proto.ColumnType_STRING, Transform: transform.FromField("Arn"), Description: "The Alibaba Cloud Resource Name (ARN) of the RAM role."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("RoleId"), Description: "The ID of the RAM role."},
			// TODO: Not avialable - {Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the RAM role."},
			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the RAM role."},
			{Name: "max_session_duration", Type: proto.ColumnType_INT, Description: "The maximum session duration of the RAM role."},
			{Name: "create_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the RAM role was created."},
			{Name: "update_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the RAM role was modified."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Arn").Transform(ensureStringArray), Description: ColumnDescriptionAkas},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: ColumnDescriptionTags},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("RoleName"), Description: ColumnDescriptionTitle},

			// alicloud standard columns
			{Name: "region", Description: ColumnDescriptionRegion, Type: proto.ColumnType_STRING, Transform: transform.FromConstant("global")},
			{Name: "account_id", Description: ColumnDescriptionAccount, Type: proto.ColumnType_STRING, Hydrate: getCommonColumns, Transform: transform.FromField("AccountID")},
		},
	}
}

func listRamRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_role.listRamRole", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListRolesRequest()
	request.Scheme = "https"

	for {
		response, err := client.ListRoles(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ram_role.listRamRole", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Roles.Role {
			plugin.Logger(ctx).Warn("listRamRole", "item", i)
			d.StreamListItem(ctx, i)
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

func getRamRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_role.getRamRole", "connection_error", err)
		return nil, err
	}

	var name string

	if h.Item != nil {
		i := h.Item.(ram.RoleInListRoles)
		name = i.RoleName
	} else {
		quals := d.KeyColumnQuals
		name = quals["name"].GetStringValue()
	}

	request := ram.CreateGetRoleRequest()
	request.Scheme = "https"
	request.RoleName = name

	response, err := client.GetRole(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "EntityNotExist.Role" {
			plugin.Logger(ctx).Warn("alicloud_ram_role.getRamRole", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_ram_role.getRamRole", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.Role, nil
}

func roleToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch d.Value.(type) {
	case ram.RoleInListRoles:
		i := d.Value.(ram.RoleInListRoles)
		return "acs:ram::" + "ACCOUNT_ID" + ":role/" + i.RoleName, nil
	case ram.RoleInGetRole:
		i := d.Value.(ram.RoleInGetRole)
		return "acs:ram::" + "ACCOUNT_ID" + ":role/" + i.RoleName, nil
	}
	return nil, nil
}
