package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsInstanceMetricCpuUtilizationHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_instance_metric_cpu_utilization_hourly",
		Description: "Alicloud RDS Instance Cloud Monitor Metrics - CPU Utilization (Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: listRdsInstances,
			Hydrate:       listRdsInstanceMetricCpuUtilizationHourly,
		},
		GetMatrixItem: BuildRegionList,
		Columns: cmMetricColumns(
			[]*plugin.Column{
				{
					Name:        "db_instance_id",
					Description: "The ID of the single instance to query.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			}),
	}
}

func listRdsInstanceMetricCpuUtilizationHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(rds.DBInstance)
	return listCMMetricStatistics(ctx, d, "HOURLY", "acs_rds_dashboard", "CpuUsage", "instanceId", data.DBInstanceId)
}
