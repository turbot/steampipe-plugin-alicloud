package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudEcsImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_image",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listEcsImage,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the image.",
				Transform: transform.FromField("ImageName")},

			{Name: "image_owner_alias", Type: proto.ColumnType_STRING, Description: "The source of the image.."},
			{Name: "image_id", Type: proto.ColumnType_STRING, Description: "The ID of the image that the instance is running."},
			{Name: "os_name", Type: proto.ColumnType_STRING, Description: "The Chinese name of the operating system."},
			{Name: "os_name_en", Type: proto.ColumnType_STRING, Description: "The English name of the operating system."},
			{Name: "image_family", Type: proto.ColumnType_STRING, Description: "The name of the image family. You can set this parameter to query images of the specified image family."},
			{Name: "architecture", Type: proto.ColumnType_STRING, Description: "The image architecture. Valid values:i386,x86_64"},

			{Name: "size", Type: proto.ColumnType_INT, Description: "The size of the disk."},
			{Name: "is_support_io_optimized", Type: proto.ColumnType_BOOL, Description: "Indicates whether the image can be used on I/O optimized instances."},

			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the image belongs."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the image."},

			{Name: "usage", Type: proto.ColumnType_STRING, Description: "Indicates whether the image has been used to create ECS instances."},

			{Name: "is_copied", Type: proto.ColumnType_BOOL, Description: "Indicates whether the image is a copy of another image."},
			{Name: "image_version", Type: proto.ColumnType_STRING, Description: "The version of the image."},
			{Name: "os_type", Type: proto.ColumnType_STRING, Description: "The type of the operating system. Valid values:windows,linux"},
			{Name: "is_subscribed", Type: proto.ColumnType_BOOL, Description: "Indicates whether you have subscribed to the image that corresponds to the specified product code."},
			{Name: "is_support_cloudinit", Type: proto.ColumnType_BOOL, Description: "Indicates whether the image supports cloud-init."},
			{Name: "creation_time", Type: proto.ColumnType_STRING, Description: "The time when the image was created."},
			{Name: "product_code", Type: proto.ColumnType_STRING, Description: "The product code of the Alibaba Cloud Marketplace image."},
			{Name: "progress", Type: proto.ColumnType_STRING, Description: "The image creation progress. Unit: percent (%)."},
			{Name: "platform", Type: proto.ColumnType_STRING, Description: "The platform of the operating system."},
			{Name: "is_self_shared", Type: proto.ColumnType_STRING, Description: "Indicates whether the image has been shared to other Alibaba Cloud accounts."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the image."},
			{Name: "disk_device_mappings", Type: proto.ColumnType_JSON, Description: "The mappings between disks and snapshots under the image."},

			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("ImageName"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listEcsImage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs.listEcsImage", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeImagesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	// if quals["is_default"] != nil {
	// 	request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	// }
	if quals["id"] != nil {
		request.ImageId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeImages(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs.listEcsImages", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Images.Image {
			plugin.Logger(ctx).Warn("alicloud_ecs.listEcsImage", "tags", i.Tags, "item", i)
			d.StreamListItem(ctx, i)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}
