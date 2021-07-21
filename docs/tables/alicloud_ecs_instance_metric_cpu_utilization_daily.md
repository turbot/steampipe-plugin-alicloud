# Table: alicloud_ecs_instance_metric_cpu_utilization_daily

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_ecs_instance_metric_cpu_utilization_daily` table provides metric statistics at 24 hour intervals for the most recent 90 days.

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
  alicloud_ecs_instance_metric_cpu_utilization_daily
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
  alicloud_ecs_instance_metric_cpu_utilization_daily
where average > 80
order by
  instance_id,
  timestamp;
```