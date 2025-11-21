package alicloud

import (
	"context"
	"slices"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudSecurityCenterAsset(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_security_center_asset",
		Description: "Alicloud Security Center Assets (ECS instances with Security Center agent status)",
		List: &plugin.ListConfig{
			Hydrate: listSecurityCenterAssets,
			Tags:    map[string]string{"service": "sas", "action": "DescribeCloudCenterInstances"},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "instance_id", Require: plugin.Optional},
				{Name: "client_status", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the ECS instance.",
			},
			{
				Name:        "instance_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the ECS instance.",
			},
			{
				Name:        "uuid",
				Type:        proto.ColumnType_STRING,
				Description: "The UUID of the asset (used by Security Center).",
			},
			{
				Name:        "client_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the Security Center agent. Values: online, offline, uninstall (not installed).",
			},
			{
				Name:        "client_version",
				Type:        proto.ColumnType_STRING,
				Description: "The version of the Security Center agent installed.",
			},
			{
				Name:        "os_name",
				Type:        proto.ColumnType_STRING,
				Description: "The operating system name.",
			},
			{
				Name:        "os",
				Type:        proto.ColumnType_STRING,
				Description: "The operating system type.",
			},
			{
				Name:        "ip",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address of the instance.",
			},
			{
				Name:        "internet_ip",
				Type:        proto.ColumnType_STRING,
				Description: "The internet IP address of the instance.",
			},
			{
				Name:        "intranet_ip",
				Type:        proto.ColumnType_STRING,
				Description: "The intranet IP address of the instance.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the instance in Security Center.",
			},
			{
				Name:        "asset_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of asset (e.g., EcsInstance).",
			},
			{
				Name:        "vpc_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The VPC instance ID.",
			},
			{
				Name:        "vul_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of vulnerabilities detected on this asset.",
			},
			{
				Name:        "risk_count",
				Type:        proto.ColumnType_STRING,
				Description: "The number of security risks detected.",
			},
			{
				Name:        "safe_event_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of security events detected.",
			},
			{
				Name:        "health_check_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of health checks performed.",
			},
			{
				Name:        "hc_status",
				Type:        proto.ColumnType_STRING,
				Description: "The health check status.",
			},
			{
				Name:        "vul_status",
				Type:        proto.ColumnType_STRING,
				Description: "The vulnerability status.",
			},
			{
				Name:        "alarm_status",
				Type:        proto.ColumnType_STRING,
				Description: "The alarm status.",
			},
			{
				Name:        "risk_status",
				Type:        proto.ColumnType_STRING,
				Description: "The risk status.",
			},
			{
				Name:        "group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The group ID of the asset.",
			},
			{
				Name:        "importance",
				Type:        proto.ColumnType_INT,
				Description: "The importance level of the asset.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("InstanceName"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: ColumnDescriptionAkas,
				Hydrate:     getSecurityCenterAssetAkas,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getSecurityCenterAssetRegion,
				Transform:   transform.FromValue(),
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

func listSecurityCenterAssets(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	// supported regions for security center are International(cn-hangzhou), Malaysia(ap-southeast-3) and Singapore(ap-southeast-1)
	supportedRegions := []string{"cn-hangzhou", "ap-southeast-1", "ap-southeast-3"}
	if !slices.Contains(supportedRegions, region) {
		return nil, nil
	}

	// Create service connection
	client, err := SecurityCenterService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listSecurityCenterAssets", "connection_error", err)
		return nil, err
	}

	request := sas.CreateDescribeCloudCenterInstancesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.CurrentPage = requests.NewInteger(1)

	// Note: DescribeCloudCenterInstances may not support instance_id filter directly
	// We'll filter in the application layer if needed

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeCloudCenterInstances(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_listSecurityCenterAssets", "query_error", err, "request", request)
			return nil, err
		}

		for _, instance := range response.Instances {
			// Apply filters if provided
			if d.EqualsQualString("instance_id") != "" && instance.InstanceId != d.EqualsQualString("instance_id") {
				continue
			}
			if d.EqualsQualString("client_status") != "" && instance.ClientStatus != d.EqualsQualString("client_status") {
				continue
			}

			instanceCopy := instance
			d.StreamListItem(ctx, &instanceCopy)
			count++
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		pageSize := 50
		if len(response.Instances) < pageSize || count >= response.PageInfo.TotalCount {
			break
		}

		// Get current page number from response and increment
		currentPage := response.PageInfo.CurrentPage + 1
		request.CurrentPage = requests.NewInteger(currentPage)
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func getSecurityCenterAssetAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSecurityCenterAssetAkas")

	data := h.Item.(*sas.Instance)
	region := d.EqualsQualString(matrixKeyRegion)

	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:security-center:" + region + ":" + accountID + ":asset/" + data.Uuid}
	return akas, nil
}

func getSecurityCenterAssetRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	return region, nil
}
