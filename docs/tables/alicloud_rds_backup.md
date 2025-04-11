---
title: "Steampipe Table: alicloud_rds_backup - Query Alicloud RDS Backups using SQL"
description: "Allows users to query Alicloud RDS Backups, specifically details about each backup such as backup status, type, method, and more, providing insights into the backup strategy and potential risks."
folder: "RDS"
---

# Table: alicloud_rds_backup - Query Alicloud RDS Backups using SQL

Alicloud RDS Backup is a feature of the Alicloud RDS service which allows users to create backups of their RDS instances. These backups can be used to restore a database instance to a previous state, ensuring data security and continuity. The backups can be created manually or scheduled according to the user's needs.

## Table Usage Guide

The `alicloud_rds_backup` table provides insights into the backups of RDS instances within Alicloud. As a Database Administrator, you can explore backup-specific details through this table, including backup status, type, method, and more. Utilize it to understand your current backup strategy, identify potential risks, and ensure the security and continuity of your data.

## Examples

### Basic info
Explore the status and details of your database backups to ensure they are functioning as expected and to understand their impact on storage capacity. This is crucial in maintaining data safety and optimizing resource usage.

```sql+postgres
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

```sql+sqlite
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
Assess the number of backups associated with each database instance. This can be useful for managing and monitoring data redundancy and recovery processes.

```sql+postgres
select
  db_instance_id,
  count(backup_id) as backup_count
from
  alicloud_rds_backup
group by
  db_instance_id;
```

```sql+sqlite
select
  db_instance_id,
  count(backup_id) as backup_count
from
  alicloud_rds_backup
group by
  db_instance_id;
```

### List manual backups
Explore which backups have been manually created in your database. This can be useful in understanding your backup strategy and ensuring that manual backups are being created as expected.

```sql+postgres
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

```sql+sqlite
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
Explore which backups are incremental in nature, allowing you to focus on specific data recovery scenarios where only changes since the last backup are needed, thus saving time and resources.

```sql+postgres
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

```sql+sqlite
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
Explore which backups are stored in a specific location to better manage storage and retrieval processes.

```sql+postgres
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

```sql+sqlite
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
This query helps to monitor and manage data security by identifying recent database backups. It is particularly useful in ensuring regular backups are being created and stored correctly, which is crucial for data recovery in case of accidental loss or system failure.

```sql+postgres
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

```sql+sqlite
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
  backup_end_time >= datetime('now', '-30 day');
```