# Table: alicloud_rds_instance_metric_cpu_utilization

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_rds_instance_metric_cpu_utilization` table provides metric statistics at 5 minute intervals for the most recent 5 days.

Note: If the instance is not older than 5 minute then we will not get any metric statistics.

## Examples

### Basic info

```sql
select
  db_instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_rds_instance_metric_cpu_utilization
order by
  db_instance_id,
  timestamp;
```

### CPU over 80% average

```sql
select
  db_instance_id,
  timestamp,
  round(minimum::numeric,2) as min_cpu,
  round(maximum::numeric,2) as max_cpu,
  round(average::numeric,2) as avg_cpu
from
  alicloud_rds_instance_metric_cpu_utilization
where
  average > 80
order by
  db_instance_id,
  timestamp;
```
