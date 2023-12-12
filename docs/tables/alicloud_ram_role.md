---
title: "Steampipe Table: alicloud_ram_role - Query Alicloud RAM Roles using SQL"
description: "Allows users to query Alicloud RAM Roles, providing insights into role-specific details, permissions, and trust policies."
---

# Table: alicloud_ram_role - Query Alicloud RAM Roles using SQL

Alicloud RAM (Resource Access Management) is a service that helps you manage user identities and access permissions. You can create and manage multiple identities under your Alibaba Cloud account, and control the resources that each identity can access. RAM allows you to grant precise access permissions to different users, user groups, and roles.

## Table Usage Guide

The `alicloud_ram_role` table provides insights into RAM roles within Alicloud Resource Access Management. As a security analyst, explore role-specific details through this table, including permissions, trust policies, and associated metadata. Utilize it to uncover information about roles, such as those with wildcard permissions, the trust relationships between roles, and the verification of trust policies.

## Examples

### List the policies attached to the roles
This query is used to gain insights into the various policies attached to different roles within your Alicloud RAM. It allows you to assess the elements within each role's policy, such as the policy's name, type, default version, and attachment date, providing a comprehensive overview of your role-based access controls.

```sql+postgres
select
  name,
  policies ->> 'PolicyName' as policy_name,
  policies ->> 'PolicyType' as policy_type,
  policies ->> 'DefaultVersion' as policy_default_version,
  policies ->> 'AttachDate' as policy_attachment_date
from
  alicloud_ram_role,
  jsonb_array_elements(attached_policy) as policies
order by name;
```

```sql+sqlite
select
  name,
  json_extract(policies.value, '$.PolicyName') as policy_name,
  json_extract(policies.value, '$.PolicyType') as policy_type,
  json_extract(policies.value, '$.DefaultVersion') as policy_default_version,
  json_extract(policies.value, '$.AttachDate') as policy_attachment_date
from
  alicloud_ram_role,
  json_each(attached_policy) as policies
order by name;
```

### Find all roles having Administrator access
Discover the segments that have Administrator access within a system. This is particularly useful for auditing purposes, ensuring only the correct roles have such high-level permissions.

```sql+postgres
select
  name,
  policies ->> 'PolicyName' as policy_name
from
  alicloud_ram_role,
  jsonb_array_elements(attached_policy) as policies
where 
  policies ->> 'PolicyName' = 'AdministratorAccess';
```

```sql+sqlite
select
  name,
  json_extract(policies.value, '$.PolicyName') as policy_name
from
  alicloud_ram_role,
  json_each(attached_policy) as policies
where 
  json_extract(policies.value, '$.PolicyName') = 'AdministratorAccess';
```



### Find all roles grant cross-account access in the Trust Policy
This query allows you to identify roles that have been granted access to other accounts within the Trust Policy, providing a way to review and manage cross-account permissions. This can be useful in maintaining security and control over data access across multiple accounts.

```sql+postgres
select
  name,
  principal,
  split_part(principal, ':', 4) as foreign_account
from
  alicloud_ram_role,
  jsonb_array_elements(assume_role_policy_document -> 'Statement') as stmt,
  jsonb_array_elements_text(stmt -> 'Principal' -> 'RAM') as principal
where 
  split_part(principal, ':',4) <> account_id;
```

```sql+sqlite
Error: SQLite does not support split or string_to_array functions.
```