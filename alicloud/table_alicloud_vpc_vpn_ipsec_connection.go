package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcVpnIpsecConnection(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_vpn_ipsec_conection",
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listVpcVpnIpsecConnection,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "The name of the IPsec-VPN connection."},
			{Name: "vpn_connection_id", Type: proto.ColumnType_STRING, Description: "The name of the IPsec-VPN connection."},
			// Other columns
			{Name: "customer_gateway_id", Type: proto.ColumnType_STRING, Description: "The ID of the customer gateway."},
			{Name: "vpn_gateway_id", Type: proto.ColumnType_STRING, Description: "The ID of the VPN gateway."},
			{Name: "create_time", Type: proto.ColumnType_DOUBLE, Description: "The time when the IPsec-VPN connection was created."},
			{Name: "local_subnet", Type: proto.ColumnType_STRING, Description: "The CIDR block of the VPC."},
			{Name: "remote_subnet", Type: proto.ColumnType_STRING, Description: "The CIDR block of the on-premises data center."},
			{Name: "EffectImmediately", Type: proto.ColumnType_BOOL, Description: "Indicates whether the connection immediately takes effect"},
			{Name: "Status", Type: proto.ColumnType_STRING, Description: "The status of the connection."},
			{Name: "enable_dpd", Type: proto.ColumnType_BOOL, Description: "Indicates whether dead peer detection (DPD) is enabled."},
			{Name: "enable_nat_traversal", Type: proto.ColumnType_BOOL, Description: "Indicates whether to enable the NAT traversal feature."},
			{Name: "ike_config", Type: proto.ColumnType_JSON, Transform: transform.FromField("IkeConfig"), Description: "The configurations of Phase 1 negotiations."},
			{Name: "ipsec_config", Type: proto.ColumnType_JSON, Transform: transform.FromField("IpsecConfig"), Description: "The configurations for Phase 2 negotiations."},
			{Name: "vco_health_check", Type: proto.ColumnType_JSON, Transform: transform.FromField("VcoHealthCheck"), Description: "The health check configurations."},
			{Name: "VpnBgpConfig", Type: proto.ColumnType_JSON, Transform: transform.FromField("VpnBgpConfig"), Description: "BGP configuration information."},

			// Resource interface
			// {Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(SslServerToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			// TODO - It appears that Tags are not returned by the go SDK?
			// {Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listVpcVpnIpsecConnection(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpcVpnIpsecConnection", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVpnConnectionsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	// if quals["is_default"] != nil {
	// 	request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	// }
	if quals["id"] != nil {
		request.VpnConnectionId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeVpnConnections(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc.listVpcVpnIpsecConnection", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.VpnConnections.VpnConnection {
			plugin.Logger(ctx).Warn("alicloud_vpc.listVpcVpnIpsecConnection", "Name", i.Name, "item", i)
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

// func SslServerToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := d.Value.(vpc.SslVpnServer)
// 	return "acs:vpc:" + i.RegionId + ":" + i.SslVpnServerId + ":vpc/" + i.Name, nil
// }
