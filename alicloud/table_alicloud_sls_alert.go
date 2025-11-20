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

func tableAlicloudSLSAlert(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_sls_alert",
		Description: "Alicloud Log Service (SLS) Alert.",
		List: &plugin.ListConfig{
			ParentHydrate: listLogProjects,
			Hydrate:       listSLSAlerts,
			Tags:          map[string]string{"service": "sls", "action": "ListAlert"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project", "name"}),
			Hydrate:    getSLSAlert,
			Tags:       map[string]string{"service": "sls", "action": "GetAlert"},
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
				Description: "The alert internal name.",
				Transform:   transform.FromField("Alert.Name"),
			},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The alert display name.",
				Transform:   transform.FromField("Alert.DisplayName"),
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Alert status, e.g., ENABLED/DISABLED.",
				Transform:   transform.FromField("Alert.Status"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Alert description.",
				Transform:   transform.FromField("Alert.Description"),
			},
			{
				Name:        "dashboard",
				Type:        proto.ColumnType_STRING,
				Description: "Dashboard associated with the alert (if any).",
				Transform:   transform.FromField("Alert.Configuration.Dashboard"),
			},
			{
				Name:        "schedule",
				Type:        proto.ColumnType_JSON,
				Description: "Schedule configuration.",
				Transform:   transform.FromField("Alert.Schedule"),
			},
			{
				Name:        "configuration",
				Type:        proto.ColumnType_JSON,
				Description: "Full alert configuration.",
				Transform:   transform.FromField("Alert.Configuration"),
			},
			{
				Name:        "query_list",
				Type:        proto.ColumnType_JSON,
				Description: "List of queries evaluated by the alert.",
				Transform:   transform.FromField("Alert.Configuration.QueryList"),
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Alert creation time.",
				Transform:   transform.FromField("Alert.CreateTime").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "last_modified_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Alert last modification time.",
				Transform:   transform.FromField("Alert.LastModifiedTime").Transform(transform.UnixToTimestamp),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: ColumnDescriptionTitle,
				Transform:   transform.FromField("Alert.DisplayName"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: ColumnDescriptionAkas,
				Hydrate:     getSLSAlertAkas,
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

type slsAlertItem struct {
	Project string
	Region  string
	Alert   *sls.Alert
}

func listSLSAlerts(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	// Get project from parent hydrate (parent items are passed as h.Item in child hydrates)
	if h.Item == nil {
		plugin.Logger(ctx).Warn("alicloud_listSLSAlerts", "parent_item_is_nil")
		return nil, nil
	}

	parentItem, ok := h.Item.(*sls.LogProject)
	if !ok || parentItem == nil {
		plugin.Logger(ctx).Error("alicloud_listSLSAlerts", "invalid_parent_item_type", "type", fmt.Sprintf("%T", h.Item))
		return nil, nil
	}

	if parentItem.Name == "" {
		plugin.Logger(ctx).Warn("alicloud_listSLSAlerts", "project_name_is_empty", "project", parentItem)
		return nil, nil
	}

	project := parentItem.Name
	plugin.Logger(ctx).Trace("alicloud_listSLSAlerts", "project", project, "region", region)

	client, err := SLSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listSLSAlerts", "connection_error", err)
		return nil, err
	}

	// List alerts for this project with pagination
	offset := 0
	size := 100
	for {
		d.WaitForListRateLimit(ctx)
		alerts, total, count, err := client.ListAlert(project, "", "", offset, size)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_listSLSAlerts", "list_alert_error", err, "project", project)
			break
		}
		for _, a := range alerts {
			d.StreamListItem(ctx, slsAlertItem{
				Project: project,
				Region:  region,
				Alert:   a,
			})
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		offset += count
		if offset >= total || count == 0 {
			break
		}
	}

	return nil, nil
}

//// GET FUNCTION

func getSLSAlert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	project := d.EqualsQualString("project")
	name := d.EqualsQualString("name")
	if project == "" || name == "" {
		return nil, nil
	}

	client, err := SLSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_getSLSAlert", "connection_error", err)
		return nil, err
	}

	alert, err := client.GetAlert(project, name)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_getSLSAlert", "get_alert_error", err, "project", project, "name", name)
		return nil, err
	}
	return slsAlertItem{
		Project: project,
		Region:  region,
		Alert:   alert,
	}, nil
}

//// TRANSFORMS

func getSLSAlertAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getSLSAlertAkas")

	data := h.Item.(slsAlertItem)

	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:log:" + data.Region + ":" + accountID + ":project/" + data.Project + "/alert/" + data.Alert.Name}
	return akas, nil
}
