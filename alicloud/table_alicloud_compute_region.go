package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudComputeRegion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_compute_region",
		Description: "Alicloud Compute Region",
		List: &plugin.ListConfig{
			Hydrate: listComputeRegions,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "region",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
				Description: ColumnDescriptionRegion,
			},

			{
				Name:        "local_name",
				Type:        proto.ColumnType_STRING,
				Description: "The local name of the region.",
			},
			{
				Name:        "region_endpoint",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionEndpoint"),
				Description: "The endpoint of the region.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the cluster is sold out.",
			},
			// steampipe standard columns
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRegionAkas,
				Transform:   transform.FromValue(),
				Description: ColumnDescriptionAkas,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
				Description: ColumnDescriptionTitle,
			},

			// alicloud common columns
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

func listComputeRegions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs.listComputeRegions", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"

	response, err := client.DescribeRegions(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs.listComputeRegions", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.Regions.Region {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getRegionAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRegionAkas")
	data := h.Item.(ecs.Region)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ecs::" + accountID + ":region/" + data.RegionId}, nil
}
