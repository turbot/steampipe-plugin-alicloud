package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcVpnGateway(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vswitch",
		Description: "VSwitches to divide the VPC network into one or more subnets.",
		List: &plugin.ListConfig{
			Hydrate: listVpnGateways,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpn_gateway_id"),
			Hydrate:    getVpnGateway,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the VPC.",
			},
			{
				Name:        "vpn_gateway_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpnGatewayId"),
				Description: "The unique ID of the VPC.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the VPC.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_STRING,
				Description: "The creation time of the VPC.",
			},
			{
				Name:        "end_time",
				Type:        proto.ColumnType_STRING,
				Description: "The creation time of the VPC.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the VPC.",
			},
			{
				Name:        "business_status",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the VPC.",
			},
			{
				Name:        "enable_bgp",
				Type:        proto.ColumnType_BOOL,
				Description: "The creation time of the VPC.",
			},
			{
				Name:        "auto_propagate",
				Type:        proto.ColumnType_BOOL,
				Description: "The creation time of the VPC.",
			},

			{
				Name:        "internet_ip",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available.",
			},
			{
				Name:        "vswitch_d",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VSwitchId"),
				Description: "The IPv4 CIDR block of the VPC.",
			},
			{
				Name:        "spec",
				Type:        proto.ColumnType_STRING,
				Description: "The IPv6 CIDR block of the VPC.",
			},
			{
				Name:        "charge_type",
				Type:        proto.ColumnType_STRING,
				Description: "The zone to which the VSwitch belongs.",
			},
			{
				Name:        "ipsec_vpn",
				Type:        proto.ColumnType_STRING,
				Description: "The zone to which the VSwitch belongs.",
			},
			{
				Name:        "ssl_vpn",
				Type:        proto.ColumnType_STRING,
				Description: "The number of available IP addresses in the VSwitch.",
			},
			{
				Name:        "ssl_max_connections",
				Type:        proto.ColumnType_STRING,
				Description: "The number of available IP addresses in the VSwitch.",
			},
			{
				Name:        "Auto_propagate",
				Type:        proto.ColumnType_BOOL,
				Description: "The number of available IP addresses in the VSwitch.",
			},
			{
				Name:        "reservation_data",
				Type:        proto.ColumnType_JSON,
				Description: "The number of available IP addresses in the VSwitch.",
			},
			{
				Name:        "tag",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the VPC.",
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag"),
				Description: ColumnDescriptionTags,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
				Description: ColumnDescriptionTitle,
			},
			// alicloud common columns
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

func listVpnGateways(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vswitch.listVSwitch", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVpnGatewaysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	if quals["id"] != nil {
		request.VpnGatewayId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeVpnGateways(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_vpn_gateway.listVpnGateways", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.VpnGateways.VpnGateway {
			plugin.Logger(ctx).Warn("alicloud_vpc_vpn_gateway.listVpnGateways", "tags", i.Tags, "item", i)
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

func getVpnGateway(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_gateway.getVpnGateway", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVpnGatewayRequest()
	request.Scheme = "https"
	var id string
	if h.Item != nil {
		data := h.Item.(vpc.VpnGateway)
		id = data.VpnGatewayId
	} else {
		id = d.KeyColumnQuals["vpn_gateway_id"].GetStringValue()
	}
	request.VpnGatewayId = id
	response, err := client.DescribeVpnGateway(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_gateway.getVpnGateway", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}
