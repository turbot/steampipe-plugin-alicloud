# Table: alicloud_ecs_disk_metric_write_iops_hourly

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_ecs_disk_metric_write_iops_hourly` table provides metric statistics at 1 hour intervals for the most recent 30 days.

Note: If the instance is not older than one day then we will not get any metric statistics.

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
  alicloud_ecs_disk_metric_write_iops_hourly
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average write iops

```sql
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_write_iops_hourly
where average > 1000
order by
  instance_id,
  timestamp;
```