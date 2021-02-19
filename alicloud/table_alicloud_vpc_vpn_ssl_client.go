package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcVpnSslClient(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_vpn_ssl_client",
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listVpcVpnSslClient,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: "The name of the SSL client certificate."},
			{Name: "ssl_vpn_server_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("SslVpnServerId"), Description: "The ID of the SSL-VPN server."},
			// Other columns
			{Name: "region_id", Type: proto.ColumnType_STRING, Description: "The region of the SSL client certificate to query"},
			{Name: "ssl_vpn_client_cert_id", Type: proto.ColumnType_STRING, Description: "The ID of the SSL client certificate."},
			{Name: "create_time", Type: proto.ColumnType_INT, Description: "The time when the SSL client certificate was created."},
			{Name: "end_time", Type: proto.ColumnType_INT, Description: "The time when the SSL client certificate expires."},
			{Name: "Status", Type: proto.ColumnType_STRING, Description: "The status of the client certificate"},

			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(SslClientToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			// TODO - It appears that Tags are not returned by the go SDK?
			// {Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listVpcVpnSslClient(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpcVpnSslClient", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeSslVpnClientCertsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	// if quals["is_default"] != nil {
	// 	request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	// }
	if quals["id"] != nil {
		request.SslVpnClientCertId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeSslVpnClientCerts(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc.listVpcVpnSslClient", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.SslVpnClientCertKeys.SslVpnClientCertKey {
			plugin.Logger(ctx).Warn("alicloud_vpc.listVpcVpnSslClient", "Name", i.Name, "item", i)
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

func SslClientToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(vpc.SslVpnClientCertKey)
	return "acs:vpc:" + i.RegionId + ":" + i.SslVpnClientCertId + ":vpc/" + i.Name, nil
}
