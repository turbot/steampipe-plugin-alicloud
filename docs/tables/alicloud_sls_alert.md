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

### CIS 2.10: Ensure monitoring and alert for RAM Policy changes (Create/Update/Delete)

This checks for an enabled alert whose query monitors RAM/ResourceManager policy events.

```sql+postgres
select
  project,
  name,
  display_name,
  status,
  query_obj ->> 'Query' as q
from
  alicloud_sls_alert,
  jsonb_array_elements(query_list) as query_obj
where
  status = 'ENABLED'
  and (
    (query_obj ->> 'Query') ilike '%\"event.serviceName\": ResourceManager%' or
    (query_obj ->> 'Query') ilike '%\"event.serviceName\": Ram%'
  )
  and (
    (query_obj ->> 'Query') ilike '%\"event.eventName\": CreatePolicy%' or
    (query_obj ->> 'Query') ilike '%\"event.eventName\": DeletePolicy%' or
    (query_obj ->> 'Query') ilike '%\"event.eventName\": CreatePolicyVersion%' or
    (query_obj ->> 'Query') ilike '%\"event.eventName\": UpdatePolicyVersion%' or
    (query_obj ->> 'Query') ilike '%\"event.eventName\": SetDefaultPolicyVersion%' or
    (query_obj ->> 'Query') ilike '%\"event.eventName\": DeletePolicyVersion%'
  );
```

```sql+sqlite
-- Expand query_list array and search for the required patterns
with expanded as (
  select
    project,
    name,
    status,
    json_extract(query_obj.value, '$.Query') as q
  from
    alicloud_sls_alert,
    json_each(query_list) as query_obj
)
select
  project,
  name,
  q
from
  expanded
where
  status = 'ENABLED'
  and (
    lower(q) like '%\"event.servicename\": resourcemanager%' or
    lower(q) like '%\"event.servicename\": ram%'
  )
  and (
    lower(q) like '%\"event.eventname\": createpolicy%' or
    lower(q) like '%\"event.eventname\": deletepolicy%' or
    lower(q) like '%\"event.eventname\": createpolicyversion%' or
    lower(q) like '%\"event.eventname\": updatepolicyversion%' or
    lower(q) like '%\"event.eventname\": setdefaultpolicyversion%' or
    lower(q) like '%\"event.eventname\": deletepolicyversion%'
  );
```


