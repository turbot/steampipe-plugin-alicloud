---
title: "Steampipe Table: alicloud_ecs_disk_metric_read_iops_hourly - Query Alibaba Cloud ECS Disk Metrics using SQL"
description: "Allows users to query ECS Disk Metrics in Alibaba Cloud, specifically the hourly read IOPS (Input/Output Operations Per Second), providing insights into disk performance and potential bottlenecks."
---

# Table: alicloud_ecs_disk_metric_read_iops_hourly - Query Alibaba Cloud ECS Disk Metrics using SQL

Elastic Compute Service (ECS) Disks in Alibaba Cloud are block-level storage devices that can be attached to ECS instances. These disks provide persistent block storage capacity and are designed for high performance and low latency. Metrics related to ECS Disks, such as read IOPS, can provide valuable insights into disk performance and usage patterns.

## Table Usage Guide

The `alicloud_ecs_disk_metric_read_iops_hourly` table provides insights into the hourly read IOPS of ECS Disks within Alibaba Cloud Elastic Compute Service. As a system administrator or DevOps engineer, explore disk-specific details through this table, including the read IOPS, which can indicate the performance of the disk and identify potential bottlenecks. Utilize it to monitor and optimize disk performance, ensuring efficient operation of your Alibaba Cloud ECS instances.

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
  alicloud_ecs_disk_metric_read_iops_hourly
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
  alicloud_ecs_disk_metric_read_iops_hourly
where average > 1000
order by
  instance_id,
  timestamp;
```