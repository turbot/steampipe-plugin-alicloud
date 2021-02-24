# Table: alicloud_ecs_image

An ECS image stores information that is required to create an ECS instance. An image works as a copy that stores data from one or more disks. An ECS instance image may store data from a system disk or from both system and data disks.

## Examples

### Image basic info

```sql
select
  name,
  id,
  size,
  status,
  usage
from
  alicloud_ecs_image;
```

### List of public images

```sql
select
  name,
  id,
  image_owner_alias
from
  alicloud_ecs_image
where
  image_owner_alias = 'system';
```

### List of custom (user-defined) images defined in this account

```sql
select
  name,
  id,
  image_owner_alias
from
  alicloud_ecs_image
where
  image_owner_alias = 'self';
```

### List of user-defined images which do not have owner tag key

```sql
select
  name,
  id
from
  alicloud_ecs_image
where
  tags -> 'owner' is null
  and image_owner_alias = 'self';
```

### List of available images older than 90 days

```sql
select
  name,
  id,
  creation_time,
  age(creation_time),
  status
from
  alicloud_ecs_image
where
  creation_time <= (current_date - interval '90' day)
  and status = 'Available'
order by
  creation_time;
```

### List of unused images

```sql
select
  name,
  id
from
  alicloud_ecs_image
where
  usage = 'none';
```

### List of running instances created from the image

```sql
select
  instance.name as instance_name,
  instance.instance_id as instance_id,
  image.name as image_name,
  instance.status as instance_status
from
  alicloud_ecs_image as image,
  alicloud_ecs_instance as instance
where
  instance.image_id = image.id
  and instance.status = 'Running';
```
