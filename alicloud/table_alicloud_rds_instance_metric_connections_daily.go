package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsInstanceMetricConnectionsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_instance_metric_connections_daily",
		Description: "Alicloud RDS Instance Cloud Monitor Metrics - Connections (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listRdsInstances,
			Hydrate:       listRdsInstanceMetricConnectionsDaily,
		},
		GetMatrixItemFunc: BuildRegionList,
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

func listRdsInstanceMetricConnectionsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(rds.DBInstance)
	return listCMMetricStatistics(ctx, d, "DAILY", "acs_rds_dashboard", "ConnectionUsage", "instanceId", data.DBInstanceId)
}
