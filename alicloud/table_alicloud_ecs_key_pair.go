package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcskeyPair(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_key_pair",
		Description: "An SSH key pair is a secure and convenient authentication method provided by Alibaba Cloud for instance logon. An SSH key pair consists of a public key and a private key. You can use SSH key pairs to log on to only Linux instances.",
		List: &plugin.ListConfig{
			Hydrate: listEcsKeypair,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getEcsKeypair,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the key pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyPairName"),
			},
			{
				Name:        "key_pair_finger_print",
				Description: "The fingerprint of the key pair.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the key pair was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the key pair belongs.",
				Type:        proto.ColumnType_STRING,
			},

			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
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
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyPairName"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsKeypairAka,
				Transform:   transform.FromValue(),
			},
			// Alibaba standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getKeyPairRegion,
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

func listEcsKeypair(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := ECSService(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_keypair.listEcsKeypair", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeKeyPairsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	count := 0
	for {
		response, err := client.DescribeKeyPairs(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_keypair.listEcsKeypair", "query_error", err, "request", request)
			return nil, err
		}
		for _, keypair := range response.KeyPairs.KeyPair {
			plugin.Logger(ctx).Warn("listEcsKeypair", "item", keypair)
			d.StreamListItem(ctx, keypair)
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

func getEcsKeypair(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSnapshot")

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_keypair.getEcsKeypair", "connection_error", err)
		return nil, err
	}

	var name string
	if h.Item != nil {
		keypair := h.Item.(ecs.KeyPair)
		name = keypair.KeyPairName
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	request := ecs.CreateDescribeKeyPairsRequest()
	request.Scheme = "https"
	request.KeyPairName = name

	response, err := client.DescribeKeyPairs(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_keypair.getEcsKeypair", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.KeyPairs.KeyPair != nil && len(response.KeyPairs.KeyPair) > 0 {
		return response.KeyPairs.KeyPair[0], nil
	}

	return nil, nil
}

func getEcsKeypairAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsKeypairAka")
	data := h.Item.(ecs.KeyPair)
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Get account details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ecs:" + region + ":" + accountID + ":keypair/" + data.KeyPairName}

	return akas, nil
}

func getKeyPairRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	return region, nil
}
