package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsSnapshot(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_snapshot",
		Description: "ECS Disk Snapshot.",
		List: &plugin.ListConfig{
			Hydrate: listEcsSnapshot,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getEcsSnapshot,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SnapshotName"),
			},
			{
				Name:        "snapshot_id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsSnapshotArn,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "type",
				Description: "The type of the snapshot. Default value: all. Possible values are: auto, user, and all.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SnapshotType"),
			},
			{
				Name:        "serial_number",
				Description: "The serial number of the snapshot.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SnapshotSN"),
			},
			{
				Name:        "status",
				Description: "Specifies the current state of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the snapshot was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "A user provided, human readable description for this resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "encrypted",
				Description: "Indicates whether the snapshot was encrypted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "instant_access",
				Description: "Indicates whether the instant access feature is enabled.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "instant_access_retention_days",
				Description: "Indicates the retention period of the instant access feature. After the retention per iod ends, the snapshot is automatically released.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "kms_key_id",
				Description: "The ID of the KMS key used by the data disk.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KMSKeyId"),
			},
			{
				Name:        "last_modified_time",
				Description: "The time when the snapshot was last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "product_code",
				Description: "The product code of the Alibaba Cloud Marketplace image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "progress",
				Description: "The progress of the snapshot creation task. Unit: percent (%).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remain_time",
				Description: "The remaining time required to create the snapshot (in seconds).",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the snapshot belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "retention_days",
				Description: "The number of days that an automatic snapshot can be retained.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "source_disk_id",
				Description: "The ID of the source disk. This parameter is retained even after the source disk of the snapshot is released.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_disk_size",
				Description: "The capacity of the source disk (in GiB).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_disk_type",
				Description: "The category of the source disk.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "usage",
				Description: "Indicates whether the snapshot has been used to create images or disks.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(ecsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsSnapshotArn,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Default:     transform.FromField("SnapshotName"),
				Transform:   transform.FromField("SnapshotId"),
			},

			// Alibaba standard columns
			{
				Name:        "region",
				Description: "The region ID where the resource is located.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getSnapshotRegion,
				Transform:   transform.FromValue(),
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
	client, err := ECSService(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_snapshot.listEcsSnapshot", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeSnapshotsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeSnapshots(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_snapshot.listEcsSnapshot", "query_error", err, "request", request)
			return nil, err
		}
		for _, snapshot := range response.Snapshots.Snapshot {
			plugin.Logger(ctx).Warn("listEcsSnapshot", "item", snapshot)
			d.StreamListItem(ctx, snapshot)
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
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_snapshot.getEcsSnapshot", "connection_error", err)
		return nil, err
	}

	var name string
	if h.Item != nil {
		snapshot := h.Item.(ecs.Snapshot)
		name = snapshot.SnapshotName
	} else {
		name = d.EqualsQuals["name"].GetStringValue()
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
		return response.Snapshots.Snapshot[0], nil
	}

	return nil, nil
}

func getEcsSnapshotArn(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsSnapshotArn")
	data := h.Item.(ecs.Snapshot)
	region := d.EqualsQualString(matrixKeyRegion)

	// Get account details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + region + ":" + accountID + ":snapshot/" + data.SnapshotId

	return arn, nil
}

func getSnapshotRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	return region, nil
}
