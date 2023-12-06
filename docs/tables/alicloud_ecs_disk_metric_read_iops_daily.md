# Table: alicloud_ecs_disk_metric_read_iops_daily

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_ecs_disk_metric_read_iops_daily` table provides metric statistics at 24 hour intervals for the most recent 30 days.

Note: If the instance is not older than one day then we will not get any metric statistics.

## Examples

### Basic info
Gain insights into daily disk read operations on your Alibaba Cloud ECS instances. This can help you monitor performance trends and identify potential issues related to disk input/output operations.

```sql+postgres
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_disk_metric_read_iops_daily
order by
  instance_id,
  timestamp;
```

```sql+sqlite
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_disk_metric_read_iops_daily
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average read iops
Explore instances where the daily average read operations per second (IOPS) on an Alibaba Cloud Elastic Compute Service (ECS) disk exceed 1000. This can be useful in identifying periods of high disk usage, which could impact system performance.

```sql+postgres
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_read_iops_daily
where average > 1000
order by
  instance_id,
  timestamp;
```

```sql+sqlite
select
  instance_id,
  timestamp,
  round(minimum,2) as min_ops,
  round(maximum,2) as max_ops,
  round(average,2) as avg_ops
from
  alicloud_ecs_disk_metric_read_iops_daily
where average > 1000
order by
  instance_id,
  timestamp;
```