---
title: "Steampipe Table: alicloud_ecs_instance_metric_cpu_utilization_daily - Query Alibaba Cloud ECS Instance Metrics using SQL"
description: "Allows users to query ECS Instance Metrics in Alibaba Cloud, specifically the daily CPU utilization, providing insights into instance performance and usage patterns."
---

# Table: alicloud_ecs_instance_metric_cpu_utilization_daily - Query Alibaba Cloud ECS Instance Metrics using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides scalable, on-demand cloud servers for secure, flexible, and efficient application environments. It allows you to operate and manage online applications in a smoother, more reliable, and more secure manner. ECS instances are the fundamental computing unit provided by Alibaba Cloud ECS.

## Table Usage Guide

The `alicloud_ecs_instance_metric_cpu_utilization_daily` table provides insights into ECS Instance Metrics within Alibaba Cloud Elastic Compute Service (ECS). As a system administrator or DevOps engineer, explore instance-specific details through this table, including daily CPU utilization. Utilize it to uncover information about instances, such as CPU usage patterns, which can help in performance optimization and capacity planning.

**Important Notes**
- If the instance is not older than one day then we will not get any metric statistics.

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
  alicloud_ecs_instance_metric_cpu_utilization_daily
order by
  instance_id,
  timestamp;
```

### CPU Over 80% average

```sql
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_cpu,
  round(maximum::numeric,2) as max_cpu,
  round(average::numeric,2) as avg_cpu
from
  alicloud_ecs_instance_metric_cpu_utilization_daily
where average > 80
order by
  instance_id,
  timestamp;
```
