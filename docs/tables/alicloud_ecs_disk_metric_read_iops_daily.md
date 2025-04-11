---
title: "Steampipe Table: alicloud_ecs_disk_metric_read_iops_daily - Query Daily Read IOPS Metrics for ECS Disks using SQL"
description: "Allows users to query daily average Read IOPS metrics for Alibaba Cloud ECS disks, including disk ID, instance ID, timestamp, and read IOPS values."
folder: "ECS"
---

# Table: alicloud_ecs_disk_metric_read_iops_daily - Query Daily Read IOPS Metrics for ECS Disks using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides persistent block storage through ECS disks, which are critical for storing data and running workloads. Monitoring Read IOPS (Input/Output Operations Per Second) helps assess disk performance and identify potential bottlenecks.

## Table Usage Guide

The `alicloud_ecs_disk_metric_read_iops_daily` table allows cloud engineers, performance analysts, and DevOps teams to query daily average Read IOPS metrics for ECS disks. Use this table to retrieve details such as disk ID, associated instance ID, timestamp, and average read IOPS. This information is useful for performance monitoring, identifying underperforming disks, and optimizing resource allocation based on workload demands.

Note: If the instance is not older than one day then we will not get any metric statistics.

## Examples

### Basic info
Gain insights into daily disk read operations on your Alibaba Cloud ECS instances. This can help you monitor performance trends and identify potential issues related to disk input/output operations.

```sql+postgres
select
  instance_id,
  timestamp,
  minimum,
  maximum,
  average
from
  alicloud_ecs_disk_metric_read_iops_daily
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
  alicloud_ecs_disk_metric_read_iops_daily
order by
  instance_id,
  timestamp;
```

### Intervals where operation exceed 1000 average read iops
Explore instances where the daily average read operations per second (IOPS) on an Alibaba Cloud Elastic Compute Service (ECS) disk exceed 1000. This can be useful in identifying periods of high disk usage, which could impact system performance.

```sql+postgres
select
  instance_id,
  timestamp,
  round(minimum::numeric,2) as min_ops,
  round(maximum::numeric,2) as max_ops,
  round(average::numeric,2) as avg_ops
from
  alicloud_ecs_disk_metric_read_iops_daily
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
  alicloud_ecs_disk_metric_read_iops_daily
where average > 1000
order by
  instance_id,
  timestamp;
```