package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type launchTemplateInfo = struct {
	ecs.LaunchTemplateSet
	Region string
}

//// TABLE DEFINITION

func tableAlicloudEcsLaunchTemplate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_launch_template",
		Description: "Alicloud ECS Launch Template",
		List: &plugin.ListConfig{
			Hydrate: listEcsLaunchTemplates,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("launch_template_id"),
			Hydrate:    getEcsLaunchTemplate,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "A friendly name for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("LaunchTemplateName"),
			},
			{
				Name:        "launch_template_id",
				Description: "An unique identifier for the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_by",
				Description: "Specifies the creator of the launch template.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "The time when the launch template was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "default_version_number",
				Description: "The default version number of the launch template.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "latest_version_number",
				Description: "The latest version number of the launch template.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "modified_time",
				Description: "The time when the launch template was modified.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the launch template belongs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "latest_version_details",
				Description: "Describes the configuration of latest launch template version.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getEcsLaunchTemplateLatestVersionDetails,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(modifyEcsSourceTags),
			},

			// steampipe standard columns
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
				Hydrate:     getEcsLaunchTemplateAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(ecsLaunchTemplateAka),
			},

			// alicloud standard columns
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

func listEcsLaunchTemplates(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_launch_template.listEcsLaunchTemplates", "connection_error", err)
		return nil, err
	}
	request := ecs.CreateDescribeLaunchTemplatesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeLaunchTemplates(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_launch_template.listEcsLaunchTemplates", "query_error", err, "request", request)
			return nil, err
		}
		for _, launchTemplate := range response.LaunchTemplateSets.LaunchTemplateSet {
			d.StreamListItem(ctx, launchTemplateInfo{launchTemplate, region})
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

func getEcsLaunchTemplate(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getEcsLaunchTemplate")

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_launch_template.getEcsLaunchTemplate", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		disk := h.Item.(ecs.Disk)
		id = disk.DiskId
	} else {
		id = d.KeyColumnQuals["launch_template_id"].GetStringValue()
	}

	request := ecs.CreateDescribeLaunchTemplatesRequest()
	request.Scheme = "https"
	request.LaunchTemplateId = &[]string{id}

	response, err := client.DescribeLaunchTemplates(request)
	if err != nil {
		return nil, err
	}

	if response.LaunchTemplateSets.LaunchTemplateSet != nil && len(response.LaunchTemplateSets.LaunchTemplateSet) > 0 {
		return launchTemplateInfo{response.LaunchTemplateSets.LaunchTemplateSet[0], region}, nil
	}

	return nil, nil
}

func getEcsLaunchTemplateLatestVersionDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsLaunchTemplateLatestVersionDetails")

	data := h.Item.(launchTemplateInfo)
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ECSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ecs_launch_template.getEcsLaunchTemplateLatestVersionDetails", "connection_error", err)
		return nil, err
	}

	request := ecs.CreateDescribeLaunchTemplateVersionsRequest()
	request.Scheme = "https"
	request.LaunchTemplateId = data.LaunchTemplateId
	request.LaunchTemplateVersion = &[]string{strconv.Itoa(int(data.LatestVersionNumber))}

	response, err := client.DescribeLaunchTemplateVersions(request)
	if err != nil {
		return nil, err
	}

	if response.LaunchTemplateVersionSets.LaunchTemplateVersionSet != nil && len(response.LaunchTemplateVersionSets.LaunchTemplateVersionSet) > 0 {
		return response.LaunchTemplateVersionSets.LaunchTemplateVersionSet[0], nil
	}

	return nil, nil
}

func getEcsLaunchTemplateAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getEcsLaunchTemplateAka")
	data := h.Item.(launchTemplateInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:ecs:" + data.Region + ":" + accountID + ":launch-template/" + data.LaunchTemplateId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func ecsLaunchTemplateAka(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(launchTemplateInfo)

	// Build resource title
	title := data.LaunchTemplateId

	if len(data.LaunchTemplateName) > 0 {
		title = data.LaunchTemplateName
	}

	return title, nil
}
