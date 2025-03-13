package alicloud

import (
	"context"
	"slices"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

var supportedRegions = []string{"cn-hangzhou", "ap-south-1", "me-east-1", "eu-central-1", "ap-northeast-1", "ap-southeast-2"}

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
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "The ID of the certificate.",
				Type:        proto.ColumnType_DOUBLE,
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

			// Steampipe standard columns
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

			// Alicloud standard columns
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
	region := d.EqualsQualString(matrixKeyRegion)

	// API does not return any error, if the request is made from an unsupported region
	// If the request is made from an unsupported region, it lists all the certificates
	// created in 'cn-hangzhou' region
	// Return nil, if unsupported region (To avoid duplicate entries, when using multi-region configuration)
	if !slices.Contains(supportedRegions, region) {
		return nil, nil
	}

	// Create service connection
	client, err := CasService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "connection_error", err)
		return nil, err
	}

	request := cas.CreateListUserCertificateOrderRequest()
	request.ShowSize = "50"
	request.CurrentPage = "1"
	request.QueryParams["RegionId"] = region

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.ListUserCertificateOrder(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "query_error", err, "request", request)
			return nil, err
		}

		for _, i := range response.CertificateOrderList {
			d.StreamListItem(ctx, i)
			count++
		}
		if count >= int(response.TotalCount) {
			break
		}
		request.CurrentPage = requests.NewInteger(int(response.CurrentPage) + 1)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getUserCertificate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificate")
	region := d.EqualsQualString(matrixKeyRegion)

	// API does not return any error, if the request is made from an unsupported region
	// If the request is made from an unsupported region, it lists all the certificates
	// created in 'cn-hangzhou' region
	// Return nil, if unsupported region (To avoid duplicate entries, when using multi-region configuration)
	if !slices.Contains(supportedRegions, region) {
		return nil, nil
	}

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
		id = d.EqualsQuals["id"].GetInt64Value()
	}

	request := cas.CreateGetUserCertificateDetailRequest()
	request.CertId = requests.NewInteger(int(id))

	response, err := client.GetUserCertificateDetail(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "query_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func getUserCertificateAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificateAka")
	region := d.EqualsQualString(matrixKeyRegion)
	data := casCertificate(h.Item)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
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
	region := d.EqualsQualString(matrixKeyRegion)
	return region, nil
}

func casCertificate(item interface{}) int64 {
	switch item := item.(type) {
	case cas.CertificateOrderListItem:
		return item.CertificateId
	case *cas.GetUserCertificateDetailResponse:
		return item.Id
	}
	return 0
}
