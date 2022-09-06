package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

type zoneInfo = struct {
	ecs.Zone
	Region string
}

//// TABLE DEFINITION

func tableAlicloudEcsZone(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_zone",
		Description: "Elastic Compute Zone",
		List: &plugin.ListConfig{
			ParentHydrate: listEcsRegions,
			Hydrate:       listEcsZones,
		},
		Columns: []*plugin.Column{
			{
				Name:        "zone_id",
				Type:        proto.ColumnType_STRING,
				Description: "The zone ID.",
			},
			{
				Name:        "local_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the zone in the local language.",
			},
			{
				Name:        "available_dedicated_host_types",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableDedicatedHostTypes.DedicatedHostType"),
				Description: "The supported types of dedicated hosts. The data type of this parameter is List.",
			},
			{
				Name:        "available_disk_categories",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableDiskCategories.DiskCategories"),
				Description: "The supported disk categories. The data type of this parameter is List.",
			},
			{
				Name:        "available_instance_types",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableInstanceTypes.InstanceTypes"),
				Description: "The instance types of instances that can be created. The data type of this parameter is List.",
			},
			{
				Name:        "available_resources",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableResources.ResourcesInfo"),
				Description: "An array consisting of ResourcesInfo data.",
			},
			{
				Name:        "available_resource_creation",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableResourceCreation.ResourceTypes"),
				Description: "The types of the resources that can be created. The data type of this parameter is List.",
			},
			{
				Name:        "available_volume_categories",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AvailableVolumeCategories.VolumeCategories"),
				Description: "The categories of available shared storage. The data type of this parameter is List.",
			},
			{
				Name:        "dedicated_host_generations",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DedicatedHostGenerations.DedicatedHostGeneration"),
				Description: "The generation numbers of dedicated hosts. The data type of this parameter is List.",
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ZoneId"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getZoneAkas,
				Transform:   transform.FromValue(),
				Description: ColumnDescriptionAkas,
			},

			// Alicloud common columns
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

func listEcsZones(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := h.Item.(ecs.Region).RegionId

	// Create service connection
	client, err := ECSRegionService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs.listEcsZones", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeZonesRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.AcceptLanguage = "en-US"

	response, err := client.DescribeZones(request)
	plugin.Logger(ctx).Trace("alicloud_ecs.listEcsZones", "network_test:", response)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs.listEcsZones", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.Zones.Zone {
		d.StreamListItem(ctx, zoneInfo{i, region})
	}
	return nil, nil
}

func getZoneAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getZoneAkas")
	data := h.Item.(zoneInfo)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ecs::" + accountID + ":zone/" + data.ZoneId}, nil
}
