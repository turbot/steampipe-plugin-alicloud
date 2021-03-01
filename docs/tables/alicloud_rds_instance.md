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


