package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

//// TABLE DEFINITION

func tableAlicloudVpcBGPPeer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_bgp_peer",
		Description: "Aliclod VPC BGP Peer",
		List: &plugin.ListConfig{
			Hydrate: listVpcBgpPeers,
			ShouldIgnoreError: isNotFoundError([]string{"InvalidRegionId.NotFound"}),
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "bgp_group_id", Require: plugin.Optional},
				{Name: "bgp_peer_id", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the BGP peer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bgp_peer_id",
				Description: "The ID of the BGP peer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bgp_group_id",
				Description: "The ID of the BGP group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The IP address of the BGP peer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "peer_ip_address",
				Description: "The boot file name of DHCP option set.",
				Type:        proto.ColumnType_IPADDR,
			},
			{
				Name:        "peer_asn",
				Description: "The autonomous system (AS) number of the BGP peer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auth_key",
				Description: "The authentication key of the BGP group.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "router_id",
				Description: "The ID of the virtual border router (VBR) that is associated with the BGP peer that you want to query.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bgp_status",
				Description: "The status of the BGP connection. Possible values are: Idle | Connect |Active | Established | Down.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the BGP peer. Valid values: Pending | Available | Modifying | Deleting | Deleted.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "keepalive",
				Description: "The Keepalive interval.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "local_asn",
				Description: "The AS number of the device on the Alibaba Cloud side.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "hold",
				Description: "The hold time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "is_fake",
				Description: "Indicates whether a fake ASN is used.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "route_limit",
				Description: "The limit on routes.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region_id",
				Description: "The region ID of the BGP group to which the BGP peer that you want to query belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_bfd",
				Description: "Indicates whether BFD is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ip_version",
				Description: "The version of the IP address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "bfd_multi_hop",
				Description: "The Bidirectional Forwarding Detection (BFD) hop count.",
				Type:        proto.ColumnType_INT,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			// {
			// 	Name:        "akas",
			// 	Description: ColumnDescriptionAkas,
			// 	Type:        proto.ColumnType_JSON,
			// 	Hydrate:     getVpcDhcpOptionSetAka,
			// 	Transform:   transform.FromValue(),
			// },

			// Alicloud common columns
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

func listVpcBgpPeers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_bgp_peer.listVpcBgpPeers", "connection_error", err)
		return nil, err
	}

	request := vpc.CreateDescribeBgpPeersRequest()
	request.Scheme = "https"

	if d.KeyColumnQualString("bgp_group_id") != "" {
		request.BgpGroupId = d.KeyColumnQualString("bgp_group_id")
	}
	if d.KeyColumnQualString("bgp_peer_id") != "" {
		request.BgpPeerId = d.KeyColumnQualString("bgp_peer_id")
	}

	response, err := client.DescribeBgpPeers(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_bgp_peer.listVpcBgpPeers", "api_error", err, "request", request)
		return nil, err
	}

	for _, peer := range response.BgpPeers.BgpPeer {
		d.StreamListItem(ctx, peer)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
