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

### DB Instance Details for a particular DB Instance

```sql
select
  db_instance_id,
  vpc_id,
  create_time,
  engine
from
  alicloud_rds_instance where db_instance_id='****';
```


### DB Instance Details from a region

```sql
select
  db_instance_id,
  vpc_id,
  create_time,
  engine
from
  alicloud_rds_instance where region_id='us-east-1';
```


### DB Instance Details where engine is MySQL

```sql
select
  db_instance_id,
  vpc_id,
  create_time,
  engine
from
  alicloud_rds_instance where engine='MySQL';
```

### DBInstance count by Instance ID

```sql
select
  db_instance_id,
  count(db_instance_id) as db_instance_count
from
  alicloud_rds_instance
group by
  vpc_id;
```
