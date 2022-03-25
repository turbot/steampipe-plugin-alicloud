# Table: alicloud_ecs_instance_metric_cpu_utilization_hourly

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_ecs_instance_metric_cpu_utilization_hourly` table provides metric statistics at 1 hour intervals for the most recent 30 days.

Note: If the instance is not older than one hour then we will not get any metric statistics.

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
  alicloud_ecs_instance_metric_cpu_utilization_hourly
order by
  instance_id,
  timestamp;
```

### CPU Over 80% average

```sql
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_cpu,
  round(maximum::numeric,2) as max_cpu,
  round(average::numeric,2) as avg_cpu
from
  alicloud_ecs_instance_metric_cpu_utilization_hourly
where average > 80
order by
  instance_id,
  timestamp;
```