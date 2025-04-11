---
title: "Steampipe Table: alicloud_rds_instance_metric_cpu_utilization_hourly - Query Hourly CPU Utilization Metrics for RDS Instances using SQL"
description: "Allows users to query hourly average CPU utilization metrics for Alibaba Cloud RDS instances, including instance ID, timestamp, and CPU usage percentage."
folder: "RDS"
---

# Table: alicloud_rds_instance_metric_cpu_utilization_hourly - Query Hourly CPU Utilization Metrics for RDS Instances using SQL

Alibaba Cloud Relational Database Service (RDS) supports high-performance and reliable databases for critical applications. Monitoring CPU utilization on an hourly basis is essential for understanding workload behavior and ensuring optimal performance.

## Table Usage Guide

The `alicloud_rds_instance_metric_cpu_utilization_hourly` table provides hourly average CPU usage metrics for RDS instances in Alibaba Cloud. Use this table to retrieve metrics such as instance ID, timestamp, and average CPU utilization percentage. This data is valuable for performance tuning, identifying spikes or inefficiencies, and making informed decisions on scaling and resource provisioning.

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