package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

type routeTableRowData = struct {
	vpc.RouterTableListType
	Region string
}

//// TABLE DEFINITION

func tableAlicloudVpcRouteTable(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_route_table",
		Description: "Alicloud VPC Route Table",
		List: &plugin.ListConfig{
			Hydrate: listVpcRouteTable,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("route_table_id"),
			Hydrate:    getVpcRouteTable,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Route Table.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RouteTableName"),
			},
			{
				Name:        "route_table_id",
				Description: "The id of the Route Table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the Route Table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the Route Table was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "route_table_type",
				Description: "The type of Route Table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "router_id",
				Description: "The ID of the region to which the VPC belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "router_type",
				Description: "The type of the VRouter to which the route table belongs. Valid Values are 'VRouter' and 'VBR'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the route table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vswitch_ids",
				Description: "The unique ID of the VPC.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VSwitchIds.VSwitchId"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC to which the route table belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the VPC belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "route_entries",
				Description: "Route entry represents a route item of one VPC route table.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcRouteTableEntryList,
				Transform:   transform.FromField("RouteEntrys.RouteEntry"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the owner of the VPC.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag"),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(vpcTurbotTags),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcRouteTableAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(vpcRouteTableTitle),
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
				Transform:   transform.FromField("OwnerId"),
			},
		},
	}
}

//// LIST FUNCTION

func listVpcRouteTable(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
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
			d.StreamListItem(ctx, routeTableRowData{i, region})
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

func getVpcRouteTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcRouteTable")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.listVpcRouteTable", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["route_table_id"].GetStringValue()

	request := vpc.CreateDescribeRouteTableListRequest()
	request.Scheme = "https"
	request.RouteTableId = id

	response, err := client.DescribeRouteTableList(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.listVpcRouteTable", "query_error", err, "request", request)
		return nil, err
	}

	if response.RouterTableList.RouterTableListType != nil && len(response.RouterTableList.RouterTableListType) > 0 {
		return routeTableRowData{response.RouterTableList.RouterTableListType[0], region}, nil
	}

	return nil, nil
}

func getVpcRouteTableEntryList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcRouteTableEntryList")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	data := h.Item.(routeTableRowData)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.getVpcRouteTableEntryList", "connection_error", err)
		return nil, err
	}

	request := vpc.CreateDescribeRouteEntryListRequest()
	request.Scheme = "https"
	request.RouteTableId = data.RouteTableId

	response, err := client.DescribeRouteEntryList(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_table.getVpcRouteTableEntryList", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func getVpcRouteTableAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcRouteTableAka")
	data := h.Item.(routeTableRowData)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + data.Region + ":" + accountID + ":route-table/" + data.RouteTableId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func vpcRouteTableTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(routeTableRowData)

	// Build resource title
	title := data.RouteTableId

	if len(data.RouteTableName) > 0 {
		title = data.RouteTableName
	}

	return title, nil
}
