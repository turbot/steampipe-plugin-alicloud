package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudVpcFlowLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_flow_log",
		Description: "Virtual Private Cloud (VPC) provides the flow log feature to capture information about inbound and outbound traffic on an elastic network interface (ENI).",
		List: &plugin.ListConfig{
			Hydrate: listVpcFlowLogs,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getVpcFlowLog,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FlowLogName"),
				Description: "The name of the flow log.",
			},
			{
				Name:        "flow_log_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the flow log.",
			},

			// Other columns
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the flow log.",
			},
			{
				Name:        "creation_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the flow log was created.",
			},
			{
				Name:        "resource_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of resource from which traffic is captured.",
			},
			{
				Name:        "resource_id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the resource from which traffic is captured.",
			},
			{
				Name:        "project_name",
				Type:        proto.ColumnType_STRING,
				Description: "The project that stores the captured traffic data.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the flow log.",
			},
			{
				Name:        "log_store_name",
				Type:        proto.ColumnType_BOOL,
				Description: "The Logstore that stores the captured traffic data.",
			},
			{
				Name:        "traffic_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of traffic that is captured.",
			},

			// steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(vpcFlowLogTitle),
				Description: ColumnDescriptionTitle,
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Hydrate:     vpcFlowLogAkas,
				Description: ColumnDescriptionAkas,
				Transform:   transform.FromValue(),
			},

			// alicloud common columns
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

func listVpcFlowLogs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_flow_log.listVpcFlowLogs", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeFlowLogsRequest()
	request.Scheme = "https"

	response, err := client.DescribeFlowLogs(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_flow_log.listVpcFlowLogs", "query_error", err, "request", request)
		return nil, err
	}
	for _, i := range response.FlowLogs.FlowLog {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getVpcFlowLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := VpcService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcFlowLog", "connection_error", err)
		return nil, err
	}
	name := d.KeyColumnQuals["name"].GetStringValue()

	request := vpc.CreateDescribeFlowLogsRequest()
	request.Scheme = "https"
	request.FlowLogName = name
	response, err := client.DescribeFlowLogs(request)
	if err != nil {
		plugin.Logger(ctx).Error("getVpcFlowLog", "query_error", err, "request", request)
		return nil, err
	}

	if response.FlowLogs.FlowLog != nil && len(response.FlowLogs.FlowLog) > 0 {
		return response.FlowLogs.FlowLog[0], nil
	}

	return nil, nil
}

func vpcFlowLogAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	i := h.Item.(vpc.FlowLog)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID
	return []string{"acs:vpc:" + i.RegionId + ":" + accountID + ":flow-log/" + i.FlowLogId}, nil
}

//// TRANSFORM FUNCTIONS

func vpcFlowLogTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.HydrateItem.(vpc.FlowLog)

	// Build resource title
	title := i.FlowLogId
	if len(i.FlowLogName) > 0 {
		title = i.FlowLogName
	}

	return title, nil
}
