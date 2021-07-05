# Table: alicloud_ram_policy

Permissions are specified by a statement within a policy that allows or denies access to a specific Alibaba Cloud resource.
A policy defines a set of permissions that are described based on the policy structure and syntax. A policy can accurately describe the authorized resource sets, authorized operation sets, and authorization conditions.

## Examples

### Basic info

```sql
select
  name,
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
  name,
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
  name,
  policy_type,
  description,
  default_version,
  policy_document
from
  alicloud_ram_policy
where
  policy_type = 'Custom';
```

### Find policy statements that grant Full Control access

```sql
select
  name,
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
