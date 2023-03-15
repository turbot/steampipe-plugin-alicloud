package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type vpnSslClientCertInfo = struct {
	Name               string
	SslVpnClientCertId string
	SslVpnServerId     string
	Status             string
	CreateTime         int64
	EndTime            int64
	CaCert             string
	ClientCert         string
	ClientKey          string
	ClientConfig       string
	Region             string
}

//// TABLE DEFINITION

func tableAlicloudVpcSslVpnClientCert(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_ssl_vpn_client_cert",
		Description: "SSL Client is responsible for managing client certificates. The client needs to first complete certificate verification in order to connect to the SSL Server.",
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("ssl_vpn_client_cert_id"),
			ShouldIgnoreError: isNotFoundError([]string{"InvalidSslVpnClientCertId.NotFound"}),
			Hydrate:           getVpcSslVpnClientCert,
		},
		List: &plugin.ListConfig{
			Hydrate: listVpcSslVpnClientCerts,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the SSL client certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ssl_vpn_client_cert_id",
				Description: "The ID of the SSL client certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ssl_vpn_server_id",
				Description: "The ID of the SSL-VPN server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the client certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "The time when the SSL client certificate was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "end_time",
				Description: "The time when the SSL client certificate expires.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("EndTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "ca_cert",
				Description: "The CA certificate.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcSslVpnClientCert,
			},
			{
				Name:        "client_cert",
				Description: "The client certificate.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcSslVpnClientCert,
			},
			{
				Name:        "client_key",
				Description: "The client key.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcSslVpnClientCert,
			},
			{
				Name:        "client_config",
				Description: "The client configuration.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcSslVpnClientCert,
			},

			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcSslVpnClientCertCertAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(sslVpnClientCertTitle),
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

func listVpcSslVpnClientCerts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_client.listVpcSslVpnClientCerts", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeSslVpnClientCertsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeSslVpnClientCerts(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_client.listVpcSslVpnClientCerts", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.SslVpnClientCertKeys.SslVpnClientCertKey {
			d.StreamListItem(ctx, vpnSslClientCertInfo{i.Name, i.SslVpnClientCertId, i.SslVpnServerId, i.Status, i.CreateTime, i.EndTime, "", "", "", "", i.RegionId})
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

func getVpcSslVpnClientCert(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcSslVpnClientCert")

	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_client.getVpcSslVpnClientCert", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(vpnSslClientCertInfo)
		id = data.SslVpnClientCertId
	} else {
		id = d.EqualsQuals["ssl_vpn_client_cert_id"].GetStringValue()
	}

	request := vpc.CreateDescribeSslVpnClientCertRequest()
	request.Scheme = "https"
	request.SslVpnClientCertId = id

	data, err := client.DescribeSslVpnClientCert(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_ssl_client.getVpcSslVpnClientCert", "query_error", err, "request", request)
		return nil, err
	}

	return vpnSslClientCertInfo{data.Name, data.SslVpnClientCertId, data.SslVpnServerId, data.Status, data.CreateTime, data.EndTime, data.CaCert, data.ClientCert, data.ClientKey, data.ClientConfig, data.RegionId}, nil
}

func getVpcSslVpnClientCertCertAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcSslVpnClientCertCertAka")

	data := h.Item.(vpnSslClientCertInfo)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + data.Region + ":" + accountID + ":sslclientcert/" + data.SslVpnClientCertId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func sslVpnClientCertTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(vpnSslClientCertInfo)

	// Build resource title
	title := data.SslVpnClientCertId

	if len(data.Name) > 0 {
		title = data.Name
	}

	return title, nil
}
