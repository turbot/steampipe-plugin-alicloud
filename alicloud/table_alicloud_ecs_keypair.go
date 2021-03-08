package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type keypairInfo = struct {
	KeyPair ecs.KeyPair
	Region  string
}

//// TABLE DEFINITION

func tableAlicloudEcskeypair(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_keypair",
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
				Transform:   transform.FromField("KeyPair.KeyPairName"),
			},
			{
				Name:        "key_pair_finger_print",
				Description: "The fingerprint of the key pair.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyPair.KeyPairFingerPrint"),
			},
			{
				Name:        "creation_time",
				Description: "The time when the key pair was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("KeyPair.CreationTime"),
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the key pair belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyPair.ResourceGroupId"),
			},

			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("KeyPair.Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("KeyPair.Tags.Tag").Transform(ecsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsKeypairAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KeyPair.KeyPairName"),
			},

			// alibaba standard columns
			{
				Name:        "region",
				Description: "The region ID where the resource is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "account_id",
				Description: "The alicloud Account ID in which the resource is located.",
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
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_keypair.listEcsKeypair", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeKeyPairsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	//	request.RegionId = region
	count := 0
	for {
		response, err := client.DescribeKeyPairs(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_keypair.listEcsKeypair", "query_error", err, "request", request)
			return nil, err
		}
		for _, keypair := range response.KeyPairs.KeyPair {
			plugin.Logger(ctx).Warn("listEcsKeypair", "item", keypair)
			d.StreamListItem(ctx, keypairInfo{keypair, region})
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

	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
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
	//request.RegionId = region

	response, err := client.DescribeKeyPairs(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_keypair.getEcsKeypair", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.KeyPairs.KeyPair != nil && len(response.KeyPairs.KeyPair) > 0 {
		return keypairInfo{response.KeyPairs.KeyPair[0], region}, nil
	}

	return nil, nil
}

func getEcsKeypairAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsKeypairAka")
	data := h.Item.(keypairInfo)

	// Get account details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ecs:" + data.Region + ":" + accountID + ":keypair/" + data.KeyPair.KeyPairName}

	return akas, nil
}
