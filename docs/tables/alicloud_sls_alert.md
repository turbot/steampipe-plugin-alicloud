---
title: "Steampipe Table: alicloud_sls_alert - Query Alibaba Cloud SLS Alerts using SQL"
description: "List and inspect Alibaba Cloud Log Service (SLS) alert rules across projects and regions."
folder: "SLS"
---

# Table: alicloud_sls_alert - Query Alibaba Cloud SLS Alerts using SQL

Alibaba Cloud Log Service (SLS) Alerts continuously evaluate saved queries and trigger notifications when conditions are met.

## Table Usage Guide

The `alicloud_sls_alert` table lets you query SLS alert rules across regions and projects. Use it to validate alert coverage, status, schedules, and the queries that power alerts.

## Examples

### List all enabled alerts
Discover all active SLS alert rules that are currently enabled across your projects and regions. This helps in monitoring your alert coverage, understanding which alerts are operational, and ensuring critical monitoring is in place for your log data.

```sql+postgres
select
  region,
  project,
  name,
  display_name,
  status
from
  alicloud_sls_alert
where
  status = 'ENABLED';
```

```sql+sqlite
select
  region,
  project,
  name,
  display_name,
  status
from
  alicloud_sls_alert
where
  status = 'ENABLED';
```

### Show alert queries
Examine the query configurations for all SLS alerts to understand what conditions are being monitored. This is useful for reviewing alert logic, validating query syntax, and ensuring alerts are configured to detect the right conditions in your log data.

```sql+postgres
select
  project,
  name,
  jsonb_pretty(query_list) as query_list
from
  alicloud_sls_alert;
```

```sql+sqlite
select
  project,
  name,
  json(query_list) as query_list
from
  alicloud_sls_alert;
```
