---
title: "Steampipe Table: alicloud_ecs_instance - Query Alibaba Cloud ECS Instances using SQL"
description: "Allows users to query Alibaba Cloud ECS Instances, including instance ID, name, status, type, region, zone, and associated network and security details."
folder: "ECS"
---

# Table: alicloud_ecs_instance - Query Alibaba Cloud ECS Instances using SQL

Alibaba Cloud Elastic Compute Service (ECS) delivers scalable virtual servers that provide a secure, high-performance computing environment. ECS instances support a variety of workloads and offer flexible configurations for CPU, memory, storage, and networking to suit a wide range of application needs.

## Table Usage Guide

The `alicloud_ecs_instance` table allows system administrators, DevOps engineers, and security teams to query detailed information about ECS instances within Alibaba Cloud. Use this table to retrieve attributes such as instance ID, name, status, instance type, creation time, region, zone, VPC and subnet associations, public and private IP addresses, and security group configurations. This information is essential for managing your compute resources, tracking utilization, enforcing security policies, and optimizing your cloud environment.

## Examples

### Basic Instance Info
Assess the elements within your Alibaba Cloud ECS instances to gain insights into their operational status, type, and network details. This can help in managing resources, ensuring optimal performance, and identifying potential issues.
```sql+postgres
select
  instance_id,
  name,
  arn,
  status,
  instance_type,
  os_name_en,
  public_ip_address,
  private_ip_address,
  zone
from
  alicloud_ecs_instance;
```

```sql+sqlite
select
  instance_id,
  name,
  arn,
  status,
  instance_type,
  os_name_en,
  public_ip_address,
  private_ip_address,
  zone
from
  alicloud_ecs_instance;
```


### List stopped instances that you are still being charged for
This query is useful for identifying instances that are in a stopped state but still incurring charges. It helps in managing costs by pinpointing areas where resources are not being optimally used.
```sql+postgres
select
  instance_id,
  name,
  status,
  stopped_mode,
  instance_type,
  os_name_en,
  public_ip_address,
  private_ip_address,
  zone
from
  alicloud_ecs_instance
where
  stopped_mode = 'KeepCharging';
```

```sql+sqlite
select
  instance_id,
  name,
  status,
  stopped_mode,
  instance_type,
  os_name_en,
  public_ip_address,
  private_ip_address,
  zone
from
  alicloud_ecs_instance
where
  stopped_mode = 'KeepCharging';
```


### List linux instances
Identify instances where the operating system is Linux. This is useful for managing resources or troubleshooting specific issues related to Linux-based instances.
```sql+postgres
select
  instance_id,
  name,
  instance_type,
  os_name_en,
  zone
from
  alicloud_ecs_instance
where
  os_type = 'linux';
```

```sql+sqlite
select
  instance_id,
  name,
  instance_type,
  os_name_en,
  zone
from
  alicloud_ecs_instance
where
  os_type = 'linux';
```


### Instance count in each zone
Determine the distribution of instances across various zones to balance and optimize resource allocation. This aids in planning infrastructure decisions and mitigating risk by avoiding over-reliance on a single zone.

```sql+postgres
select
  zone as az,
  count(*)
from
  alicloud_ecs_instance
group by
  zone;
```

```sql+sqlite
select
  zone as az,
  count(*)
from
  alicloud_ecs_instance
group by
  zone;
```

### Count the number of instances by instance type
Assess the distribution of different instance types within your Alicloud Elastic Compute Service to better understand your resource usage. This can aid in optimizing your allocation strategy and managing costs more effectively.

```sql+postgres
select
  instance_type,
  count(instance_type) as count
from
  alicloud_ecs_instance
group by
  instance_type;
```

```sql+sqlite
select
  instance_type,
  count(instance_type) as count
from
  alicloud_ecs_instance
group by
  instance_type;
```

### List of instances without application tag key
Identify instances where the 'application' tag key is missing, which could indicate a lack of necessary metadata for proper resource management and classification. This can be crucial for maintaining organized and efficient infrastructure in a cloud environment.

```sql+postgres
select
  instance_id,
  tags
from
  alicloud_ecs_instance
where
  tags ->> 'application' is null;
```

```sql+sqlite
select
  instance_id,
  tags
from
  alicloud_ecs_instance
where
  json_extract(tags, '$.application') is null;
```

### List of ECS instances provisioned with undesired(for example ecs.t5-lc2m1.nano and ecs.t6-c2m1.large is desired) instance type(s)
Identify instances where ECS instances are provisioned with undesired types, helping to manage resources and maintain preferred configurations.

```sql+postgres
select
  instance_type,
  count(*) as count
from
  alicloud_ecs_instance
where
  instance_type not in ('ecs.t5-lc2m1.nano', 'ecs.t6-c2m1.large')
group by
  instance_type;
```

```sql+sqlite
select
  instance_type,
  count(*) as count
from
  alicloud_ecs_instance
where
  instance_type not in ('ecs.t5-lc2m1.nano', 'ecs.t6-c2m1.large')
group by
  instance_type;
```

### List ECS instances having deletion protection safety feature disabled
Determine the areas in which Elastic Compute Service (ECS) instances might be at risk due to disabled deletion protection. This query is useful to identify potential vulnerabilities and ensure data safety.

```sql+postgres
select
  instance_id,
  deletion_protection
from
  alicloud_ecs_instance
where
  not deletion_protection;
```

```sql+sqlite
select
  instance_id,
  deletion_protection
from
  alicloud_ecs_instance
where
  deletion_protection = 0;
```