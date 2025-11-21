---
title: "Steampipe Table: alicloud_security_center_asset - Query Alibaba Cloud Security Center Assets using SQL"
description: "Allows users to query Security Center Assets in Alibaba Cloud, providing insights into ECS instances and their Security Center agent installation status."
folder: "Security Center"
---

# Table: alicloud_security_center_asset - Query Alibaba Cloud Security Center Assets using SQL

Alibaba Cloud Security Center monitors and protects your ECS instances through an agent that must be installed on each instance. This table provides information about assets (primarily ECS instances) that are registered with Security Center and their agent status.

## Table Usage Guide

The `alicloud_security_center_asset` table provides insights into assets monitored by Security Center within Alibaba Cloud. As a security engineer, explore asset-specific details through this table, including agent installation status, agent version, vulnerability counts, and security status. Utilize it to identify instances without endpoint protection, track agent health, and ensure compliance with security policies.

## Examples

### List all instances with Security Center agent installed
Find all ECS instances that have the Security Center agent installed and are being monitored.

```sql+postgres
select
  instance_id,
  instance_name,
  client_status,
  client_version,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  client_status != 'uninstall';
```

```sql+sqlite
select
  instance_id,
  instance_name,
  client_status,
  client_version,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  client_status != 'uninstall';
```

### List instances without Security Center agent
Identify ECS instances that do not have the Security Center agent installed, which need endpoint protection.

```sql+postgres
select
  instance_id,
  instance_name,
  client_status,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  client_status = 'uninstall';
```

```sql+sqlite
select
  instance_id,
  instance_name,
  client_status,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  client_status = 'uninstall';
```

### List instances with offline agents
Find instances where the Security Center agent is installed but currently offline.

```sql+postgres
select
  instance_id,
  instance_name,
  client_status,
  client_version,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  client_status = 'offline';
```

```sql+sqlite
select
  instance_id,
  instance_name,
  client_status,
  client_version,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  client_status = 'offline';
```

### Count instances by agent status
Get a summary of how many instances have agents installed, online, or not installed.

```sql+postgres
select
  client_status,
  count(*) as instance_count
from
  alicloud_security_center_asset
group by
  client_status
order by
  instance_count desc;
```

```sql+sqlite
select
  client_status,
  count(*) as instance_count
from
  alicloud_security_center_asset
group by
  client_status
order by
  instance_count desc;
```
