---
title: "Steampipe Table: alicloud_log_store - Query Alibaba Cloud SLS Logstores using SQL"
description: "List and inspect Alibaba Cloud Log Service (SLS) logstores across projects and regions."
folder: "SLS"
---

# Table: alicloud_log_store - Query Alibaba Cloud SLS Logstores using SQL

Alibaba Cloud Log Service (SLS) Logstores are containers within SLS projects that store and manage log data. Each logstore has configurable retention periods, shard counts, and other settings that control how log data is stored and accessed.

## Table Usage Guide

The `alicloud_log_store` table lets you query SLS logstores across regions and projects. Use it to discover logstores, check their retention periods, shard configurations, and ensure compliance with data retention policies.

## Examples

### List all logstores
Explore all SLS logstores across projects and regions to gain insights into their configuration, retention settings, and operational status. This helps in managing your log storage resources and understanding the distribution of logstores across your Alibaba Cloud environment.

```sql+postgres
select
  project,
  name,
  region,
  ttl,
  shard_count,
  mode
from
  alicloud_log_store;
```

```sql+sqlite
select
  project,
  name,
  region,
  ttl,
  shard_count,
  mode
from
  alicloud_log_store;
```

### Find logstores with retention period less than 365 days
Identify logstores that do not meet the recommended 365-day retention period for compliance purposes. This is useful for ensuring your log data is retained long enough for incident response and audit requirements, as specified in CIS benchmarks.

```sql+postgres
select
  project,
  name,
  region,
  ttl,
  case
    when ttl = -1 then 'Permanent'
    else ttl::text || ' days'
  end as retention_period
from
  alicloud_log_store
where
  ttl < 365
  and ttl != -1;
```

```sql+sqlite
select
  project,
  name,
  region,
  ttl,
  case
    when ttl = -1 then 'Permanent'
    else cast(ttl as text) || ' days'
  end as retention_period
from
  alicloud_log_store
where
  ttl < 365
  and ttl != -1;
```

### List logstores by project
Discover all logstores within a specific SLS project to understand the log storage structure and configuration. This helps in project-level resource management and planning.

```sql+postgres
select
  name,
  ttl,
  shard_count,
  web_tracking,
  auto_split,
  create_time
from
  alicloud_log_store
where
  project = 'my-project';
```

```sql+sqlite
select
  name,
  ttl,
  shard_count,
  web_tracking,
  auto_split,
  create_time
from
  alicloud_log_store
where
  project = 'my-project';
```

### Get logstore details by name
Retrieve comprehensive information about a specific logstore including its retention settings, shard configuration, and operational parameters. This is essential for detailed logstore management and troubleshooting.

```sql+postgres
select
  project,
  name,
  region,
  ttl,
  shard_count,
  web_tracking,
  auto_split,
  max_split_shard,
  mode,
  create_time,
  last_modify_time
from
  alicloud_log_store
where
  project = 'my-project'
  and name = 'my-logstore';
```

```sql+sqlite
select
  project,
  name,
  region,
  ttl,
  shard_count,
  web_tracking,
  auto_split,
  max_split_shard,
  mode,
  create_time,
  last_modify_time
from
  alicloud_log_store
where
  project = 'my-project'
  and name = 'my-logstore';
```

### Find logstores with permanent storage enabled
Identify logstores that have permanent storage enabled (TTL = -1), which means log data will be stored indefinitely. This is useful for identifying critical logstores that require long-term retention.

```sql+postgres
select
  project,
  name,
  region,
  shard_count,
  create_time
from
  alicloud_log_store
where
  ttl = -1;
```

```sql+sqlite
select
  project,
  name,
  region,
  shard_count,
  create_time
from
  alicloud_log_store
where
  ttl = -1;
```

### List logstores with auto-split enabled
Discover logstores that have automatic shard splitting enabled, which allows the system to automatically increase shard count based on data volume. This helps in understanding which logstores are configured for automatic scaling.

```sql+postgres
select
  project,
  name,
  region,
  shard_count,
  max_split_shard,
  auto_split
from
  alicloud_log_store
where
  auto_split = true;
```

```sql+sqlite
select
  project,
  name,
  region,
  shard_count,
  max_split_shard,
  auto_split
from
  alicloud_log_store
where
  auto_split = 1;
```

