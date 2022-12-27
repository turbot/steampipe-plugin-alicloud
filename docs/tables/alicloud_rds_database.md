# Table: alicloud_rds_database

Alibaba Cloud ApsaraDB for RDS (Relational Database Service) is a stable and reliable online database service that scales elastically.

## Examples

### Basic info

```sql
select
  db_name,
  db_instance_id,
  db_status,
  engine
from
  alicloud_rds_database;
```

### Count databases per instance

```sql
select
  db_instance_id,
  count("db_name") as database_count
from
  alicloud_rds_database
group by
  db_instance_id;
```

### List databases of engine MySQL

```sql
select
  db_name,
  db_instance_id,
  db_status,
  engine
from
  alicloud_rds_database
where
  engine = 'MySQL';
```

### Get DB instance details for each database

```sql
select
  d.db_name,
  d.db_instance_id,
  i.vpc_id,
  i.creation_time,
  i.engine,
  i.db_instance_net_type
from
  alicloud_rds_database as d,
  alicloud_rds_instance as i
where
  d.db_instance_id = i.db_instance_id;
```

### List databases that are not running

```sql
select
  db_name,
  db_instance_id,
  db_status
from
  alicloud_rds_instance
where
  db_status <> 'Running';
```