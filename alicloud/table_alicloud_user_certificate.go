package alicloud

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type certInfo struct {
	CertificateList []certificateInfo
}

type certificateInfo struct {
	country     string
	orgName     string
	city        string
	endDate     string
	sans        string
	issuer      string
	expired     bool
	common      string
	province    string
	name        string
	fingerprint string
	buyInAliyun bool
	id          string
	startDate   string
	Cert        string
	Key         string
}

//// TABLE DEFINITION

func tableAlicloudUserCertificate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_user_certificate",
		Description: "Alicloud User Certificate",
		List: &plugin.ListConfig{
			Hydrate: listUserCertificate,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cert_id"),
			Hydrate:    getUserCertificate,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cert_id",
				Description: "The ID of the certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("id"),
			},
			{
				Name:        "start_date",
				Description: "The issuance date of the certificate.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "sans",
				Description: "All domain names bound to the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "province",
				Description: "The province where the organization that purchases the certificate is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_name",
				Description: "The name of the organization that purchases the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "end_date",
				Description: "The expiration date of the certificate.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "issuer",
				Description: "The certificate authority.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "fingerprint",
				Description: "The certificate fingerprint.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "expired",
				Description: "Indicates whether the certificate has expired.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "Country",
				Description: "The country where the organization that purchases the certificate is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "common",
				Description: "The common name (CN) attribute of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "City",
				Description: "The city where the organization that purchases the certificate is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "buy_in_aliyun",
				Description: "Indicates whether the certificate was purchased from Alibaba Cloud.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "cert",
				Description: "The certificate content, in PEM format.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getUserCertificate,
			},
			{
				Name:        "key",
				Description: "The private key of the certificate, in PEM format.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getUserCertificate,
			},
			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getUserCertificateAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
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

func listUserCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := CommonService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "connection_error", err)
		return nil, err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "cas.aliyuncs.com"
	request.Version = "2018-07-13"
	request.ApiName = "DescribeUserCertificateList"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["ShowSize"] = "50"
	request.QueryParams["CurrentPage"] = "1"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "query_error", err, "request", request)
		return nil, err
	}
	cert := response.GetHttpContentString()
	cert1 := strings.Replace(cert, "\"{", "{", 1)
	cert2 := strings.Replace(cert1, "}\"", "}", 1)
	res := certInfo{}
	json.Unmarshal([]byte(cert2), &res)
	plugin.Logger(ctx).Warn("alicloud_user_certificate.listUserCertificate", "cert_value", res)
	for _, i := range res.CertificateList {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

func getUserCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := CommonService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "connection_error", err)
		return nil, err
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "cas.aliyuncs.com"
	request.Version = "2018-07-13"
	request.ApiName = "DescribeUserCertificateDetail"
	request.QueryParams["RegionId"] = "cn-hangzhou"

	var id string
	if h.Item != nil {
		data := h.Item.(certificateInfo)
		id = data.id
	} else {
		id = d.KeyColumnQuals["cert_id"].GetStringValue()
	}
	request.QueryParams["CertId"] = id

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "query_error", err, "request", request)
		return nil, err
	}
	cert := response.GetHttpContentString()
	res := certificateInfo{}
	json.Unmarshal([]byte(cert), &res)
	plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "get_value", res)
	return res, nil
}

//// HYDRATE FUNCTIONS

func getUserCertificateAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificateAka")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	//cert := h.Item.(certInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + region + ":" + accountID + ":disk/"}

	return akas, nil
}
