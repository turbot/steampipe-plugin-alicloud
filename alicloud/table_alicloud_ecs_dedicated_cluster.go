package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func tableAlicloudEcsDedicatedCluster(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_dedicated_cluster",
		Description: "Alicloud Elastic Dedicated Cluster",
		List: &plugin.ListConfig{
			ParentHydrate: listComputeRegions,
			Hydrate:       listEcsDedicatedCluster,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("region"),
			Hydrate:    getEcsDedicatedCluster,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the dedicated host cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostClusterName"),
			},
			{
				Name:        "dedicatedHost_cluster_id",
				Description: "The ID of the dedicated host cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region_id",
				Description: "The region ID of the dedicated host cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "zone_id",
				Description: "The zone ID of the dedicated host cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the dedicated host cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the dedicated host cluster belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "dedicated_host_ids",
				Description: "The IDs of dedicated hosts in the dedicated host cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DedicatedHostIds"),
			},
			{
				Name:        "dedicated_host_cluster_capacity",
				Description: "The capacity of the dedicated host cluster.",
				Type:        proto.ColumnType_JSON,
			},

			// steampipe standard columns
			{
				Name:        "tags_src",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyEcsSourceTags),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(ecsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsClusterAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DedicatedHostClusterName"),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
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

func listEcsDedicatedCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_cluster.listEcsDedicatedCluster", "connection_error", err)
		return nil, err
	}
	regionList := h.Item.(ecs.Region)
	request := ecs.CreateDescribeDedicatedHostClustersRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	request.RegionId = regionList.RegionId
	count := 0
	for {
		response, err := client.DescribeDedicatedHostClusters(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_cluster.listEcsDedicatedCluster", "query_error", err, "request", request)
			return nil, err
		}
		for _, cluster := range response.DedicatedHostClusters.DedicatedHostCluster {
			d.StreamLeafListItem(ctx, cluster)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getEcsDedicatedCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsDedicatedCluster")

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_cluster.getEcsDedicatedCluster", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		cluster := h.Item.(ecs.DedicatedHostCluster)
		id = cluster.RegionId
	} else {
		id = d.KeyColumnQuals["region"].GetStringValue()
	}

	request := ecs.CreateDescribeDedicatedHostClustersRequest()
	request.Scheme = "https"
	request.RegionId = id
	response, err := client.DescribeDedicatedHostClusters(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_instance.getEcsDedicatedCluster", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.DedicatedHostClusters.DedicatedHostCluster != nil && len(response.DedicatedHostClusters.DedicatedHostCluster) > 0 {
		return response.DedicatedHostClusters.DedicatedHostCluster[0], nil
	}

	return nil, nil
}

func getEcsClusterAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsClusterAka")
	cluster := h.Item.(ecs.DedicatedHostCluster)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + cluster.RegionId + ":" + accountID + ":cluster/" + cluster.DedicatedHostClusterId}

	return akas, nil
}
