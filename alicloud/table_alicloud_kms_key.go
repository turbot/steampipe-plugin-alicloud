package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudKmsKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_user",
		Description: "Resource Access Management users who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listKmsKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getKmsKey,
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
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(userToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("UserName"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listKmsKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectKms(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.listKmsKey", "connection_error", err)
		return nil, err
	}
	request := kms.CreateListKeysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.ListKeys(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_kms_key.listKmsKey", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Keys.Key {
			plugin.Logger(ctx).Warn("listKmsKey", "item", i)
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

func getKmsKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := connectKms(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "connection_error", err)
		return nil, err
	}

	var name string

	request := kms.CreateDescribeKeyRequest()
	request.Scheme = "https"
	request.KeyId = name

	response, err := client.DescribeKey(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "EntityNotExist.Key" {
			plugin.Logger(ctx).Warn("alicloud_kms_key.getKmsKey", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.KeyMetadata, nil
}

// func userToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	switch d.Value.(type) {
// 	case ram.UserInListUsers:
// 		i := d.Value.(ram.UserInListUsers)
// 		return "acs:ram::" + "ACCOUNT_ID" + ":user/" + i.UserName, nil
// 	case ram.UserInGetUser:
// 		i := d.Value.(ram.UserInGetUser)
// 		return "acs:ram::" + "ACCOUNT_ID" + ":user/" + i.UserName, nil
// 	}
// 	return nil, nil
// }
