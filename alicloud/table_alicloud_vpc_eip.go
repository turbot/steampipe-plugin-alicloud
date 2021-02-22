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
		Name:        "alicloud_vpc_eip",
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listVpcEip,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the EIP."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("AllocationId"), Description: "The unique ID of the EIP."},
			// Other columns
			{Name: "descritpion", Type: proto.ColumnType_STRING, Description: "The description of the EIP."},
			{Name: "ip_address", Type: proto.ColumnType_STRING, Description: "The IP address of the EIP."},
			{Name: "expired_time", Type: proto.ColumnType_STRING, Description: "The expiry time of the EIP."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the EIP."},
			{Name: "instance_id", Type: proto.ColumnType_STRING, Description: "The ID of the instance associated to the EIP."},
			{Name: "instance_region_id", Type: proto.ColumnType_STRING, Description: "The region ID of the instance that attached to the EIP."},
			{Name: "instance_type", Type: proto.ColumnType_STRING, Description: "The type of the instance to which you want to bind the EIP."},
			{Name: "internet_charge_type", Type: proto.ColumnType_STRING, Description: "The billing method of the EIP."},
			{Name: "isp", Type: proto.ColumnType_STRING, Transform: transform.FromField("ISP"), Description: "The type of connection."},
			{Name: "region_id", Type: proto.ColumnType_STRING, Description: "The ID of the region where the EIPs are created."},
			{Name: "allocation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The time when EIP was created"},
			{Name: "available_regions", Type: proto.ColumnType_JSON, Description: "The regions where EIP is available."},
			{Name: "bandwidth", Type: proto.ColumnType_STRING, Description: "The data transfer rate of EIP."},
			{Name: "bandwidth_package_bandwidth", Type: proto.ColumnType_STRING, Description: "The maximum bandwidth of the EIP in Mbit/s."},
			{Name: "bandwidth_package_type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "business_status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "charge_type", Type: proto.ColumnType_STRING, Description: "The billing method of the EIP"},
			{Name: "hd_monitor_status", Type: proto.ColumnType_STRING, Transform: transform.FromField("HDMonitorStatus"), Description: ""},
			{Name: "has_reservation_data", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "mode", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "netmode", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "operation_locks", Type: proto.ColumnType_JSON, Description: ""},
			{Name: "private_ip_address", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "second_limited", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "segment_instance_id", Type: proto.ColumnType_STRING, Description: "The ID of the instance with which the contiguous EIP is associated."},
			{Name: "service_managed", Type: proto.ColumnType_INT, Transform: transform.FromField("ServiceManaged"), Description: ""},
			{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the EIP belongs."},
			// TODO - It appears that Tags are not returned by the go SDK?
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")}},
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

// func getEip(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	client, err := connectVpc(ctx)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("alicloud_eip.getEipAttributes", "connection_error", err)
// 		return nil, err
// 	}
// 	request := vpc.CreateDescribeEipAddressesRequest()
// 	request.Scheme = "https"
// 	i := h.Item.(vpc.EipAddress)
// 	request.AllocationId = i.AllocationId
// 	response, err := client.DescribeEipAddresses(request)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("alicloud_eip.getEip", "query_error", err, "request", request)
// 		return nil, err
// 	}
// 	return response, nil
// }

// func vpcToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
// 	i := d.Value.(vpc.Vpc)
// 	return "acs:vpc:" + i.RegionId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vpc/" + i.VpcName, nil
// }
