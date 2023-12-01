---
title: "Steampipe Table: alicloud_ram_policy - Query Alicloud RAM Policies using SQL"
description: "Allows users to query Alicloud RAM Policies, specifically to retrieve information about the policy name, policy type, description, and creation time."
---

# Table: alicloud_ram_policy - Query Alicloud RAM Policies using SQL

Alicloud RAM Policy is a service within Alibaba Cloud that allows you to manage access permissions to your Alicloud resources. It provides a centralized way to set up and manage policies for various Alicloud resources, including ECS instances, databases, web applications, and more. Alicloud RAM Policy helps you control who has authorization to access and manage your Alicloud resources.

## Table Usage Guide

The `alicloud_ram_policy` table provides insights into RAM policies within Alibaba Cloud Resource Access Management (RAM). As a security administrator, delve into policy-specific details through this table, including policy names, types, descriptions, and creation times. Utilize it to uncover information about policies, such as those with specific permissions, the resources they apply to, and when they were created.

## Examples

### Basic info

```sql
select
  policy_name,
  policy_type,
  description,
  default_version,
  policy_document
from
  alicloud_ram_policy;
```

### List system policies

```sql
select
  policy_name,
  policy_type,
  description,
  default_version,
  policy_document
from
  alicloud_ram_policy
where
  policy_type = 'System';
```

### List custom policies

```sql
select
  policy_name,
  policy_type,
  description,
  default_version,
  policy_document
from
  alicloud_ram_policy
where
  policy_type = 'Custom';
```

### List policies with statements granting full access

```sql
select
  policy_name,
  policy_type,
  action,
  s ->> 'Effect' as effect
from
  alicloud_ram_policy,
  jsonb_array_elements(policy_document_std -> 'Statement') as s,
  jsonb_array_elements_text(s -> 'Action') as action
where
  action in ('*', '*:*')
  and s ->> 'Effect' = 'Allow';
```
