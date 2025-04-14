---
title: "Steampipe Table: alicloud_ecs_disk_metric_read_iops - Query Alibaba Cloud ECS Disk Metrics using SQL"
description: "Allows users to query ECS Disk Metrics in Alibaba Cloud, specifically the read IOPS (Input/Output Operations Per Second), providing insights into disk performance and potential bottlenecks."
folder: "ECS"
---

# Table: alicloud_ecs_disk_metric_read_iops - Query Alibaba Cloud ECS Disk Metrics using SQL

Alibaba Cloud Elastic Compute Service (ECS) is a high-performance, stable, reliable, and scalable IaaS-level service provided by Alibaba Cloud. ECS eliminates the need to invest in IT hardware up front and allows you to quickly scale computing resources on demand, making ECS more convenient and efficient than physical servers. ECS provides a variety of instance types that suit different business needs and help boost business growth.

## Table Usage Guide

The `alicloud_ecs_disk_metric_read_iops` table provides insights into the read IOPS of disks within Alibaba Cloud Elastic Compute Service (ECS). As a system administrator, explore disk-specific details through this table, including performance metrics, potential bottlenecks, and associated metadata. Utilize it to uncover information about disk performance, such as those with high read IOPS, and the verification of disk performance policies.

## Examples

### Basic info
Explore the performance of various instances over time by analyzing their minimum, maximum, and average read IOPS (Input/Output Operations Per Second). This can help in identifying instances that may need optimization or troubleshooting.

```sql+postgres
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_disk_metric_read_iops
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
  alicloud_ecs_disk_metric_read_iops
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average read iops
Pinpoint specific intervals where the average read operations per second (iops) on your Alicloud Elastic Compute Service (ECS) disk exceed 1000. This helps in identifying potential performance bottlenecks and planning for capacity upgrades.

```sql+postgres
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_read_iops
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
  alicloud_ecs_disk_metric_read_iops
where average > 1000
order by
  instance_id,
  timestamp;
```