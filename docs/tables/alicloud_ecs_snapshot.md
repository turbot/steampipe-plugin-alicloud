# Table: alicloud_ecs_snapshot

The Alibaba Cloud snapshot service allows you to create crash-consistent snapshots for all disk categories. Crash-consistent snapshots are an effective solution to disaster recovery and are used to back up data, create custom images, and implement disaster recovery for applications.

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