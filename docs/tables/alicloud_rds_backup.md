---
title: "Steampipe Table: alicloud_rds_backup - Query Alicloud RDS Backups using SQL"
description: "Allows users to query Alicloud RDS Backups, specifically details about each backup such as backup status, type, method, and more, providing insights into the backup strategy and potential risks."
---

# Table: alicloud_rds_backup - Query Alicloud RDS Backups using SQL

Alicloud RDS Backup is a feature of the Alicloud RDS service which allows users to create backups of their RDS instances. These backups can be used to restore a database instance to a previous state, ensuring data security and continuity. The backups can be created manually or scheduled according to the user's needs.

## Table Usage Guide

The `alicloud_rds_backup` table provides insights into the backups of RDS instances within Alicloud. As a Database Administrator, you can explore backup-specific details through this table, including backup status, type, method, and more. Utilize it to understand your current backup strategy, identify potential risks, and ensure the security and continuity of your data.

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

### List backups that were created in the last 30 days

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