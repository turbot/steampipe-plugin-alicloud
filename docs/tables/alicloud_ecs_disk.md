# Table: alicloud_ecs_disk

Cloud disks are block-level Elastic Block Storage (EBS) products provided by Alibaba Cloud for ECS and provide low latency, high performance, high durability, and high reliability. Cloud disks use a distributed triplicate mechanism to ensure data durability for ECS instances. If service disruptions occur within a zone (for example, due to hardware failures), data within the zone is copied to an available disk in another zone to help ensure data availability.

## Examples

### Basic info

```sql
select
  name,
  disk_id,
  arn,
  size,
  type,
  billing_method,
  zone,
  region
from
  alicloud_ecs_disk;
```

### Unencrypted Disks
```sql
select
  name,
  disk_id,
  encrypted,
  zone,
  status,
  size,
  instance_id,
  kms_key_id
from
  alicloud_ecs_disk
where
  not encrypted;
```

### List of disks Encrypted with Default Service CMK

 ```sql
 select
   name,
   disk_id,
  encrypted,
   zone,
  status,
  size,
  instance_id,
   kms_key_id
 from
  alicloud_ecs_disk
where
  encrypted
  and kms_key_id = '';
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
  disk_id,
  tags
from
  alicloud_ecs_disk
where
  tags ->> 'owner' is null;
```

### List disks attached to a specific instance
```sql
select
  name,
  disk_id,
  size,
  type,
  billing_method,
  zone,
  region,
  encrypted
from
  alicloud_ecs_disk
where
  instance_id = 'i-0xickpvpsaih9w7s4zrq';
```


### List of disks not attached to any instances

```sql
select
  name,
  disk_id,
  status,
  attached_time,
  detached_time
from
  alicloud_ecs_disk
where
  status = 'Available';
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


### Top 10 largest Disks
```sql
select
  name,
  disk_id,
  size,
  status,
  instance_id
from
  alicloud_ecs_disk
order by
  size desc
limit 10;
```

### List of disks having no attached running instances

```sql
select
  i.instance_id as "Instance ID",
  i.name as "Name",
  i.arn as "Instance ARN",
  i.status as "Instance State",
  attachment ->> 'AttachedTime' as "Attachment Time"
from
  alicloud_ecs_disk as v,
  jsonb_array_elements(attachments) as attachment,
  alicloud_ecs_instance as i
where
  i.instance_id = attachment ->> 'InstanceId'
  and i.status <> 'Running'
order by
  i.instance_id;
```