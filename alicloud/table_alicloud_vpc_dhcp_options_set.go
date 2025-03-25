package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

//// TABLE DEFINITION

func tableAlicloudVpcDhcpOptionsSet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_dhcp_options_set",
		Description: "Aliclod VPC DHCP Options Set",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("dhcp_options_set_id"),
			Hydrate:    getVpcDhcpOptionsSet,
			Tags:       map[string]string{"service": "vpc", "action": "DescribeDhcpOptionsSet"},
		},
		List: &plugin.ListConfig{
			Hydrate: listVpcDhcpOptionsSets,
			Tags:    map[string]string{"service": "vpc", "action": "DescribeDhcpOptionsSets"},
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "name", Require: plugin.Optional},
				{Name: "domain_name", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the DHCP option set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptionsSetName"),
			},
			{
				Name:        "dhcp_options_set_id",
				Description: "The ID of the DHCP option set.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the DHCP option set.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "associate_vpc_count",
				Description: "The number of VPCs associated with DHCP option set.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "boot_file_name",
				Description: "The boot file name of DHCP option set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptions.BootFileName"),
			},
			{
				Name:        "description",
				Description: "The description for the DHCP option set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptionsSetDescription"),
			},
			{
				Name:        "domain_name",
				Description: "The root domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptions.DomainName"),
			},
			{
				Name:        "domain_name_servers",
				Description: "The IP addresses of your DNS servers.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptions.DomainNameServers"),
			},
			{
				Name:        "owner_id",
				Description: "The ID of the account to which the DHCP options set belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tftp_server_name",
				Description: "The tftp server name of the DHCP option set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptions.TFTPServerName"),
			},
			{
				Name:        "associate_vpcs",
				Description: "The information of the VPC network that is associated with the DHCP options set.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcDhcpOptionsSet,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DhcpOptionsSetName"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcDhcpOptionSetAka,
				Transform:   transform.FromValue(),
			},

			// Alicloud common columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     vpcDhcpOptionsetRegion,
				Transform:   transform.FromValue(),
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

func listVpcDhcpOptionsSets(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_dhcp_options_set.listVpcDhcpOptionsSets", "connection_error", err)
		return nil, err
	}

	request := vpc.CreateListDhcpOptionsSetsRequest()
	request.Scheme = "https"
	request.MaxResults = requests.NewInteger(100)

	if d.EqualsQualString("name") != "" {
		request.DhcpOptionsSetName = d.EqualsQualString("name")
	}
	if d.EqualsQualString("domain_name") != "" {
		request.DomainName = d.EqualsQualString("domain_name")
	}

	// Limiting the results
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < 100 {
			if *limit < 1 {
				request.MaxResults = requests.NewInteger(1)
			} else {
				request.MaxResults = requests.NewInteger(int(*limit))
			}
		}
	}

	pageLeft := true
	for pageLeft {
		d.WaitForListRateLimit(ctx)
		response, err := client.ListDhcpOptionsSets(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_dhcp_options_set.listVpcDhcpOptionsSets", "query_error", err, "request", request)
			return nil, err
		}

		for _, dhcpOptionSet := range response.DhcpOptionsSets {
			d.StreamListItem(ctx, dhcpOptionSet)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if response.NextToken != "" {
			request.NextToken = response.NextToken
		} else {
			pageLeft = false
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getVpcDhcpOptionsSet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcDhcpOptionsSet")

	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_dhcp_options_set.getVpcDhcpOptionsSet", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		id = h.Item.(vpc.DhcpOptionsSet).DhcpOptionsSetId
	} else {
		id = d.EqualsQuals["dhcp_options_set_id"].GetStringValue()
	}

	request := vpc.CreateGetDhcpOptionsSetRequest()
	request.Scheme = "https"
	request.DhcpOptionsSetId = id

	response, err := client.GetDhcpOptionsSet(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_dhcp_options_set.getVpcDhcpOptionsSet", "query_error", err, "request", request)
		return nil, nil
	}
	if response.DhcpOptionsSetId != "" {
		return response, nil
	}

	return nil, nil
}

func getVpcDhcpOptionSetAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getVpcDhcpOptionSetAka")

	var id, region string
	region = d.EqualsQualString(matrixKeyRegion)

	switch item := h.Item.(type) {
	case *vpc.GetDhcpOptionsSetResponse:
		id = item.DhcpOptionsSetId
	case vpc.DhcpOptionsSet:
		id = item.DhcpOptionsSetId
	}

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + region + ":" + accountID + ":dhcpoptionset/" + id}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func vpcDhcpOptionsetRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	return region, nil
}
