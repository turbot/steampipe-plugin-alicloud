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
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listComputeRegions,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("region_id"),
			Hydrate:    getComputeRegion,
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
				Transform:   transform.FromField("LocalName"),
				Description: "",
			},
			{
				Name:        "region_endpoint",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionEndpoint"),
				Description: "",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Status"),
				Description: "",
			},

			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRegionAkas,
				Transform:   transform.FromValue(),
				Description: ColumnDescriptionAkas,
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
		plugin.Logger(ctx).Error("alicloud_vpc.listComputeRegions", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.Regions.Region {
		plugin.Logger(ctx).Warn("alicloud_vpc.listComputeRegions", "item", i)
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getComputeRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getComputeRegion", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		ecs := h.Item.(ecs.Region)
		id = ecs.RegionId
	} else {
		id = d.KeyColumnQuals["region_id"].GetStringValue()
	}
	request := ecs.CreateDescribeRegionsRequest()
	request.Scheme = "https"
	request.RegionId = id
	response, err := client.DescribeRegions(request)
	if err != nil {
		plugin.Logger(ctx).Error("getComputeRegion", "query_error", err, "request", request)
		return nil, err
	}

	if response.Regions.Region != nil && len(response.Regions.Region) > 0 {
		return response.Regions.Region, nil
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

//// TRANSFORM FUNCTIONS
