# Table: alicloud_rds_instance_metric_cpu_utilization_hourly

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_rds_instance_metric_cpu_utilization_hourly` table provides metric statistics at 1 hour intervals for the most recent 30 days.

**Important Notes**
- If the instance is not older than 1 hour then we will not get any metric statistics.

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
  alicloud_rds_instance_metric_cpu_utilization_hourly
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
  alicloud_rds_instance_metric_cpu_utilization_hourly
where
  average > 80
order by
  db_instance_id,
  timestamp;
```
