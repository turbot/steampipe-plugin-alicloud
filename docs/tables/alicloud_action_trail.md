---
title: "Steampipe Table: alicloud_action_trail - Query Alibaba Cloud Action Trails using SQL"
description: "Allows users to query Action Trails in Alibaba Cloud, providing insights into user activity and operations within the platform."
folder: "ActionTrail"
---

# Table: alicloud_action_trail - Query Alibaba Cloud Action Trails using SQL

Action Trail is a feature within Alibaba Cloud that records and audits user activity within an Alibaba Cloud account. It provides detailed information about API calls, including the caller identities, IP addresses, time of the calls, request parameters, and response elements. Action Trail helps in security analysis, resource change tracking, compliance auditing, and troubleshooting.

## Table Usage Guide

The `alicloud_action_trail` table provides insights into user activity within Alibaba Cloud. As a security analyst, you can explore trail-specific details through this table, including the name, home region, role name, and other associated metadata. Utilize it to uncover information about trails, such as delivery location, trail status, and the creation time, which can aid in auditing and compliance tasks.

## Examples

### Basic info
Explore the status and regional details of your Alicloud action trails to understand which trails are active and where they are operating. This can help in managing and optimizing your security audit trails.

```sql+postgres
select
  name,
  home_region,
  event_rw,
  status,
  trail_region
from
  alicloud_action_trail;
```

```sql+sqlite
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
Discover the segments that are actively monitoring your Alibaba Cloud resources. This query will help you understand which of your action trails are currently enabled and actively logging events, providing insights into your system's security and compliance.

```sql+postgres
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

```sql+sqlite
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
This query is useful for identifying all the action trails that are set up across multiple accounts in your organization. It helps in understanding the configuration and status of these trails, which can be beneficial for auditing and compliance purposes.

```sql+postgres
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

```sql+sqlite
select
  name,
  home_region,
  is_organization_trail,
  status,
  trail_region
from
  alicloud_action_trail
where
  is_organization_trail = 1;
```

### List shadow trails
Determine the areas in which Alicloud's action trails are active across all regions, but their home region is different. This can be useful for understanding the distribution and operation of action trails in different regions.

```sql+postgres
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

```sql+sqlite
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