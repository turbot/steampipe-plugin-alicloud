package alicloud

import (
	"context"
	"strings"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudOssBucket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_oss_bucket",
		Description: "Object Storage Bucket",
		List: &plugin.ListConfig{
			Hydrate: listBucket,
			Tags:    map[string]string{"service": "oss", "action": "ListBuckets"},
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func: getBucketInfo,
				Tags: map[string]string{"service": "oss", "action": "GetBucketInfo"},
			},
			{
				Func: getBucketLifecycle,
				Tags: map[string]string{"service": "oss", "action": "GetBucketLifecycle"},
			},
			{
				Func: getBucketLogging,
				Tags: map[string]string{"service": "oss", "action": "GetBucketLogging"},
			},
			{
				Func: getBucketTagging,
				Tags: map[string]string{"service": "oss", "action": "GetBucketTags"},
			},
			{
				Func: getBucketPolicy,
				Tags: map[string]string{"service": "oss", "action": "GetBucketPolicy"},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the Bucket.",
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the OSS bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(bucketARN),
			},
			{
				Name:        "location",
				Type:        proto.ColumnType_STRING,
				Description: "Location of the Bucket.",
			},
			{
				Name:        "creation_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Date when the bucket was created.",
			},
			{
				Name:        "storage_class",
				Type:        proto.ColumnType_STRING,
				Description: "The storage class of objects in the bucket.",
			},
			{
				Name:        "redundancy_type",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBucketInfo,
				Transform:   transform.FromField("BucketInfo.RedundancyType"),
				Description: "The type of disaster recovery for a bucket. Valid values: LRS and ZRS",
			},
			{
				Name:        "versioning",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBucketInfo,
				Transform:   transform.FromField("BucketInfo.Versioning"),
				Description: "The status of versioning for the bucket. Valid values: Enabled and Suspended.",
			},
			{
				Name:        "acl",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBucketInfo,
				Transform:   transform.FromField("BucketInfo.ACL"),
				Description: "The access control list setting for bucket. Valid values: public-read-write, public-read, and private. public-read-write: Any users, including anonymous users can read and write objects in the bucket. Exercise caution when you set the ACL of a bucket to public-read-write. public-read: Only the owner or authorized users of this bucket can write objects in the bucket. Other users, including anonymous users can only read objects in the bucket. Exercise caution when you set the ACL of a bucket to public-read. private: Only the owner or authorized users of this bucket can read and write objects in the bucket. Other users, including anonymous users cannot access the objects in the bucket without authorization.",
			},
			{
				Name:        "server_side_encryption",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBucketInfo,
				Transform:   transform.FromField("BucketInfo.SseRule").Transform(bucketSSEConfiguration),
				Description: "The server-side encryption configuration for bucket",
			},
			{
				Name:        "lifecycle_rules",
				Type:        proto.ColumnType_JSON,
				Description: "A list of lifecycle rules for a bucket.",
				Hydrate:     getBucketLifecycle,
				Transform:   transform.FromField("Rules"),
			},
			{
				Name:        "logging",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBucketLogging,
				Transform:   transform.FromField("LoggingEnabled"),
				Description: "Indicates the container used to store access logging configuration of a bucket.",
			},
			{
				Name:        "policy",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBucketPolicy,
				Transform:   transform.FromValue().Transform(transform.UnmarshalYAML),
				Description: "Allows you to grant permissions on OSS resources to RAM users from your Alibaba Cloud and other Alibaba Cloud accounts. You can also control access based on the request source.",
			},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBucketTagging,
				Transform:   transform.FromField("Tags").Transform(ossBucketTagsSrc),
				Description: "A list of tags assigned to bucket",
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getBucketTagging,
				Transform:   transform.FromField("Tags").Transform(ossBucketTags),
				Description: ColumnDescriptionTags,
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(bucketARN).Transform(transform.EnsureStringArray),
				Description: ColumnDescriptionAkas,
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(bucketRegion),
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

func listBucket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := GetDefaultRegion(d.Connection)
	client, err := OssService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listBucket", "connection_error", err)
		return nil, err
	}

	param := &oss.ListBucketsRequest{
		MaxKeys: int32(1000),
	}

	page := client.NewListBucketsPaginator(param)

	for page.HasNext() {
		d.WaitForListRateLimit(ctx)
		p, err := page.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("listBucket", "paging_error", err)
			return nil, err
		}

		for _, i := range p.Buckets {
			d.StreamListItem(ctx, i)
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getBucketTagging(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	bucket := h.Item.(oss.BucketProperties)
	client, err := OssService(ctx, d, removeSuffixFromLocation(*bucket.Location))
	if err != nil {
		logger.Error("GetBucketTagging", "connection_error", err)
		return nil, err
	}

	param := &oss.GetBucketTagsRequest{
		Bucket: bucket.Name,
	}
	// Get bucket encryption
	response, err := client.GetBucketTags(ctx, param)
	if err != nil {
		logger.Error("GetBucketTagging", "query_error", err, "bucket", bucket.Name)
		return nil, err
	}
	return response, nil
}

func getBucketPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	bucket := h.Item.(oss.BucketProperties)
	client, err := OssService(ctx, d, removeSuffixFromLocation(*bucket.Location))
	if err != nil {
		logger.Error("GetBucketPolicy", "connection_error", err)
		return nil, err
	}
	param := &oss.GetBucketPolicyRequest{
		Bucket: bucket.Name,
	}
	// Get bucket encryption
	response, err := client.GetBucketPolicy(ctx, param)
	if err != nil {
		if a, ok := err.(*oss.ServiceError); ok {
			if a.Code == "NoSuchBucketPolicy" {
				logger.Debug("GetBucketPolicy", "query_error", a, "bucket", bucket.Name)
				return nil, nil
			}
			return nil, err
		}
	}
	return response, nil
}

func getBucketLogging(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	bucket := h.Item.(oss.BucketProperties)
	client, err := OssService(ctx, d, removeSuffixFromLocation(*bucket.Location))
	if err != nil {
		logger.Error("getBucketLogging", "connection_error", err)
		return nil, err
	}

	param := &oss.GetBucketLoggingRequest{
		Bucket: bucket.Name,
	}

	// Get bucket encryption
	response, err := client.GetBucketLogging(ctx, param)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Gives out the Bucket ACL and Encryption info
func getBucketInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	bucket := h.Item.(oss.BucketProperties)

	client, err := OssService(ctx, d, removeSuffixFromLocation(*bucket.Location))
	if err != nil {
		logger.Error("getBucketLogging", "connection_error", err)
		return nil, err
	}

	param := &oss.GetBucketInfoRequest{
		Bucket: bucket.Name,
	}
	// Get bucket encryption
	response, err := client.GetBucketInfo(ctx, param)
	if err != nil {
		logger.Error("getBucketInfo", "query_error", err, "bucket", bucket.Name)
		return nil, err
	}
	return response, nil
}

func getBucketLifecycle(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getBucketLifecycle")

	bucket := h.Item.(oss.BucketProperties)
	client, err := OssService(ctx, d, removeSuffixFromLocation(*bucket.Location))
	if err != nil {
		return nil, err
	}

	param := &oss.GetBucketLifecycleRequest{
		Bucket: bucket.Name,
	}

	// Get bucket encryption
	response, err := client.GetBucketLifecycle(ctx, param)
	if a, ok := err.(*oss.ServiceError); ok {
		if a.Code == "NoSuchLifecycle" {
			return nil, nil
		}
		return nil, err
	}
	return response, nil
}

//// TRANSFORM FUNCTIONS

func ossBucketTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}

	tags := d.Value.([]oss.Tag)
	var turbotTagsMap map[string]string

	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func ossBucketTagsSrc(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	tags := d.Value.([]oss.Tag)
	var turbotTagsMap []map[string]string

	for _, i := range tags {
		turbotTagsMap = append(turbotTagsMap, map[string]string{"Key": *i.Key, "Value": *i.Value})
	}

	return turbotTagsMap, nil
}

func bucketSSEConfiguration(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	sse := d.Value.(oss.SSERule)

	result := make(map[string]string, 0)

	if sse.KMSDataEncryption != nil {
		result["KMSDataEncryption"] = *sse.KMSDataEncryption
	}
	if sse.KMSMasterKeyID != nil {
		result["KMSMasterKeyID"] = *sse.KMSMasterKeyID
	}
	if sse.SSEAlgorithm != nil {
		result["SSEAlgorithm"] = *sse.SSEAlgorithm
	}

	return result, nil
}

func bucketARN(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("bucketARN")
	bucket := d.HydrateItem.(oss.BucketProperties)

	return "arn:acs:oss:::" + *bucket.Name, nil
}

func bucketRegion(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("bucketRegion")
	bucket := d.HydrateItem.(oss.BucketProperties)
	return strings.TrimPrefix(*bucket.Location, "oss-"), nil
}

func removeSuffixFromLocation(location string) string {
	return strings.TrimPrefix(location, "oss-")
}
