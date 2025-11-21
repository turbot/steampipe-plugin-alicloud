package alicloud

import (
	"context"
	"fmt"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudLogStore(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_log_store",
		Description: "Alicloud Log Service (SLS) Logstore.",
		List: &plugin.ListConfig{
			ParentHydrate: listLogProjects,
			Hydrate:       listLogstores,
			Tags:          map[string]string{"service": "sls", "action": "ListLogStoreV2"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project", "name"}),
			Hydrate:    getLogstore,
			Tags:       map[string]string{"service": "sls", "action": "GetLogStore"},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "project",
				Type:        proto.ColumnType_STRING,
				Description: "The SLS project name.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The logstore name.",
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "ttl",
				Type:        proto.ColumnType_INT,
				Description: "The data retention period in days. -1 indicates permanent storage.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.TTL"),
			},
			{
				Name:        "shard_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of shards in the logstore.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.ShardCount"),
			},
			{
				Name:        "web_tracking",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether web tracking is enabled.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.WebTracking"),
			},
			{
				Name:        "auto_split",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether auto-split is enabled.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.AutoSplit"),
			},
			{
				Name:        "max_split_shard",
				Type:        proto.ColumnType_INT,
				Description: "The maximum number of split shards.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.MaxSplitShard"),
			},
			{
				Name:        "append_meta",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates whether metadata is appended.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.AppendMeta"),
			},
			{
				Name:        "telemetry_type",
				Type:        proto.ColumnType_STRING,
				Description: "The telemetry type of the logstore.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.TelemetryType"),
			},
			{
				Name:        "hot_ttl",
				Type:        proto.ColumnType_INT,
				Description: "The hot data retention period in days.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.HotTTL"),
			},
			{
				Name:        "mode",
				Type:        proto.ColumnType_STRING,
				Description: "The logstore mode (query or standard).",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.Mode"),
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the logstore was created.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.CreateTime").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "last_modify_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the logstore was last modified.",
				Hydrate:     getLogstore,
				Transform:   transform.FromField("Logstore.LastModifyTime").Transform(transform.UnixToTimestamp),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: ColumnDescriptionAkas,
				Hydrate:     getLogstoreAkas,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
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

type logstoreItem struct {
	Project  string
	Region   string
	Name     string
	Logstore *sls.LogStore
}

func listLogstores(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	// Get project from parent hydrate (parent items are passed as h.Item in child hydrates)
	if h.Item == nil {
		plugin.Logger(ctx).Warn("alicloud_listLogstores", "parent_item_is_nil")
		return nil, nil
	}

	parentItem, ok := h.Item.(*sls.LogProject)
	if !ok || parentItem == nil {
		plugin.Logger(ctx).Error("alicloud_listLogstores", "invalid_parent_item_type", "type", fmt.Sprintf("%T", h.Item))
		return nil, nil
	}

	if parentItem.Name == "" {
		plugin.Logger(ctx).Warn("alicloud_listLogstores", "project_name_is_empty", "project", parentItem)
		return nil, nil
	}

	project := parentItem.Name
	plugin.Logger(ctx).Trace("alicloud_listLogstores", "project", project, "region", region)

	client, err := SLSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listLogstores", "connection_error", err)
		return nil, err
	}

	// List logstores for this project with pagination
	offset := 0
	size := 100
	for {
		d.WaitForListRateLimit(ctx)
		logstoreNames, err := client.ListLogStoreV2(project, offset, size, "")
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_listLogstores", "list_logstore_error", err, "project", project)
			break
		}

		if len(logstoreNames) == 0 {
			break
		}

		// Stream list items with basic info only (full details fetched via hydrate when needed)
		for _, logstoreName := range logstoreNames {
			d.StreamListItem(ctx, logstoreItem{
				Project: project,
				Region:  region,
				Name:    logstoreName,
			})
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// If we got fewer results than requested, we've reached the end
		if len(logstoreNames) < size {
			break
		}
		offset += size
	}

	return nil, nil
}

//// GET FUNCTION

func getLogstore(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getLogstore")

	var project, name, region string

	// If called from list, get data from h.Item
	if h.Item != nil {
		data := h.Item.(logstoreItem)
		project = data.Project
		name = data.Name
		region = data.Region
	} else {
		// If called from get, get data from quals
		region = d.EqualsQualString(matrixKeyRegion)
		project = d.EqualsQualString("project")
		name = d.EqualsQualString("name")
	}

	if project == "" || name == "" {
		return nil, nil
	}

	client, err := SLSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_getLogstore", "connection_error", err)
		return nil, err
	}

	logstore, err := client.GetLogStore(project, name)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_getLogstore", "get_logstore_error", err, "project", project, "name", name)
		return nil, err
	}

	// Update the item with full logstore details
	item := logstoreItem{
		Project:  project,
		Region:   region,
		Name:     name,
		Logstore: logstore,
	}

	// If called from list, update the existing item
	if h.Item != nil {
		return item, nil
	}

	return item, nil
}

//// TRANSFORMS

func getLogstoreAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getLogstoreAkas")

	data := h.Item.(logstoreItem)

	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:log:" + data.Region + ":" + accountID + ":project/" + data.Project + "/logstore/" + data.Name}
	return akas, nil
}
