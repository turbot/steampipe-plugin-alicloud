---
title: "Steampipe Table: alicloud_ecs_instance_metric_cpu_utilization_hourly - Query Alibaba Cloud ECS Instance Metrics using SQL"
description: "Allows users to query ECS Instance Metrics in Alibaba Cloud, specifically the hourly CPU utilization, providing insights into resource usage and performance trends."
folder: "ECS"
---

# Table: alicloud_ecs_instance_metric_cpu_utilization_hourly - Query Alibaba Cloud ECS Instance Metrics using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides scalable, on-demand computing resources for secure, flexible, and efficient applications. ECS Instance Metrics are part of the monitoring service of Alibaba Cloud ECS, which collects and analyzes the performance and operational status of your ECS instances. It helps you monitor the usage of your instances, allowing you to optimize resource allocation and troubleshoot system issues.

## Table Usage Guide

The `alicloud_ecs_instance_metric_cpu_utilization_hourly` table provides insights into the hourly CPU utilization of ECS instances within Alibaba Cloud. As a system administrator or DevOps engineer, explore instance-specific details through this table, including CPU usage trends, peak usage times, and overall performance. Utilize it to uncover information about instances, such as those with high CPU usage, the correlation between usage and performance, and the need for resource optimization.

## Examples

### Basic info
Explore which instances have the highest average CPU utilization over time, allowing you to identify potential areas for performance optimization and resource management.

```sql+postgres
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_instance_metric_cpu_utilization_hourly
order by
  instance_id,
  timestamp;
```

```sql+sqlite
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_instance_metric_cpu_utilization_hourly
order by
  instance_id,
  timestamp;
```

### CPU Over 80% average
Determine the areas in which the average CPU utilization exceeds 80% for hourly intervals. This query is useful for identifying potential performance issues or bottlenecks in your system.

```sql+postgres
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_cpu,
  round(maximum::numeric,2) as max_cpu,
  round(average::numeric,2) as avg_cpu
from
  alicloud_ecs_instance_metric_cpu_utilization_hourly
where average > 80
order by
  instance_id,
  timestamp;
```

```sql+sqlite
select
  instance_id,
  timestamp,
  round(minimum,2) as min_cpu,
  round(maximum,2) as max_cpu,
  round(average,2) as avg_cpu
from
  alicloud_ecs_instance_metric_cpu_utilization_hourly
where average > 80
order by
  instance_id,
  timestamp;
```