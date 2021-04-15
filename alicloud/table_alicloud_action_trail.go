package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/actiontrail"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudActionTrail(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_action_trail",
		Description: "Alicloud Action Trail",
		List: &plugin.ListConfig{
			Hydrate: listActionTrails,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getActionTrail,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the trail.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "mns_topic_arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the Message Service (MNS) topic to which ActionTrail sends messages.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "home_region",
				Description: "The home region of the trail.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role_name",
				Description: "The name of the Resource Access Management (RAM) role that ActionTrail is allowed to assume.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the trail.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_organization_trail",
				Description: "Indicates whether the trail was created as a multi-account trail.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "oss_bucket_name",
				Description: "The name of the OSS bucket to which events are delivered.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "The time when the trail was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreateTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "event_rw",
				Description: "The read/write type of the delivered events.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EventRW"),
			},
			{
				Name:        "oss_key_prefix",
				Description: "The prefix of log files stored in the OSS bucket.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sls_project_arn",
				Description: "The ARN of the Log Service project to which events are delivered.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sls_write_role_arn",
				Description: "The ARN of the RAM role assumed by ActionTrail for delivering logs to the destination Log Service project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start_logging_time",
				Description: "The most recent date and time when logging was enabled for the trail.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("StartLoggingTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "stop_logging_time",
				Description: "The most recent date and time when logging was disabled for the trail.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("StopLoggingTime").Transform(transform.NullIfZeroValue).Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "trail_region",
				Description: "The regions to which the trail is applied.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "update_time",
				Description: "The most recent time when the configuration of the trail was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("UpdateTime").Transform(transform.UnixMsToTimestamp),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getActionTrailAka,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HomeRegion"),
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

func listActionTrails(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ActionTrailService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listActionTrails", "connection_error", err)
		return nil, err
	}
	request := actiontrail.CreateDescribeTrailsRequest()
	request.Scheme = "https"

	response, err := client.DescribeTrails(request)
	if err != nil {
		plugin.Logger(ctx).Error("listActionTrails", "query_error", err, "request", request)
		return nil, err
	}
	for _, trail := range response.TrailList {
		d.StreamListItem(ctx, trail)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getActionTrail(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getActionTrail")

	// Create service connection
	client, err := ActionTrailService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getActionTrail", "connection_error", err)
		return nil, err
	}

	name := d.KeyColumnQuals["name"].GetStringValue()

	request := actiontrail.CreateDescribeTrailsRequest()
	request.Scheme = "https"
	request.NameList = name
	response, err := client.DescribeTrails(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("getActionTrail", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if response.TrailList != nil && len(response.TrailList) > 0 {
		return response.TrailList[0], nil
	}

	return nil, nil
}

func getActionTrailAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getActionTrailAka")
	data := h.Item.(actiontrail.TrailListItem)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:actiontrail:" + data.HomeRegion + ":" + accountID + ":actiontrail/" + data.Name}

	return akas, nil
}
