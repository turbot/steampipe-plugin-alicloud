package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type CustomerGatewayInfo = struct {
	CustomerGateway vpc.CustomerGateway
}

func tableAlicloudVpcVpnCustomerGateway(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_vpc_customer_gateway",
		Description: "NAT gateways are.",
		List: &plugin.ListConfig{
			Hydrate: listCustomerGateway,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("customer_gateway_id"),
			Hydrate:    getCustomerGateway,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the customer gateway.",
			},
			{
				Name:        "customer_gateway_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CustomerGatewayId"),
				Description: "The ID of the customer gateway.",
			},
			// Other columns
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the customer gateway.",
			},
			{
				Name:        "ip_address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address of the customer gateway.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp),
				Description: "The time when the customer gateway was created.",
			},
			{
				Name:        "asn",
				Type:        proto.ColumnType_STRING,
				Description: "The IPv4 CIDR block of the VPC.",
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcSustomerGatewayAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID"),
				Description: "The alicloud Account ID in which the resource is located.",
			},
		},
	}
}

//// LIST FUNCTION

func listCustomerGateway(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.listCustomerGateway", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeCustomerGatewaysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeCustomerGateways(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.listCustomerGateway", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.CustomerGateways.CustomerGateway {
			plugin.Logger(ctx).Warn("alicloud_vpc_vpn_customer_gateway.listCustomerGateway", "item", i)
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

//// GET FUNCTION

func getCustomerGateway(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.getCustomerGatewayAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeCustomerGatewayRequest()
	request.Scheme = "https"

	var id string
	if h.Item != nil {
		data := h.Item.(vpc.CustomerGateway)
		id = data.CustomerGatewayId
	} else {
		id = d.KeyColumnQuals["customer_gateway_id"].GetStringValue()
	}
	request.CustomerGatewayId = id

	response, err := client.DescribeCustomerGateway(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.getCustomerGatewayAttributes", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}
