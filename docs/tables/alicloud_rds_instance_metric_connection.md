# Table: alicloud_rds_instance_metric_connection

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_rds_instance_metric_connection` table provides metric statistics at 5 minute intervals for the most recent 5 days.

Note: If the instance is not older than one day then we will not get any metric statistices.
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
  alicloud_rds_instance_metric_connection
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
  round(average::numeric,2) as avg_conn
from
  alicloud_rds_instance_metric_connection
where average > 1000
order by
  instance_id,
  timestamp;
```