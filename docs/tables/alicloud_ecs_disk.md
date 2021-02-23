# Table: alicloud_ecs_disk

Cloud disks are block-level Elastic Block Storage (EBS) products provided by Alibaba Cloud for ECS and provide low latency, high performance, high durability, and high reliability. Cloud disks use a distributed triplicate mechanism to ensure data durability for ECS instances. If service disruptions occur within a zone (for example, due to hardware failures), data within the zone is copied to an available disk in another zone to help ensure data availability.

## Examples

### Basic info

```sql
select
  name,
  id,
  size,
  type,
  billing_method,
  zone,
  region
from
  alicloud_ecs_disk;
```

### List of disks with Default Service CMK

```sql
select
  name,
  id,
  zone,
  kms_key_id
from
  alicloud_ecs_disk
where
  kms_key_id = '';
```

### List Auto Snapshot Policy details applied to disk

```sql
select
  name,
  auto_snapshot_policy_id,
  auto_snapshot_policy_name,
  auto_snapshot_policy_creation_time,
  auto_snapshot_policy_enable_cross_region_copy,
  auto_snapshot_policy_repeat_week_days,
  auto_snapshot_policy_retention_days,
  auto_snapshot_policy_status,
  auto_snapshot_policy_time_points,
  auto_snapshot_policy_tags
from
  alicloud_ecs_disk;
```

### List of disks without owner tag key

```sql
select
  name,
  id,
  tags
from
  alicloud_ecs_disk
where
  not tags :: JSONB ? 'owner';
```

### List of disks not attached with any instances

```sql
select
  name,
  id,
  attached_time
from
  alicloud_ecs_disk
where
  attached_time is null;
```

### Disk count in each availability zone

```sql
select
  zone,
  count(*)
from
  alicloud_ecs_disk
group by
  zone
order by
  count desc;
```
