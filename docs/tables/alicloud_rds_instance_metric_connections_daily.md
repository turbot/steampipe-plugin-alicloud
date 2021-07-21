# Table: alicloud_rds_instance_metric_connections_daily

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_rds_instance_metric_connections_daily` table provides metric statistics at 24 hour intervals for the most recent 90 days.

## Examples

### Basic info

```sql
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_rds_instance_metric_connections_daily
order by
  instance_id,
  timestamp;
```

### Intervals where connection exceed 1000 average

```sql
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_conn,
  round(maximum::numeric,2) as max_conn,
  round(average::numeric,2) as avg_conn,
from
  alicloud_rds_instance_metric_connections_daily
where average > 1000
order by
  instance_id,
  timestamp;
```