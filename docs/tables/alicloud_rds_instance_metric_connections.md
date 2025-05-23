---
title: "Steampipe Table: alicloud_rds_instance_metric_connections - Query RDS Instance Connection Metrics using SQL"
description: "Allows users to query connection metrics for Alibaba Cloud RDS instances, including instance ID, timestamp, and number of active connections."
folder: "RDS"
---

# Table: alicloud_rds_instance_metric_connections - Query RDS Instance Connection Metrics using SQL

Alibaba Cloud Relational Database Service (RDS) provides scalable and managed database solutions for various database engines. Monitoring the number of active connections to an RDS instance helps assess database load, troubleshoot performance issues, and ensure availability.

## Table Usage Guide

The `alicloud_rds_instance_metric_connections` table enables database administrators and DevOps engineers to query real-time and historical connection metrics for Alibaba Cloud RDS instances. Use this table to retrieve attributes such as the RDS instance ID, metric timestamp, and the number of active connections. This information supports proactive monitoring, capacity planning, and performance optimization of your database infrastructure.

Note: If the instance is not older than 5 minute then we will not get any metric statistics.

## Examples

### Basic info
Explore the connection metrics of database instances to gain insights into their performance and usage over time. This can help in identifying potential bottlenecks, planning capacity, and understanding the overall health of your databases.

```sql+postgres
select
  db_instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_rds_instance_metric_connections
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
  alicloud_rds_instance_metric_connections
order by
  db_instance_id,
  timestamp;
```

### Intervals where connection exceed 1000 average
Determine the intervals in which the average number of connections to your Alicloud RDS instances exceeds 1000. This can help identify potential performance issues or periods of heavy usage.

```sql+postgres
select
  db_instance_id,
  timestamp,
  round(minimum::numeric,2) as min_conn,
  round(maximum::numeric,2) as max_conn,
  round(average::numeric,2) as avg_conn
from
  alicloud_rds_instance_metric_connections
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
  alicloud_rds_instance_metric_connections
where
  average > 1000
order by
  db_instance_id,
  timestamp;
```