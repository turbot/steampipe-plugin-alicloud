package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

//// TABLE DEFINITION

func tableAlicloudVpcRouteEntry(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_route_entry",
		Description: "Alicloud VPC Route Entry",
		List: &plugin.ListConfig{
			ParentHydrate: listVpcRouteTable,
			Hydrate:       listVpcRouteEntries,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the route entry.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RouteEntryName"),
			},
			{
				Name:        "route_entry_id",
				Description: "The ID of the route entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "route_table_id",
				Description: "The ID of the route table.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the route entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of the route entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the VRouter.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance associated with the next hop.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "private_ip_address",
				Description: "Specifies the private ip address for the route entry.",
				Type:        proto.ColumnType_IPADDR,
			},
			{
				Name:        "next_hop_oppsite_instance_id",
				Description: "The ID of the instance associated with the next hop.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "next_hop_type",
				Description: "The type of the next hop.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ip_version",
				Description: "The version of the IP protocol.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "destination_cidr_block",
				Description: "The destination Classless Inter-Domain Routing (CIDR) block of the route entry.",
				Type:        proto.ColumnType_CIDR,
			},
			{
				Name:        "next_hop_region_id",
				Description: "The region where the next hop instance is deployed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "next_hop_oppsite_type",
				Description: "The type of the next hop.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "next_hop_oppsite_region_id",
				Description: "The region where the next hop instance is deployed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "next_hops",
				Description: "The information about the next hop.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("NextHops.NextHop"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcRouteEntryTurbotData,
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcRouteEntryTurbotData,
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcRouteEntryTurbotData,
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

func listVpcRouteEntries(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_entry.listVpcRouteEntries", "connection_error", err)
		return nil, err
	}

	// Get VPC Route Table details
	routeTable := h.Item.(vpc.RouterTableListType)

	request := vpc.CreateDescribeRouteEntryListRequest()
	request.Scheme = "https"
	request.RouteTableId = routeTable.RouteTableId

	response, err := client.DescribeRouteEntryList(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_route_entry.listVpcRouteEntries", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.RouteEntrys.RouteEntry {
		d.StreamLeafListItem(ctx, i)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getVpcRouteEntryTurbotData(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcRouteEntryTurbotData")
	region := d.EqualsQualString(matrixKeyRegion)
	data := h.Item.(vpc.RouteEntry)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	var title string
	var akas []string
	if len(data.RouteEntryId) > 0 {
		akas = []string{"acs:vpc:" + region + ":" + accountID + ":route-entry/" + data.RouteEntryId}
		title = data.RouteEntryName
	} else {
		akas = []string{"acs:vpc:" + region + ":" + accountID + ":route-entry/" + data.RouteTableId}
		if len(data.NextHops.NextHop[0].NextHopId) > 0 {
			title = data.RouteTableId + ":" + data.DestinationCidrBlock + ":" + data.NextHops.NextHop[0].NextHopType + ":" + data.NextHops.NextHop[0].NextHopId
		} else {
			title = data.RouteTableId + ":" + data.DestinationCidrBlock + ":" + data.NextHops.NextHop[0].NextHopType
		}
	}

	// Mapping all turbot defined properties
	turbotData := map[string]interface{}{
		"Akas":   akas,
		"Title":  title,
		"Region": region,
	}

	return turbotData, nil
}
