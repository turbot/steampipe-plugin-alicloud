package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRdsInstanceMetricConnections(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_instance_metric_connections",
		Description: "Alicloud RDS Instance Cloud Monitor Metrics - Connections",
		List: &plugin.ListConfig{
			Hydrate: listRdsInstanceMetricConnections,
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

func listRdsInstanceMetricConnections(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return listCMMetricStatistics(ctx, d, "5_MIN", "acs_rds_dashboard", "ConnectionUsage", "instanceId")
}
