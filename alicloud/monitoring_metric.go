package alicloud

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// append the common cloud monitoring metric columns onto the column list
func cmMetricColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonCMMetricColumns()...)
}

func commonCMMetricColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "metric_name",
			Description: "The name of the metric.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "namespace",
			Description: "The metric namespace.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "average",
			Description: "The average of the metric values that correspond to the data point.",
			Type:        proto.ColumnType_DOUBLE,
		},
		{
			Name:        "maximum",
			Description: "The maximum metric value for the data point.",
			Type:        proto.ColumnType_DOUBLE,
		},
		{
			Name:        "minimum",
			Description: "The minimum metric value for the data point.",
			Type:        proto.ColumnType_DOUBLE,
		},
		{
			Name:        "sample_count",
			Description: "The number of metric values that contributed to the aggregate value of this data point.",
			Type:        proto.ColumnType_DOUBLE,
		},
		{
			Name:        "sum",
			Description: "The sum of the metric values for the data point.",
			Type:        proto.ColumnType_DOUBLE,
		},
		{
			Name:        "unit",
			Description: "The standard unit for the data point.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "timestamp",
			Description: "The time stamp used for the data point.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
	}
}

type CMMetricRow struct {
	// The (single) metric Dimension name
	DimensionName *string

	// The value for the (single) metric Dimension
	DimensionValue *string

	// The namespace of the metric
	Namespace *string

	// The name of the metric
	MetricName *string

	// The average of the metric values that correspond to the data point.
	Average float64

	// The percentile statistic for the data point.
	//ExtendedStatistics map[string]*float64 `type:"map"`

	// The maximum metric value for the data point.
	Maximum float64

	// The minimum metric value for the data point.
	Minimum float64

	// The number of metric values that contributed to the aggregate value of this
	// data point.
	SampleCount *float64

	// The sum of the metric values for the data point.
	Sum *float64

	// The time stamp used for the data point.
	Timestamp *time.Time

	// The standard unit for the data point.
	Unit *string
}

func getCMStartDateForGranularity(granularity string) string {
	switch strings.ToUpper(granularity) {
	case "DAILY":
		// 1 year
		return time.Now().AddDate(-1, 0, 0).Format(time.RFC3339)
	case "HOURLY":
		// 60 days
		return time.Now().AddDate(0, 0, -60).Format(time.RFC3339)
	}
	// else 5 days
	return time.Now().AddDate(0, 0, -5).Format(time.RFC3339)
}

func getCMPeriodForGranularity(granularity string) string {
	switch strings.ToUpper(granularity) {
	case "DAILY":
		// 24 hours
		return "86400"
	case "HOURLY":
		// 1 hour
		return "3600"
	}
	// else 5 minutes
	return "300"
}

func listCMMetricStatistics(ctx context.Context, d *plugin.QueryData, granularity string, namespace string, metricName string, dimensionName string, dimensionValue string) (*cms.DescribeMetricListResponse, error) {
	region := GetDefaultRegion(d.Connection)

	// Create service connection
	client, err := CmsService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listCMMetricStatistics", "connection_error", err)
		return nil, err
	}
	request := cms.CreateDescribeMetricListRequest()

	request.MetricName = metricName
	request.StartTime = "2021-07-20T05:57:35Z" //getCMStartDateForGranularity(granularity)
	request.EndTime = "2021-07-21T05:57:35Z"//time.Now().Format(time.RFC3339)
	request.Namespace = namespace
	// request.Dimensions = "[{\"" + dimensionName + "\":\"" + dimensionValue + "\"}]"
	request.Period = "60" //getCMPeriodForGranularity(granularity)

	// count := 0
	// for {
	stats, err := client.DescribeMetricList(request)
	if err != nil {
		return nil, err
	}
	plugin.Logger(ctx).Trace("My Result => ", stats)



	var results []map[string]interface{}
	json.Unmarshal([]byte(stats.Datapoints), &results)
	plugin.Logger(ctx).Trace("Point Values => ", results)
	for _, pointValue := range results {
		d.StreamListItem(ctx, &CMMetricRow{
			DimensionName:  &dimensionName,
			DimensionValue: &dimensionValue,
			Namespace:      &namespace,
			MetricName:     &metricName,
			Average:        pointValue["Average"].(float64),
			Maximum:        pointValue["Maximum"].(float64),
			Minimum:        pointValue["Minimum"].(float64),
			// Sum: "",
			// SampleCount: "",
			// Timestamp: pointValue["timestamp"].(*time.Time),
			// Unit: "",
		})
	}
	//[{"timestamp":1548777660000,"userId":"120886317861****","instanceId":"i-abc","Minimum":9.92,"Average":9.92,"Maximum":9.92}]

	if stats.NextToken != "" {
		request.NextToken = stats.NextToken
	}

	// }

	return nil, nil
}
