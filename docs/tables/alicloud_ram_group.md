---
title: "Steampipe Table: alicloud_ram_group - Query Alibaba Cloud RAM Groups using SQL"
description: "Allows users to query RAM Groups in Alibaba Cloud, specifically providing details on group name, id, comments, and the creation time."
---

# Table: alicloud_ram_group - Query Alibaba Cloud RAM Groups using SQL

Alibaba Cloud Resource Access Management (RAM) is a service that helps manage user identities and resource access permissions. RAM allows you to create and manage multiple identities under one Alibaba Cloud account, and control the access of these identities to your resources. You can grant different permissions to different identities to ensure that your resources can only be accessed by trusted entities.

## Table Usage Guide

The `alicloud_ram_group` table provides insights into RAM Groups within Alibaba Cloud Resource Access Management (RAM). As a system administrator, explore group-specific details through this table, including group name, id, comments, and the creation time. Utilize it to manage and control access to your resources, ensuring that only trusted entities have the necessary permissions.

## Examples

### User details associated with each RAM group

```sql
select
  name as group_name,
  iam_user ->> 'UserName' as user_name,
  iam_user ->> 'DisplayName' as display_name,
  iam_user ->> 'JoinDate' as user_join_date
from
  alicloud_ram_group
  cross join jsonb_array_elements(users) as iam_user;
```

### List the policies attached to each RAM group

```sql
select
  name as group_name,
  policies ->> 'PolicyName' as policy_name,
  policies ->> 'PolicyType' as policy_type,
  policies ->> 'DefaultVersion' as policy_default_version,
  policies ->> 'AttachDate' as policy_attachment_date
from
  alicloud_ram_group,
  jsonb_array_elements(attached_policy) as policies;
```

### List of RAM groups with no users added to it

```sql
select
  name as group_name,
  create_date,
  users
from
  alicloud_ram_group
where
  users = '[]';
```
