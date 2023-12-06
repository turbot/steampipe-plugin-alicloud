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
Gain insights into the basic information of users within your Alicloud resource access management system. This is beneficial for managing user identities and controlling access to your resources.

```sql+postgres
select
  user_id,
  name,
  display_name
from
  alicloud_ram_user;
```

```sql+sqlite
select
  user_id,
  name,
  display_name
from
  alicloud_ram_user;
```

### Users who have not logged in for 30 days
Identify instances where users have been inactive for a month. This can be useful to monitor user engagement and potentially re-engage inactive users.

```sql+postgres
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date < current_date - interval '30 days';
```

```sql+sqlite
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date < date('now','-30 day');
```

### Users who have never logged in
Identify users who have yet to log in for the first time. This can be useful for understanding user engagement and identifying potentially inactive accounts.

```sql+postgres
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date is null;
```

```sql+sqlite
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date is null;
```

### Groups details to which the RAM user belongs
This query is useful for identifying which groups a particular RAM user belongs to and when they joined those groups. This could be beneficial for managing user permissions and access within an Alicloud environment.

```sql+postgres
select
  name as user_name,
  iam_group ->> 'GroupName' as group_name,
  iam_group ->> 'JoinDate' as join_date
from
  alicloud_ram_user,
  jsonb_array_elements(groups) as iam_group;
```

```sql+sqlite
select
  name as user_name,
  json_extract(iam_group.value, '$.GroupName') as group_name,
  json_extract(iam_group.value, '$.JoinDate') as join_date
from
  alicloud_ram_user,
  json_each(groups) as iam_group;
```

### List all the users having Administrator access
Determine the areas in which users have been granted administrative access. This is useful for auditing security and ensuring that only authorized individuals have high-level permissions.

```sql+postgres
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

```sql+sqlite
select
  name as user_name,
  json_extract(policies.value, '$.PolicyName') as policy_name,
  json_extract(policies.value, '$.PolicyType') as policy_type,
  json_extract(policies.value, '$.DefaultVersion') as policy_default_version,
  json_extract(policies.value, '$.AttachDate') as policy_attachment_date
from
  alicloud_ram_user,
  json_each(attached_policy) as policies
where
  json_extract(policies.value, '$.PolicyName') = 'AdministratorAccess';
```

### List all the users for whom MFA is not enabled
Explore which users have not enabled multi-factor authentication, a crucial security feature, to identify potential vulnerabilities in your system. This can be particularly useful in prioritizing security improvements and ensuring compliance with best practices.

```sql+postgres
select
  name as user_name,
  user_id as user_id,
  mfa_enabled
from
  alicloud_ram_user
where
  not mfa_enabled;
```

```sql+sqlite
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
Discover the users who have been granted permissions for role-based access control in the Container Service for Kubernetes. This is particularly useful for managing user access and ensuring only authorized users have certain permissions.

```sql+postgres
select
  name as user_name,
  user_id as user_id
from
  alicloud_ram_user
where
  cs_user_permission <> '[]';
```

```sql+sqlite
select
  name as user_name,
  user_id as user_id
from
  alicloud_ram_user
where
  cs_user_permission != '[]';
```