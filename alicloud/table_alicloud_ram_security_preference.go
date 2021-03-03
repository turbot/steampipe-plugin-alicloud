package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRAMSecurityPreference(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_security_preference",
		Description: "Alibaba Cloud RAM Security Preference",
		// Avoid NullIfZero since may columns in this table are 0 or false (zero values)
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: listRAMSecurityPreference,
		},
		Columns: []*plugin.Column{
			{
				Name:        "allow_user_to_change_password",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("LoginProfilePreference.AllowUserToChangePassword"),
				Description: "Indicates whether RAM users can change their passwords.",
			},
			{
				Name:        "allow_user_to_manage_access_keys",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("AccessKeyPreference.AllowUserToManageAccessKeys"),
				Description: "Indicates whether RAM users can manage their AccessKey pairs.",
			},
			{
				Name:        "allow_user_to_manage_mfa_devices",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("MFAPreference.AllowUserToManageMFADevices"),
				Description: "Indicates whether RAM users can manage their MFA devices.",
			},
			{
				Name:        "allow_user_to_manage_public_keys",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("PublicKeyPreference.AllowUserToManagePublicKeys"),
				Description: "Indicates whether RAM users can manage their public keys.",
			},
			{
				Name:        "enable_save_mfa_ticket",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("LoginProfilePreference.EnableSaveMFATicket"),
				Description: "Indicates whether RAM users can save security codes for multi-factor authentication (MFA) during logon. Each security code is valid for seven days.",
			},
			{
				Name:        "login_network_masks",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("LoginProfilePreference.LoginNetworkMasks").TransformP(csvToStringArray, ";"),
				Description: "The subnet mask that indicates the IP addresses from which logon to the Alibaba Cloud Management Console is allowed. This parameter applies to password-based logon and single sign-on (SSO). However, this parameter does not apply to API calls that are authenticated based on AccessKey pairs. May be more than one CIDR range. If empty then login is allowed from any source.",
			},
			{
				Name:        "login_session_duration",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("LoginProfilePreference.LoginSessionDuration"),
				Description: "The validity period of a logon session of a RAM user. Unit: hours.",
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

func listRAMSecurityPreference(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listRamSecurityPreference", "connection_error", err)
		return nil, err
	}
	request := ram.CreateGetSecurityPreferenceRequest()
	request.Scheme = "https"
	response, err := client.GetSecurityPreference(request)
	if err != nil {
		plugin.Logger(ctx).Error("listRamSecurityPreference", "query_error", err, "request", request)
		return nil, err
	}
	d.StreamListItem(ctx, response.SecurityPreference)
	return nil, nil
}
