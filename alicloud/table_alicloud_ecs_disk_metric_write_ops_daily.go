package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION
func tableAlicloudEcsDiskMetricWriteOpsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_disk_metric_write_ops_daily",
		Description: "Alicloud ECS Disk Cloud Monitor Metrics - Write Ops (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listEcsInstance,
			Hydrate:       listEcsDisksMetricWriteIopsDaily,
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

func listEcsDisksMetricWriteIopsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(ecs.Instance)
	return listCMMetricStatistics(ctx, d, "DAILY", "acs_ecs_dashboard", "DiskWriteIOPS", "instanceId", data.InstanceId)
}
