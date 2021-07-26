package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
			KeyColumns: plugin.SingleColumn("cert_name"),
			Hydrate:    getUserCertificate,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "cert_name",
				Description: "The name of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "common_name",
				Description: "The common name (CN) attribute of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "CertIdentifier",
				Description: "The certificate identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "issuer",
				Description: "The certificate authority.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "serial_no",
				Description: "The certificate serial no.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "key_size",
				Description: "The certificate key size.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "algorithm",
				Description: "The certificate algorithm.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sign_algorithm",
				Description: "The sign algorithm of the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "after_date",
				Description: "",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "before_date",
				Description: "",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "sans",
				Description: "All domain names bound to the certificate.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sha2",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain_match_cert",
				Description: "",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "md5",
				Description: "",
				Type:        proto.ColumnType_STRING,
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
				Transform:   transform.FromField("CertName"),
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
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// API does not return any error, if the request is made from an unsupported region
	// If the request is made from an unsupported region, it lists all the certificates
	// created in 'cn-hangzhou' region
	// Return nil, if unsupported region (To avoid duplicate entries, when using multi-region configuration)
	if !helpers.StringSliceContains(supportedRegions, region) {
		return nil, nil
	}

	// Create service connection
	client, err := CasService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "connection_error", err)
		return nil, err
	}

	request := cas.CreateDescribeSSLCertificateListRequest()
	request.ShowSize = "50"
	request.CurrentPage = "1"
	request.QueryParams["RegionId"] = region

	count := 0
	for {
		response, err := client.DescribeSSLCertificateList(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_user_certificate.listUserCertificate", "query_error", err, "request", request)
			return nil, err
		}

		for _, i := range response.CertMetaList {
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

	// API does not return any error, if the request is made from an unsupported region
	// If the request is made from an unsupported region, it lists all the certificates
	// created in 'cn-hangzhou' region
	// Return nil, if unsupported region (To avoid duplicate entries, when using multi-region configuration)
	if !helpers.StringSliceContains(supportedRegions, region) {
		return nil, nil
	}

	// Create service connection
	client, err := CasService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "connection_error", err)
		return nil, err
	}

	var name string
	if h.Item != nil {
		name = h.Item.(cas.CertificateInfo).CertName
	} else {
		name = d.KeyColumnQuals["cert_name"].GetStringValue()
	}

	request := cas.CreateDescribeSSLCertificateListRequest()
	request.SearchValue = name

	response, err := client.DescribeSSLCertificateList(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_user_certificate.getUserCertificate", "query_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func getUserCertificateAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificateAka")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	data := h.Item.(cas.CertificateInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cas:" + region + ":" + accountID + ":certificate/" + data.CertName}

	return akas, nil
}

func getUserCertificateRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserCertificateRegion")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	return region, nil
}
