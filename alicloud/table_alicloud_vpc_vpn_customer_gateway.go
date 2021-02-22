package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcVpnCustomerGateway(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_nat_gateway",
		Description: "NAT gateways are enterprise-class Internet gateways that provide SNAT and DNAT functions for VPCs.",
		List: &plugin.ListConfig{
			Hydrate: listCustomerGateway,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCustomerGatewayAttributes,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the customer gateway."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("CustomerGatewayId"), Description: "The ID of the customer gateway."},
			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the customer gateway."},
			{Name: "ip_address", Type: proto.ColumnType_STRING, Description: "The IP address of the customer gateway."},
			{Name: "create_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp), Description: "The time when the customer gateway was created."},
			{Name: "asn", Type: proto.ColumnType_STRING, Description: "The IPv4 CIDR block of the VPC."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listCustomerGateway(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.listCustomerGateway", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeCustomerGatewaysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	if quals["id"] != nil {
		request.CustomerGatewayId = quals["id"].GetStringValue()
	}

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

func getCustomerGatewayAttributes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.getCustomerGatewayAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeCustomerGatewayRequest()
	request.Scheme = "https"
	i := h.Item.(vpc.CustomerGateway)
	request.CustomerGatewayId = i.CustomerGatewayId
	response, err := client.DescribeCustomerGateway(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_vpn_customer_gateway.getCustomerGatewayAttributes", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}
