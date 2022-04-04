package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

type versionInfo struct {
	sas.DescribeVersionConfigResponse
	Region string
}

//// TABLE DEFINITION

func tableAlicloudSecurityCenterVersion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_security_center_version",
		Description: "Alicloud Security Center Version",
		List: &plugin.ListConfig{
			Hydrate: listSecurityCenterVersions,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the purchased Security Center instance.",
			},
			{
				Name:        "version",
				Type:        proto.ColumnType_STRING,
				Description: "The purchased edition of Security Center.",
			},
			{
				Name:        "is_trial_version",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether Security Center is the free trial edition.",
			},
			{
				Name:        "app_white_list",
				Type:        proto.ColumnType_INT,
				Description: "Indicates whether the application whitelist is enabled.",
			},
			{
				Name:        "app_white_list_auth_count",
				Type:        proto.ColumnType_INT,
				Description: "The quota on the servers to which you can apply your application whitelist.",
			},
			{
				Name:        "asset_level",
				Type:        proto.ColumnType_INT,
				Description: "The purchased quota for Security Center.",
			},
			{
				Name:        "is_over_balance",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the number of existing servers exceeds your quota.",
			},
			{
				Name:        "last_trail_end_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the last free trial ends.",
				Transform:   transform.FromField("LastTrailEndTime").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "release_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the Security Center instance expired.",
				Transform:   transform.FromField("ReleaseTime").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "sas_log",
				Type:        proto.ColumnType_INT,
				Description: "Indicates whether log analysis is purchased.",
			},
			{
				Name:        "sas_screen",
				Type:        proto.ColumnType_INT,
				Description: "Indicates whether the security dashboard is purchased.",
			},
			{
				Name:        "sls_capacity",
				Type:        proto.ColumnType_INT,
				Description: "The purchased capacity of log storage.",
			},
			{
				Name:        "user_defined_alarms",
				Type:        proto.ColumnType_INT,
				Description: "Indicates whether the custom alert feature is enabled.",
			},
			{
				Name:        "web_lock",
				Type:        proto.ColumnType_INT,
				Description: "Indicates whether web tamper proofing is enabled.",
			},
			{
				Name:        "web_lock_auth_count",
				Type:        proto.ColumnType_INT,
				Description: "The quota on the servers that web tamper proofing protects.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Version"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: ColumnDescriptionAkas,
				Hydrate:     getSecurityCenterVersionAkas,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
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

func listSecurityCenterVersions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// supported regions for security center are International(cn-hangzhou), Malaysia(ap-southeast-3) and Singapore(ap-southeast-1)
	supportedRegions := []string{"cn-hangzhou", "ap-southeast-1", "ap-southeast-3"}
	if !helpers.StringSliceContains(supportedRegions, region) {
		return nil, nil
	}
	// Create service connection
	client, err := SecurityCenterService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listSecurityCenterVersions", "connection_error", err)
		return nil, err
	}

	request := sas.CreateDescribeVersionConfigRequest()
	request.Scheme = "https"

	response, err := client.DescribeVersionConfig(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listSecurityCenterVersions", "query_error", err, "request", request)
		return nil, err
	}
	d.StreamListItem(ctx, versionInfo{*response, region})

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getSecurityCenterVersionAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSecurityCenterVersionAkas")

	data := h.Item.(versionInfo)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:security-center:" + data.Region + ":" + accountID + ":version/" + strconv.Itoa(data.Version)}

	return akas, nil
}
