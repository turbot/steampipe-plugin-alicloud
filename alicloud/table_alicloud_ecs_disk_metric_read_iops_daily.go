package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudEcsDiskMetricReadIopsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ecs_disk_metric_read_iops_daily",
		Description: "Alicloud ECS Disk Cloud Monitor Metrics - Read IOPS (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listEcsInstance,
			Hydrate:       listEcsDisksMetricReadIopsDaily,
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

func listEcsDisksMetricReadIopsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(ecs.Instance)
	return listCMMetricStatistics(ctx, d, "DAILY", "acs_ecs_dashboard", "DiskReadIOPS", "instanceId", data.InstanceId)
}
