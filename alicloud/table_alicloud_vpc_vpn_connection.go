package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

type vpnConnectionInfo = struct {
	VpnConnection vpc.VpnConnection
	Region        string
}

//// TABLE DEFINITION

func tableAlicloudVpcVpnConnection(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_vpn_connection",
		Description: "VPN Connection is an Internet-based tunnel between VPN Gateway and User Gateway.",
		List: &plugin.ListConfig{
			Hydrate: listVpcVpnConnections,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpn_connection_id"),
			Hydrate:    getVpcVpnConnection,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the IPsec-VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpnConnection.Name"),
			},
			{
				Name:        "vpn_connection_id",
				Description: "The ID of the IPsec-VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpnConnection.VpnConnectionId"),
			},
			{
				Name:        "status",
				Description: "The status of the IPsec-VPN connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpnConnection.Status"),
			},
			{
				Name:        "create_time",
				Description: "The time when the IPsec-VPN connection was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("VpnConnection.CreateTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "customer_gateway_id",
				Description: "The ID of the customer gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpnConnection.CustomerGatewayId"),
			},
			{
				Name:        "effect_immediately",
				Description: "Indicates whether IPsec-VPN negotiations are initiated immediately.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("VpnConnection.EffectImmediately"),
			},
			{
				Name:        "enable_dpd",
				Description: "Indicates whether dead peer detection (DPD) is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("VpnConnection.EnableDpd"),
			},
			{
				Name:        "enable_nat_traversal",
				Description: "Indicates whether to enable the NAT traversal feature.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("VpnConnection.EnableNatTraversal"),
			},
			{
				Name:        "local_subnet",
				Description: "The CIDR block of the virtual private cloud (VPC).",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("VpnConnection.LocalSubnet"),
			},
			{
				Name:        "remote_subnet",
				Description: "The CIDR block of the on-premises data center.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("VpnConnection.RemoteSubnet"),
			},
			{
				Name:        "vpn_gateway_id",
				Description: "The ID of the VPN gateway.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpnConnection.VpnGatewayId"),
			},
			{
				Name:        "ike_config",
				Description: "The configurations of Phase 1 negotiations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VpnConnection.IkeConfig"),
			},
			{
				Name:        "ipsec_config",
				Description: "The configurations for Phase 2 negotiations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VpnConnection.IpsecConfig"),
			},
			{
				Name:        "vco_health_check",
				Description: "The health check configurations.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VpnConnection.VcoHealthCheck"),
			},
			{
				Name:        "vpn_bgp_config",
				Description: "BGP configuration information.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VpnConnection.VpnBgpConfig"),
			},

			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpnConnectionAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(vpnConnectionTitle),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
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

func listVpcVpnConnections(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_connection.listVpcVpnConnections", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVpnConnectionsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeVpnConnections(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_vpn_connection.listVpcVpnConnections", "query_error", err, "request", request)
			return nil, err
		}
		for _, vpnConnection := range response.VpnConnections.VpnConnection {
			d.StreamListItem(ctx, vpnConnectionInfo{vpnConnection, region})
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

func getVpcVpnConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcVpnConnection")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_connection.getVpcVpnConnection", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(vpnConnectionInfo)
		id = data.VpnConnection.VpnConnectionId
	} else {
		id = d.KeyColumnQuals["vpn_connection_id"].GetStringValue()
	}

	request := vpc.CreateDescribeVpnConnectionsRequest()
	request.Scheme = "https"
	request.VpnConnectionId = id

	response, err := client.DescribeVpnConnections(request)
	if err != nil {
		return nil, err
	}

	if response.VpnConnections.VpnConnection != nil && len(response.VpnConnections.VpnConnection) > 0 {
		return vpnConnectionInfo{response.VpnConnections.VpnConnection[0], region}, nil
	}

	return nil, nil
}

func getVpnConnectionAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpnConnectionAka")
	data := h.Item.(vpnConnectionInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:vpc:" + data.Region + ":" + accountID + ":vpnconnection/" + data.VpnConnection.VpnConnectionId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func vpnConnectionTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(vpnConnectionInfo)

	// Build resource title
	title := data.VpnConnection.VpnConnectionId

	if len(data.VpnConnection.Name) > 0 {
		title = data.VpnConnection.Name
	}

	return title, nil
}
