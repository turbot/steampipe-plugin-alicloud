package alicloud

import (
	"context"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsDisk(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_disk",
		Description: "Elastic Compute Disk",
		List: &plugin.ListConfig{
			Hydrate: listEcsDisk,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("disk_id"),
			Hydrate:    getEcsDisk,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DiskName"),
			},
			{
				Name:        "disk_id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the ECS disk.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsDiskARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "status",
				Description: "Specifies the current state of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "size",
				Description: "Specifies the size of the disk.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "type",
				Description: "Specifies the type of the disk. Possible values are: 'system' and 'data'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "billing_method",
				Description: "The billing method of the disk. Possible values are: PrePaid and PostPaid.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DiskChargeType"),
			},
			{
				Name:        "attached_time",
				Description: "The time when the disk was attached.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "auto_snapshot_policy_id",
				Description: "The ID of the automatic snapshot policy applied to the disk.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auto_snapshot_policy_name",
				Description: "The name of the automatic snapshot policy applied to the disk.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
			},
			{
				Name:        "auto_snapshot_policy_creation_time",
				Description: "The time when the auto snapshot policy was created.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("CreationTime"),
			},
			{
				Name:        "auto_snapshot_policy_enable_cross_region_copy",
				Description: "The ID of the automatic snapshot policy applied to the disk.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("EnableCrossRegionCopy"),
			},
			{
				Name:        "auto_snapshot_policy_repeat_week_days",
				Description: "The days of a week on which automatic snapshots are created. Valid values: 1 to 7, which corresponds to the days of the week. 1 indicates Monday. One or more days can be specified.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("RepeatWeekdays"),
			},
			{
				Name:        "auto_snapshot_policy_retention_days",
				Description: "The retention period of the automatic snapshot.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("RetentionDays"),
			},
			{
				Name:        "auto_snapshot_policy_status",
				Description: "The status of the automatic snapshot policy.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "auto_snapshot_policy_time_points",
				Description: "The points in time at which automatic snapshots are created. The least interval at which snapshots can be created is one hour. Valid values: 0 to 23, which corresponds to the hours of the day from 00:00 to 23:00. 1 indicates 01:00. You can specify multiple points in time.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("TimePoints"),
			},
			{
				Name:        "auto_snapshot_policy_tags",
				Description: "The days of a week on which automatic snapshots are created. Valid values: 1 to 7, which corresponds to the days of the week. 1 indicates Monday. One or more days can be specified.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsDiskAutoSnapshotPolicy,
				Transform:   transform.FromField("Tags.Tag"),
			},
			{
				Name:        "category",
				Description: "The category of the disk.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the disk was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "delete_auto_snapshot",
				Description: "Indicates whether the automatic snapshots of the disk are deleted when the disk is released.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "delete_with_instance",
				Description: "Indicates whether the disk is released when its associated instance is released.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "description",
				Description: "A user provided, human readable description for this resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "detached_time",
				Description: "The time when the disk was detached.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "device",
				Description: "The device name of the disk on its associated instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "enable_auto_snapshot",
				Description: "Indicates whether the automatic snapshot policy feature was enabled for the disk.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "enable_automated_snapshot_policy",
				Description: "Indicates whether an automatic snapshot policy was applied to the disk.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "encrypted",
				Description: "Indicates whether the disk was encrypted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "expired_time",
				Description: "The time when the subscription disk expires.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "iops",
				Description: "The number of input/output operations per second (IOPS).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("IOPS"),
			},
			{
				Name:        "iops_read",
				Description: "The number of I/O reads per second.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("IOPSRead"),
			},
			{
				Name:        "iops_write",
				Description: "The number of I/O writes per second.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("IOPSWrite"),
			},
			{
				Name:        "image_id",
				Description: "The ID of the image used to create the instance. This parameter is empty unless the disk was created from an image. The value of this parameter remains unchanged throughout the lifecycle of the disk.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance to which the disk is attached. This parameter has a value only when the value of Status is In_use.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kms_key_id",
				Description: "The device name of the disk on its associated instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("KMSKeyId"),
			},
			{
				Name:        "mount_instance_num",
				Description: "The number of instances to which the Shared Block Storage device is attached.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "performance_level",
				Description: "The performance level of the ESSD.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "portable",
				Description: "Indicates whether the disk is removable.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "product_code",
				Description: "The product code in Alibaba Cloud Marketplace.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the disk belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "serial_number",
				Description: "The serial number of the disk.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "source_snapshot_id",
				Description: "The ID of the snapshot used to create the disk. This parameter is empty unless the disk was created from a snapshot. The value of this parameter remains unchanged throughout the lifecycle of the disk.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_set_id",
				Description: "The ID of the storage set.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "storage_set_partition_number",
				Description: "The maximum number of partitions in a storage set.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "zone",
				Description: "The zone name in which the resource is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ZoneId"),
			},
			{
				Name:        "attachments",
				Description: "The attachment information of the cloud disk.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Attachments.Attachment"),
			},
			{
				Name:        "mount_instances",
				Description: "The attaching information of the disk.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("MountInstances.MountInstance"),
			},
			{
				Name:        "operation_lock",
				Description: "The reasons why the disk was locked.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("OperationLocks.OperationLock"),
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
				Hydrate:     getEcsDiskARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ecsDiskTitle),
			},

			// Alicloud standard columns
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

func listEcsDisk(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_disk.listEcsDisk", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeDisksRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeDisks(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_disk.listEcsDisk", "query_error", err, "request", request)
			return nil, err
		}
		for _, disk := range response.Disks.Disk {
			plugin.Logger(ctx).Warn("listEcsDisk", "item", disk)
			d.StreamListItem(ctx, disk)
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

func getEcsDisk(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsDisk")

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_disk.getEcsDisk", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		disk := h.Item.(ecs.Disk)
		id = disk.DiskId
	} else {
		id = d.EqualsQuals["disk_id"].GetStringValue()
	}

	// In SDK, the Datatype of DiskIds is string, though the value should be passed as
	// ["d-bp67acfmxazb4p****", "d-bp67acfmxazb4g****", ... "d-bp67acfmxazb4d****"]
	input, err := json.Marshal([]string{id})
	if err != nil {
		return nil, err
	}

	request := ecs.CreateDescribeDisksRequest()
	request.Scheme = "https"
	request.DiskIds = string(input)

	response, err := client.DescribeDisks(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_disk.getEcsDisk", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.Disks.Disk) > 0 {
		return response.Disks.Disk[0], nil
	}

	return nil, nil
}

func getEcsDiskAutoSnapshotPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsDiskAutomaticSnapshotPolicy")
	disk := h.Item.(ecs.Disk)

	// Create service connection
	client, err := ECSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_disk.getEcsDisk", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeAutoSnapshotPolicyExRequest()
	request.Scheme = "https"
	request.AutoSnapshotPolicyId = disk.AutoSnapshotPolicyId

	response, err := client.DescribeAutoSnapshotPolicyEx(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ecs_disk.getEcsDiskAutoSnapshotPolicy", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.AutoSnapshotPolicies.AutoSnapshotPolicy) > 0 {
		return response.AutoSnapshotPolicies.AutoSnapshotPolicy[0], nil
	}

	return nil, nil
}

func getEcsDiskARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsDiskARN")
	disk := h.Item.(ecs.Disk)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	arn := "arn:acs:ecs:" + disk.RegionId + ":" + accountID + ":disk/" + disk.DiskId

	return arn, nil
}

//// TRANSFORM FUNCTIONS

func ecsDiskTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	disk := d.HydrateItem.(ecs.Disk)

	// Build resource title
	title := disk.DiskId

	if len(disk.DiskName) > 0 {
		title = disk.DiskName
	}

	return title, nil
}
