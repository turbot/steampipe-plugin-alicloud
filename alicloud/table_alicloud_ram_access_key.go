package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableAlicloudRamAccessKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_access_key",
		Description: "Alibaba Cloud RAM User Access Key",
		List: &plugin.ListConfig{
			ParentHydrate: listRamUser,
			Hydrate:       listRamUserAccessKey,
		},
		Columns: []*plugin.Column{
			{Name: "user_name", Type: proto.ColumnType_STRING, Description: "Name of the User that the access key belongs to."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the AccessKey pair. Valid values: Active and Inactive."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The AccessKey ID."},
			{Name: "create_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the AccessKey pair was created."},
		},
	}
}

type accessKeyRow struct {
	ram.AccessKeyInListAccessKeys
	UserName string
}

func listRamUserAccessKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectRam(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_access_key.listRamAccessKey", "connection_error", err)
		return nil, err
	}

	var name string

	if h.Item != nil {
		switch h.Item.(type) {
		case ram.UserInListUsers:
			i := h.Item.(ram.UserInListUsers)
			name = i.UserName
		case accessKeyRow:
			i := h.Item.(accessKeyRow)
			name = i.UserName
		}
	} else {
		quals := d.KeyColumnQuals
		name = quals["name"].GetStringValue()
	}

	request := ram.CreateListAccessKeysRequest()
	request.Scheme = "https"
	request.UserName = name

	response, err := client.ListAccessKeys(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_access_key.listRamAccessKey", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.AccessKeys.AccessKey {
		plugin.Logger(ctx).Warn("listRamAccessKey", "item", i)
		ak := accessKeyRow{i, name}
		d.StreamListItem(ctx, ak)
	}
	return nil, nil
}
