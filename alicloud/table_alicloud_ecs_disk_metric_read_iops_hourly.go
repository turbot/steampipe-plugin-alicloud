package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsDiskMetricReadIopsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_disk_metric_read_iops_hourly",
		Description: "Alicloud ECS Disk Cloud Monitor Metrics - Read IOPS (Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: listEcsInstance,
			Hydrate:       listEcsDisksMetricReadIopsHourly,
		},
		GetMatrixItemFunc: BuildRegionList,
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

func listEcsDisksMetricReadIopsHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(ecs.Instance)
	return listCMMetricStatistics(ctx, d, "HOURLY", "acs_ecs_dashboard", "DiskReadIOPS", "instanceId", data.InstanceId)
}
