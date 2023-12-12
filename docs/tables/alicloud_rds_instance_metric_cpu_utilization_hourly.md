# Table: alicloud_rds_instance_metric_cpu_utilization_hourly

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_rds_instance_metric_cpu_utilization_hourly` table provides metric statistics at 1 hour intervals for the most recent 30 days.

Note: If the instance is not older than 1 hour then we will not get any metric statistics.

## Examples

### Basic info
Analyze the settings to understand the CPU utilization of each database instance over time. This can help in assessing the performance and identifying potential bottlenecks.

```sql+postgres
select
  db_instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_rds_instance_metric_cpu_utilization_hourly
order by
  db_instance_id,
  timestamp;
```

```sql+sqlite
select
  db_instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_rds_instance_metric_cpu_utilization_hourly
order by
  db_instance_id,
  timestamp;
```

### CPU over 80% average
Determine the areas in which database instances are experiencing high CPU utilization, specifically where the average CPU usage exceeds 80%. This can assist in identifying potential performance issues and optimizing resource allocation.

```sql+postgres
select
  db_instance_id,
  timestamp,
  round(minimum::numeric,2) as min_cpu,
  round(maximum::numeric,2) as max_cpu,
  round(average::numeric,2) as avg_cpu
from
  alicloud_rds_instance_metric_cpu_utilization_hourly
where
  average > 80
order by
  db_instance_id,
  timestamp;
```

```sql+sqlite
select
  db_instance_id,
  timestamp,
  round(minimum,2) as min_cpu,
  round(maximum,2) as max_cpu,
  round(average,2) as avg_cpu
from
  alicloud_rds_instance_metric_cpu_utilization_hourly
where
  average > 80
order by
  db_instance_id,
  timestamp;
```