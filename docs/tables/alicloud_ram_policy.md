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
Explore the various policies in your Alicloud RAM to understand their types, descriptions, and default versions. This can be beneficial in managing and reviewing your security settings.

```sql+postgres
select
  policy_name,
  policy_type,
  description,
  default_version,
  policy_document
from
  alicloud_ram_policy;
```

```sql+sqlite
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
Determine the areas in which system policies are implemented for better understanding of the default versions and descriptions. This can aid in assessing the elements within your Alicloud RAM policy, offering insights into your system's security configuration.

```sql+postgres
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

```sql+sqlite
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
Explore which custom policies are in place within your system. This allows you to gain insights into the policy name, type, description, default version, and policy document, helping you better manage and understand your system's security measures.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which policies are granting full access. This is useful for assessing potential security vulnerabilities and ensuring that access permissions align with your organization's security protocols.

```sql+postgres
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

```sql+sqlite
select
  policy_name,
  policy_type,
  action.value as action,
  json_extract(s.value, '$.Effect') as effect
from
  alicloud_ram_policy,
  json_each(policy_document_std, '$.Statement') as s,
  json_each(s.value, '$.Action') as action
where
  action.value in ('*', '*:*')
  and json_extract(s.value, '$.Effect') = 'Allow';
```