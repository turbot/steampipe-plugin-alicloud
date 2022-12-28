# Table: alicloud_rds_backup

ApsaraDB RDS Backup is a policy expression that defines when and how you want to back up your DB Instances.

## Examples

### Basic info

```sql
select
  backup_id,
  db_instance_id,
  backup_status,
  backup_mode,
  backup_size,
  backup_start_time
from
  alicloud_rds_backup;
```

### Count backup by instance

```sql
select
  db_instance_id,
  count(backup_id) as backup_count
from
  alicloud_rds_backup
group by
  db_instance_id;
```

### List manual backups

```sql
select
  backup_id,
  db_instance_id,
  backup_status,
  backup_mode,
  backup_size,
  backup_start_time
from
  alicloud_rds_backup
where
  backup_mode = 'Manual';
```

### List backups of type incremental

```sql
select
  backup_id,
  db_instance_id,
  backup_status,
  backup_type,
  backup_size,
  backup_start_time,
  backup_end_time
from
  alicloud_rds_backup
where
  backup_type = 'IncrementalBackup';
```

### List backups by location

```sql
select
  backup_id,
  db_instance_id,
  backup_status,
  backup_type,
  backup_size,
  backup_start_time,
  backup_end_time,
  backup_location
from
  alicloud_rds_backup
where
  backup_location = 'OSS';
```

### List backups that are completed in the last 30 days

```sql
select
  backup_id,
  db_instance_id,
  backup_status,
  backup_type,
  backup_size,
  backup_start_time,
  backup_end_time,
  backup_location
from
  alicloud_rds_backup
where
  backup_end_time >= now() - interval '30' day;
```