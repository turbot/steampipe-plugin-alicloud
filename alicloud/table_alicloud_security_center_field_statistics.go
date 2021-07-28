package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type FieldInfo struct {
	sas.GroupedFields
	Region string
}

//// TABLE DEFINITION

func tableAlicloudSecurityCenterFieldStatistics(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_security_center_field_statistics",
		Description: "Alicloud Security Center Field Statistics",
		List: &plugin.ListConfig{
			Hydrate: listSecurityCenterFieldStatistics,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "category_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of assets category.",
			},
			{
				Name:        "general_asset_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of general assets.",
			},
			{
				Name:        "group_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of asset groups.",
			},
			{
				Name:        "important_asset_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of important assets.",
			},
			{
				Name:        "instance_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The total number of assets of the specified type.",
			},
			{
				Name:        "new_instance_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of new servers.",
			},
			{
				Name:        "not_running_status_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of inactive servers.",
			},
			{
				Name:        "offline_instance_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of offline servers.",
			},
			{
				Name:        "region_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of regions to which the servers belong.",
			},
			{
				Name:        "risk_instance_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of assets that are at risk.",
			},
			{
				Name:        "test_asset_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of test assets.",
			},
			{
				Name:        "unprotected_instance_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of unprotected assets.",
			},
			{
				Name:        "vpc_count",
				Type:        proto.ColumnType_INT,
				Default:     0,
				Description: "The number of VPCs.",
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

func listSecurityCenterFieldStatistics(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// supported regions for security center are International(cn-hangzhou), Malaysia(ap-southeast-3) and Singapore(ap-southeast-1)
	supportedRegions := []string{"cn-hangzhou", "ap-southeast-1", "ap-southeast-3"}
	if !helpers.StringSliceContains(supportedRegions, region) {
		return nil, nil
	}
	// Create service connection
	client, err := SecurityCenterService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listSecurityCenterFieldStatistics", "connection_error", err)
		return nil, err
	}

	request := sas.CreateDescribeFieldStatisticsRequest()
	request.Scheme = "https"

	response, err := client.DescribeFieldStatistics(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listSecurityCenterFieldStatistics", "query_error", err, "request", request)
		return nil, err
	}
	d.StreamListItem(ctx, FieldInfo{response.GroupedFields, region})

	return nil, nil
}
