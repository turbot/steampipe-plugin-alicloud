package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/privatelink"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudVpcEndpoint(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_endpoint",
		Description: "PrivateLink-Endpoint allows you to access resources in a VPC from another VPC over secure and private networks. PrivateLink simplifies network architecture and avoids security risks from the Internet.",
		List: &plugin.ListConfig{
			Hydrate: listVpcEndpoints,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("endpoint_id"),
			Hydrate:    getVpcEndpoint,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EndpointName"),
				Description: "The name of the Endpoint.",
			},
			{
				Name:        "endpoint_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EndpointId"),
				Description: "The unique ID of the Endpoint.",
			},

			// Other columns
			{
				Name:        "endpoint_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the endpoint. Valid values:Creating: The endpoint is being created.Active: The endpoint is available.Pending: The endpoint is being modified.Deleting: The endpoint is being deleted.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the EndPoint.",
			},
			{
				Name:        "endpoint_description",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "endpoint_business_status",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "endpoint_domain",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "bandwidth",
				Type:        proto.ColumnType_DOUBLE,
				Description: "",
			},
			{
				Name:        "connection_status",
				Type:        proto.ColumnType_STRING,
				Description: "The Status of the Endpoint Connection",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the endpoint service that is associated with the endpoint to be queried.",
			},
			{
				Name:        "service_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the endpoint service that is associated with the endpoint to be queried.",
			},
			{
				Name:        "vpc_id",
				Type:        proto.ColumnType_STRING,
				Description: "The virtual private cloud (VPC) to which the endpoint belongs.",
			},
			{
				Name:        "region_id",
				Type:        proto.ColumnType_STRING,
				Description: "",
			},
			{
				Name:        "resource_owner",
				Type:        proto.ColumnType_BOOL,
				Description: "",
			},

			// {
			// 	Name:        "tags_src",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   transform.FromField("Tags.Tag"),
			// 	Description: ColumnDescriptionTags,
			// },

			// Resource interface
			// {
			// 	Name:        "tags",
			// 	Type:        proto.ColumnType_JSON,
			// 	Transform:   transform.FromField("Tags.Tag").Transform(vpcTurbotTags),
			// 	Description: ColumnDescriptionTags,
			// },
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(endpointTitle),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEndpointAkas,
				Transform:   transform.FromValue(),
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
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID"),
			},
		},
	}
}

//// LIST FUNCTION

func listVpcEndpoints(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcEndpointService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.listVpcEndpoints", "connection_error", err)
		return nil, err
	}
	request := privatelink.CreateListVpcEndpointsRequest()
	request.Scheme = "https"

	response, err := client.ListVpcEndpoints(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc.ListVpcEndpoints", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.Endpoints {
		plugin.Logger(ctx).Warn("alicloud_vpc.ListVpcEndpoints", "item", i)
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getVpcEndpoint(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcEndpointService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcEndPoint", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		privatelink := h.Item.(privatelink.Endpoint)
		id = privatelink.EndpointId
	} else {
		id = d.KeyColumnQuals["endpoint_id"].GetStringValue()
	}

	request := privatelink.CreateListVpcEndpointsRequest()
	request.Scheme = "https"
	request.EndpointId = id
	response, err := client.ListVpcEndpoints(request)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcEndPoint", "query_error", err, "request", request)
		return nil, err
	}

	if response.Endpoints != nil && len(response.Endpoints) > 0 {
		return response.Endpoints[0], nil
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func getEndpointAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	i := h.Item.(privatelink.Endpoint)
	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID
	return []string{"acs:endpoint:" + i.RegionId + ":" + accountID + ":endpoint/" + i.EndpointId}, nil
}

func endpointTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.HydrateItem.(privatelink.Endpoint)

	// Build resource title
	title := i.EndpointId
	if len(i.EndpointName) > 0 {
		title = i.EndpointName
	}

	return title, nil
}
