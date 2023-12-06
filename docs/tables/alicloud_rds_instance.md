---
title: "Steampipe Table: alicloud_rds_instance - Query Alibaba Cloud RDS Instances using SQL"
description: "Allows users to query Alibaba Cloud Relational Database Service (RDS) Instances, providing detailed information about each instance's configuration and status."
---

# Table: alicloud_rds_instance - Query Alibaba Cloud RDS Instances using SQL

Alibaba Cloud Relational Database Service (RDS) is a stable and reliable online database service that supports MySQL, SQL Server, PostgreSQL, and PPAS. RDS handles routine database tasks such as database backup, patch upgrades, and failure detection and recovery. It provides automatic monitoring, backup, and disaster recovery capabilities, freeing up developers to focus on their applications rather than managing databases.

## Table Usage Guide

The `alicloud_rds_instance` table provides insights into RDS instances within Alibaba Cloud Relational Database Service (RDS). As a database administrator, explore instance-specific details through this table, including the instance's ID, creation time, status, and associated metadata. Utilize it to uncover information about instances, such as their storage and memory usage, the network type they are using, and their security settings.

## Examples

### Basic info
Discover the segments that provide insights into the instances of database services, including when they were created and the type of engine used. This information can be useful in assessing the overall setup of your cloud services and ensuring they align with your operational requirements.

```sql+postgres
select
  db_instance_id,
  arn,
  vpc_id,
  creation_time,
  engine
from
  alicloud_rds_instance;
```

```sql+sqlite
select
  db_instance_id,
  arn,
  vpc_id,
  creation_time,
  engine
from
  alicloud_rds_instance;
```

### DB instance count in each region
Determine the distribution of database instances across different regions to better understand resource allocation and usage patterns.

```sql+postgres
select
  region_id as region,
  db_instance_class,
  count(*)
from
  alicloud_rds_instance
group by
  region_id,
  db_instance_class;
```

```sql+sqlite
select
  region_id as region,
  db_instance_class,
  count(*)
from
  alicloud_rds_instance
group by
  region_id,
  db_instance_class;
```

### List DB instances whose engine is MySQL
Explore which database instances are running on the MySQL engine. This can be useful for assessing your infrastructure's dependencies on this particular database engine.

```sql+postgres
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine
from
  alicloud_rds_instance
where
  engine = 'MySQL';
```

```sql+sqlite
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine
from
  alicloud_rds_instance
where
  engine = 'MySQL';
```

### List DB instances that are currently running
Identify instances where your database is currently active. This is useful for monitoring system performance and ensuring resources are not being wasted on idle databases.

```sql+postgres
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine
from
  alicloud_rds_instance
where
  db_instance_status = 'Running';
```

```sql+sqlite
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine
from
  alicloud_rds_instance
where
  db_instance_status = 'Running';
```

### List DB instances that allow 0.0.0.0/0
Determine the areas in which database instances are set to allow all IP addresses. This query is useful in identifying potential security risks, as allowing all IP addresses may expose your database to unwanted access.

```sql+postgres
select
  db_instance_id,
  security_ips
from
  alicloud_rds_instance
where
  security_ips :: jsonb ? '0.0.0.0/0';
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### List DB instances with SSL encryption disabled
Identify instances where the SSL encryption is disabled on your database, allowing you to pinpoint potential security vulnerabilities and improve your database's protection against unauthorized access.

```sql+postgres
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine,
  ssl_encryption
from
  alicloud_rds_instance
where
  ssl_status = 'Disabled';
```

```sql+sqlite
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine,
  ssl_encryption
from
  alicloud_rds_instance
where
  ssl_status = 'Disabled';
```

### List DB instances with TDE disabled
Identify instances where the Transparent Data Encryption (TDE) feature is disabled within your database instances. This can be useful for enhancing data security by pinpointing potential vulnerabilities and areas that need attention.

```sql+postgres
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine,
  tde_status
from
  alicloud_rds_instance
where
  tde_status = 'Disabled';
```

```sql+sqlite
select
  db_instance_id,
  vpc_id,
  creation_time,
  engine,
  tde_status
from
  alicloud_rds_instance
where
  tde_status = 'Disabled';
```

### Get security group configuration details for instances
Explore the security configurations of your database instances to understand their network types and geographical regions. This can help in assessing your security measures and ensuring they align with your organization's standards and requirements.

```sql+postgres
select
  i.db_instance_id,
  s ->> 'NetworkType' as network_type,
  s ->> 'RegionId' as security_group_region_id,
  s ->> 'SecurityGroupId' as security_group_id
from
  alicloud_rds_instance as i,
  jsonb_array_elements(security_group_configuration) as s;
```

```sql+sqlite
select
  i.db_instance_id,
  json_extract(s.value, '$.NetworkType') as network_type,
  json_extract(s.value, '$.RegionId') as security_group_region_id,
  json_extract(s.value, '$.SecurityGroupId') as security_group_id
from
  alicloud_rds_instance as i,
  json_each(security_group_configuration) as s;
```

### Get encryption details for all the instances
Identify the encryption details for all instances to enhance your understanding of your database security. This is useful for auditing your security measures and ensuring necessary precautions are in place.

```sql+postgres
select
  i.arn as instance_arn,
  i.title as instance_name,
  encryption_key,
  k.title as kms_key_name 
from
  alicloud_rds_instance i 
  left join
    alicloud_kms_key k 
    on encryption_key = key_id;
```

```sql+sqlite
select
  i.arn as instance_arn,
  i.title as instance_name,
  encryption_key,
  k.title as kms_key_name 
from
  alicloud_rds_instance i 
  left join
    alicloud_kms_key k 
    on encryption_key = k.key_id;
```