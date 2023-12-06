# Table: alicloud_rds_instance_metric_connections_daily

Alicloud Monitoring metrics provide data about the performance of your systems. The `alicloud_rds_instance_metric_connections_daily` table provides metric statistics at 24 hour intervals for the most recent 30 days.

Note: If the instance is not older than one day then we will not get any metric statistics.

## Examples

### Basic info
Explore the variation in daily connection metrics for your RDS instances to understand usage patterns and potential bottlenecks. This can be useful for capacity planning and performance optimization.

```sql+postgres
select
  db_instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_rds_instance_metric_connections_daily
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
  alicloud_rds_instance_metric_connections_daily
order by
  db_instance_id,
  timestamp;
```

### Intervals where connection exceed 1000 average
Discover the instances where the average number of connections exceeds 1000 in your Alicloud RDS, providing a detailed overview of database usage and potential performance issues.

```sql+postgres
select
  db_instance_id,
  timestamp,
  round(minimum::numeric,2) as min_conn,
  round(maximum::numeric,2) as max_conn,
  round(average::numeric,2) as avg_conn
from
  alicloud_rds_instance_metric_connections_daily
where
  average > 1000
order by
  db_instance_id,
  timestamp;
```

```sql+sqlite
select
  db_instance_id,
  timestamp,
  round(minimum,2) as min_conn,
  round(maximum,2) as max_conn,
  round(average,2) as avg_conn
from
  alicloud_rds_instance_metric_connections_daily
where
  average > 1000
order by
  db_instance_id,
  timestamp;
```