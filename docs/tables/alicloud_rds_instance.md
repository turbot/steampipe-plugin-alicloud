# Table: alicloud_rds_instance

Provides an RDS instance resource. A DB instance is an isolated database environment in the cloud. A DB instance can contain multiple user-created databases.

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
  security_ips :: jsonb ? '0.0.0.0/0'
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
  left join alicloud_kms_key k 
    on encryption_key = key_id
```