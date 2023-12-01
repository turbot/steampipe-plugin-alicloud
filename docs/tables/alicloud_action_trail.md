---
title: "Steampipe Table: alicloud_action_trail - Query Alibaba Cloud Action Trails using SQL"
description: "Allows users to query Action Trails in Alibaba Cloud, providing insights into user activity and operations within the platform."
---

# Table: alicloud_action_trail - Query Alibaba Cloud Action Trails using SQL

Action Trail is a feature within Alibaba Cloud that records and audits user activity within an Alibaba Cloud account. It provides detailed information about API calls, including the caller identities, IP addresses, time of the calls, request parameters, and response elements. Action Trail helps in security analysis, resource change tracking, compliance auditing, and troubleshooting.

## Table Usage Guide

The `alicloud_action_trail` table provides insights into user activity within Alibaba Cloud. As a security analyst, you can explore trail-specific details through this table, including the name, home region, role name, and other associated metadata. Utilize it to uncover information about trails, such as delivery location, trail status, and the creation time, which can aid in auditing and compliance tasks.

## Examples

### Basic info

```sql
select
  name,
  home_region,
  event_rw,
  status,
  trail_region
from
  alicloud_action_trail;
```

### List enabled trails

```sql
select
  name,
  home_region,
  event_rw,
  status,
  trail_region
from
  alicloud_action_trail
where
  status = 'Enable';
```

### List multi-account trails

```sql
select
  name,
  home_region,
  is_organization_trail,
  status,
  trail_region
from
  alicloud_action_trail
where
  is_organization_trail;
```

### List shadow trails

```sql
select
  name,
  region,
  home_region
from
  alicloud_action_trail
where
  trail_region = 'All'
  and home_region <> region;
```
