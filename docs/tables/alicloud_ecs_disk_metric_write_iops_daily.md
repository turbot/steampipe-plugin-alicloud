---
title: "Steampipe Table: alicloud_ecs_disk_metric_write_iops_daily - Query AliCloud ECS Disk Metrics using SQL"
description: "Allows users to query ECS Disk Metrics in AliCloud, specifically the daily write IOPS (Input/Output Operations Per Second), providing insights into disk performance and potential bottlenecks."
---

# Table: alicloud_ecs_disk_metric_write_iops_daily - Query AliCloud ECS Disk Metrics using SQL

AliCloud Elastic Compute Service (ECS) provides scalable, on-demand cloud servers for secure, flexible, and efficient application environments. ECS Disk Metrics provide detailed performance metrics for ECS disks, including write IOPS, which measures the number of write operations to a disk in a second. These metrics can be used to monitor the performance of ECS disks and identify potential issues.

## Table Usage Guide

The `alicloud_ecs_disk_metric_write_iops_daily` table provides insights into the daily write performance of ECS disks in AliCloud. As a system administrator or DevOps engineer, explore disk-specific details through this table, including daily write IOPS, to monitor and optimize disk performance. Utilize it to uncover information about disk usage patterns, identify potential bottlenecks, and ensure optimal resource allocation.

## Examples

### Basic info
Explore the daily write operations per second (IOPS) on your Alibaba Cloud ECS disks. This can help you understand the performance trends and identify any potential bottlenecks or high usage periods.

```sql+postgres
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_disk_metric_write_iops_daily
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
  alicloud_ecs_disk_metric_write_iops_daily
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average write iops
Explore instances where the average write operations per second on your Alibaba Cloud ECS disk exceed 1000. This can help you identify potential performance issues and manage your resource usage effectively.

```sql+postgres
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_write_iops_daily
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
  alicloud_ecs_disk_metric_write_iops_daily
where average > 1000
order by
  instance_id,
  timestamp;
```