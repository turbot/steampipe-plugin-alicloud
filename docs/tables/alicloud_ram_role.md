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

```sql
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

### Find all roles having Administrator access

```sql
select
  name,
  policies ->> 'PolicyName' as policy_name
from
  alicloud_ram_role,
  jsonb_array_elements(attached_policy) as policies
where 
  policies ->> 'PolicyName' = 'AdministratorAccess';
```



### Find all roles grant cross-account access in the Trust Policy

```sql
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
