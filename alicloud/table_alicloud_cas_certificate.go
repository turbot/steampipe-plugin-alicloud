package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

// var supportedRegion = []string{"cn-hangzhou", "ap-south-1", "me-east-1", "eu-central-1", "ap-northeast-1", "ap-southeast-2"}

//// TABLE DEFINITION

func tableAlicloudUserCertificate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_cas_certificate",
		Description: "Alicloud CAS Certificate",
		List: &plugin.ListConfig{
			Hydrate: listUserCertificate,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
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
				Name:        "id",
				Description: "The ID of the certificate.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "org_name",
				Description: "The name of the organization that purchases the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "issuer",
				Description: "The certificate authority.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "buy_in_aliyun",
				Description: "Indicates whether the certificate was purchased from Alibaba Cloud.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "common",
				Description: "The common name (CN) attribute of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "expired",
				Description: "Indicates whether the certificate has expired.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "fingerprint",
				Description: "The certificate fingerprint.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start_date",
				Description: "The issuance date of the certificate.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "end_date",
				Description: "The expiration date of the certificate.",
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
				Name:        "country",
				Description: "The country where the organization that purchases the certificate is located.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "city",
				Description: "The city where the organization that purchases the certificate is located.",
				Type:        proto.ColumnType_STRING,
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

			//steampipe standard columns
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
				Hydrate:     getUserCertificateRegion,
				Transform:   transform.FromValue(),
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
	client, err := CasService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "connection_error", err)
		return nil, err
	}

	request := cas.CreateDescribeUserCertificateListRequest()
	request.ShowSize = "50"
	request.CurrentPage = "1"
	request.QueryParams["RegionId"] = region

	count := 0
	for {
		response, err := client.DescribeUserCertificateList(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "query_error", err, "request", request)
			return nil, err
		}

		for _, i := range response.CertificateList {
			d.StreamListItem(ctx, i)
			count++
		}
		if count >= response.TotalCount {
			break
		}
		request.CurrentPage = requests.NewInteger(response.CurrentPage + 1)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getUserCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificate")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := CasService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "connection_error", err)
		return nil, err
	}

	var id int64
	if h.Item != nil {
		data := casCertificate(h.Item)
		id = data
	} else {
		id = d.KeyColumnQuals["id"].GetInt64Value()
	}

	request := cas.CreateDescribeUserCertificateDetailRequest()
	request.CertId = requests.NewInteger(int(id))

	response, err := client.DescribeUserCertificateDetail(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "query_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func getUserCertificateAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificateAka")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	data := casCertificate(h.Item)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cas:" + region + ":" + accountID + ":certificate/" + strconv.Itoa(int(data))}

	return akas, nil
}

func getUserCertificateRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificateRegion")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	return region, nil
}

func casCertificate(item interface{}) int64 {
	switch item.(type) {
	case cas.Certificate:
		return item.(cas.Certificate).Id
	case *cas.DescribeUserCertificateDetailResponse:
		return item.(*cas.DescribeUserCertificateDetailResponse).Id
	}
	return 0
}
