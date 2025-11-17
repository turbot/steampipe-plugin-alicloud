---
title: "Steampipe Table: alicloud_log_project - Query Alibaba Cloud SLS Projects using SQL"
description: "List and inspect Alibaba Cloud Log Service (SLS) projects across regions."
folder: "SLS"
---

# Table: alicloud_log_project - Query Alibaba Cloud SLS Projects using SQL

Alibaba Cloud Log Service (SLS) Projects are containers for log data. Each project can contain multiple logstores and is region-specific.

## Table Usage Guide

The `alicloud_log_project` table lets you query SLS projects across regions. Use it to discover projects, check their status, redundancy configuration, and creation details.

## Examples

### List all SLS projects

```sql+postgres
select
  name,
  description,
  status,
  region,
  data_redundancy_type,
  location,
  create_time
from
  alicloud_log_project;
```

```sql+sqlite
select
  name,
  description,
  status,
  region,
  data_redundancy_type,
  location,
  create_time
from
  alicloud_log_project;
```

### List projects with ZRS redundancy

```sql+postgres
select
  name,
  region,
  data_redundancy_type,
  location
from
  alicloud_log_project
where
  data_redundancy_type = 'ZRS';
```

```sql+sqlite
select
  name,
  region,
  data_redundancy_type,
  location
from
  alicloud_log_project
where
  data_redundancy_type = 'ZRS';
```

### List projects by status

```sql+postgres
select
  name,
  status,
  region,
  create_time
from
  alicloud_log_project
where
  status = 'Normal';
```

```sql+sqlite
select
  name,
  status,
  region,
  create_time
from
  alicloud_log_project
where
  status = 'Normal';
```

### Find projects created in the last 30 days

```sql+postgres
select
  name,
  description,
  region,
  create_time
from
  alicloud_log_project
where
  create_time >= now() - interval '30 days';
```

```sql+sqlite
select
  name,
  description,
  region,
  create_time
from
  alicloud_log_project
where
  create_time >= datetime('now', '-30 days');
```

### Get project details by name

```sql+postgres
select
  name,
  description,
  status,
  owner,
  region,
  data_redundancy_type,
  location,
  create_time,
  last_modify_time
from
  alicloud_log_project
where
  name = 'my-project';
```

```sql+sqlite
select
  name,
  description,
  status,
  owner,
  region,
  data_redundancy_type,
  location,
  create_time,
  last_modify_time
from
  alicloud_log_project
where
  name = 'my-project';
```

