package alicloud

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/sethvargo/go-retry"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudKmsKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_key",
		Description: "Alicloud KMS Key",
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("key_id"),
			Hydrate:           getKmsKey,
			ShouldIgnoreError: isNotFoundError([]string{"EntityNotExist.Key", "Forbidden.KeyNotFound"}),
		},
		List: &plugin.ListConfig{
			Hydrate: listKmsKey,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "key_state", Require: plugin.Optional},
				{Name: "key_usage", Require: plugin.Optional},
				{Name: "key_spec", Require: plugin.Optional},
				{Name: "protection_level", Require: plugin.Optional},
			},
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
	request.PageSize = requests.NewInteger(100)
	request.PageNumber = requests.NewInteger(1)

	// https://partners-intl.aliyun.com/help/doc-detail/28951.htm
	var queryFilters QueryFilters
	if value, ok := GetStringQualValueList(d.Quals, "key_state"); ok {
		if !(len(helpers.StringSliceDiff(value, []string{"Enabled", "Disabled", "PendingDeletion", "PendingImport"})) > 0) {
			queryFilters = append(queryFilters, QueryFilterItem{Key: "KeyState", Values: value})
		}
	}
	if value, ok := GetStringQualValueList(d.Quals, "key_usage"); ok {
		if !(len(helpers.StringSliceDiff(value, []string{"ENCRYPT/DECRYPT", "SIGN/VERIFY"})) > 0) {
			queryFilters = append(queryFilters, QueryFilterItem{Key: "KeyUsage", Values: value})
		}
	}
	if value, ok := GetStringQualValueList(d.Quals, "key_spec"); ok {
		if !(len(helpers.StringSliceDiff(value, []string{"Aliyun_AES_256", "Aliyun_SM4", "RSA_2048", "EC_P256", "EC_P256K", "EC_SM2"})) > 0) {
			queryFilters = append(queryFilters, QueryFilterItem{Key: "KeySpec", Values: value})
		}
	}
	if value, ok := GetStringQualValueList(d.Quals, "protection_level"); ok {
		if !(len(helpers.StringSliceDiff(value, []string{"SOFTWARE", "HSM"})) > 0) {
			queryFilters = append(queryFilters, QueryFilterItem{Key: "ProtectionLevel", Values: value})
		}
	}

	if len(queryFilters) > 0 {
		filter, err := queryFilters.String()
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_kms_key.listKmsKey", "filter_string_error", err)
			return nil, err
		}
		request.Filters = filter
	}

	plugin.Logger(ctx).Info("alicloud_kms_key.listKmsKey", "filter", request.Filters)

	// If the request no of items is less than the paging max limit
	// update limit to requested no of results.
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		pageSize, err := request.PageSize.GetValue64()
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_instance.listEcsInstance", "page_size_error", err)
			return nil, err
		}
		if *limit < pageSize {
			request.PageSize = requests.NewInteger(int(*limit))
		}
	}

	count := 0
	for {
		response, err := client.ListKeys(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_kms_key.listKmsKey", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Keys.Key {
			d.StreamListItem(ctx, kms.KeyMetadata{Arn: i.KeyArn, KeyId: i.KeyId})
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
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
	var response *kms.DescribeKeyResponse
	if h.Item != nil {
		data := h.Item.(kms.KeyMetadata)
		id = data.KeyId
	} else {
		id = d.KeyColumnQuals["key_id"].GetStringValue()
	}

	request := kms.CreateDescribeKeyRequest()
	request.Scheme = "https"
	request.KeyId = id

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeKey(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "query_error", err, "request", request)
				return err
			}
		}
		return nil
	})

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "query_retry_error", err, "request", request)
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
	var response *kms.ListAliasesByKeyIdResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.ListAliasesByKeyId(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				plugin.Logger(ctx).Error("alicloud_kms_key.getKeyAlias", "query_error", err, "request", request)
				return err
			}
		}
		return nil
	})

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyAlias", "query_retry_error", err, "request", request)
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
	var response *kms.ListResourceTagsResponse

	request := kms.CreateListResourceTagsRequest()
	request.Scheme = "https"
	request.KeyId = data.KeyId

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.ListResourceTags(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				plugin.Logger(ctx).Error("alicloud_kms_key.getKeyTags", "query_error", err, "request", request)
				return err
			}
		}
		return nil
	})

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKeyTags", "query_retry_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func getKmsKeyRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKmsKeyRegion")
	region := d.KeyColumnQualString(matrixKeyRegion)

	return region, nil
}
