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
  agent_installed,
  agent_online,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  agent_installed = true;
```

```sql+sqlite
select
  instance_id,
  instance_name,
  client_status,
  client_version,
  agent_installed,
  agent_online,
  os_name,
  ip,
  region
from
  alicloud_security_center_asset
where
  agent_installed = 1;
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
  agent_installed = false
  or client_status = 'uninstall';
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
  agent_installed = 0
  or client_status = 'uninstall';
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
  agent_installed = true
  and agent_online = false;
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
  agent_installed = 1
  and agent_online = 0;
```

### Count instances by agent status

Get a summary of how many instances have agents installed, online, or not installed.

```sql+postgres
select
  client_status,
  count(*) as instance_count,
  count(*) filter (where agent_installed = true) as agent_installed_count,
  count(*) filter (where agent_online = true) as agent_online_count
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
  count(*) as instance_count,
  sum(case when agent_installed = 1 then 1 else 0 end) as agent_installed_count,
  sum(case when agent_online = 1 then 1 else 0 end) as agent_online_count
from
  alicloud_security_center_asset
group by
  client_status
order by
  instance_count desc;
```

### CIS 4.6: Ensure that the endpoint protection for all Virtual Machines is installed

This query checks if all ECS instances have the Security Center agent installed for endpoint protection.

```sql+postgres
select
  'acs:ecs:' || i.region || ':' || i.account_id || ':instance/' || i.instance_id as resource,
  case
    when exists (
      select 1
      from alicloud_security_center_asset sca
      where sca.instance_id = i.instance_id
        and sca.agent_installed = true
        and sca.region = i.region
        and sca.account_id = i.account_id
    )
    then 'ok'
    else 'alarm'
  end as status,
  case
    when exists (
      select 1
      from alicloud_security_center_asset sca
      where sca.instance_id = i.instance_id
        and sca.agent_installed = true
        and sca.agent_online = true
        and sca.region = i.region
        and sca.account_id = i.account_id
    )
    then i.name || ' has Security Center agent installed and online'
    when exists (
      select 1
      from alicloud_security_center_asset sca
      where sca.instance_id = i.instance_id
        and sca.agent_installed = true
        and sca.region = i.region
        and sca.account_id = i.account_id
    )
    then i.name || ' has Security Center agent installed but is offline'
    else i.name || ' does not have Security Center agent installed. Install the agent in Security Center Console > Settings > Agent'
  end as reason
from
  alicloud_ecs_instance i
where
  i.status = 'Running';
```

```sql+sqlite
select
  'acs:ecs:' || i.region || ':' || i.account_id || ':instance/' || i.instance_id as resource,
  case
    when exists (
      select 1
      from alicloud_security_center_asset sca
      where sca.instance_id = i.instance_id
        and sca.agent_installed = 1
        and sca.region = i.region
        and sca.account_id = i.account_id
    )
    then 'ok'
    else 'alarm'
  end as status,
  case
    when exists (
      select 1
      from alicloud_security_center_asset sca
      where sca.instance_id = i.instance_id
        and sca.agent_installed = 1
        and sca.agent_online = 1
        and sca.region = i.region
        and sca.account_id = i.account_id
    )
    then i.name || ' has Security Center agent installed and online'
    when exists (
      select 1
      from alicloud_security_center_asset sca
      where sca.instance_id = i.instance_id
        and sca.agent_installed = 1
        and sca.region = i.region
        and sca.account_id = i.account_id
    )
    then i.name || ' has Security Center agent installed but is offline'
    else i.name || ' does not have Security Center agent installed. Install the agent in Security Center Console > Settings > Agent'
  end as reason
from
  alicloud_ecs_instance i
where
  i.status = 'Running';
```

### List instances with vulnerabilities and agent status

Find instances that have vulnerabilities detected and check their agent status.

```sql+postgres
select
  instance_id,
  instance_name,
  agent_installed,
  agent_online,
  vul_count,
  risk_count,
  safe_event_count,
  vul_status,
  risk_status,
  region
from
  alicloud_security_center_asset
where
  vul_count > 0
  or risk_count::int > 0
order by
  vul_count desc,
  risk_count desc;
```

```sql+sqlite
select
  instance_id,
  instance_name,
  agent_installed,
  agent_online,
  vul_count,
  risk_count,
  safe_event_count,
  vul_status,
  risk_status,
  region
from
  alicloud_security_center_asset
where
  vul_count > 0
  or cast(risk_count as integer) > 0
order by
  vul_count desc,
  risk_count desc;
```

