package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsInstanceMetricCpuUtilization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_instance_metric_cpu_utilization",
		Description: "Alicloud RDS Instance Cloud Monitor Metrics - CPU Utilization",
		List: &plugin.ListConfig{
			ParentHydrate: listRdsInstances,
			Hydrate:       listRdsInstanceMetricCpuUtilization,
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

func listRdsInstanceMetricCpuUtilization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(rds.DBInstance)
	return listCMMetricStatistics(ctx, d, "5_MIN", "acs_rds_dashboard", "CpuUsage", "instanceId", data.DBInstanceId)
}
