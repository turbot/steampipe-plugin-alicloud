package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

//// TABLE DEFINITION

func tableAlicloudVpcSslVpnServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_ssl_vpn_server",
		Description: "SSL Server refers to the SSL-VPN server within the VPC. It authenticates clients and manages configurations.",
		List: &plugin.ListConfig{
			Hydrate: listVpcVpnSslServers,
			Tags:    map[string]string{"service": "vpc", "action": "DescribeSslVpnServers"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ssl_vpn_server_id"),
			Hydrate:    getVpnSslServer,
			Tags:       map[string]string{"service": "vpc", "action": "DescribeSslVpnServers"},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the SSL-VPN server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ssl_vpn_server_id",
				Description: "The ID of the SSL-VPN server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vpn_gateway_id",
				Description: "The ID of the VPN gateway.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cipher",
				Description: "The encryption algorithm.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "client_ip_pool",
				Description: "The client IP address pool.",
				Type:        proto.ColumnType_CIDR,
			},
			{
				Name:        "connections",
				Description: "The total number of current connections.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "create_time",
				Description: "The time when the SSL-VPN server was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "enable_multi_factor_auth",
				Description: "Indicates whether the multi factor authenticaton is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "internet_ip",
				Description: "The public IP address.",
				Type:        proto.ColumnType_IPADDR,
			},
			{
				Name:        "is_compressed",
				Description: "Indicates whether the transmitted data is compressed.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Compress"),
			},
			{
				Name:        "local_subnet",
				Description: "The CIDR block of the client.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_connections",
				Description: "The maximum number of connections.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "port",
				Description: "The port used by the SSL-VPN server.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "proto",
				Description: "The protocol used by the SSL-VPN server.",
				Type:        proto.ColumnType_STRING,
			},

			// steampipe standard column
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpnSslServerAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ecsVpnSslServerTitle),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
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

func listVpcVpnSslServers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_server.listVpcVpnSslServers", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeSslVpnServersRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeSslVpnServers(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_server.listVpcVpnSslServers", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.SslVpnServers.SslVpnServer {
			plugin.Logger(ctx).Warn("alicloud_vpc_vpn_ssl_server.listVpcVpnSslServers", "Name", i.Name, "item", i)
			d.StreamListItem(ctx, i)
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
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

func getVpnSslServer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpnSslServer")

	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_server.getVpnSslServer", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		sslServer := h.Item.(vpc.SslVpnServer)
		id = sslServer.SslVpnServerId
	} else {
		id = d.EqualsQuals["ssl_vpn_server_id"].GetStringValue()
	}

	request := vpc.CreateDescribeSslVpnServersRequest()
	request.Scheme = "https"
	request.SslVpnServerId = id

	response, err := client.DescribeSslVpnServers(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_server.getVpnSslServer", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.SslVpnServers.SslVpnServer) > 0 {
		return response.SslVpnServers.SslVpnServer[0], nil
	}

	return nil, nil
}

func getVpnSslServerAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpnSslServerAka")
	sslServer := h.Item.(vpc.SslVpnServer)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + sslServer.RegionId + ":" + accountID + ":sslVpnServer/" + sslServer.SslVpnServerId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func ecsVpnSslServerTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	sslServer := d.HydrateItem.(vpc.SslVpnServer)

	// Build resource title
	title := sslServer.SslVpnServerId

	if len(sslServer.Name) > 0 {
		title = sslServer.Name
	}
	return title, nil
}
