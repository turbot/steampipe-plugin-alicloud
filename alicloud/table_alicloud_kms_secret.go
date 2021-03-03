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

func tableAlicloudKmsSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_secret",
		Description: "Secret enables to manage secrets in a centralized manner throughout their lifecycle (creation, retrieval, updating, and deletion).",
		List: &plugin.ListConfig{
			Hydrate: listKmsSecret,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getKmsSecret,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name: "name",
				Type: proto.ColumnType_STRING,
				// Hydrate:     getKmsSecret,
				Transform:   transform.FromField("SecretName"),
				Description: "The name of the secret.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the secret.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "arn",
				Type:        proto.ColumnType_STRING,
				Description: "The Alibaba Cloud Resource Name (ARN).",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "secret_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the secret.",
			},
			{
				Name:        "encryption_key_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the KMS customer master key (CMK) that is used to encrypt the secret value.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the KMS Secret was created.",
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the KMS Secret was modifies.",
			},
			{
				Name:        "planned_delete_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the KMS Secret is planned to delete.",
			},
			{
				Name:        "automatic_rotation",
				Type:        proto.ColumnType_STRING,
				Description: "Specifies whether automatic key rotation is enabled.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "last_rotation_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Date of last rotation of Secret.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "rotation_interval",
				Type:        proto.ColumnType_STRING,
				Description: "The rotation perion of Secret.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "next_rotation_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date of next rotation of Secret.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "extended_config",
				Type:        proto.ColumnType_JSON,
				Description: "The extended configuration of Secret.",
				Hydrate:     getKmsSecret,
			},
			{
				Name:        "version_ids",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listKmsSecretVersionIds,
				Transform:   transform.FromField("VersionId"),
				Description: "The list of secret versions.",
			},
			{
				Name:      "tags_src",
				Type:      proto.ColumnType_JSON,
				Hydrate:   getKmsSecret,
				Transform: transform.FromField("Tags.Tag").Transform(modifyKmsSourceTags),
			},
			// steampipe standard columns
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

			// alicloud standard columns
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

func listKmsSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := KMSService(ctx, d, region)
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
			plugin.Logger(ctx).Warn("alicloud_kms_secret.listKmsSecret", "tags", i.Tags, "item", i)
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

func getKmsSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := KMSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "connection_error", err)
		return nil, err
	}

	var name string
	if h.Item != nil {
		data := h.Item.(kms.Secret)
		name = data.SecretName
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}
	// panic(name)

	request := kms.CreateDescribeSecretRequest()
	request.Scheme = "https"
	request.SecretName = name
	request.FetchTags = "true"

	response, err := client.DescribeSecret(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "Forbidden.ResourceNotFound" {
			plugin.Logger(ctx).Warn("alicloud_kms_secret.getKmsSecret", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}

func listKmsSecretVersionIds(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := KMSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "connection_error", err)
		return nil, err
	}

	secretData := h.Item.(kms.Secret)

	request := kms.CreateListSecretVersionIdsRequest()
	request.Scheme = "https"
	request.SecretName = secretData.SecretName
	request.IncludeDeprecated = "true"

	response, err := client.ListSecretVersionIds(request)
	if err != nil {
		plugin.Logger(ctx).Error("listKmsSecretVersionIds", "query_error", err, "request", request)
		return nil, err
	}

	if response.VersionIds.VersionId != nil && len(response.VersionIds.VersionId) > 0 {
		return response.VersionIds, nil
	}

	return nil, nil
}
