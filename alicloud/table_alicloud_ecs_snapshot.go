package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type snapshotInfo = struct {
	Snapshot ecs.Snapshot
	Region   string
}

//// TABLE DEFINITION

func tableAlicloudEcsSnapshot(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_snapshot",
		Description: "Elastic Compute Service Snapshot.",
		List: &plugin.ListConfig{
			Hydrate: listEcsSnapshot,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getEcsSnapshot,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SnapshotName"),
			},
			{
				Name:        "id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SnapshotId"),
			},
			{
				Name:        "type",
				Description: "The type of the snapshot. Default value: all. Possible values are: auto, user, and all.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SnapshotType"),
			},
			{
				Name:        "serial_number",
				Description: "The serial number of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SnapshotSN"),
			},
			{
				Name:        "status",
				Description: "Specifies the current state of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.Status"),
			},
			{
				Name:        "creation_time",
				Description: "The time when the snapshot was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Snapshot.CreationTime"),
			},
			{
				Name:        "description",
				Description: "A user provided, human readable description for this resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.Description"),
			},
			{
				Name:        "encrypted",
				Description: "Indicates whether the snapshot was encrypted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Snapshot.Encrypted"),
			},
			{
				Name:        "instant_access",
				Description: "Indicates whether the instant access feature is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Snapshot.InstantAccess"),
			},
			{
				Name:        "instant_access_retention_days",
				Description: "Indicates the retention period of the instant access feature. After the retention per iod ends, the snapshot is automatically released.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Snapshot.InstantAccessRetentionDays"),
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the KMS key used by the data disk.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.KMSKeyId"),
			},
			{
				Name:        "last_modified_time",
				Description: "The time when the snapshot was last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Snapshot.LastModifiedTime"),
			},
			{
				Name:        "product_code",
				Description: "The product code of the Alibaba Cloud Marketplace image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.ProductCode"),
			},
			{
				Name:        "progress",
				Description: "The progress of the snapshot creation task. Unit: percent (%).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.Progress"),
			},
			{
				Name:        "remain_time",
				Description: "The remaining time required to create the snapshot (in seconds).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Snapshot.RemainTime"),
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the snapshot belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.ResourceGroupId"),
			},
			{
				Name:        "retention_days",
				Description: "The number of days that an automatic snapshot can be retained.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Snapshot.RetentionDays"),
			},
			{
				Name:        "source_disk_id",
				Description: "The ID of the source disk. This parameter is retained even after the source disk of the snapshot is released.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SourceDiskId"),
			},
			{
				Name:        "source_disk_size",
				Description: "The capacity of the source disk (in GiB).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SourceDiskSize"),
			},
			{
				Name:        "source_disk_type",
				Description: "The category of the source disk.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.SourceDiskType"),
			},
			{
				Name:        "usage",
				Description: "Indicates whether the snapshot has been used to create images or disks.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Snapshot.Usage"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Snapshot.Tags.Tag"),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(ecsSnapshotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsSnapshotAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Default:     transform.FromField("Snapshot.SnapshotName"),
				Transform:   transform.FromField("Snapshot.SnapshotId"),
			},

			// alibaba standard columns
			{
				Name:        "region_id",
				Description: "The region ID where the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Region"),
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

//// LIST FUNCTION

func listEcsSnapshot(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_snapshot.listEcsSnapshot", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeSnapshotsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	// Get the region details
	region, _, _, err := getEnv(ctx)
	if err != nil {
		return nil, err
	}

	count := 0
	for {
		response, err := client.DescribeSnapshots(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_snapshot.listEcsSnapshot", "query_error", err, "request", request)
			return nil, err
		}
		for _, snapshot := range response.Snapshots.Snapshot {
			plugin.Logger(ctx).Warn("listEcsSnapshot", "item", snapshot)
			d.StreamListItem(ctx, snapshotInfo{snapshot, region})
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

func getEcsSnapshot(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSnapshot")

	// Create service connection
	client, err := connectEcs(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_snapshot.getEcsSnapshot", "connection_error", err)
		return nil, err
	}

	// Get the region details
	region, _, _, err := getEnv(ctx)
	if err != nil {
		return nil, err
	}

	var name string
	if h.Item != nil {
		snapshot := h.Item.(ecs.Snapshot)
		name = snapshot.SnapshotName
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	request := ecs.CreateDescribeSnapshotsRequest()
	request.Scheme = "https"
	request.SnapshotName = name

	response, err := client.DescribeSnapshots(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_snapshot.getEcsSnapshot", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.Snapshots.Snapshot != nil && len(response.Snapshots.Snapshot) > 0 {
		return snapshotInfo{response.Snapshots.Snapshot[0], region}, nil
	}

	return nil, nil
}

func getEcsSnapshotAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSnapshotAka")
	data := h.Item.(snapshotInfo)

	// Get account details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:ecs:" + data.Region + ":" + accountID + ":snapshot/" + data.Snapshot.SnapshotId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func ecsSnapshotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(snapshotInfo)
	return ecsTagsToMap(data.Snapshot.Tags.Tag)
}
