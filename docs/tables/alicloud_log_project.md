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
Explore all SLS projects across regions to gain insights into their operational status, redundancy configuration, and creation details. This can help in managing your log service resources, understanding their distribution, and identifying projects that may need attention.

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
Discover SLS projects that are configured with Zone Redundant Storage (ZRS) for enhanced data durability and availability. This is useful for identifying projects with high-availability configurations and ensuring critical log data is properly protected across multiple availability zones.

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
Identify SLS projects that are in a normal operational state. This helps in monitoring project health and ensuring all projects are functioning correctly, which is essential for maintaining reliable log collection and analysis capabilities.

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
Determine which SLS projects have been recently created within the last month. This is useful for tracking new log service deployments, auditing recent changes, and managing project lifecycle across your Alibaba Cloud environment.

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
Retrieve comprehensive information about a specific SLS project by its name. This helps in getting detailed insights into project configuration, ownership, redundancy settings, and modification history, which is essential for project management and troubleshooting.

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

