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

```sql
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

```sql
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

```sql
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

```sql
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

```sql
select
  db_instance_id,
  security_ips
from
  alicloud_rds_instance
where
  security_ips :: jsonb ? '0.0.0.0/0';
```

### List DB instances with SSL encryption disabled

```sql
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

```sql
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

```sql
select
  i.db_instance_id,
  s ->> 'NetworkType' as network_type,
  s ->> 'RegionId' as security_group_region_id,
  s ->> 'SecurityGroupId' as security_group_id
from
  alicloud_rds_instance as i,
  jsonb_array_elements(security_group_configuration) as s;
```

### Get encryption details for all the instances

```sql
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
