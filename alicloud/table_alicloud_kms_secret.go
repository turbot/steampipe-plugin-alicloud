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
		Name:        "alicloud_vpc",
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
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
			{Name: "secret_name", Type: proto.ColumnType_STRING, Description: "The name of the secret."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the secret."},
			{Name: "secret_type", Type: proto.ColumnType_STRING, Description: "The type of the secret."},
			// Other columns
			{Name: "create_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the KMS Secret was created."},
			{Name: "update_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the KMS Secret was modifies."},
			{Name: "planned_delete_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the KMS Secret is planned to delete."},
			{Name: "encryption_key_id", Type: proto.ColumnType_STRING, Description: "The key id with which the KMS secret is encrypted"},
			{Name: "automatic_rotation", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "last_rotation_date", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "rotation_interval", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "next_rotation_date", Type: proto.ColumnType_STRING, Description: "To be update"},
			{Name: "extended_config", Type: proto.ColumnType_STRING, Description: "To be update"},
			// Resource interface
			// {Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(vpcToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			// TODO - It appears that Tags are not returned by the go SDK?
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("VpcName"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listKmsSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectKms(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpc", "connection_error", err)
		return nil, err
	}
	request := kms.CreateListSecretsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	// quals := d.KeyColumnQuals
	// if quals["is_default"] != nil {
	// 	request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	// }
	// if quals["id"] != nil {
	// 	request.VpcId = quals["id"].GetStringValue()
	// }

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
	// i := h.Item.(vpc.Vpc)
	// request.VpcId = i.VpcId
	response, err := client.DescribeSecret(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_kms_secret.getKmsSecret", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

// func vpcToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := d.Value.(vpc.Vpc)
// 	return "acs:vpc:" + i.RegionId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vpc/" + i.VpcName, nil
// }
