---
title: "Steampipe Table: alicloud_ecs_disk_metric_write_iops_hourly - Query Alibaba Cloud ECS Disk Metrics using SQL"
description: "Allows users to query ECS Disk Metrics in Alibaba Cloud, specifically the hourly write IOPS (input/output operations per second), providing insights into disk performance and potential issues."
---

# Table: alicloud_ecs_disk_metric_write_iops_hourly - Query Alibaba Cloud ECS Disk Metrics using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides scalable, on-demand computing resources for secure, flexible, and efficient applications. ECS Disk Metrics is a feature within ECS that allows you to monitor and analyze disk performance and usage. It provides data such as read and write IOPS, throughput, and latency, which can be useful for capacity planning and troubleshooting.

## Table Usage Guide

The `alicloud_ecs_disk_metric_write_iops_hourly` table provides insights into the hourly write IOPS of ECS disks in Alibaba Cloud. As a system administrator or a DevOps engineer, explore disk-specific details through this table, including write IOPS, which can be quite useful for performance tuning, capacity planning, and troubleshooting. Utilize it to uncover information about disk performance, such as identifying disks with high write operations, and the verification of disk usage patterns.

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
  alicloud_ecs_disk_metric_write_iops_hourly
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average write iops

```sql
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_write_iops_hourly
where average > 1000
order by
  instance_id,
  timestamp;
```