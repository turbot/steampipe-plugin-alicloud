package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcNatGateway(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_nat_gateway",
		Description: "NAT gateways are enterprise-class Internet gateways that provide SNAT and DNAT functions for VPCs.",
		List: &plugin.ListConfig{
			Hydrate: listNatGateway,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getNatGatewayAttributes,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the VPN gateway."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("NatGatewayId"), Description: "The ID of the VPN gateway."},
			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the VPN gateway."},
			{Name: "vpc_id", Type: proto.ColumnType_STRING, Description: "The ID of the VPC for which the VPN gateway is created."},
			{Name: "vswitch_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("VSwitchId"), Description: "The ID of the VSwitch to which the VPN gateway belongs."},
			{Name: "spec", Type: proto.ColumnType_STRING, Description: "The IPv4 CIDR block of the VPC."},
			{Name: "instance_charge_type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "expired_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the VPN gateway expires."},
			{Name: "auto_pay", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "business_status", Type: proto.ColumnType_STRING, Description: "The payment state of the VPN gateway. Valid values: Normal and FinancialLocked."},
			{Name: "creation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the VPN gateway was created."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the VPN gateway."},
			{Name: "nat_type", Type: proto.ColumnType_STRING, Description: "Indicates whether the VPC is attached to any Cloud Enterprise Network (CEN) instance."},
			{Name: "internet_charge_type", Type: proto.ColumnType_STRING, Description: "The billing method of the VPN gateway."},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "deletion_protection", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "ecs_metric_enabled", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "forward_table", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "snat_table", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "private_info", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "bandwidth_package_ids", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "nat_gateway_private_info", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "ip_list_item", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listNatGateway(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpc", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeNatGatewaysRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	if quals["id"] != nil {
		request.NatGatewayId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeNatGateways(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_nat_gateway.listVpcNatGateway", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.NatGateways.NatGateway {
			plugin.Logger(ctx).Warn("alicloud_vpc_nat_gateway.listVpcNatGateway", "item", i)
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

func getNatGatewayAttributes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_nat_gateway.getNatGatewayAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateGetNatGatewayAttributeRequest()
	request.Scheme = "https"
	i := h.Item.(vpc.NatGateway)
	request.NatGatewayId = i.NatGatewayId
	response, err := client.GetNatGatewayAttribute(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_nat_gateway.getNatGatewayAttributes", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

