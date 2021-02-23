# Table: alicloud_ecs_snapshot

The Alibaba Cloud snapshot service allows you to create crash-consistent snapshots for all disk categories. Crash-consistent snapshots are an effective solution to disaster recovery and are used to back up data, create custom images, and implement disaster recovery for applications.

## Examples

### List of snapshots which are not encrypted

```sql
select
  name,
  id,
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
  id,
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
  count(id) as snapshot
from
  alicloud_ecs_snapshot
group by
  source_disk_id;
```

### List of snapshots without owner tag key

```sql
select
  name,
  id,
  tags
from
  alicloud_ecs_snapshot
where
  not tags :: JSONB ? 'owner';
```
