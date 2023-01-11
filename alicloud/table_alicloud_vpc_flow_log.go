package alicloud

import (
	"context"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
)

//// TABLE DEFINITION

func tableAlicloudVpcFlowLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_vpc_flow_log",
		Description: "Alicloud VPC Flow Log",
		List: &plugin.ListConfig{
			Hydrate: listVpcFlowLogs,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "name", Require: plugin.Optional},
				{Name: "log_store_name", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
				{Name: "project_name", Require: plugin.Optional},
				{Name: "traffic_type", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("flow_log_id"),
			Hydrate:    getVpcFlowLog,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the flow log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FlowLogName"),
			},
			{
				Name:        "flow_log_id",
				Description: "The id of the flow log.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the flow log.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the flow log was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "resource_type",
				Description: "The resource type of traffic to capture.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "resource_id",
				Description: "The resource ID of the traffic to capture.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_name",
				Description: "Project that manages captured traffic.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "log_store_name",
				Description: "Logstore for storing captured traffic.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of the flow log.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "traffic_type",
				Description: "The collected traffic type. ",
				Type:        proto.ColumnType_STRING,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FlowLogName"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getVpcFlowLogAka,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getVpcFlowLogRegion,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "account_id",
				Description: ColumnDescriptionAccount,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OwnerId"),
			},
		},
	}
}

//// LIST FUNCTION

func listVpcFlowLogs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_flow_log.listVpcFlowLogs", "connection_error", err)
		return nil, err
	}
	request := vpc.CreateDescribeFlowLogsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	if d.KeyColumnQualString("name") != "" {
		request.FlowLogName = d.KeyColumnQualString("name")
	}
	if d.KeyColumnQualString("log_store_name") != "" {
		request.LogStoreName = d.KeyColumnQualString("log_store_name")
	}
	if d.KeyColumnQualString("resource_id") != "" {
		request.ResourceId = d.KeyColumnQualString("resource_id")
	}
	if d.KeyColumnQualString("status") != "" {
		request.Status = d.KeyColumnQualString("status")
	}
	if d.KeyColumnQualString("project_name") != "" {
		request.ProjectName = d.KeyColumnQualString("project_name")
	}
	if d.KeyColumnQualString("traffic_type") != "" {
		request.TrafficType = d.KeyColumnQualString("traffic_type")
	}

	count := 0
	for {
		response, err := client.DescribeFlowLogs(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_vpc_flow_log.listVpcFlowLogs", "api_error", err, "request", request)
			return nil, err
		}
		for _, flowLog := range response.FlowLogs.FlowLog {
			d.StreamListItem(ctx, flowLog)
			count++
		}
		totalCount, _ := strconv.Atoi(response.TotalCount)
		if count >= totalCount {
			break
		}
		pageNumber, _ := strconv.Atoi(response.PageNumber)
		request.PageNumber = requests.NewInteger(pageNumber + 1)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getVpcFlowLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := VpcService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_flow_log.getVpcFlowLog", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["flow_log_id"].GetStringValue()

	request := vpc.CreateDescribeFlowLogsRequest()
	request.Scheme = "https"
	request.FlowLogId = id

	response, err := client.DescribeFlowLogs(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_vpc_flow_log.getVpcFlowLog", "api_error", err, "request", request)
		return nil, err
	}

	if response.FlowLogs.FlowLog != nil && len(response.FlowLogs.FlowLog) > 0 {
		return response.FlowLogs.FlowLog[0], nil
	}

	return nil, nil
}

func getVpcFlowLogAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(vpc.FlowLog)
	region := d.KeyColumnQualString(matrixKeyRegion)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:vpc:" + region + ":" + accountID + ":flowlog/" + data.FlowLogId}

	return akas, nil
}

func getVpcFlowLogRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.KeyColumnQualString(matrixKeyRegion)

	return region, nil
}
