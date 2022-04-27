package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRamPasswordPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_password_policy",
		Description: "Alibaba Cloud RAM Password Policy",
		// Avoid NullIfZero since may columns in this table are 0 or false (zero values)
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: listRAMPasswordPolicy,
		},
		Columns: []*plugin.Column{
			{
				Name:        "hard_expiry",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the password has expired.",
			},

			// It's spelt incorrectly in the official API, fix it for Steampipe
			{
				Name:        "max_login_attempts",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("MaxLoginAttemps"),
				Description: "The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password.",
			},
			{
				Name:        "max_password_age",
				Type:        proto.ColumnType_INT,
				Description: "The number of days for which a password is valid. Default value: 0. The default value indicates that the password never expires.",
			},
			{
				Name:        "minimum_password_length",
				Type:        proto.ColumnType_INT,
				Description: "The minimum required number of characters in a password.",
			},
			{
				Name:        "password_reuse_prevention",
				Type:        proto.ColumnType_INT,
				Description: "The number of previous passwords that the user is prevented from reusing. Default value: 0. The default value indicates that the RAM user is not prevented from reusing previous passwords.",
			},
			{
				Name:        "require_lowercase_characters",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether a password must contain one or more lowercase letters.",
			},
			{
				Name:        "require_numbers",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether a password must contain one or more digits.",
			},
			{
				Name:        "require_symbols",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether a password must contain one or more special characters.",
			},
			{
				Name:        "require_uppercase_characters",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether a password must contain one or more uppercase letters.",
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromConstant("global")},
			{
				Name:        "account_id",
				Description: ColumnDescriptionAccount,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID")},
		},
	}
}

//// LIST FUNCTION

func listRAMPasswordPolicy(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listRamPasswordPolicy", "connection_error", err)
		return nil, err
	}
	request := ram.CreateGetPasswordPolicyRequest()
	request.Scheme = "https"
	response, err := client.GetPasswordPolicy(request)
	if err != nil {
		plugin.Logger(ctx).Error("listRamPasswordPolicy", "query_error", err, "request", request)
		return nil, err
	}
	d.StreamListItem(ctx, response.PasswordPolicy)
	return nil, nil
}
