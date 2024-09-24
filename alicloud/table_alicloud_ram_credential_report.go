package alicloud

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/gocarina/gocsv"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type alicloudRamCredentialReportResult struct {
	GeneratedTime                   *string `csv:"-"`
	UserName                        *string `csv:"user"`
	UserCreationTime                *string `csv:"user_creation_time"`
	UserLastLogon                   *string `csv:"user_last_logon"`
	PasswordExist                   *string `csv:"password_exist"`
	PasswordActive                  *string `csv:"password_active"`
	PasswordLastChanged             *string `csv:"password_last_changed"`
	PasswordNextRotation            *string `csv:"password_next_rotation"`
	MfaActive                       *string `csv:"mfa_active"`
	AccessKey1Exist                 *string `csv:"access_key_1_exist"`
	AccessKey1Active                *string `csv:"access_key_1_active"`
	AccessKey1LastRotated           *string `csv:"access_key_1_last_rotated"`
	AccessKey1LastUsed              *string `csv:"access_key_1_last_used"`
	AccessKey2Exist                 *string `csv:"access_key_2_exist"`
	AccessKey2Active                *string `csv:"access_key_2_active"`
	AccessKey2LastRotated           *string `csv:"access_key_2_last_rotated"`
	AccessKey2LastUsed              *string `csv:"access_key_2_last_used"`
	AdditionalAccessKey1Exist       *string `csv:"additional_access_key_1_exist"`
	AdditionalAccessKey1Active      *string `csv:"additional_access_key_1_active"`
	AdditionalAccessKey1LastRotated *string `csv:"additional_access_key_1_last_rotated"`
	AdditionalAccessKey1LastUsed    *string `csv:"additional_access_key_1_last_used"`
	AdditionalAccessKey2Exist       *string `csv:"additional_access_key_2_exist"`
	AdditionalAccessKey2Active      *string `csv:"additional_access_key_2_active"`
	AdditionalAccessKey2LastRotated *string `csv:"additional_access_key_2_last_rotated"`
	AdditionalAccessKey2LastUsed    *string `csv:"additional_access_key_2_last_used"`
	AdditionalAccessKey3Exist       *string `csv:"additional_access_key_3_exist"`
	AdditionalAccessKey3Active      *string `csv:"additional_access_key_3_active"`
	AdditionalAccessKey3LastRotated *string `csv:"additional_access_key_3_last_rotated"`
	AdditionalAccessKey3LastUsed    *string `csv:"additional_access_key_3_last_used"`
}

//// TABLE DEFINITION

func tableAlicloudRAMCredentialReport(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_credential_report",
		Description: "Alicloud RAM Credential Report",
		List: &plugin.ListConfig{
			Hydrate: listRAMCredentialReports,
		},
		Columns: []*plugin.Column{
			{
				Name:        "user_name",
				Description: "The email of the RAM user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "password_exist",
				Description: "Indicates whether the user have any password for logging in, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfEqual("LOGIN_DISABLED").Transform(transform.ToBool),
			},
			{
				Name:        "password_active",
				Description: "Indicates whether the password is active, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("LOGIN_DISABLED").NullIfEqual("N/A").Transform(transform.ToBool),
			},
			{
				Name:        "mfa_active",
				Description: "Indicates whether multi-factor authentication (MFA) device has been enabled for the user.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfEqual("LOGIN_DISABLED").Transform(transform.ToBool),
			},
			{
				Name:        "user_creation_time",
				Description: "Specifies the time when the user is created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "user_last_logon",
				Description: "Specifies the time when the user last logged in to the console.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfEqual("LOGIN_DISABLED").NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "password_last_changed",
				Description: "Specifies the time when the password has been updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfEqual("LOGIN_DISABLED").NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "password_next_rotation",
				Description: "Specifies the time when the password will be rotated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfEqual("LOGIN_DISABLED").NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "access_key_1_exist",
				Description: "Indicates whether the user have access key, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "access_key_1_active",
				Description: "Indicates whether the user access key is active, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "access_key_1_last_rotated",
				Description: "Specifies the time when the access key has been rotated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "access_key_1_last_used",
				Description: "Specifies the time when the access key was most recently used to sign an Alicloud API request.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "access_key_2_exist",
				Description: "Indicates whether the user have access key, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "access_key_2_active",
				Description: "Indicates whether the user access key is active, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "access_key_2_last_rotated",
				Description: "Specifies the time when the access key has been rotated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "access_key_2_last_used",
				Description: "Specifies the time when the access key was most recently used to sign an Alicloud API request.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "additional_access_key_1_exist",
				Description: "Indicates whether the user have access key, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "additional_access_key_1_active",
				Description: "Indicates whether the user access key is active, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "additional_access_key_1_last_rotated",
				Description: "Specifies the time when the access key has been rotated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "additional_access_key_1_last_used",
				Description: "Specifies the time when the access key was most recently used to sign an Alicloud API request.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "additional_access_key_2_exist",
				Description: "Indicates whether the user have access key, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "additional_access_key_2_active",
				Description: "Indicates whether the user access key is active, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "additional_access_key_2_last_rotated",
				Description: "Specifies the time when the access key has been rotated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "additional_access_key_2_last_used",
				Description: "Specifies the time when the access key was most recently used to sign an Alicloud API request.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "additional_access_key_3_exist",
				Description: "Indicates whether the user have access key, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "additional_access_key_3_active",
				Description: "Indicates whether the user access key is active, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-").Transform(transform.ToBool),
			},
			{
				Name:        "additional_access_key_3_last_rotated",
				Description: "Specifies the time when the access key has been rotated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "additional_access_key_3_last_used",
				Description: "Specifies the time when the access key was most recently used to sign an Alicloud API request.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero().NullIfEqual("N/A").NullIfEqual("-"),
			},
			{
				Name:        "generated_time",
				Description: "Specifies the time when the credential report has been generated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},

			// alicloud standard columns
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

func listRAMCredentialReports(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_credential_report.listRAMCredentialReports", "connection_error", err)
		return nil, err
	}

	// Create a new request to retrieve the credential report.
	// To maintain consistency within the authentication module, we chose not to use
	// the "github.com/alibabacloud-go/ims-20190815/client" and "github.com/alibabacloud-go/tea-rpc/client" packages.
	// Instead, we opted for the "github.com/aliyun/alibaba-cloud-sdk-go/sdk" package to handle authentication and make the API call.

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "ims.aliyuncs.com"
	request.Version = "2019-08-15"
	request.ApiName = "GetCredentialReport"

	// Make the API call
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_credential_report.listRAMCredentialReports", "Api_error", err)
		return nil, err
	}

	// Unmarshal the response to the struct
	var credentialReportResponse GetCredentialReportResponse
	err = json.Unmarshal(response.GetHttpContentBytes(), &credentialReportResponse)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_credential_report.listRAMCredentialReports", "unmarshal_error", err)
		return nil, err
	}

	// The report is Base64-encoded. After decoding the report, the credential report is in the CSV format.
	data, err := base64.StdEncoding.DecodeString(*credentialReportResponse.Content)
	if err != nil {
		return nil, err
	}
	content := string(data[:])

	rows := []*alicloudRamCredentialReportResult{}
	if err := gocsv.UnmarshalString(content, &rows); err != nil {
		return nil, err
	}

	for _, row := range rows {
		row.GeneratedTime = credentialReportResponse.GeneratedTime
		d.StreamListItem(ctx, row)
	}

	return nil, nil
}

type GetCredentialReportResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
	GeneratedTime *string `json:"GeneratedTime,omitempty" xml:"GeneratedTime,omitempty" require:"true"`
}
