package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsInstanceMetricCpuUtilizationDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_instance_metric_cpu_utilization_daily",
		Description: "Alicloud ECS Instance Cloud Monitor Metrics - CPU Utilization (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listEcsInstance,
			Tags:          map[string]string{"service": "ecs", "monitoring": "CPUUtilization", "daily": "true"},
			Hydrate:       listEcsInstanceMetricCpuUtilizationDaily,
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: cmMetricColumns(
			[]*plugin.Column{
				{
					Name:        "instance_id",
					Description: "The ID of the instance.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			}),
	}
}

func listEcsInstanceMetricCpuUtilizationDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(ecs.Instance)
	return listCMMetricStatistics(ctx, d, "DAILY", "acs_ecs_dashboard", "CPUUtilization", "instanceId", data.InstanceId)
}
