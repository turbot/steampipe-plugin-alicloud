---
title: "Steampipe Table: alicloud_ecs_snapshot - Query Alibaba Cloud ECS Snapshots using SQL"
description: "Allows users to query Alibaba Cloud ECS Snapshots, specifically the snapshot details, providing insights into snapshot usage and potential anomalies."
---

# Table: alicloud_ecs_snapshot - Query Alibaba Cloud ECS Snapshots using SQL

Elastic Compute Service (ECS) Snapshots in Alibaba Cloud are a point-in-time copy of ECS disk data. Snapshots are used for data backup and restoration, disaster recovery, and migration across regions and zones. They provide a cost-effective and efficient way to create copies of data at a specific point in time.

## Table Usage Guide

The `alicloud_ecs_snapshot` table provides insights into ECS Snapshots within Alibaba Cloud Elastic Compute Service (ECS). As a DevOps engineer, explore snapshot-specific details through this table, including snapshot status, creation time, and associated metadata. Utilize it to uncover information about snapshots, such as those that are unused, the relationships between snapshots and disks, and the verification of snapshot policies.

## Examples

### List of snapshots which are not encrypted

```sql
select
  name,
  snapshot_id,
  arn,
  encrypted
from
  alicloud_ecs_snapshot
where
  not encrypted;
```

### List of unused snapshots

```sql
select
  name,
  snapshot_id,
  type
from
  alicloud_ecs_snapshot
where
  usage = 'none';
```

### Find the snapshot count per disk

```sql
select
  source_disk_id,
  count(*) as snapshot
from
  alicloud_ecs_snapshot
group by
  source_disk_id;
```

### List of snapshots without owner tag key

```sql
select
  name,
  snapshot_id,
  tags
from
  alicloud_ecs_snapshot
where
  tags ->> 'owner' is null;
```

### List of snapshots older than 90 days

```sql
select
  name,
  snapshot_id,
  type,
  creation_time,
  age(creation_time),
  retention_days
from
  alicloud_ecs_snapshot
where
  creation_time <= (current_date - interval '90' day)
order by
  creation_time;
```