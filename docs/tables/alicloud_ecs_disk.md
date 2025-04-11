---
title: "Steampipe Table: alicloud_ecs_disk - Query Alibaba Cloud Elastic Compute Service Disks using SQL"
description: "Allows users to query Elastic Compute Service Disks in Alibaba Cloud, specifically the disk details such as its status, type, and size, providing insights into disk usage and management."
folder: "ECS"
---

# Table: alicloud_ecs_disk - Query Alibaba Cloud Elastic Compute Service Disks using SQL

The Alibaba Cloud Elastic Compute Service (ECS) provides scalable, on-demand cloud servers for secure, flexible, and efficient application environments. ECS supports both Linux and Windows OS and offers a variety of instance types optimized to fit different workloads and scenarios. It allows users to manage the underlying physical resources while focusing on developing and deploying applications.

## Table Usage Guide

The `alicloud_ecs_disk` table provides insights into the Elastic Compute Service Disks within Alibaba Cloud. As a system administrator, explore disk-specific details through this table, including status, type, and size. Utilize it to uncover information about disks, such as those with high usage, the types of disks in use, and the verification of disk sizes.

## Examples

### Basic info
Explore which Alibaba Cloud Elastic Compute Service (ECS) disks are being used, their sizes, types, and billing methods. This information can help you understand your usage and costs, and make informed decisions about resource allocation and budgeting.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which there are unencrypted disks in your Alicloud ECS instances. This can help in identifying potential security risks and ensuring that all data storage devices comply with encryption standards.
```sql+postgres
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

```sql+sqlite
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
  encrypted = 0;
```

### List of disks Encrypted with Default Service CMK

 ```sql+postgres
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
This query is used to explore the details of auto snapshot policies applied to encrypted disks in the Alibaba Cloud ECS service. It helps in managing and understanding the security aspects of disk storage by identifying those without a specified Key Management Service (KMS) key ID.
```


```sql+sqlite
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
  encrypted = 1
  and kms_key_id = '';
```
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
Discover the segments that lack an assigned 'owner' in the disk resource data, allowing for prompt identification and rectification of unassigned resources.

```sql+postgres
select
  name,
  disk_id,
  tags
from
  alicloud_ecs_disk
where
  tags ->> 'owner' is null;
```

```sql+sqlite
select
  name,
  disk_id,
  tags
from
  alicloud_ecs_disk
where
  json_extract(tags, '$.owner') is null;
```

### List disks attached to a specific instance
Determine the specifics of disks attached to a particular instance, such as their size, type, billing method, and encryption status. This can be useful to assess storage usage, cost implications, and security measures.
```sql+postgres
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

```sql+sqlite
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
Identify the disks that are currently not in use within your system. This can help manage resources more effectively by highlighting potential areas for storage optimization.

```sql+postgres
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

```sql+sqlite
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
Uncover the details of disk distribution across different availability zones to optimize resource allocation and balance load. This can be beneficial for improving system performance and resilience.

```sql+postgres
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

```sql+sqlite
select
  zone,
  count(*)
from
  alicloud_ecs_disk
group by
  zone
order by
  count(*) desc;
```


### Top 10 largest Disks
Analyze your cloud storage to pinpoint the specific locations where the ten largest disks are in use. This is particularly useful for managing storage resources and planning for capacity upgrades.
```sql+postgres
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

```sql+sqlite
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
Determine the areas in which disks are not being utilized by identifying instances where disks are not attached to any running instances. This query is useful in pinpointing potential resources that may be underutilized or misallocated, helping optimize resource allocation and potentially reducing costs.

```sql+postgres
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

```sql+sqlite
select
  i.instance_id as "Instance ID",
  i.name as "Name",
  i.arn as "Instance ARN",
  i.status as "Instance State",
  json_extract(attachment.value, '$.AttachedTime') as "Attachment Time"
from
  alicloud_ecs_disk as v,
  json_each(attachments) as attachment,
  alicloud_ecs_instance as i
where
  i.instance_id = json_extract(attachment.value, '$.InstanceId')
  and i.status <> 'Running'
order by
  i.instance_id;
```