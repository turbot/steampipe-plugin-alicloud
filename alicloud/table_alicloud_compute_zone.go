package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudComputeZone(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_compute_zone",
		Description: "",
		List: &plugin.ListConfig{
			ParentHydrate: listComputeRegions,
			Hydrate:       listComputeZones,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:      "zone_no",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("ZoneNo"),
			},
			{
				Name:        "zone_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ZoneId"),
				Description: "The zone ID.",
			},
			{
				Name:        "local_name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LocalName"),
				Description: "The name of the zone in the local language.",
			},

			// Other columns
			{
				Name:        "available_resource_creation",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableResourceCreation"),
				Description: "The types of the resources that can be created. The data type of this parameter is List.",
			},
			{
				Name:        "available_volume_categories",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableVolumeCategories"),
				Description: "The categories of available shared storage. The data type of this parameter is List.",
			},
			{
				Name:        "available_instance_types",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableInstanceTypes"),
				Description: "The instance types of instances that can be created. The data type of this parameter is List.",
			},
			{
				Name:        "available_dedicated_host_types",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableDedicatedHostTypes"),
				Description: "The supported types of dedicated hosts. The data type of this parameter is List.",
			},
			{
				Name:      "network_types",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("NetworkTypes"),
			},
			{
				Name:        "available_disk_categories",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableDiskCategories"),
				Description: "The supported disk categories. The data type of this parameter is List.",
			},
			{
				Name:        "dedicated_host_generations",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DedicatedHostGenerations"),
				Description: "The generation numbers of dedicated hosts. The data type of this parameter is List.",
			},
			{
				Name:        "available_resources",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableResources"),
				Description: "An array consisting of ResourcesInfo data.",
			},

			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getZoneAkas,
				Transform:   transform.FromValue(),
				Description: ColumnDescriptionAkas,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ZoneId"),
				Description: ColumnDescriptionTitle,
			},

			//	alicloud common columns
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

func listComputeZones(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs.listComputeZones", "connection_error", err)
		return nil, err
	}
	regionList := h.Item.(ecs.Region)
	request := ecs.CreateDescribeZonesRequest()
	request.Scheme = "https"
	request.RegionId = regionList.RegionId
	response, err := client.DescribeZones(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listComputeZones", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.Zones.Zone {
		plugin.Logger(ctx).Warn("alicloud_vpc.listComputeZones", "item", i)
		d.StreamLeafListItem(ctx, i)
	}
	return nil, nil
}

func getZoneAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getZoneAkas")
	data := h.Item.(ecs.Zone)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ecs::" + accountID + ":zone/" + data.ZoneId}, nil
}

//// TRANSFORM FUNCTIONS
