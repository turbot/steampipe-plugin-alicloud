package alicloud

import (
	"context"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableAlicloudOssBucket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_bucket",
		Description: "",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.AnyColumn([]string{"is_default", "id"}),
			Hydrate: listBucket,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the Bucket."},
			{Name: "xml_name", Type: proto.ColumnType_JSON, Transform: transform.FromField("XMLName"), Description: "XML name of the Bucket."},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "Location of the Bucket."},
			{Name: "creation_date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the bucket was created."},
			{Name: "storage_class", Type: proto.ColumnType_STRING, Description: "The storage class of objects in the bucket."},
			// {Name: "versioning_status", Type: proto.ColumnType_STRING, Hydrate: getBucketVersioning, Transform: transform.FromField("Status"), Description: ""},
			/*
				{Name: "bucket_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("BucketId"), Description: "The unique ID of the Bucket."},
				// Other columns
				{Name: "status", Type: proto.ColumnType_STRING, Description: "The status of the Bucket. Pending: The Bucket is being configured. Available: The Bucket is available."},
				{Name: "cidr_block", Type: proto.ColumnType_CIDR, Description: "The IPv4 CIDR block of the Bucket."},
				{Name: "ipv6_cidr_block", Type: proto.ColumnType_CIDR, Transform: transform.FromField("Ipv6CidrBlock"), Description: "The IPv6 CIDR block of the Bucket."},
				{Name: "zone_id", Type: proto.ColumnType_STRING, Description: "The zone to which the Bucket belongs."},
				{Name: "available_ip_address_count", Type: proto.ColumnType_INT, Description: "The number of available IP addresses in the Bucket."},
				{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the Bucket."},
				{Name: "creation_time", Type: proto.ColumnType_TIMESTAMP, Description: "The creation time of the Bucket."},
				{Name: "is_default", Type: proto.ColumnType_BOOL, Description: "True if the Bucket is the default Bucket in the region."},
				{Name: "resource_group_id", Type: proto.ColumnType_STRING, Description: "The ID of the resource group to which the Bucket belongs."},
				{Name: "network_acl_id", Type: proto.ColumnType_STRING, Description: "A list of IDs of NAT Gateways."},
				{Name: "owner_id", Type: proto.ColumnType_STRING, Description: "The ID of the owner of the Bucket."},
				{Name: "share_type", Type: proto.ColumnType_STRING, Description: ""},
				{Name: "route_table", Type: proto.ColumnType_JSON, Description: "Details of the route table."},
				{Name: "cloud_resources", Type: proto.ColumnType_JSON, Hydrate: getBucketAttributes, Transform: transform.FromField("CloudResourceSetType"), Description: "The list of resources in the Bucket."},
				// Resource interface
				{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(bucketToURN).Transform(ensureStringArray), Description: ColumnDescriptionAkas},
				// TODO - It appears that Tags are not returned by the go SDK?
				{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags.Tag"), Description: ColumnDescriptionTags},
			*/
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: ColumnDescriptionTitle},
		},
	}
}

func listBucket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connectOss(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_bucket.listBucket", "connection_error", err)
		return nil, err
	}
	pre := oss.Prefix("")
	marker := oss.Marker("")
	for {
		response, err := client.ListBuckets(oss.MaxKeys(50), pre, marker)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_oss_bucket.listBucket", "query_error", err, "marker", marker)
			return nil, err
		}
		for _, i := range response.Buckets {
			plugin.Logger(ctx).Warn("listBucket", "item", i)
			d.StreamListItem(ctx, i)
		}
		if !response.IsTruncated {
			break
		}
		pre = oss.Prefix(response.Prefix)
		marker = oss.Marker(response.NextMarker)
	}
	return nil, nil
}

func getBucketVersioning(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectOss(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_bucket.getBucketVersioning", "connection_error", err)
		return nil, err
	}
	bucket := h.Item.(oss.BucketProperties)
	// Get bucket encryption
	response, err := client.GetBucketVersioning(bucket.Name)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_bucket.getBucketVersioning", "query_error", err, "bucket", bucket)
		return nil, err
	}
	return response, nil
}

/*
func getBucketAttributes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connectOss(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_bucket.getBucketAttributes", "connection_error", err)
		return nil, err
	}
	request := oss.CreateDescribeBucketAttributesRequest()
	request.Scheme = "https"
	i := h.Item.(oss.Bucket)
	request.BucketId = i.BucketId
	response, err := client.DescribeBucketAttributes(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_bucket.getBucketAttributes", "query_error", err, "request", request)
		return nil, err
	}
	return response, nil
}

func bucketToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(oss.Bucket)
	return "acs:bucket:" + i.ZoneId + ":" + strconv.FormatInt(i.OwnerId, 10) + ":bucket/" + i.BucketId, nil
}
*/
