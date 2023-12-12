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
Explore which databases are active and their corresponding engines in your Alicloud RDS service. This can be beneficial for a quick overview of your database's status and understanding which database engines are in use.

```sql+postgres
select
  db_name,
  db_instance_id,
  db_status,
  engine
from
  alicloud_rds_database;
```

```sql+sqlite
select
  db_name,
  db_instance_id,
  db_status,
  engine
from
  alicloud_rds_database;
```

### Count databases per instance
Assess the number of databases within each instance to better understand your resource distribution and usage. This can be useful in managing your resources effectively and planning for capacity or scaling needs.

```sql+postgres
select
  db_instance_id,
  count("db_name") as database_count
from
  alicloud_rds_database
group by
  db_instance_id;
```

```sql+sqlite
select
  db_instance_id,
  count("db_name") as database_count
from
  alicloud_rds_database
group by
  db_instance_id;
```

### List databases of engine MySQL
Explore the status and instance IDs of your MySQL databases. This is useful for managing and tracking your MySQL databases across your Alicloud RDS service.

```sql+postgres
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

```sql+sqlite
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
This example helps you discover the segments that include specific details about each database in your Alicloud RDS instances. It assists in gaining insights into the network type, creation time, and engine used, which can be beneficial for database management and optimization.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which certain databases are not operational. This is particularly useful for identifying potential issues and ensuring all databases are functioning as expected.

```sql+postgres
select
  db_name,
  db_instance_id,
  db_status
from
  alicloud_rds_instance
where
  db_status <> 'Running';
```

```sql+sqlite
select
  db_name,
  db_instance_id,
  db_status
from
  alicloud_rds_instance
where
  db_status <> 'Running';
```