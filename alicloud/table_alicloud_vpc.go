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

func tableAlicloudVpc(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc",
		Description: "A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.",
		List: &plugin.ListConfig{
			Hydrate: listVpcs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("vpc_id"),
			Hydrate:    getVpc,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpcName"),
				Description: "The name of the VPC.",
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VpcId"),
				Description: "The unique ID of the VPC.",
			},

			// Other columns
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the VPC. Pending: The VPC is being configured. Available: The VPC is available.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the VPC.",
			},
			{
				Name:        "cidr_block",
				Type:        proto.ColumnType_CIDR,
				Description: "The IPv4 CIDR block of the VPC.",
			},
			{
				Name:        "ipv6_cidr_block",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Ipv6CidrBlock"),
				Description: "The IPv6 CIDR block of the VPC.",
			},
			{
				Name:        "vrouter_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VRouterId"),
				Description: "The ID of the VRouter.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the VPC.",
			},
			{
				Name:        "is_default",
				Type:        proto.ColumnType_BOOL,
				Description: "True if the VPC is the default VPC in the region.",
			},
			{
				Name:        "network_acl_num",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "resource_group_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the resource group to which the VPC belongs.",
			},
			{
				Name:        "cen_status",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the VPC is attached to any Cloud Enterprise Network (CEN) instance.",
			},
			{
				Name:        "owner_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the owner of the VPC.",
			},
			{
				Name:        "support_advanced_feature",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "advanced_resource",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},
			{
				Name:        "dhcp_options_set_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the DHCP options set associated to vpc.",
			},
			{
				Name:        "dhcp_options_set_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the VPC network that is associated with the DHCP options set. Valid values: InUse and Pending",
			},
			{
				Name:        "associated_cens",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcAttributes,
				Transform:   transform.FromField("AsssociatedCens"),
				Description: "The list of Cloud Enterprise Network (CEN) instances to which the VPC is attached. No value is returned if the VPC is not attached to any CEN instance.",
			},
			{
				Name:        "classic_link_enabled",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getVpcAttributes,
				Description: "True if the ClassicLink function is enabled.",
			},
			{
				Name:        "cloud_resources",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcAttributes,
				Transform:   transform.FromField("CloudResourceSetType"),
				Description: "The list of resources in the VPC.",
			},
			{
				Name:        "vswitch_ids",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("VSwitchIds.VSwitchId"),
				Description: "A list of VSwitches in the VPC.",
			},
			{
				Name:        "user_cidrs",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("UserCidrs.UserCidr"),
				Description: "A list of user CIDRs.",
			},
			{
				Name:        "nat_gateway_ids",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("NatGatewayIds.NatGatewayIds"),
				Description: "A list of IDs of NAT Gateways.",
			},
			{
				Name:        "route_table_ids",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("RouterTableIds.RouterTableIds"),
				Description: "A list of IDs of route tables.",
			},
			{
				Name:        "secondary_cidr_blocks",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SecondaryCidrBlocks.SecondaryCidrBlock"),
				Description: "A list of secondary IPv4 CIDR blocks of the VPC.",
			},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag"),
				Description: ColumnDescriptionTags,
			},

			// Resource interface
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(vpcTurbotTags),
				Description: ColumnDescriptionTags,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(vpcTitle),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(vpcAkas),
				Description: ColumnDescriptionAkas,
			},

			// alicloud common columns
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
				Transform:   transform.FromField("OwnerId"),
			},
		},
	}
}

func listVpcs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpc", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVpcsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	quals := d.KeyColumnQuals
	if quals["is_default"] != nil {
		request.IsDefault = requests.NewBoolean(quals["is_default"].GetBoolValue())
	}
	if quals["id"] != nil {
		request.VpcId = quals["id"].GetStringValue()
	}

	count := 0
	for {
		response, err := client.DescribeVpcs(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc.listVpc", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Vpcs.Vpc {
			plugin.Logger(ctx).Warn("alicloud_vpc.listVpc", "tags", i.Tags, "item", i)
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

//// Hydrate Functions

func getVpc(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcAttributes", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		vpc := h.Item.(vpc.Vpc)
		id = vpc.VpcId
	} else {
		id = d.KeyColumnQuals["vpc_id"].GetStringValue()
	}

	request := vpc.CreateDescribeVpcsRequest()
	request.Scheme = "https"
	request.VpcId = id
	response, err := client.DescribeVpcs(request)
	if err != nil {
		plugin.Logger(ctx).Error("getVpc", "query_error", err, "request", request)
		return nil, err
	}

	if response.Vpcs.Vpc != nil && len(response.Vpcs.Vpc) > 0 {
		return response.Vpcs.Vpc[0], nil
	}

	return nil, nil
}

func getVpcAttributes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcAttributes", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeVpcAttributeRequest()
	request.Scheme = "https"
	i := h.Item.(vpc.Vpc)
	request.VpcId = i.VpcId
	response, err := client.DescribeVpcAttribute(request)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcAttributes", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func vpcAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.HydrateItem.(vpc.Vpc)
	return []string{"acs:vpc:" + i.RegionId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":vpc/" + i.VpcId}, nil
}

func vpcTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.HydrateItem.(vpc.Vpc)

	// Build resource title
	title := i.VpcId
	if len(i.VpcName) > 0 {
		title = i.VpcName
	}

	return title, nil
}
