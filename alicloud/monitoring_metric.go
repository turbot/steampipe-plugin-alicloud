package alicloud

import (
	"math"
	"strings"
	"time"

	// cms "github.com/alibabacloud-go/cms-20190101/client" // Temporarily disabled due to code generation issues

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			Default:     0,
		},
		{
			Name:        "maximum",
			Description: "The maximum metric value for the data point.",
			Type:        proto.ColumnType_DOUBLE,
			Default:     0,
		},
		{
			Name:        "minimum",
			Description: "The minimum metric value for the data point.",
			Type:        proto.ColumnType_DOUBLE,
			Default:     0,
		},
		{
			Name:        "timestamp",
			Description: "The timestamp used for the data point.",
			Type:        proto.ColumnType_TIMESTAMP,
		},
		{
			Name:        "account_id",
			Description: ColumnDescriptionAccount,
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Transform:   transform.FromField("AccountID"),
		},
	}
}

type CMMetricRow struct {
	// The (single) metric Dimension name
	DimensionName string

	// The value for the (single) metric Dimension
	DimensionValue string

	// The namespace of the metric
	Namespace string

	// The name of the metric
	MetricName string

	// The average of the metric values that correspond to the data point.
	Average float64

	// The percentile statistic for the data point.
	// ExtendedStatistics map[string]*float64 `type:"map"`

	// The maximum metric value for the data point.
	Maximum float64

	// The minimum metric value for the data point.
	Minimum float64

	// The timestamp used for the data point.
	Timestamp string
}

func getCMStartDateForGranularity(granularity string) string {
	str := "2006-01-02T15:04:05Z"
	switch strings.ToUpper(granularity) {
	case "DAILY":
		// 30 days
		return time.Now().AddDate(0, 0, -30).Format(str)
	case "HOURLY":
		// 30 days
		return time.Now().AddDate(0, 0, -30).Format(str)
	}
	// else 5 days
	return time.Now().AddDate(0, 0, -5).Format(str)
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

func getCustomError(errorMessage string) error {
	return errors.NewServerError(500, errorMessage, "")
}

// Temporarily disabled due to CMS service issues
// func listCMMetricStatistics(...) { ... }

func formatTime(timestamp float64) string {
	timeInSec := math.Floor(timestamp / 1000)
	unixTimestamp := time.Unix(int64(timeInSec), 0)
	timestampRFC3339Format := unixTimestamp.Format(time.RFC3339)
	return timestampRFC3339Format
}
