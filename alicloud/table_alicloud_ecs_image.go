package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type imageInfo = struct {
	Image  ecs.Image
	Region string
}

//// TABLE DEFINITION

func tableAlicloudEcsImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_image",
		Description: "AliCloud ECS Image.",
		List: &plugin.ListConfig{
			Hydrate: listEcsImages,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getEcsImage,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.ImageName"),
			},
			{
				Name:        "id",
				Description: "The ID of the image that the instance is running.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.ImageId"),
			},
			{
				Name:        "size",
				Description: "The size of the image (in GiB).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Image.Size"),
			},
			{
				Name:        "status",
				Description: "The status of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.Status"),
			},
			{
				Name:        "image_owner_alias",
				Type:        proto.ColumnType_STRING,
				Description: "The alias of the image owner. Possible values are: system, self, others, marketplace.",
				Transform:   transform.FromField("Image.ImageOwnerAlias"),
			},
			{
				Name:        "architecture",
				Description: "The image architecture. Possible values are: 'i386', and 'x86_64'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.Architecture"),
			},
			{
				Name:        "creation_time",
				Description: "The time when the image was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Image.CreationTime"),
			},
			{
				Name:        "description",
				Description: "A user-defined, human readable description for the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.Description"),
			},
			{
				Name:        "image_family",
				Description: "The name of the image family.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.ImageFamily"),
			},
			{
				Name:        "image_version",
				Description: "The version of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.ImageVersion"),
			},
			{
				Name:        "is_copied",
				Description: "Indicates whether the image is a copy of another image.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Image.IsCopied"),
			},
			{
				Name:        "is_self_shared",
				Description: "Indicates whether the image has been shared to other Alibaba Cloud accounts.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Image.IsSelfShared").NullIfZero().Transform(transform.ToBool),
			},
			{
				Name:        "is_subscribed",
				Description: "Indicates whether you have subscribed to the image that corresponds to the specified product code.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Image.IsSubscribed"),
			},
			{
				Name:        "is_support_cloud_init",
				Description: "Indicates whether the image supports cloud-init.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Image.IsSupportCloudinit"),
			},
			{
				Name:        "is_support_io_optimized",
				Description: "Indicates whether the image can be used on I/O optimized instances.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Image.IsSupportIoOptimized"),
			},
			{
				Name:        "os_name",
				Description: "The Chinese name of the operating system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.OSName"),
			},
			{
				Name:        "os_name_en",
				Description: "The English name of the operating system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.OSNameEn"),
			},
			{
				Name:        "os_type",
				Description: "The type of the operating system. Possible values are: windows,and linux",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.OSType"),
			},
			{
				Name:        "platform",
				Description: "The platform of the operating system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.Platform"),
			},
			{
				Name:        "product_code",
				Description: "The product code of the Alibaba Cloud Marketplace image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.ProductCode"),
			},
			{
				Name:        "progress",
				Description: "The image creation progress, in percent(%).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.Progress"),
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the image belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.ResourceGroupId"),
			},
			{
				Name:        "usage",
				Description: "Indicates whether the image has been used to create ECS instances.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Image.Usage"),
			},
			{
				Name:        "disk_device_mappings",
				Description: "The mappings between disks and snapshots under the image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Image.DiskDeviceMappings.DiskDeviceMapping"),
			},
			{
				Name:        "share_permissions",
				Description: "A list of groups and accounts that the image can be shared.",
				Hydrate:     getEcsImageSharePermission,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Image.Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Image.Tags.Tag").Transform(ecsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsImageAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Default:     transform.FromField("Image.ImageId"),
				Transform:   transform.FromField("Image.ImageName"),
			},

			// alicloud standard columns
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

func listEcsImages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_image.listEcsImages", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeImagesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeImages(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_image.listEcsImages", "query_error", err, "request", request)
			return nil, err
		}
		for _, image := range response.Images.Image {
			plugin.Logger(ctx).Warn("listEcsDisk", "item", image)
			d.StreamListItem(ctx, imageInfo{image, response.RegionId})
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

func getEcsImage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsImage")

	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_image.getEcsImage", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(imageInfo)
		id = data.Image.ImageId
	} else {
		id = d.KeyColumnQuals["id"].GetStringValue()
	}

	request := ecs.CreateDescribeImagesRequest()
	request.Scheme = "https"
	request.ImageId = id

	response, err := client.DescribeImages(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_image.getEcsImage", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.Images.Image != nil && len(response.Images.Image) > 0 {
		return imageInfo{response.Images.Image[0], response.RegionId}, nil
	}

	return nil, nil
}

func getEcsImageSharePermission(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsImageSharePermission")

	data := h.Item.(imageInfo)

	// This operation can only be performed on a custom image
	if data.Image.ImageOwnerAlias != "self" {
		return nil, nil
	}

	id := data.Image.ImageId

	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_image.getEcsImage", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeImageSharePermissionRequest()
	request.Scheme = "https"
	request.ImageId = id

	var groups []ecs.ShareGroup
	var accounts []ecs.Account

	count := 0
	for {
		response, err := client.DescribeImageSharePermission(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_image.getEcsImageSharePermission", "query_error", err, "request", request)
			return nil, err
		}
		for _, group := range response.ShareGroups.ShareGroup {
			plugin.Logger(ctx).Warn("listEcsDisk", "item", group)
			groups = append(groups, group)
			count++
		}
		for _, account := range response.Accounts.Account {
			plugin.Logger(ctx).Warn("listEcsDisk", "item", account)
			accounts = append(accounts, account)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}

	result := map[string]interface{}{
		"ShareGroups": groups,
		"Accounts":    accounts,
	}

	return result, nil
}

func getEcsImageAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsImageAka")

	data := h.Item.(imageInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + data.Region + ":" + accountID + ":image/" + data.Image.ImageId}

	return akas, nil
}
