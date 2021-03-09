package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

//// TABLE DEFINITION

func tableAlicloudVpcNetworkACL(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_network_acl",
		Description: "Alicloud VPC Network ACL.",
		List: &plugin.ListConfig{
			Hydrate: listNetworkACLs,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("network_acl_id"),
			Hydrate:           getNetworkACLAttribute,
			ShouldIgnoreError: isNotFoundError([]string{"InvalidNetworkAcl.NotFound", "MissingParameter"}),
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the network ACL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkAclName"),
			},
			{
				Name:        "network_acl_id",
				Description: "The ID of the network ACL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the network ACL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC associated with the network ACL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the network ACL was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description of the network ACL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner_id",
				Description: "The ID of the owner of the resource.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "region_id",
				Description: "The name of the region where the resource resides.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ingress_acl_entries",
				Description: "A list of inbound rules of the network ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("IngressAclEntries.IngressAclEntry"),
			},
			{
				Name:        "egress_acl_entries",
				Description: "A list of outbound rules of the network ACL.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("EgressAclEntries.EgressAclEntry"),
			},
			{
				Name:        "resources",
				Description: "A list of associated resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Resources.Resource"),
			},

			// steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(vpcNetworkACLTitle),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getNetworkACLAka,
				Transform:   transform.FromValue(),
			},

			// alicloud standard columns
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

func listNetworkACLs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_network_acl.listNetworkACLs", "connection_error", err)
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
			plugin.Logger(ctx).Error("alicloud_vpc_network_acl.listNetworkACLs", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.NetworkAcls.NetworkAcl {
			d.StreamListItem(ctx, vpc.NetworkAclAttribute{
				NetworkAclName: i.NetworkAclName,
				RegionId:       i.RegionId,
				NetworkAclId:   i.NetworkAclId,
				CreationTime:   i.CreationTime,
				VpcId:          i.VpcId,
				Status:         i.Status,
				OwnerId:        i.OwnerId,
				Description:    i.Description,
				IngressAclEntries: vpc.IngressAclEntriesInCreateNetworkAcl{
					IngressAclEntry: i.IngressAclEntries.IngressAclEntry,
				},
				Resources: vpc.ResourcesInCreateNetworkAcl{
					Resource: i.Resources.Resource,
				},
				EgressAclEntries: vpc.EgressAclEntriesInCreateNetworkAcl{
					EgressAclEntry: i.EgressAclEntries.EgressAclEntry,
				},
			})
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

//// HYDRATE FUNCTIONS

func getNetworkACLAttribute(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getNetworkACLAttribute")
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_network_acl.getNetworkACLAttribute", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["network_acl_id"].GetStringValue()

	request := vpc.CreateDescribeNetworkAclAttributesRequest()
	request.Scheme = "https"
	request.NetworkAclId = id

	response, err := client.DescribeNetworkAclAttributes(request)
	if err != nil {
		return nil, err
	}
	return response.NetworkAclAttribute, nil
}

func getNetworkACLAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getNetworkACLAka")
	data := h.Item.(vpc.NetworkAclAttribute)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + data.RegionId + ":" + accountID + ":network-acl/" + data.NetworkAclId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func vpcNetworkACLTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(vpc.NetworkAclAttribute)

	// Build resource title
	title := data.NetworkAclId

	if len(data.NetworkAclName) > 0 {
		title = data.NetworkAclName
	}

	return title, nil
}
