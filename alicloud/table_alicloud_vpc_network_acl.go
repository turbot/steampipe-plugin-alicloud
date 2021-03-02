package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudVpcNetworkAcl(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_network_acl",
		Description: "VPC network ACL.",
		List: &plugin.ListConfig{
			Hydrate: listNetworkAcl,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("network_acl_id"),
			Hydrate:    getNetworkAclAttribute,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkAclName"),
				Description: "The name of the network ACL.",
			},
			{
				Name:        "network_acl_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkAclId"),
				Description: "The ID of the network ACL.",
			},
			// Other columns
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the network ACL.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the network ACL.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the network ACL was created.",
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Ipv6CidrBlock"),
				Description: "The ID of the associated VPC.",
			},
			{
				Name:        "region_id",
				Type:        proto.ColumnType_STRING,
				Description: "The region of the network ACL.",
			},
			{
				Name:        "ingress_acl_entries",
				Type:        proto.ColumnType_JSON,
				Description: "The inbound rule information.",
				Transform:   transform.FromField("IngressAclEntries.IngressAclEntry"),
			},
			{
				Name:        "egress_acl_entries",
				Type:        proto.ColumnType_JSON,
				Description: "The outbound rule information.",
				Transform:   transform.FromField("EgressAclEntries.EgressAclEntry"),
			},
			{
				Name:        "resources",
				Type:        proto.ColumnType_JSON,
				Description: "The associated resources.",
				Transform:   transform.FromField("Resources.Resource"),
			},
			{
				Name:        "owner_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the owner of the VPC.",
			},
			// steampipe standard columns
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag"),
				Description: ColumnDescriptionTags,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkAclName"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getNetworkAclAka,
				Transform:   transform.FromValue(),
			},
			// alicloud standard columns
			{
				Name:        "region",
				Description: "The name of the region where the resource resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
			},
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

func listNetworkAcl(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_network_acl.listNetworkAcl", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeNetworkAclsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	count := 0
	for {
		response, err := client.DescribeNetworkAcls(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_network_acl.listNetworkAcl", "query_error", err, "request", request)
			return nil, err
		}
		for _, NetworkAcl := range response.NetworkAcls.NetworkAcl {
			plugin.Logger(ctx).Warn("alicloud_vpc_network_acl.listNetworkAcl", "tags", NetworkAcl, "item", NetworkAcl)
			d.StreamListItem(ctx, NetworkAcl)
			count++
		}
		totalcount, err := strconv.Atoi(response.TotalCount)
		pageNumber, err := strconv.Atoi(response.PageNumber)
		if count >= totalcount {
			break
		}
		request.PageNumber = requests.NewInteger(pageNumber + 1)
	}
	return nil, nil
}

func getNetworkAclAttribute(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_network_acl.getNetworkAclAttribute", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		networkAcl := h.Item.(vpc.NetworkAcl)
		id = networkAcl.NetworkAclId
	} else {
		id = d.KeyColumnQuals["network_acl_id"].GetStringValue()
	}
	request := vpc.CreateDescribeNetworkAclAttributesRequest()
	request.Scheme = "https"
	request.NetworkAclId = id

	response, err := client.DescribeNetworkAclAttributes(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		if serverErr.ErrorCode() == "InvalidNetworkAcl.NotFound" {
			plugin.Logger(ctx).Warn("alicloud_vpc_network_acl.getNetworkAclAttribute", "not_found_error", serverErr, "request", request)
			return nil, nil
		}
		plugin.Logger(ctx).Error("alicloud_vpc_network_acl.getNetworkAclAttribute", "query_error", serverErr, "request", request)
		return nil, serverErr
	}
	return response.NetworkAclAttribute, nil
}

func getNetworkAclAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsDiskAka")
	networkAcl := h.Item.(vpc.NetworkAcl)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:vpc:" + networkAcl.RegionId + ":" + accountID + ":networkAcl/" + networkAcl.NetworkAclId}

	return akas, nil
}
