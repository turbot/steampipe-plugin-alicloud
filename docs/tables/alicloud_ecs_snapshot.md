---
title: "Steampipe Table: alicloud_ecs_snapshot - Query Alibaba Cloud ECS Snapshots using SQL"
description: "Allows users to query Alibaba Cloud ECS Snapshots, specifically the snapshot details, providing insights into snapshot usage and potential anomalies."
folder: "ECS"
---

# Table: alicloud_ecs_snapshot - Query Alibaba Cloud ECS Snapshots using SQL

Elastic Compute Service (ECS) Snapshots in Alibaba Cloud are a point-in-time copy of ECS disk data. Snapshots are used for data backup and restoration, disaster recovery, and migration across regions and zones. They provide a cost-effective and efficient way to create copies of data at a specific point in time.

## Table Usage Guide

The `alicloud_ecs_snapshot` table provides insights into ECS Snapshots within Alibaba Cloud Elastic Compute Service (ECS). As a DevOps engineer, explore snapshot-specific details through this table, including snapshot status, creation time, and associated metadata. Utilize it to uncover information about snapshots, such as those that are unused, the relationships between snapshots and disks, and the verification of snapshot policies.

## Examples

### List of snapshots which are not encrypted
Determine the areas in which snapshots lack encryption, allowing you to enhance your system's security by identifying potential vulnerabilities.

```sql+postgres
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

```sql+sqlite
select
  name,
  snapshot_id,
  arn,
  encrypted
from
  alicloud_ecs_snapshot
where
  encrypted = 0;
```

### List of unused snapshots
Discover the snapshots which are currently not in use within your Alicloud Elastic Compute Service. This can help in managing resources efficiently by identifying and removing unused elements.

```sql+postgres
select
  name,
  snapshot_id,
  type
from
  alicloud_ecs_snapshot
where
  usage = 'none';
```

```sql+sqlite
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
Uncover the details of how many snapshots each disk holds in your Alicloud ECS environment. This is useful in understanding the frequency of snapshots taken and can aid in storage management and cost optimization.

```sql+postgres
select
  source_disk_id,
  count(*) as snapshot
from
  alicloud_ecs_snapshot
group by
  source_disk_id;
```

```sql+sqlite
select
  source_disk_id,
  count(*) as snapshot
from
  alicloud_ecs_snapshot
group by
  source_disk_id;
```

### List of snapshots without owner tag key
Discover the segments that consist of snapshots lacking an 'owner' tag. This is particularly useful for identifying untagged resources that may lead to management issues or unnecessary costs.

```sql+postgres
select
  name,
  snapshot_id,
  tags
from
  alicloud_ecs_snapshot
where
  tags ->> 'owner' is null;
```

```sql+sqlite
select
  name,
  snapshot_id,
  tags
from
  alicloud_ecs_snapshot
where
  json_extract(tags, '$.owner') is null;
```

### List of snapshots older than 90 days
Determine the areas in which snapshots are older than 90 days in order to identify potential areas for data cleanup or archival. This can help optimize storage use and manage costs effectively.

```sql+postgres
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

```sql+sqlite
select
  name,
  snapshot_id,
  type,
  creation_time,
  julianday('now') - julianday(creation_time) as age,
  retention_days
from
  alicloud_ecs_snapshot
where
  julianday(creation_time) <= julianday(date('now','-90 day'))
order by
  creation_time;
```