package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudKmsKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_key",
		Description: "Alicloud KMS Key",
		List: &plugin.ListConfig{
			Hydrate: listKmsKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("key_id"),
			Hydrate:           getKmsKey,
			ShouldIgnoreError: isNotFoundError([]string{"EntityNotExist.Key", "Forbidden.KeyNotFound"}),
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "key_id",
				Description: "The globally unique ID of the CMK.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the CMK.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key_state",
				Description: "The status of the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "creator",
				Description: "The creator of the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "creation_date",
				Description: "The date and time the CMK was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "description",
				Description: "The description of the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "key_usage",
				Description: "The purpose of the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "key_spec",
				Description: "The type of the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "last_rotation_date",
				Description: "The date and time the last rotation was performed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "automatic_rotation",
				Description: "Indicates whether automatic key rotation is enabled.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "delete_date",
				Description: "The date and time the CMK is scheduled for deletion.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "material_expire_time",
				Description: "The time and date the key material for the CMK expires.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "origin",
				Description: "The source of the key material for the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "protection_level",
				Description: "The protection level of the CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "primary_key_version",
				Description: "The ID of the current primary key version of the symmetric CMK.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
			},
			{
				Name:        "key_aliases",
				Description: "A list of aliases bound to a CMK.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKeyAlias,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the key.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKeyTags,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyKmsSourceTags),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKeyTags,
				Transform:   transform.FromField("Tags.Tag").Transform(kmsTurbotTags),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyId"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Arn").Transform(ensureStringArray),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKeyRegion,
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

func listKmsKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.listKmsKey", "connection_error", err)
		return nil, err
	}

	request := kms.CreateListKeysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.ListKeys(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_kms_key.listKmsKey", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Keys.Key {
			plugin.Logger(ctx).Warn("listKmsKey", "item", i)
			d.StreamListItem(ctx,
				kms.KeyMetadata{
					Arn:   i.KeyArn,
					KeyId: i.KeyId,
				},
			)
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

func getKmsKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKmsKey")

	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(kms.KeyMetadata)
		id = data.KeyId
	} else {
		id = d.KeyColumnQuals["key_id"].GetStringValue()
	}

	request := kms.CreateDescribeKeyRequest()
	request.Scheme = "https"
	request.KeyId = id

	response, err := client.DescribeKey(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "query_error", err, "request", request)
		return nil, err
	}

	return response.KeyMetadata, nil
}

func getKeyAlias(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKeyAlias")

	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyAlias", "connection_error", err)
		return nil, err
	}

	data := h.Item.(kms.KeyMetadata)

	request := kms.CreateListAliasesByKeyIdRequest()
	request.Scheme = "https"
	request.KeyId = data.KeyId

	response, err := client.ListAliasesByKeyId(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyAlias", "query_error", err, "request", request)
		return nil, err
	}

	if response.Aliases.Alias != nil {
		return response.Aliases.Alias, nil
	}

	return nil, nil
}

func getKeyTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKeyTags")

	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyTags", "connection_error", err)
		return nil, err
	}

	data := h.Item.(kms.KeyMetadata)

	request := kms.CreateListResourceTagsRequest()
	request.Scheme = "https"
	request.KeyId = data.KeyId

	response, err := client.ListResourceTags(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyTags", "query_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func getKmsKeyRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKmsKeyRegion")
	region := d.KeyColumnQualString(matrixKeyRegion)

	return region, nil
}
