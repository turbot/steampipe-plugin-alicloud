---
title: "Steampipe Table: alicloud_rds_database - Query Alicloud RDS Databases using SQL"
description: "Allows users to query Alicloud RDS Databases, specifically details about each database hosted within the Alicloud RDS service."
---

# Table: alicloud_rds_database - Query Alicloud RDS Databases using SQL

Alicloud RDS (Relational Database Service) is a web service that makes it easier to set up, operate, and scale a relational database in the cloud. It provides cost-efficient, resizable capacity for an industry-standard relational database and manages common database administration tasks. Alicloud RDS supports multiple types of databases, including MySQL, SQL Server, and PostgreSQL, among others.

## Table Usage Guide

The `alicloud_rds_database` table provides insights into databases within Alicloud RDS. As a database administrator or a DevOps engineer, explore database-specific details through this table, including database names, character sets, and associated metadata. Utilize it to uncover information about databases, such as their status, creation times, and engine versions.

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

### Get DB instance details of each database

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