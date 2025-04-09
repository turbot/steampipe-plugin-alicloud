package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAlicloudVpcEip(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_eip",
		Description: "An independent public IP resource that decouples ECS and public IP resources, allowing you to flexibly manage public IP resources.",
		List: &plugin.ListConfig{
			Hydrate: listVpcEip,
			Tags:    map[string]string{"service": "vpc", "action": "DescribeEipAddresses"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("allocation_id"),
			Hydrate:    getEip,
			Tags:       map[string]string{"service": "vpc", "action": "DescribeEipAddresses"},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Description: "The name of the EIP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "allocation_id",
				Description: "The unique ID of the EIP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the EIP.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcEipArn,
				Transform:   transform.FromValue(),
			},

			// Other columns
			{
				Name:        "description",
				Description: "The description of the EIP.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Descritpion"),
			},
			{
				Name:        "ip_address",
				Description: "The IP address of the EIP.",
				Type:        proto.ColumnType_IPADDR,
			},
			{
				Name:        "expired_time",
				Description: "The expiration time of the EIP.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "status",
				Description: "The status of the EIP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance to which the EIP is bound.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_region_id",
				Description: "The region ID of the bound resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_type",
				Description: "The type of the instance to which the EIP is bound.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "internet_charge_type",
				Description: "The metering method of the EIP can be one of PayByBandwidth or PayByTraffic.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "isp",
				Description: "The Internet service provider.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ISP"),
			},
			{
				Name:        "allocation_time",
				Description: "The time when the EIP was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "bandwidth",
				Type:        proto.ColumnType_STRING,
				Description: "The peak bandwidth of the EIP. Unit: Mbit/s.",
			},
			{
				Name:        "bandwidth_package_bandwidth",
				Description: "The maximum bandwidth of the EIP in Mbit/s.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bandwidth_package_type",
				Description: "The bandwidth value of the EIP Bandwidth Plan to which the EIP is added.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "charge_type",
				Description: "The billing method of the EIP",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hd_monitor_status",
				Description: "Indicates whether fine-grained monitoring is enabled for the EIP.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HDMonitorStatus"),
			},
			{
				Name:        "has_reservation_data",
				Description: "Indicates whether renewal data is included.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "mode",
				Description: "The type of the instance to which you want to bind the EIP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "netmode",
				Description: "The network type of the EIP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name: "private_ip_address",
				Type: proto.ColumnType_BOOL,
			},
			{
				Name:        "second_limited",
				Description: "Indicates whether level-2 traffic throttling is configured.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "segment_instance_id",
				Description: "The ID of the instance with which the contiguous EIP is associated.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name: "service_managed",
				Type: proto.ColumnType_INT,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "available_regions",
				Description: "The ID of the region to which the EIP belongs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "operation_locks_reason",
				Description: "The reason why the EIP is locked. Valid values: financial security.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OperationLocks.LockReason"),
			},

			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcEipArn,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getVpcEipTitle),
			},

			// alibaba standard columns
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

func listVpcEip(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
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
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeEipAddresses(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_eip.listEip", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.EipAddresses.EipAddress {
			plugin.Logger(ctx).Warn("alicloud_eip.listEip", "tags", i.Tags, "item", i)
			d.StreamListItem(ctx, i)
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
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
	// Create service connection
	client, err := VpcService(ctx, d)
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
		id = d.EqualsQuals["allocation_id"].GetStringValue()
	}
	request.AllocationId = id
	response, err := client.DescribeEipAddresses(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_eip.getEip", "query_error", err, "request", request)
		return nil, err
	}

	if len(response.EipAddresses.EipAddress) > 0 {
		return response.EipAddresses.EipAddress[0], nil
	}
	return nil, nil
}

func getVpcEipArn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcEipArn")
	data := h.Item.(vpc.EipAddress)

	// Get account details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:vpc:" + data.RegionId + ":" + accountID + ":eip/" + data.AllocationId

	return arn, nil
}

//// TRANSFORM FUNCTION

func getVpcEipTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcEipTitle")
	eip := d.HydrateItem.(vpc.EipAddress)

	if eip.Name != "" {
		return eip.Name, nil
	}

	return eip.AllocationId, nil
}
