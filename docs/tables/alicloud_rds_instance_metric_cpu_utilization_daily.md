---
title: "Steampipe Table: alicloud_rds_instance_metric_cpu_utilization_daily - Query Alibaba Cloud RDS Instance Metrics using SQL"
description: "Allows users to query Alibaba Cloud RDS Instance Metrics, specifically the daily CPU utilization, providing insights into resource usage and potential performance bottlenecks."
---

# Table: alicloud_rds_instance_metric_cpu_utilization_daily - Query Alibaba Cloud RDS Instance Metrics using SQL

Alibaba Cloud Relational Database Service (RDS) is a stable and reliable online database service that supports MySQL, SQL Server, and PostgreSQL engines. It provides a complete set of solutions to handle disaster recovery, backup, restoration, monitoring, and migration, allowing users to focus on business innovation. RDS Instance Metrics provide detailed performance and health insights for instances within the RDS service.

## Table Usage Guide

The `alicloud_rds_instance_metric_cpu_utilization_daily` table provides insights into the daily CPU utilization of RDS instances within Alibaba Cloud. As a database administrator or DevOps engineer, you can explore instance-specific details through this table, including CPU usage patterns, peak usage times, and potential performance bottlenecks. Utilize it to monitor and optimize resource usage, ensuring the efficient operation of your databases.

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
  alicloud_rds_instance_metric_cpu_utilization_daily
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
  alicloud_rds_instance_metric_cpu_utilization_daily
where
  average > 80
order by
  db_instance_id,
  timestamp;
```
