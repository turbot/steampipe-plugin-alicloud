package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_account",
		Description: "Alicloud Account",
		List: &plugin.ListConfig{
			Hydrate: listAccountAlias,
		},
		Columns: []*plugin.Column{
			{
				Name:        "alias",
				Type:        proto.ColumnType_STRING,
				Description: "Specify the alias associated with the account.",
				Transform:   transform.FromField("AccountAlias"),
			},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAccountAkas,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccountAlias"),
			},

			// Alicloud standard columns
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

func listAccountAlias(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := RAMService(ctx, d)
	if err != nil {
		return nil, err
	}
	request := ram.CreateGetAccountAliasRequest()
	request.Scheme = "https"

	response, err := client.GetAccountAlias(request)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, response)

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAccountAkas")

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"arn:acs:::" + accountID}, nil
}
