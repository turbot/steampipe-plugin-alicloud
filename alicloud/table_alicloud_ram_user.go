package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudRamUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_user",
		Description: "Resource Access Management users who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listRamUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getRamUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("UserName"), Description: "The username of the RAM user."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("UserId"), Description: "The unique ID of the RAM user."},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "The display name of the RAM user."},
			// Other columns
			{Name: "email", Type: proto.ColumnType_STRING, Hydrate: getRamUser, Description: "The email address of the RAM user."},
			{Name: "last_login_date", Type: proto.ColumnType_TIMESTAMP, Hydrate: getRamUser, Description: "The time when the RAM user last logged on to the console by using the password."},
			{Name: "mobile_phone", Type: proto.ColumnType_STRING, Hydrate: getRamUser, Description: "The mobile phone number of the RAM user."},
			{Name: "comments", Type: proto.ColumnType_STRING, Description: "The description of the RAM user."},
			{Name: "create_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the RAM user was created."},
			{Name: "update_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the RAM user was modified."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Hydrate: getUserAkas, Transform: transform.FromValue(), Description: ColumnDescriptionAkas},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: ColumnDescriptionTags},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("UserName"), Description: ColumnDescriptionTitle},

			// alicloud standard columns
			{Name: "region", Description: ColumnDescriptionRegion, Type: proto.ColumnType_STRING, Transform: transform.FromConstant("global")},
			{Name: "account_id", Description: ColumnDescriptionAccount, Type: proto.ColumnType_STRING, Hydrate: getCommonColumns, Transform: transform.FromField("AccountID")},
		},
	}
}

func listRamUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_user.listRamUser", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListUsersRequest()
	request.Scheme = "https"
	for {
		response, err := client.ListUsers(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ram_user.listRamUser", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Users.User {
			plugin.Logger(ctx).Warn("listRamUser", "item", i)
			d.StreamListItem(ctx, i)
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

func getRamUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_user.getRamUser", "connection_error", err)
		return nil, err
	}

	var name string

	if h.Item != nil {
		i := h.Item.(ram.UserInListUsers)
		name = i.UserName
	} else {
		quals := d.KeyColumnQuals
		name = quals["name"].GetStringValue()
	}

	request := ram.CreateGetUserRequest()
	request.Scheme = "https"
	request.UserName = name

	response, err := client.GetUser(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "EntityNotExist.User" {
			plugin.Logger(ctx).Warn("alicloud_ram_user.getRamUser", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_ram_user.getRamUser", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.User, nil
}

func getUserAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var name string
	if h.Item != nil {
		i := h.Item.(ram.UserInListUsers)
		name = i.UserName
	} else {
		quals := d.KeyColumnQuals
		name = quals["name"].GetStringValue()
	}

	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}

	accountCommonData := commonData.(*alicloudCommonColumnData)
	return []string{"acs:ram::" + accountCommonData.AccountID + ":user/" + name}, nil
}
