package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

//// TABLE DEFINITION

func tableAlicloudEcsImage(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_image",
		Description: "AliCloud ECS Image.",
		List: &plugin.ListConfig{
			Hydrate: listEcsImages,
		},
		Get: &plugin.GetConfig{
			// We must include both image_id and region in the where clause else we will receive numerous rows. Which causes Error: get call returned 2 results - the key column is not globally unique (SQLSTATE HV000)
			KeyColumns: plugin.AllColumns([]string{"image_id", "region"}),
			Hydrate:    getEcsImage,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ImageName"),
			},
			{
				Name:        "image_id",
				Description: "The ID of the image that the instance is running.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the ECS image.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsImageARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "size",
				Description: "The size of the image (in GiB).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "status",
				Description: "The status of the image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "image_owner_alias",
				Type:        proto.ColumnType_STRING,
				Description: "The alias of the image owner. Possible values are: system, self, others, marketplace.",
			},
			{
				Name:        "architecture",
				Description: "The image architecture. Possible values are: 'i386', and 'x86_64'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the image was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "A user-defined, human readable description for the image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "image_family",
				Description: "The name of the image family.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "image_version",
				Description: "The version of the image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_copied",
				Description: "Indicates whether the image is a copy of another image.",
				Type:        proto.ColumnType_BOOL,
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
			},
			{
				Name:        "is_support_cloud_init",
				Description: "Indicates whether the image supports cloud-init.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("IsSupportCloudinit"),
			},
			{
				Name:        "is_support_io_optimized",
				Description: "Indicates whether the image can be used on I/O optimized instances.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("IsSupportIoOptimized"),
			},
			{
				Name:        "os_name",
				Description: "The Chinese name of the operating system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OSName"),
			},
			{
				Name:        "os_name_en",
				Description: "The English name of the operating system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OSNameEn"),
			},
			{
				Name:        "os_type",
				Description: "The type of the operating system. Possible values are: windows,and linux",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OSType"),
			},
			{
				Name:        "platform",
				Description: "The platform of the operating system.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "product_code",
				Description: "The product code of the Alibaba Cloud Marketplace image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "progress",
				Description: "The image creation progress, in percent(%).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the image belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "usage",
				Description: "Indicates whether the image has been used to create ECS instances.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "disk_device_mappings",
				Description: "The mappings between disks and snapshots under the image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DiskDeviceMappings.DiskDeviceMapping"),
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
				Transform:   transform.FromField("Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// Steampipe standard columns
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
				Hydrate:     getEcsImageARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Default:     transform.FromField("ImageId"),
				Transform:   transform.FromField("ImageName"),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsImageRegion,
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

func listEcsImages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_image.listEcsImages", "connection_error", err)
		return nil, err
	}

	// regionName := d.KeyColumnQuals["region"].GetStringValue()

	request := ecs.CreateDescribeImagesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	// request.RegionId = regionName

	count := 0
	for {
		response, err := client.DescribeImages(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_image.listEcsImages", "query_error", err, "request", request)
			return nil, err
		}
		for _, image := range response.Images.Image {
			plugin.Logger(ctx).Warn("listEcsDisk", "item", image)
			d.StreamListItem(ctx, image)
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
	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_image.getEcsImage", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["image_id"].GetStringValue()
	regionName := d.KeyColumnQuals["region"].GetStringValue()

	// Handle empty name or region
	if id == "" || regionName == "" {
		return nil, nil
	}

	request := ecs.CreateDescribeImagesRequest()
	request.Scheme = "https"
	request.ImageId = id
	request.RegionId = regionName

	response, err := client.DescribeImages(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_image.getEcsImage", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.Images.Image != nil && len(response.Images.Image) > 0 {
		return response.Images.Image[0], nil
	}

	return nil, nil
}

func getEcsImageSharePermission(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsImageSharePermission")

	data := h.Item.(ecs.Image)

	// This operation can only be performed on a custom image
	if data.ImageOwnerAlias != "self" {
		return nil, nil
	}

	id := data.ImageId

	// Create service connection
	client, err := ECSService(ctx, d)
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

func getEcsImageARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsImageARN")

	data := h.Item.(ecs.Image)
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + region + ":" + accountID + ":image/" + data.ImageId

	return arn, nil
}

func getEcsImageRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	return region, nil
}
