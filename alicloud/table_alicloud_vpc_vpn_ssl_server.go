package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcVpnSslServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_vpn_ssl_server",
		Description: "SSL Server refers to the SSL-VPN server within the VPC. It authenticates clients and manages configurations.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listVpcVpnSslServer,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "The name of the SSL-VPN server."},
			{Name: "ssl_vpn_server_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SslVpnServerId"), Description: "The ID of the SSL-VPN server."},
			// Other columns
			{Name: "region_id", Type: proto.ColumnType_STRING, Description: "The ID of the region where the SSL-VPN server is created."},
			{Name: "vpn_gateway_id", Type: proto.ColumnType_STRING, Description: "The ID of the VPN gateway."},
			{Name: "create_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp), Description: "The time when the SSL-VPN server was created."},
			{Name: "local_subnet", Type: proto.ColumnType_STRING, Description: "The CIDR block of the client."},
			{Name: "client_ip_pool", Type: proto.ColumnType_STRING, Description: "The client IP address pool."},
			{Name: "cipher", Type: proto.ColumnType_STRING, Description: "The encryption algorithm."},
			{Name: "proto", Type: proto.ColumnType_STRING, Description: "The protocol used by the SSL-VPN server."},
			{Name: "port", Type: proto.ColumnType_INT, Description: "The port used by the SSL-VPN server."},
			{Name: "Compress", Type: proto.ColumnType_BOOL, Description: "Indicates whether the transmitted data is compressed."},
			{Name: "connections", Type: proto.ColumnType_INT, Description: "The total number of current connections."},
			{Name: "max_connections", Type: proto.ColumnType_INT, Description: "The maximum number of connections."},
			{Name: "internet_ip", Type: proto.ColumnType_STRING, Description: "The public IP address."},
			{Name: "enable_multi_factor_auth", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "idaas_instance_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("IDaaSInstanceId"), Description: ""},

			// Resource interface
			// {Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(SslServerToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			// TODO - It appears that Tags are not returned by the go SDK?
			// {Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listVpcVpnSslServer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpcVpnSslServer", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeSslVpnServersRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	// if quals["is_default"] != nil {
	// 	request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	// }
	if quals["id"] != nil {
		request.SslVpnServerId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeSslVpnServers(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc.listVpcVpnSslServer", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.SslVpnServers.SslVpnServer {
			plugin.Logger(ctx).Warn("alicloud_vpc.listVpcVpnSslServer", "Name", i.Name, "item", i)
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
