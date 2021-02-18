package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcEip(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc",
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listVpcEip,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("EipAddress"), Description: "The name of the VPC."},
			{Name: "allocation_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("AllocationId"), Description: "The unique ID of the VPC."},
			// Other columns
			{Name: "isp", Type: proto.ColumnType_STRING, Description: "The ID of the region to which the VPC belongs."},
			{Name: "region_id", Type: proto.ColumnType_STRING, Description: "The zone to which the VSwitch belongs."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available."},
			{Name: "filter_1_key", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "filter_1_value", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "filter_2_key", Type: proto.ColumnType_STRING, Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available."},
			{Name: "filter_2_value", Type: proto.ColumnType_STRING, Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available."},
			{Name: "IncludeReservationData", Type: proto.ColumnType_BOOL, Description: "True if the VPC is the default VPC in the region."},
			{Name: "lock_reason", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "associated_instance_id", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "associated_instance_type", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "segment_instance_id", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "owner_account", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "resource_owner_account", Type: proto.ColumnType_STRING, Description: "True if the VPC is the default VPC in the region."},
			{Name: "resource_owner_id", Type: proto.ColumnType_INT, Description: "The ID of the resource group to which the VPC belongs."},
			{Name: "dry_run", Type: proto.ColumnType_BOOL, Description: "True if the VPC is the default VPC in the region."},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the VPC belongs."},
			{Name: "chargetype", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the VPC."}},
	}
}

func listVpcEip(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.listEip", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeEipAddressesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	if quals["allocation_id"] != nil {
		request.AllocationId = quals["allocation_id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeEipAddresses(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_eip.listEip", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.EipAddresses.EipAddress {
			plugin.Logger(ctx).Warn("alicloud_eip.listEip", "tags", i.Tags, "item", i)
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

func getEip(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectVpc(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.getEipAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeEipAddressesRequest()
	request.Scheme = "https"
	i := h.Item.(vpc.EipAddress)
	request.AllocationId = i.AllocationId
	response, err := client.DescribeEipAddresses(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.getEip", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

// func vpcToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := d.Value.(vpc.Vpc)
// 	return "acs:vpc:" + i.RegionId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vpc/" + i.VpcName, nil
// }
