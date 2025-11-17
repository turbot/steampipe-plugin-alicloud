package alicloud

import (
	"context"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudLogProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_log_project",
		Description: "Alicloud Log Service (SLS) Project.",
		List: &plugin.ListConfig{
			Hydrate: listLogProjects,
			Tags:    map[string]string{"service": "sls", "action": "ListProjectV2"},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getLogProject,
			Tags:       map[string]string{"service": "sls", "action": "GetProject"},
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The SLS project name.",
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The project description.",
				Transform:   transform.FromField("Description"),
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The project status (e.g., Normal).",
				Transform:   transform.FromField("Status"),
			},
			{
				Name:        "owner",
				Type:        proto.ColumnType_STRING,
				Description: "The project owner.",
				Transform:   transform.FromField("Owner"),
			},
			{
				Name:        "data_redundancy_type",
				Type:        proto.ColumnType_STRING,
				Description: "The data redundancy type. Valid values: LRS (Locally Redundant Storage) and ZRS (Zone Redundant Storage).",
				Transform:   transform.FromField("DataRedundancyType"),
			},
			{
				Name:        "location",
				Type:        proto.ColumnType_STRING,
				Description: "The project location (e.g., cn-beijing-b).",
				Transform:   transform.FromField("Location"),
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the project was created.",
				Transform:   transform.FromField("CreateTime").Transform(transform.UnixToTimestamp),
			},
			{
				Name:        "last_modify_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the project was last modified.",
				Transform:   transform.FromField("LastModifyTime").Transform(transform.UnixToTimestamp),
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
				Hydrate:     getLogProjectAkas,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(projectRegion),
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

func listLogProjects(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	plugin.Logger(ctx).Trace("alicloud_listLogProjects", "region", region)

	client, err := SLSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_listLogProjects", "connection_error", err)
		return nil, err
	}

	// List all projects in the region with pagination
	offset := 0
	size := 100
	projectCount := 0
	for {
		d.WaitForListRateLimit(ctx)
		projects, count, total, err := client.ListProjectV2(offset, size)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_listLogProjects", "list_project_error", err)
			return nil, err
		}

		for _, project := range projects {
			projectCopy := project
			d.StreamListItem(ctx, &projectCopy)
			projectCount++
			if d.RowsRemaining(ctx) == 0 {
				plugin.Logger(ctx).Trace("alicloud_listLogProjects", "streamed_projects", projectCount)
				return nil, nil
			}
		}

		offset += count
		if offset >= total || count == 0 {
			break
		}
	}

	plugin.Logger(ctx).Trace("alicloud_listLogProjects", "total_projects_streamed", projectCount)
	return nil, nil
}

//// GET FUNCTION

func getLogProject(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	name := d.EqualsQualString("name")
	if name == "" {
		return nil, nil
	}

	client, err := SLSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_getLogProject", "connection_error", err)
		return nil, err
	}

	project, err := client.GetProject(name)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_getLogProject", "get_project_error", err, "name", name)
		return nil, err
	}

	return project, nil
}

//// HYDRATE FUNCTIONS

func getLogProjectAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getLogProjectAkas")

	data := h.Item.(*sls.LogProject)

	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:log:" + data.Region + ":" + accountID + ":project/" + data.Name}
	return akas, nil
}

//// TRANSFORMS

func projectRegion(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("projectRegion")
	project := d.HydrateItem.(*sls.LogProject)
	return project.Region, nil
}
