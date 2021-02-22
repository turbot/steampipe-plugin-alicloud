package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/resourcemanager"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudResource(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_resource",
		Description: "List All Resources.",
		List: &plugin.ListConfig{
			Hydrate: listResources,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "resource_id",
				Description: "The ID of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service",
				Description: "The name of the service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_date",
				Description: "The time when the resource is created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "resource_group_id",
				Description: "The name of the resource group.",
				Type:        proto.ColumnType_STRING,
			},

			// alibaba standard columns
			{
				Name:        "region",
				Description: "The name of the region where the resource resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
			},
			// {
			// 	Name:        "account_id",
			// 	Description: "The alicloud Account ID in which the resource is located.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Hydrate:     getCommonColumns,
			// 	Transform:   transform.FromField("AccountID"),
			// },
		},
	}
}

//// LIST FUNCTION

func listResources(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := connectResourceManager(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_resource.listResources", "connection_error", err)
		return nil, err
	}
	request := resourcemanager.CreateListResourcesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.ListResources(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_resource.listResources", "query_error", err, "request", request)
			return nil, err
		}
		for _, resource := range response.Resources.Resource {
			plugin.Logger(ctx).Warn("listEcsDisk", "item", resource)
			d.StreamListItem(ctx, resource)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}
