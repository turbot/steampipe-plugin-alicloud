---
title: "Steampipe Table: alicloud_ecs_image - Query Alibaba Cloud Elastic Compute Service Images using SQL"
description: "Allows users to query Elastic Compute Service Images in Alibaba Cloud, specifically the image details, providing insights into image configurations and usage."
folder: "ECS"
---

# Table: alicloud_ecs_image - Query Alibaba Cloud Elastic Compute Service Images using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides fast memory and the latest Intel CPUs to help you to power your cloud applications and achieve faster results with low latency. All ECS instances come with Anti-DDoS protection to safeguard your data and applications from DDoS and Trojan attacks. ECS offers a variety of instance types optimized to fit different use cases and provides the flexibility to choose the appropriate mix of resources for your applications.

## Table Usage Guide

The `alicloud_ecs_image` table provides insights into the Elastic Compute Service Images within Alibaba Cloud. As a Cloud Engineer, explore image-specific details through this table, including image configurations, usage, and associated metadata. Utilize it to uncover information about images, such as those with specific configurations, the relationships between images and instances, and the verification of image usage.

## Examples

### Image basic info
Explore the basic details of your images in the Alibaba Cloud ECS service. This query helps you manage your resources by providing insights into image names, sizes, statuses, and usage patterns.

```sql+postgres
select
  name,
  image_id,
  arn,
  size,
  status,
  usage
from
  alicloud_ecs_image;
```

```sql+sqlite
select
  name,
  image_id,
  arn,
  size,
  status,
  usage
from
  alicloud_ecs_image;
```

### List of public images
Discover the segments that are using system-owned images in your environment. This is useful for understanding what resources are publicly available and potentially vulnerable.

```sql+postgres
select
  name,
  image_id,
  image_owner_alias
from
  alicloud_ecs_image
where
  image_owner_alias = 'system';
```

```sql+sqlite
select
  name,
  image_id,
  image_owner_alias
from
  alicloud_ecs_image
where
  image_owner_alias = 'system';
```

### List of custom (user-defined) images defined in this account
Explore the custom images created within your account to gain insights into their details such as size, status, architecture, and operating system. This can be beneficial in managing resources and ensuring optimal usage.

```sql+postgres
select
  name,
  image_id,
  image_owner_alias,
  is_self_shared,
  size,
  status,
  architecture,
  os_name_en
from
  alicloud_ecs_image
where
  image_owner_alias = 'self';
```

```sql+sqlite
select
  name,
  image_id,
  image_owner_alias,
  is_self_shared,
  size,
  status,
  architecture,
  os_name_en
from
  alicloud_ecs_image
where
  image_owner_alias = 'self';
```

### List of user-defined images which do not have owner tag key
Discover the user-defined images that lack an owner tag key. This can be useful in managing and organizing your images, ensuring that each has an owner associated with it for accountability and easy tracking.

```sql+postgres
select
  name,
  image_id,
  tags
from
  alicloud_ecs_image
where
  tags -> 'owner' is null
  and image_owner_alias = 'self';
```

```sql+sqlite
select
  name,
  image_id,
  tags
from
  alicloud_ecs_image
where
  json_extract(tags, '$.owner') is null
  and image_owner_alias = 'self';
```

### List of available images older than 90 days
Discover the segments that contain images available for over 90 days. This query is beneficial in managing resources by identifying older, potentially unused images that may be taking up unnecessary storage space.

```sql+postgres
select
  name,
  image_id,
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

```sql+sqlite
select
  name,
  image_id,
  creation_time,
  julianday('now') - julianday(creation_time) as age,
  status
from
  alicloud_ecs_image
where
  julianday(creation_time) <= julianday(date('now','-90 day'))
  and status = 'Available'
order by
  creation_time;
```

### List of unused images
Discover the segments that contain unused images in your Alicloud ECS, allowing you to manage your resources more efficiently by identifying and removing unnecessary data.

```sql+postgres
select
  name,
  image_id
from
  alicloud_ecs_image
where
  usage = 'none';
```

```sql+sqlite
select
  name,
  image_id
from
  alicloud_ecs_image
where
  usage = 'none';
```

### List of instances created from an image older than 90 days old
Identify instances that were created from an image that's over 90 days old. This can be useful for maintaining up-to-date system images and ensuring that instances are not running on outdated software.

```sql+postgres
select
  instance.name as instance_name,
  instance.instance_id as instance_id,
  image.name as image_name,
  instance.status as instance_status,
  age(image.creation_time) as image_age,
  age(instance.creation_time) as instance_age
from
  alicloud_ecs_image as image,
  alicloud_ecs_instance as instance
where
  instance.image_id = image.image_id
  and image.creation_time <= (current_date - interval '90' day)
```

```sql+sqlite
select
  instance.name as instance_name,
  instance.instance_id as instance_id,
  image.name as image_name,
  instance.status as instance_status,
  julianday('now') - julianday(image.creation_time) as image_age,
  julianday('now') - julianday(instance.creation_time) as instance_age
from
  alicloud_ecs_image as image,
  alicloud_ecs_instance as instance
where
  instance.image_id = image.image_id
  and date(image.creation_time) <= date(julianday('now','-90 day'))
```

### Get details of a particular image
Determine the specifics of a particular image, such as its size and status, within a specified region. This is useful for assessing the usage and availability of an image in a specific geographical location.

```sql+postgres
select
  name,
  image_id,
  arn,
  size,
  status,
  usage
from
  alicloud_ecs_image
where
  image_id = 'centos_7_9_x64_20G_alibase_20221129.vhd'
  and region = 'us-east-1';
```

```sql+sqlite
select
  name,
  image_id,
  arn,
  size,
  status,
  usage
from
  alicloud_ecs_image
where
  image_id = 'centos_7_9_x64_20G_alibase_20221129.vhd'
  and region = 'us-east-1';
```