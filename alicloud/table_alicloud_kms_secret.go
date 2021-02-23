package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudKmsSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_secret",
		Description: "Secret enables to manage secrets in a centralized manner throughout their lifecycle (creation, retrieval, updating, and deletion.)",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listKmsSecret,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getKmsSecret,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecretName"), Description: "The name of the secret."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the secret."},
			{Name: "arn", Type: proto.ColumnType_STRING, Description: "The Alibaba Cloud Resource Name (ARN)."},
			{Name: "secret_type", Type: proto.ColumnType_STRING, Description: "The type of the secret."},
			{Name: "encryption_key_id", Type: proto.ColumnType_STRING, Description: "The ID of the KMS customer master key (CMK) that is used to encrypt the secret value."},
			{Name: "create_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the KMS Secret was created."},
			{Name: "update_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the KMS Secret was modifies."},
			{Name: "planned_delete_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the KMS Secret is planned to delete."},
			{Name: "automatic_rotation", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "last_rotation_date", Type: proto.ColumnType_TIMESTAMP, Description: "To be update"},
			{Name: "rotation_interval", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "next_rotation_date", Type: proto.ColumnType_TIMESTAMP, Description: "To be update"},
			{Name: "extended_config", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("SecretName"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listKmsSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectKms(ctx)
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
	client, err := connectKms(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "connection_error", err)
		return nil, err
	}
	request := kms.CreateDescribeSecretRequest()
	request.Scheme = "https"
	i := h.Item.(kms.Secret)
	request.SecretName = i.SecretName

	quals := d.KeyColumnQuals
	if quals["name"] != nil {
		request.SecretName = quals["name"].GetStringValue()
	}
	response, err := client.DescribeSecret(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}
