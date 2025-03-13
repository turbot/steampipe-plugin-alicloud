package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudVpcVpnCustomerGateway(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_vpn_customer_gateway",
		Description: "Alicloud VPC VPN Customer Gateway.",
		List: &plugin.ListConfig{
			Hydrate: listVpcCustomerGateways,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("customer_gateway_id"),
			Hydrate:    getVpcCustomerGateway,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the customer gateway.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "customer_gateway_id",
				Description: "The ID of the customer gateway.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "asn",
				Description: "Specifies the ASN of the customer gateway.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "create_time",
				Description: "The time when the customer gateway was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "description",
				Description: "The description of the customer gateway.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ip_address",
				Description: "The IP address of the customer gateway.",
				Type:        proto.ColumnType_IPADDR,
			},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcCustomerGatewayAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(vpcCustomerGatewayTitle),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpnCustomerGatewayRegion,
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

func listVpcCustomerGateways(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.listVpcCustomerGateways", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeCustomerGatewaysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeCustomerGateways(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.listVpcCustomerGateways", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.CustomerGateways.CustomerGateway {
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

//// HYDRATE FUNCTIONS

func getVpcCustomerGateway(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.getVpcCustomerGateway", "connection_error", err)
		return nil, err
	}
	id := d.EqualsQuals["customer_gateway_id"].GetStringValue()

	request := vpc.CreateDescribeCustomerGatewaysRequest()
	request.Scheme = "https"
	request.CustomerGatewayId = id

	response, err := client.DescribeCustomerGateways(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.getVpcCustomerGateway", "query_error", err, "request", request)
		return nil, err
	}

	if len(response.CustomerGateways.CustomerGateway) > 0 {
		return response.CustomerGateways.CustomerGateway[0], nil
	}

	return nil, nil
}

func getVpcCustomerGatewayAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcCustomerGatewayAka")
	data := h.Item.(vpc.CustomerGateway)
	region := d.EqualsQualString(matrixKeyRegion)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + region + ":" + accountID + ":customergateway/" + data.CustomerGatewayId}

	return akas, nil
}

func getVpnCustomerGatewayRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpnCustomerGatewayRegion")
	region := d.EqualsQualString(matrixKeyRegion)

	return region, nil
}

//// TRANSFORM FUNCTIONS

func vpcCustomerGatewayTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(vpc.CustomerGateway)

	// Build resource title
	title := data.CustomerGatewayId

	if len(data.Name) > 0 {
		title = data.Name
	}

	return title, nil
}
