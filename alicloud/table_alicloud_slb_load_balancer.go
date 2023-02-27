package alicloud

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/sethvargo/go-retry"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudSlbLoadBalancer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_slb_load_balancer",
		Description: "Alicloud Server Load Balancer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("load_balancer_id"),
			Hydrate:    getSlbLoadBalancer,
		},
		List: &plugin.ListConfig{
			Hydrate: listSlbLoadBalancers,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "load_balancer_name", Require: plugin.Optional},
				{Name: "network_type", Require: plugin.Optional},
				{Name: "resource_group_id", Require: plugin.Optional},
				{Name: "master_zone_id", Require: plugin.Optional},
				{Name: "address_ip_version", Require: plugin.Optional},
				{Name: "v_switch_id", Require: plugin.Optional},
				{Name: "vpc_id", Require: plugin.Optional},
				{Name: "load_balancer_status", Require: plugin.Optional},
				{Name: "address_type", Require: plugin.Optional},
				{Name: "internet_charge_type", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "load_balancer_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the CLB instance.",
			},
			{
				Name:        "load_balancer_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the CLB instance.",
			},
			{
				Name:        "load_balancer_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the CLB instance. Valid values: inactive|active|locked.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_IPADDR,
				Description: "The service IP address of the CLB instance.",
			},
			{
				Name:        "address_type",
				Type:        proto.ColumnType_STRING,
				Description: "The network type of the CLB instance. Valid values: internet|intranet.",
			},
			{
				Name:        "v_switch_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the vSwitch to which the CLB instance belongs.",
				Transform:   transform.FromField("VSwitchId"),
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the virtual private cloud (VPC) to which the CLB instance belongs.",
			},
			{
				Name:        "network_type",
				Type:        proto.ColumnType_STRING,
				Description: "The network type of the internal-facing CLB instance. Valid values: vpc|classic.",
			},
			{
				Name:        "master_zone_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the primary zone to which the CLB instance belongs.",
			},
			{
				Name:        "slave_zone_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the secondary zone to which the CLB instance belongs.",
			},
			{
				Name:        "internet_charge_type",
				Type:        proto.ColumnType_STRING,
				Description: "The metering method of Internet data transfer. Valid values: paybybandwidth|paybytraffic.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the CLB instance was created.",
			},
			{
				Name:        "create_time_stamp",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The timestamp when the instance was created.",
				Transform:   transform.FromField("CreateTimeStamp").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "pay_type",
				Type:        proto.ColumnType_STRING,
				Description: "The billing method of the CLB instance. Valid values: PayOnDemand.",
			},
			{
				Name:        "resource_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the resource group.",
			},
			{
				Name:        "address_ip_version",
				Type:        proto.ColumnType_STRING,
				Description: "The IP version. Valid values: ipv4 and ipv6.",
				Transform:   transform.FromField("AddressIPVersion"),
			},
			{
				Name:        "modification_protection_status",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the configuration read-only mode is enabled for the CLB instance. ",
			},
			{
				Name:        "modification_protection_reason",
				Type:        proto.ColumnType_STRING,
				Description: "The reason why the configuration read-only mode is enabled.",
			},
			{
				Name:        "bandwidth",
				Type:        proto.ColumnType_INT,
				Description: "The maximum bandwidth of the listener. Unit: Mbit/s.",
			},
			{
				Name:        "internet_charge_type_alias",
				Type:        proto.ColumnType_STRING,
				Description: "The alias for metering method of Internet data transfer.",
			},
			{
				Name:        "load_balancer_spec",
				Type:        proto.ColumnType_STRING,
				Description: "The specification of the CLB instance.",
			},
			{
				Name:        "delete_protection",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether deletion protection is enabled for the CLB instance.",
			},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Description: "A list of tags.",
				Transform:   transform.FromField("Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("LoadBalancerName"),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTags,
				Transform:   transform.From(slbLoadbalancerTagMap),
			},

			// Alicloud standard columns
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

//// LIST FUNCTION

func listSlbLoadBalancers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := SLBService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_slb_load_balancer.listSlbLoadBalancers", "connection_error", err)
		return nil, err
	}

	request := slb.CreateDescribeLoadBalancersRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	if d.KeyColumnQualString("load_balancer_name") != "" {
		request.LoadBalancerName = d.KeyColumnQualString("load_balancer_name")
	}
	if d.KeyColumnQualString("network_type") != "" {
		request.NetworkType = d.KeyColumnQualString("network_type")
	}
	if d.KeyColumnQualString("resource_group_id") != "" {
		request.ResourceGroupId = d.KeyColumnQualString("resource_group_id")
	}
	if d.KeyColumnQualString("master_zone_id") != "" {
		request.MasterZoneId = d.KeyColumnQualString("master_zone_id")
	}
	if d.KeyColumnQualString("address_ip_version") != "" {
		request.AddressIPVersion = d.KeyColumnQualString("address_ip_version")
	}
	if d.KeyColumnQualString("v_switch_id") != "" {
		request.VSwitchId = d.KeyColumnQualString("v_switch_id")
	}
	if d.KeyColumnQualString("vpc_id") != "" {
		request.VpcId = d.KeyColumnQualString("vpc_id")
	}
	if d.KeyColumnQualString("load_balancer_status") != "" {
		request.LoadBalancerStatus = d.KeyColumnQualString("load_balancer_status")
	}
	if d.KeyColumnQualString("address_type") != "" {
		request.AddressType = d.KeyColumnQualString("address_type")
	}
	if d.KeyColumnQualString("internet_charge_type") != "" {
		request.InternetChargeType = d.KeyColumnQualString("internet_charge_type")
	}

	count := 0
	for {
		response, err := client.DescribeLoadBalancers(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_slb_load_balancer.listSlbLoadBalancers", "api_error", err, "request", request)
			return nil, err
		}
		for _, loadBalancer := range response.LoadBalancers.LoadBalancer {
			d.StreamListItem(ctx, loadBalancer)
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

func getSlbLoadBalancer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Create service connection
	client, err := SLBService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_slb_load_balancer.getSlbLoadBalancer", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["load_balancer_id"].GetStringValue()

	// Empty check
	if id == "" {
		return nil, nil
	}

	request := slb.CreateDescribeLoadBalancersRequest()
	request.Scheme = "https"
	request.LoadBalancerId = id
	var response *slb.DescribeLoadBalancersResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(5, b), func(ctx context.Context) error {
		var err error
		response, err = client.DescribeLoadBalancers(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling" {
					return retry.RetryableError(err)
				}
				return err
			}
			plugin.Logger(ctx).Error("alicloud_slb_load_balancer.getSlbLoadBalancer", "api_error", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if response.LoadBalancers.LoadBalancer != nil && len(response.LoadBalancers.LoadBalancer) > 0 {
		if response.LoadBalancers.LoadBalancer[0].RegionId == region {
			return response.LoadBalancers.LoadBalancer[0], nil
		}
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func slbLoadbalancerTagMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(slb.LoadBalancer).Tags
	if tags.Tag == nil {
		return nil, nil
	}

	if len(tags.Tag) == 0 {
		return nil, nil
	}
	turbotTagsMap := map[string]string{}
	for _, i := range tags.Tag {
		turbotTagsMap[i.TagKey] = i.TagValue
	}

	return turbotTagsMap, nil
}
