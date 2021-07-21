package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION
func tableAlicloudRdsInstanceMetricConnectionDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_rds_instance_metric_connection_daily",
		Description: "Alicloud RDS Instance Cloud Monitor Metrics - connection (Daily)",
		List: &plugin.ListConfig{
			Hydrate: listRdsInstanceMetricconnectionDaily,
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

func listRdsInstanceMetricconnectionDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return listCMMetricStatistics(ctx, d, "DAILY", "acs_rds_dashboard", "ConnectionUsage", "instanceId")
}
