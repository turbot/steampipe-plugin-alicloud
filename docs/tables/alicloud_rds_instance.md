# Table: alicloud_rds_instance

Provides an RDS instance resource. A DB instance is an isolated database environment in the cloud. A DB instance can contain multiple user-created databases.

## Examples

### Basic info

```sql
select
  db_instance_id,
  vpc_id,
  create_time,
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
  create_time,
  engine
from
  alicloud_rds_instance
where
  engine='MySQL';
```


### List DB instances that are currently running

```sql
select
  db_instance_id,
  vpc_id,
  create_time,
  engine
from
  alicloud_rds_instance
where
  db_instance_status='Running';
```

### List DB instances that allow 0.0.0.0/0.

```sql
select
	db_instance_id,
	security_ips
from
	alicloud_rds_instance
where
	jsonb_path_exists(
		security_ips,
		'$.** ? (@.type() == "string" && @ like_regex "0.0.0.0/0")'
	)
```