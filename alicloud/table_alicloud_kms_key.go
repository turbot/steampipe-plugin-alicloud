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

func tableAlicloudKmsKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_kms_key",
		Description: "Key Management Service (KMS) provides secure and compliant key management and cryptography services to help you encrypt and protect sensitive data assets.",
		List: &plugin.ListConfig{
			Hydrate: listKmsKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getKmsKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("KeyId"), Description: "The globally unique ID of the CMK."},
			{Name: "key_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("KeyState"), Description: "The status of the CMK."},
			// Other columns
			{Name: "arn", Type: proto.ColumnType_TIMESTAMP, Description: "The Alibaba Cloud Resource Name (ARN) of the CMK."},
			{Name: "creation_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the CMK was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the CMK."},
			{Name: "creator", Type: proto.ColumnType_STRING, Description: "The creator of the CMK."},
			{Name: "key_usage", Type: proto.ColumnType_STRING, Description: "The purpose of the CMK."},
			{Name: "key_spec", Type: proto.ColumnType_STRING, Description: "The type of the CMK."},
			{Name: "last_rotation_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the last rotation was performed."},
			{Name: "next_rotation_date", Type: proto.ColumnType_TIMESTAMP, Description: "The time the next rotation is scheduled for execution."},
			{Name: "automatic_rotation", Type: proto.ColumnType_STRING, Description: "Indicates whether automatic key rotation is enabled."},
			{Name: "delete_date", Type: proto.ColumnType_TIMESTAMP, Description: "The date and time the CMK is scheduled for deletion."},
			{Name: "origin", Type: proto.ColumnType_STRING, Description: "The source of the key material for the CMK."},
			{Name: "protection_level", Type: proto.ColumnType_STRING, Description: "The protection level of the CMK."},
			{Name: "primary_key_version", Type: proto.ColumnType_STRING, Description: "The ID of the current primary key version of the symmetric CMK. The primary key version of a CMK is an active encryption key. KMS uses the primary key version of a specified CMK to encrypt data. "},
			{Name: "rotation_interval", Type: proto.ColumnType_STRING, Description: "The period of automatic key rotation. Unit: seconds."},
			{Name: "material_expire_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time and date the key material for the CMK expires."},
		},
	}
}

func listKmsKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectKms(ctx)
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

func getKmsKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := connectKms(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "connection_error", err)
		return nil, err
	}

	request := kms.CreateDescribeKeyRequest()
	request.Scheme = "https"
	i := h.Item.(kms.KeyMetadata)
	request.KeyId = i.KeyId

	quals := d.KeyColumnQuals
	if quals["id"] != nil {
		request.KeyId = quals["id"].GetStringValue()
	}

	response, err := client.DescribeKey(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "EntityNotExist.Key" {
			plugin.Logger(ctx).Warn("alicloud_kms_key.getKmsKey", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_kms_key.getKmsKey", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.KeyMetadata, nil
}
