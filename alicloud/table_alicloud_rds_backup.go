package alicloud

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsBackup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_backup",
		Description: "ApsaraDB RDS Backup is a policy expression that defines when and how you want to back up your DB Instances.",
		List: &plugin.ListConfig{
			ParentHydrate: listRdsInstances,
			Hydrate:       listRdsBackups,
			Tags:          map[string]string{"service": "rds", "action": "DescribeBackups"},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "backup_id", Require: plugin.Optional},
				{Name: "db_instance_id", Require: plugin.Optional},
				{Name: "backup_status", Require: plugin.Optional},
				{Name: "backup_mode", Require: plugin.Optional},
				{Name: "backup_start_time", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "backup_end_time", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "backup_location", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "backup_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the backup set.",
			},
			{
				Name:        "db_instance_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DBInstanceId"),
				Description: "The ID of the single instance to query.",
			},
			{
				Name:        "backup_status",
				Description: "The status of the backup. Valid values: Success, Failed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "backup_mode",
				Type:        proto.ColumnType_STRING,
				Description: "The backup mode.",
			},
			{
				Name:        "backup_db_names",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the database that has been backed up.",
			},
			{
				Name:        "backup_size",
				Type:        proto.ColumnType_INT,
				Description: "The size of the backup set. Unit: bytes.",
			},
			{
				Name:        "backup_start_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The beginning of the backup time range. The time is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.",
			},
			{
				Name:        "backup_end_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The end of the backup time range. The time is in the yyyy-MM-ddTHH:mm:ssZ format and displayed in UTC.",
			},
			{
				Name:        "backup_download_url",
				Type:        proto.ColumnType_STRING,
				Description: "The internet download URL of the backup set. If the download URL is unavailable, this parameter is an empty string.",
				Transform:   transform.FromField("BackupDownloadURL"),
			},
			{
				Name:        "backup_extraction_status",
				Type:        proto.ColumnType_STRING,
				Description: "The backup extraction status.",
			},
			{
				Name:        "backup_intranet_download_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BackupIntranetDownloadURL"),
				Description: "The internal download URL of the backup set.",
			},
			{
				Name:        "backup_method",
				Type:        proto.ColumnType_STRING,
				Description: "The backup method. Valid values: Snapshot, Physical and Logical.",
			},
			{
				Name:        "backup_type",
				Type:        proto.ColumnType_STRING,
				Description: "The backup type.",
			},
			{
				Name:        "backup_initiator",
				Type:        proto.ColumnType_STRING,
				Description: "The initiator of the backup task.",
			},
			{
				Name:        "backup_location",
				Type:        proto.ColumnType_STRING,
				Description: "The location where backup is available.",
			},
			{
				Name:        "consistent_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The point in time at which the data in the data backup is consistent. ",
				Transform:   transform.FromField("ConsistentTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "copy_only_backup",
				Type:        proto.ColumnType_STRING,
				Description: "The backup mode of the data backup. Valid values: 0, 1.",
			},
			{
				Name:        "encryption",
				Type:        proto.ColumnType_STRING,
				Description: "The encryption information of the data backup.",
			},
			{
				Name:        "host_instance_id",
				Type:        proto.ColumnType_STRING,
				Description: "The number of the instance that generates the data backup. This parameter is used to indicate whether the instance that generates the data backup file is a primary instance or a secondary instance.",
				Transform:   transform.FromField("HostInstanceID"),
			},
			{
				Name:        "is_avail",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the data backup is available. Valid values: 0, 1.",
			},
			{
				Name:        "meta_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the data backup file that is used to restore individual databases or tables.",
			},
			{
				Name:        "slave_status",
				Type:        proto.ColumnType_STRING,
				Description: "The slave status of the backup.",
			},
			{
				Name:        "storage_class",
				Type:        proto.ColumnType_STRING,
				Description: "The storage class of the data backup. Valid values: 0(regular storage), 1 (archive storage).",
			},
			{
				Name:        "store_status",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates whether the data backup can be deleted. Valid values: Enabled, Disabled.",
			},
			{
				Name:        "total_backup_size",
				Type:        proto.ColumnType_STRING,
				Description: "The total backup size.",
			},

			// Steampipe standard column
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("BackupId"),
				Description: ColumnDescriptionTitle,
			},

			// Alicloud common columns
			// The API response doesn't contain region id, so removed the column region.
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

func listRdsBackups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	dbInstance := h.Item.(rds.DBInstance)

	if d.EqualsQualString("db_instance_id") != "" {
		if d.EqualsQualString("db_instance_id") != dbInstance.DBInstanceId {
			return nil, nil
		}
	}

	// Create service connection
	client, err := RDSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_rds_backup.listRdsBackups", "connection_error", err)
		return nil, err
	}
	request := rds.CreateDescribeBackupsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	request.DBInstanceId = dbInstance.DBInstanceId

	// Optional Qulas
	if d.EqualsQualString("backup_id") != "" {
		request.BackupId = d.EqualsQualString("backup_id")
	}
	if d.EqualsQualString("backup_status") != "" {
		request.BackupStatus = d.EqualsQualString("backup_status")
	}
	if d.EqualsQualString("backup_mode") != "" {
		request.BackupMode = d.EqualsQualString("backup_mode")
	}
	if d.EqualsQualString("backup_location") != "" {
		request.BackupLocation = d.EqualsQualString("backup_location")
	}
	quals := d.Quals
	if quals["backup_start_time"] != nil {
		for _, q := range quals["backup_start_time"].Quals {
			startTime := q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
			r := regexp.MustCompile(`:[0-9]+Z`)
			if q.Operator == "=" {
				request.StartTime = fmt.Sprint(r.ReplaceAllString(startTime, "Z"))
			}
		}
	}
	if quals["backup_end_time"] != nil {
		for _, q := range quals["backup_end_time"].Quals {
			endTime := q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
			r := regexp.MustCompile(`:[0-9]+Z`)
			if q.Operator == "=" {
				request.EndTime = fmt.Sprint(r.ReplaceAllString(endTime, "Z"))
			}
		}
	}

	count := 0
	for {
		d.WaitForListRateLimit(ctx)
		response, err := client.DescribeBackups(request)
		if err != nil {
			// Not found eror code could not be captured in ignore config so need to handle it here.
			if serverErr, ok := err.(*errors.ServerError); ok {
				if slices.Contains([]string{"InvalidBackupId.NotFound"}, serverErr.ErrorCode()) {
					return nil, nil
				}
			}
			plugin.Logger(ctx).Error("alicloud_rds_backup.listRdsBackups", "query_error", err, "request", request)
			return nil, err
		}
		if len(response.Items.Backup) == 0 {
			break
		}
		for _, i := range response.Items.Backup {
			d.StreamListItem(ctx, i)
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
			count++
		}
		totalRecord, _ := strconv.Atoi(response.TotalRecordCount)
		pageNumber, _ := strconv.Atoi(response.PageNumber)
		if count >= totalRecord {
			break
		}
		request.PageNumber = requests.NewInteger(pageNumber + 1)
	}
	return nil, nil
}
