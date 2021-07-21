# Table: alicloud_ecs_disk_metric_write_ops_daily

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_ecs_disk_metric_write_ops_daily` table provides metric statistics at 24 hour intervals for the most recent 90 days.

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
  alicloud_ecs_disk_metric_write_ops_daily
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average write ops

```sql
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_write_ops_daily
where average > 1000
order by
  instance_id,
  timestamp;
```