package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type kmsKeyInfo = struct {
	kms.KeyMetadata
	Region string
}

func tableAlicloudKmsKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_key",
		Description: "Key Management Service (KMS) provides secure and compliant key management and cryptography services to help you encrypt and protect sensitive data assets.",
		List: &plugin.ListConfig{
			Hydrate: listKmsKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("key_id"),
			Hydrate:    getKmsKey,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "key_id",
				Type:        proto.ColumnType_STRING,
				Description: "The globally unique ID of the CMK.",
			},
			{
				Name:        "key_state",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The status of the CMK.",
			},
			// Other columns
			{
				Name:        "key_arn",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Arn"),
				Description: "The Alibaba Cloud Resource Name (ARN) of the CMK.",
			},
			{
				Name:        "creation_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
				Description: "The date and time the CMK was created.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The description of the CMK.",
			},
			{
				Name:        "creator",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The creator of the CMK.",
			},
			{
				Name:        "key_usage",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The purpose of the CMK.",
			},
			{
				Name:        "key_spec",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The type of the CMK.",
			},
			{
				Name:        "last_rotation_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
				Description: "The date and time the last rotation was performed.",
			},
			{
				Name:        "automatic_rotation",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "Indicates whether automatic key rotation is enabled.",
			},
			{
				Name:        "delete_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
				Description: "The date and time the CMK is scheduled for deletion.",
			},
			{
				Name:        "origin",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The source of the key material for the CMK.",
			},
			{
				Name:        "protection_level",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The protection level of the CMK.",
			},
			{
				Name:        "primary_key_version",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsKey,
				Description: "The ID of the current primary key version of the symmetric CMK.",
			},
			{
				Name:        "key_aliases",
				Description: "A list of aliases bound to a CMK.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKeyAlias,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "material_expire_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsKey,
				Description: "The time and date the key material for the CMK expires.",
			},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKeyTags,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyKmsSourceTags),
				Description: "A list of tags assigned to bucket",
			},
			// steampipe standard columns
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKeyTags,
				Transform:   transform.FromField("Tags.Tag").Transform(kmsTagsToMap),
				Description: ColumnDescriptionTags,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyId"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Arn").Transform(ensureStringArray),
				Description: ColumnDescriptionAkas,
			},
			// alicloud standard columns
			{
				Name:        "region",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionRegion,
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
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := KMSService(ctx, d, region)
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
			d.StreamListItem(ctx, kmsKeyInfo{
				kms.KeyMetadata{
					Arn:   i.KeyArn,
					KeyId: i.KeyId,
				},
				region,
			})
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
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := KMSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(kmsKeyInfo)
		id = data.KeyId
	} else {
		id = d.KeyColumnQuals["key_id"].GetStringValue()
	}

	request := kms.CreateDescribeKeyRequest()
	request.Scheme = "https"
	request.KeyId = id

	response, err := client.DescribeKey(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "EntityNotExist.Key" {
			plugin.Logger(ctx).Warn("alicloud_kms_key.getKmsKey", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return kmsKeyInfo{response.KeyMetadata, region}, nil
}

func getKeyAlias(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := KMSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyAlias", "connection_error", err)
		return nil, err
	}

	data := h.Item.(kmsKeyInfo)

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
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := KMSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyTags", "connection_error", err)
		return nil, err
	}

	data := h.Item.(kmsKeyInfo)

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
