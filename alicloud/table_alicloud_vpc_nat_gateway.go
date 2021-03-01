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
		GetMatrixItem: BuildRegionList,
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("nat_gateway_id"),
			Hydrate:    getNatGateway,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the NAT gateway.",
			},
			{
				Name:        "nat_gateway_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the NAT gateway.",
			},
			// Other columns
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the NAT gateway.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The payment state of the VPN gateway. Valid values: Normal and FinancialLocked.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the NAT gateway was created.",
			},
			{
				Name:        "expired_ime",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the NAT gateway expires.",
			},
			{
				Name:        "forward_table",
				Type:        proto.ColumnType_JSON,
				Description: "The ID of the DNAT table.",
				Hydrate:     getNatGateway,
				// Transform:   transform.FromField("ForwardTableIds.ForwardTableId"),
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the virtual private cloud (VPC) to which the NAT gateway belongs.",
			},
			{
				Name:        "nat_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the NAT gateway. Valid values: 'Normal' and 'Enhanced'.",
			},
			{
				Name:        "region_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the region where the NAT gateway is deployed.",
			},
			{
				Name:        "resource_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the resource group.",
			},
			{
				Name:        "snat_table",
				Type:        proto.ColumnType_JSON,
				Description: "The ID of the SNAT table that is associated with the NAT gateway.",
				Hydrate:     getNatGateway,
				// Transform:   transform.FromField("SnatTableIds.SnatTableId"),
			},
			{
				Name:        "business_status",
				Type:        proto.ColumnType_STRING,
				Description: "The state of the NAT gateway. Valid values: 'Normal' and 'FinancialLocked'",
			},
			{
				Name:        "deletion_protection",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether deletion protection is enabled. Valid values:",
				// Hydrate:     getNatGateway,
			},
			{
				Name:        "ecs_metric_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether the traffic monitoring feature is enabled",
			},
			{
				Name:        "billing_config",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getNatGateway,
				Description: "The billing type the NAT gateway.",
			},
			{
				Name:        "ip_list",
				Type:        proto.ColumnType_JSON,
				Description: "The elastic IP address (EIP) that is associated with the NAT gateway.",
				Hydrate:     getNatGateway,
				// Transform:   transform.FromField("IpLists.IpList"),
			},
			{
				Name:        "private_info",
				Type:        proto.ColumnType_JSON,
				Description: "Information about the private network to which the enhanced NAT gateway belongs.",
				Hydrate:     getNatGateway,
			},
			// Resource interface
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
				Hydrate:     getNatGateway,
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getNatGatewayAka,
				Transform:   transform.FromValue(),
			},
			// alicloud common columns
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

func listNatGateway(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
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

func getNatGateway(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.getNatGateway", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		data := h.Item.(vpc.NatGateway)
		id = data.NatGatewayId
	} else {
		id = d.KeyColumnQuals["nat_gateway_id"].GetStringValue()
	}

	request := vpc.CreateGetNatGatewayAttributeRequest()
	request.Scheme = "https"
	request.NatGatewayId = id

	response, err := client.GetNatGatewayAttribute(request)
	if err != nil {
		plugin.Logger(ctx).Error("getNatGateway", "query_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

func getNatGatewayAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getNatGatewayAka")
	ngw := h.Item.(vpc.NatGateway)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:vpc:" + ngw.RegionId + ":" + accountID + ":natgateway/" + ngw.NatGatewayId}

	return akas, nil
}
