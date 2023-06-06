package alicloud

import (
	"context"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/sethvargo/go-retry"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudKmsSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_secret",
		Description: "Alicloud KMS Secret",
		List: &plugin.ListConfig{
			Hydrate: listKmsSecret,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getKmsSecret,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the secret.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecretName"),
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN).",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "secret_type",
				Description: "The type of the secret.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "automatic_rotation",
				Description: "Specifies whether automatic key rotation is enabled.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "create_time",
				Description: "The time when the KMS Secret was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description of the secret.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "encryption_key_id",
				Description: "The ID of the KMS customer master key (CMK) that is used to encrypt the secret value.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "last_rotation_date",
				Description: "Date of last rotation of Secret.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "next_rotation_date",
				Description: "The date of next rotation of Secret.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "planned_delete_time",
				Description: "The time when the KMS Secret is planned to delete.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "rotation_interval",
				Description: "The rotation perion of Secret.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "update_time",
				Description: "The time when the KMS Secret was modifies.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "extended_config",
				Description: "The extended configuration of Secret.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "version_ids",
				Description: "The list of secret versions.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listKmsSecretVersionIds,
				Transform:   transform.FromField("VersionId"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKmsSecret,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyKmsSourceTags),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKmsSecret,
				Transform:   transform.FromField("Tags.Tag").Transform(kmsTurbotTags),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getKmsSecret,
				Transform:   transform.FromField("Arn").Transform(ensureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecretName"),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKmsSecret,
				Transform:   transform.From(fetchRegionFromArn),
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

func listKmsSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.listKmsSecret", "connection_error", err)
		return nil, err
	}

	request := kms.CreateListSecretsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.ListSecrets(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_kms_secret.listKmsSecret", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.SecretList.Secret {
			d.StreamListItem(ctx, &kms.DescribeSecretResponse{
				CreateTime:        i.CreateTime,
				PlannedDeleteTime: i.PlannedDeleteTime,
				SecretName:        i.SecretName,
				UpdateTime:        i.UpdateTime,
				SecretType:        i.SecretType,
				Tags: kms.TagsInDescribeSecret{
					Tag: i.Tags.Tag,
				},
			})
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
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

func getKmsSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getKmsSecret")

	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "connection_error", err)
		return nil, err
	}

	var name string
	var response *kms.DescribeSecretResponse
	if h.Item != nil {
		data := h.Item.(*kms.DescribeSecretResponse)
		name = data.SecretName
	} else {
		name = d.EqualsQuals["name"].GetStringValue()
	}

	request := kms.CreateDescribeSecretRequest()
	request.Scheme = "https"
	request.SecretName = name
	request.FetchTags = "true"

	b := retry.NewFibonacci(100 * time.Millisecond)

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeSecret(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				plugin.Logger(ctx).Error("alicloud_kms_key.getKmsSecret", "query_error", err, "request", request)
				return err
			}
		}
		return nil
	})

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "query_retry_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func listKmsSecretVersionIds(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("listKmsSecretVersionIds")

	// Create service connection
	client, err := KMSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "connection_error", err)
		return nil, err
	}
	secretData := h.Item.(*kms.DescribeSecretResponse)
	var response *kms.ListSecretVersionIdsResponse

	request := kms.CreateListSecretVersionIdsRequest()
	request.Scheme = "https"
	request.SecretName = secretData.SecretName
	request.IncludeDeprecated = "true"

	b := retry.NewFibonacci(100 * time.Millisecond)

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.ListSecretVersionIds(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				plugin.Logger(ctx).Error("alicloud_kms_key.listKmsSecretVersionIds", "query_error", err, "request", request)
				return err
			}
		}
		return nil
	})

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.listKmsSecretVersionIds", "retry_query_error", err, "request", request)
		return nil, err
	}

	if response.VersionIds.VersionId != nil && len(response.VersionIds.VersionId) > 0 {
		return response.VersionIds, nil
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func fetchRegionFromArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*kms.DescribeSecretResponse)

	resourceArn := data.Arn
	region := strings.Split(resourceArn, ":")[2]
	return region, nil
}
