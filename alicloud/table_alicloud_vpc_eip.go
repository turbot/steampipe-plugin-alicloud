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
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("allocation_id"),
			Hydrate:    getEip,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the EIP.",
			},
			{
				Name:        "allocation_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique ID of the EIP.",
			},
			// Other columns
			{
				Name:        "descritpion",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the EIP.",
			},
			{
				Name:        "ip_address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address of the EIP.",
			},
			{
				Name:        "expired_time",
				Type:        proto.ColumnType_STRING,
				Description: "The expiration time of the EIP.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the EIP.",
			},
			{
				Name:        "instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the instance to which the EIP is bound.",
			},
			{
				Name:        "instance_region_id",
				Type:        proto.ColumnType_STRING,
				Description: "The region ID of the bound resource.",
			},
			{
				Name:        "instance_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the instance to which the EIP is bound.",
			},
			{
				Name:        "internet_charge_type",
				Type:        proto.ColumnType_STRING,
				Description: "The metering method of the EIP.",
			},
			{
				Name:        "isp",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ISP"),
				Description: "The Internet service provider.",
			},
			{
				Name:        "allocation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the EIP was created.",
			},
			{
				Name:        "available_regions",
				Type:        proto.ColumnType_JSON,
				Description: "The ID of the region to which the EIP belongs.",
			},
			{
				Name:        "bandwidth",
				Type:        proto.ColumnType_STRING,
				Description: "The peak bandwidth of the EIP. Unit: Mbit/s.",
			},
			{
				Name:        "bandwidth_package_bandwidth",
				Type:        proto.ColumnType_STRING,
				Description: "The maximum bandwidth of the EIP in Mbit/s.",
			},
			{
				Name:        "bandwidth_package_type",
				Type:        proto.ColumnType_STRING,
				Description: "The bandwidth value of the EIP Bandwidth Plan to which the EIP is added.",
			},
			{
				Name:        "business_status",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "charge_type",
				Type:        proto.ColumnType_STRING,
				Description: "The billing method of the EIP",
			},
			{
				Name:        "hd_monitor_status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HDMonitorStatus"),
				Description: "Indicates whether fine-grained monitoring is enabled for the EIP.",
			},
			{
				Name:        "has_reservation_data",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether renewal data is included.",
			},
			{
				Name:        "mode",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "netmode",
				Type:        proto.ColumnType_STRING,
				Description: "The network type of the EIP.",
			},
			{
				Name:        "operation_locks",
				Type:        proto.ColumnType_JSON,
				Description: "The details about the lock.",
			},
			{
				Name:        "private_ip_address",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "second_limited",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether level-2 traffic throttling is configured.",
			},
			{
				Name:        "segment_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the instance with which the contiguous EIP is associated.",
			},
			{
				Name:        "service_managed",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ServiceManaged"),
				Description: "The ID of the instance to which the contiguous EIP is bound.",
			},
			{
				Name:        "resource_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the resource group.",
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "region_id",
				Description: "The name of the region where the resource belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
			},
			// alicloud standard columns
			{
				Name:        "account_id",
				Description: "The alicloud Account ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID"),
			},
		},
	}
}

func listVpcEip(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.listEip", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeEipAddressesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

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
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.getEipAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeEipAddressesRequest()
	request.Scheme = "https"

	var id string
	if h.Item != nil {
		data := h.Item.(vpc.EipAddress)
		id = data.AllocationId
	} else {
		id = d.KeyColumnQuals["allocation_id"].GetStringValue()
	}
	request.AllocationId = id
	response, err := client.DescribeEipAddresses(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.getEip", "query_error", err, "request", request)
		return nil, err
	}

	if response.EipAddresses.EipAddress != nil && len(response.EipAddresses.EipAddress) > 0 {
		return response.EipAddresses.EipAddress[0], nil
	}
	return nil, nil
}
