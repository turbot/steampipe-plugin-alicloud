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

//// TABLE DEFINITION

func tableAlicloudEcsEni(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_eni",
		Description: "Elastic Compute Service Eni",
		List: &plugin.ListConfig{
			Hydrate: listEcsEni,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("network_interface_id"),
			Hydrate:    getEcsEni,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Description: "The name of the ENI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkInterfaceName"),
			},
			{
				Name:        "network_interface_id",
				Description: "An unique identifier for the ENI.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkInterfaceId"),
			},
			{
				Name:        "type",
				Description: "The type of the ENI. Valid values: 'Primary' and 'Secondary'",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the ENI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC to which the ENI belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vswitch_id",
				Description: "The ID of the VSwitch to which the ENI is connected.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VSwitchId"),
			},
			{
				Name:        "description",
				Description: "The description of the ENI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the ENI was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "zone_id",
				Description: "The zone ID of the ENI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the ENI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "mac_address",
				Description: "The MAC address of the ENI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance to which the ENI is bound.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_id",
				Description: "The ID of the distributor to which the ENI belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ServiceID"),
			},
			{
				Name:        "service_managed",
				Description: "Indicates whether the user is an Alibaba Cloud service or a distributor.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the ENI belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner_id",
				Description: "The ID of the account that owns the ENI.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "security_groupIds",
				Description: "The IDs of the security groups to which the ENI belongs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "associated_public_ip",
				Description: "The public IP address associated with the secondary private IP address of the ENI.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "attachment",
				Description: "Attachments of the ENI",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "private_ip_sets",
				Description: "The private IP addresses of the ENI.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "ipv6_sets",
				Description: "The IPv6 addresses assigned to the ENI.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Ipv6Sets"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(ecsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsEniAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkInterfaceName"),
			},

			// alibaba standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ZoneId").Transform(zoneToRegion),
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

func listEcsEni(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_eni.listEcsEni", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeNetworkInterfacesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeNetworkInterfaces(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_eni.listEcsEni", "query_error", err, "request", request)
			return nil, err
		}
		for _, eni := range response.NetworkInterfaceSets.NetworkInterfaceSet {
			plugin.Logger(ctx).Warn("listEcsEni", "item", eni)
			d.StreamListItem(ctx, eni)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}

//// GET FUNCTION

func getEcsEni(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsEni")

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_eni.getEcsEni", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		NetworkInterfaceSet := h.Item.(ecs.NetworkInterfaceSet)
		id = NetworkInterfaceSet.NetworkInterfaceId
	} else {
		id = d.KeyColumnQuals["network_interface_id"].GetStringValue()
	}

	request := ecs.CreateDescribeNetworkInterfaceAttributeRequest()
	request.Scheme = "https"
	request.NetworkInterfaceId = id

	response, err := client.DescribeNetworkInterfaceAttribute(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_eni.getEcsEni", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}

//// TRANSFORM FUNCTIONS

func ecsEniTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	eni := d.HydrateItem.(ecs.NetworkInterfaceSet)

	var turbotTagsMap map[string]string

	if eni.Tags.Tag == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range eni.Tags.Tag {
		turbotTagsMap[i.TagKey] = i.TagValue
	}

	return turbotTagsMap, nil
}

func getEcsEniAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsEniAka")
	eni := h.Item.(ecs.NetworkInterfaceSet)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + eni.ZoneId + ":" + accountID + ":eni/" + eni.NetworkInterfaceId}

	return akas, nil
}
