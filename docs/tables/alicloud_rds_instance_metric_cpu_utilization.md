---
title: "Steampipe Table: alicloud_rds_instance_metric_cpu_utilization - Query Alicloud RDS Instance Metrics using SQL"
description: "Allows users to query Alicloud RDS Instance Metrics, specifically the CPU utilization, providing insights into database performance and potential bottlenecks."
---

# Table: alicloud_rds_instance_metric_cpu_utilization - Query Alicloud RDS Instance Metrics using SQL

Alicloud Relational Database Service (RDS) is a stable and reliable online database service that supports MySQL, SQL Server, and PostgreSQL. RDS handles routine database tasks such as patching and backup, freeing up time to focus on application development. It provides high performance and high availability with automatic failover.

## Table Usage Guide

The `alicloud_rds_instance_metric_cpu_utilization` table provides insights into the CPU utilization of Alicloud RDS instances. As a database administrator, you can gain detailed information about the CPU usage of your RDS instances, helping you to monitor performance and identify potential bottlenecks or over-utilization. This table is particularly useful for optimizing resource allocation and maintaining efficient database operations.

**Important Notes**
- If the instance is not older than 5 minute then we will not get any metric statistics.

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
