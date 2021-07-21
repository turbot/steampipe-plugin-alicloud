package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION
func tableAlicloudEcsDiskMetricReadOps(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_disk_metric_read_ops",
		Description: "Alicloud ECS Disk Cloud Monitor Metrics - Read Ops",
		List: &plugin.ListConfig{
			ParentHydrate: listEcsDisk,
			Hydrate:       listRdsInstanceMetricReadIopsHourly,
		},
		GetMatrixItem: BuildRegionList,
		Columns: cmMetricColumns(
			[]*plugin.Column{
				{
					Name:        "disk_id",
					Description: "An unique identifier for the resource.",
					Type:        proto.ColumnType_STRING,
					Transform:   transform.FromField("DimensionValue"),
				},
			}),
	}
}

func listRdsInstanceMetricReadIopsHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	disk := h.Item.(ecs.Disk)
	return listCMMetricStatistics(ctx, d, "5_MIN", "acs_ecs_dashboard", "DiskReadIOPS", "instanceId", disk.DiskId)
}