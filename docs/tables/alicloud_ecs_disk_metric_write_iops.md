---
title: "Steampipe Table: alicloud_ecs_disk_metric_write_iops - Query Alibaba Cloud ECS Disk Metrics using SQL"
description: "Allows users to query Alibaba Cloud Elastic Compute Service (ECS) Disk Metrics, specifically the write input/output operations per second (IOPS), providing insights into disk performance and potential bottlenecks."
folder: "ECS"
---

# Table: alicloud_ecs_disk_metric_write_iops - Query Alibaba Cloud ECS Disk Metrics using SQL

Alibaba Cloud Elastic Compute Service (ECS) is a high-performance, stable, reliable, and scalable IaaS-level service provided by Alibaba Cloud. ECS eliminates the need to invest in IT hardware upfront and allows you to quickly scale computing resources on demand, making ECS more convenient and efficient than physical servers. ECS provides a variety of instance types that suit different business needs and help boost business growth.

## Table Usage Guide

The `alicloud_ecs_disk_metric_write_iops` table provides insights into the write operations performance of disks within Alibaba Cloud Elastic Compute Service (ECS). As a system administrator or a DevOps engineer, explore disk-specific details through this table, including the write input/output operations per second (IOPS). Utilize it to uncover information about disk performance, such as potential bottlenecks, and to ensure optimal resource allocation and performance tuning.

## Examples

### Basic info
Analyze the write operations per second on your Alibaba Cloud Elastic Compute Service (ECS) disks to understand their performance over time. This can help in identifying any potential bottlenecks or performance issues.

```sql+postgres
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_disk_metric_write_iops
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
  alicloud_ecs_disk_metric_write_iops
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average write iops
Determine the instances and times when the average write operations per second exceeded 1000. This can be useful for identifying periods of high disk activity and potential performance issues.

```sql+postgres
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_write_iops
where average > 1000
order by
  instance_id,
  timestamp;
```

```sql+sqlite
select
  instance_id,
  timestamp,
  round(minimum,2) as min_ops,
  round(maximum,2) as max_ops,
  round(average,2) as avg_ops
from
  alicloud_ecs_disk_metric_write_iops
where average > 1000
order by
  instance_id,
  timestamp;
```