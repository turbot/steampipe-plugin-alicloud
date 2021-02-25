package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcRouteTable(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_route_table",
		Description: "A route table contains a set of rules, called routes, that are used to determine where network traffic from your subnet or gateway is directed.",
		List: &plugin.ListConfig{
			Hydrate: listVpcRouteTable,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVpcRouteTable,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("RouteTableName"), Description: "The name of the Route Table."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("RouteTableId"), Description: "The id of the Route Table."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the Route Table."},
			{Name: "creation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when the Route Table was created.."},
			{Name: "route_table_type", Type: proto.ColumnType_STRING, Description: "The type of Route Table"},
			{Name: "router_id", Type: proto.ColumnType_STRING, Description: "The ID of the region to which the VPC belongs."},
			{Name: "router_type", Type: proto.ColumnType_STRING, Description: "The type of the VRouter to which the route table belongs. Valid Values are 'VRouter' and 'VBR'"},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the route table."},
			{Name: "vswitch_ids", Type: proto.ColumnType_JSON, Transform: transform.FromField("VSwitchIds.VSwitchId"), Description: "The unique ID of the VPC."},
			{Name: "vpc_id", Type: proto.ColumnType_STRING, Description: "The ID of the VPC to which the route table belongs."},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the VPC belongs."},
			{Name: "route_entries", Type: proto.ColumnType_JSON, Hydrate: getVpcRouteTableEntryList, Transform: transform.FromField("RouteEntrys.RouteEntry"), Description: "Route entry represents a route item of one VPC route table."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the VPC."},
			// Other columns
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: ColumnDescriptionTitle},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("RouteTableName"), Description: ColumnDescriptionAkas},
		},
	}
}

//// BUILD HYDRATE INPUT

func RouteTableIDFromRouteTable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item := &vpc.RouteTable{
		RouteTableId: id,
	}
	return item, nil
}

func listVpcRouteTable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.listVpcRouteTable", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeRouteTableListRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeRouteTableList(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_route_table.listVpcRouteTable", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.RouterTableList.RouterTableListType {
			plugin.Logger(ctx).Warn("alicloud_vpc_route_table.listVpcRouteTable", "tags", i.Tags, "item", i)
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

func getVpcRouteTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.listVpcRouteTable", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		vpc := h.Item.(vpc.RouteTable)
		id = vpc.RouteTableId
	} else {
		id = d.KeyColumnQuals["id"].GetStringValue()
	}

	request := vpc.CreateDescribeRouteTableListRequest()
	request.Scheme = "https"
	request.RouteTableId = id

	response, err := client.DescribeRouteTableList(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.listVpcRouteTable", "query_error", err, "request", request)
		return nil, err
	}

	if response.RouterTableList.RouterTableListType != nil && len(response.RouterTableList.RouterTableListType) > 0 {
		return response.RouterTableList.RouterTableListType[0], nil
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getVpcRouteTableEntryList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.getVpcRouteTableEntryList", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeRouteEntryListRequest()
	request.Scheme = "https"
	i := h.Item.(vpc.RouterTableListType)
	request.RouteTableId = i.RouteTableId
	response, err := client.DescribeRouteEntryList(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.getVpcRouteTableEntryList", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

// func vpcToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := d.Value.(vpc.RouteTables)
// 	return "acs:vpc:"  + ":"  + ":routetable/" + i.RouteTableName, nil
// }
