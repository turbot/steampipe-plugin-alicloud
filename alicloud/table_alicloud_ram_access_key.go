package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

type accessKeyRow = struct {
	AccessKeyId string
	Status      string
	CreateDate  string
	UserName    string
}

//// TABLE DEFINITION

func tableAlicloudRAMAccessKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_access_key",
		Description: "Alibaba Cloud RAM User Access Key.",
		List: &plugin.ListConfig{
			ParentHydrate: listRAMUser,
			Hydrate:       listRAMUserAccessKeys,
		},
		Columns: []*plugin.Column{
			{
				Name:        "user_name",
				Description: "Name of the User that the access key belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "access_key_id",
				Description: "The AccessKey ID.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the AccessKey pair. Valid values: Active and Inactive.",
			},
			{
				Name:        "create_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the AccessKey pair was created.",
			},

			// steampipe common columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccessKeyId"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAccessKeyArn,
				Transform:   transform.FromValue(),
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

func listRAMUserAccessKeys(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_access_key.listRAMUserAccessKeys", "connection_error", err)
		return nil, err
	}

	user := h.Item.(userInfo)

	request := ram.CreateListAccessKeysRequest()
	request.Scheme = "https"
	request.UserName = user.UserName

	response, err := client.ListAccessKeys(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_access_key.listRAMUserAccessKeys", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.AccessKeys.AccessKey {
		plugin.Logger(ctx).Warn("listRAMUserAccessKeys", "item", i)
		d.StreamLeafListItem(ctx, accessKeyRow{i.AccessKeyId, i.Status, i.CreateDate, user.UserName})
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccessKeyArn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAccessKeyArn")

	i := h.Item.(accessKeyRow)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ram::" + accountID + ":user/" + i.UserName + "/accesskey/" + i.AccessKeyId}

	return akas, nil
}
