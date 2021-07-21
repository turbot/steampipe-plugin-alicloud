package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION
func tableAlicloudEcsInstanceMetricCpuUtilizationDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_instance_metric_cpu_utilization_daily",
		Description: "Alicloud ECS Disk Cloud Monitor Metrics - CPU Utilization (Daily)",
		List: &plugin.ListConfig{
			Hydrate: listEcsInstanceMetricCpuUtilization,
		},
		GetMatrixItem: BuildRegionList,
		Columns: cmMetricColumns(
			[]*plugin.Column{
				{
					Name:        "instance_id",
					Description: "An unique identifier for the resource.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			}),
	}
}

func listEcsInstanceMetricCpuUtilization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return listCMMetricStatistics(ctx, d, "DAILY", "acs_ecs_dashboard", "CPUUtilization", "instanceId")
}
