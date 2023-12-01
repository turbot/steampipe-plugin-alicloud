---
title: "Steampipe Table: alicloud_ram_user - Query Alibaba Cloud RAM Users using SQL"
description: "Allows users to query RAM Users in Alibaba Cloud, specifically the user details and associated policies, providing insights into user access management and security."
---

# Table: alicloud_ram_user - Query Alibaba Cloud RAM Users using SQL

Alibaba Cloud RAM (Resource Access Management) is a service that helps you centrally manage your Alibaba Cloud resources. RAM allows you to control who (users and systems) has what permissions to which resources by setting policies. This aids in achieving least privilege, thereby enhancing the security of your Alibaba Cloud resources.

## Table Usage Guide

The `alicloud_ram_user` table provides insights into RAM users within Alibaba Cloud RAM. As a Security Analyst, explore user-specific details through this table, including associated policies, user creation time, and last login time. Utilize it to uncover information about users, such as those with excessive permissions, the policies associated with each user, and the verification of user activity.

## Examples

### Basic user info

```sql
select
  user_id,
  name,
  display_name
from
  alicloud_ram_user;
```

### Users who have not logged in for 30 days

```sql
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date < current_date - interval '30 days';
```

### Users who have never logged in

```sql
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date is null;
```

### Groups details to which the RAM user belongs

```sql
select
  name as user_name,
  iam_group ->> 'GroupName' as group_name,
  iam_group ->> 'JoinDate' as join_date
from
  alicloud_ram_user,
  jsonb_array_elements(groups) as iam_group;
```

### List all the users having Administrator access

```sql
select
  name as user_name,
  policies ->> 'PolicyName' as policy_name,
  policies ->> 'PolicyType' as policy_type,
  policies ->> 'DefaultVersion' as policy_default_version,
  policies ->> 'AttachDate' as policy_attachment_date
from
  alicloud_ram_user,
  jsonb_array_elements(attached_policy) as policies
where
  policies ->> 'PolicyName' = 'AdministratorAccess';
```

### List all the users for whom MFA is not enabled

```sql
select
  name as user_name,
  user_id as user_id,
  mfa_enabled
from
  alicloud_ram_user
where
  not mfa_enabled;
```

### List users with Container Service for Kubernetes role-based access control (RBAC) permissions

```sql
select
  name as user_name,
  user_id as user_id
from
  alicloud_ram_user
where
  cs_user_permission <> '[]';
```
