package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcVSwitch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vswitch",
		Description: "VSwitches to divide the VPC network into one or more subnets.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listVSwitch,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("VSwitchName"), Description: "The name of the VPC."},
			{Name: "vswitch_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("VSwitchId"), Description: "The unique ID of the VPC."},
			// Other columns
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available."},
			{Name: "cidr_block", Type: proto.ColumnType_CIDR, Description: "The IPv4 CIDR block of the VPC."},
			{Name: "ipv6_cidr_block", Type: proto.ColumnType_CIDR, Transform: transform.FromField("Ipv6CidrBlock"), Description: "The IPv6 CIDR block of the VPC."},
			{Name: "zone_id", Type: proto.ColumnType_STRING, Description: "The zone to which the VSwitch belongs."},
			{Name: "available_ip_address_count", Type: proto.ColumnType_INT, Description: "The number of available IP addresses in the VSwitch."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the VPC."},
			{Name: "creation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The creation time of the VPC."},
			{Name: "is_default", Type: proto.ColumnType_BOOL, Description: "True if the VPC is the default VPC in the region."},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the VPC belongs."},
			{Name: "network_acl_id", Type: proto.ColumnType_STRING, Description: "A list of IDs of NAT Gateways."},
			{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the VPC."},
			{Name: "share_type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "vpc_id", Type: proto.ColumnType_STRING, Description: "The ID of the VPC to which the VSwitch belongs."},
			{Name: "route_table", Type: proto.ColumnType_JSON, Description: "Details of the route table."},
			{Name: "vpc_id", Type: proto.ColumnType_STRING, Description: "The VPC ID to which the VSwitch belongs."},
			{Name: "cloud_resources", Type: proto.ColumnType_JSON, Hydrate: getVSwitchAttributes, Transform: transform.FromField("CloudResourceSetType"), Description: "The list of resources in the VSwitch."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(vswitchToURN).Transform(ensureStringArray), Description: ColumnDescriptionAkas},
			// TODO - It appears that Tags are not returned by the go SDK?
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: ColumnDescriptionTags},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("VSwitchName"), Description: ColumnDescriptionTitle},
			{Name: "account_id", Description: ColumnDescriptionAccount, Type: proto.ColumnType_STRING, Hydrate: getCommonColumns, Transform: transform.FromField("AccountID")},
		},
	}
}

func listVSwitch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vswitch.listVSwitch", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVSwitchesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	if quals["is_default"] != nil {
		request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	}
	if quals["id"] != nil {
		request.VSwitchId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeVSwitches(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vswitch.listVSwitch", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.VSwitches.VSwitch {
			plugin.Logger(ctx).Warn("alicloud_vswitch.listVSwitch", "tags", i.Tags, "item", i)
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

func getVSwitchAttributes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vswitch.getVSwitchAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVSwitchAttributesRequest()
	request.Scheme = "https"
	i := h.Item.(vpc.VSwitch)
	request.VSwitchId = i.VSwitchId
	response, err := client.DescribeVSwitchAttributes(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vswitch.getVSwitchAttributes", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func vswitchToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(vpc.VSwitch)
	return "acs:vswitch:" + i.ZoneId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vswitch/" + i.VSwitchId, nil
}
