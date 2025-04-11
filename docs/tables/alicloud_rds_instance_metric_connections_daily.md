---
title: "Steampipe Table: alicloud_rds_instance_metric_connections_daily - Query Daily Connection Metrics for RDS Instances using SQL"
description: "Allows users to query daily average connection metrics for Alibaba Cloud RDS instances, including instance ID, timestamp, and average number of active connections."
folder: "RDS"
---

# Table: alicloud_rds_instance_metric_connections_daily - Query Daily Connection Metrics for RDS Instances using SQL

Alibaba Cloud Relational Database Service (RDS) provides reliable and scalable database management. Monitoring connection metrics over time helps detect usage patterns, optimize performance, and plan for scaling.

## Table Usage Guide

The `alicloud_rds_instance_metric_connections_daily` table enables database administrators and DevOps teams to query daily average connection metrics for RDS instances in Alibaba Cloud. Use this table to retrieve values such as instance ID, date (timestamp), and the average number of active connections. This data is useful for identifying trends, spotting anomalies, and making informed decisions about database scaling and resource allocation.

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