---
title: "Steampipe Table: alicloud_ram_group - Query Alibaba Cloud RAM Groups using SQL"
description: "Allows users to query RAM Groups in Alibaba Cloud, specifically providing details on group name, id, comments, and the creation time."
folder: "RAM"
---

# Table: alicloud_ram_group - Query Alibaba Cloud RAM Groups using SQL

Alibaba Cloud Resource Access Management (RAM) is a service that helps manage user identities and resource access permissions. RAM allows you to create and manage multiple identities under one Alibaba Cloud account, and control the access of these identities to your resources. You can grant different permissions to different identities to ensure that your resources can only be accessed by trusted entities.

## Table Usage Guide

The `alicloud_ram_group` table provides insights into RAM Groups within Alibaba Cloud Resource Access Management (RAM). As a system administrator, explore group-specific details through this table, including group name, id, comments, and the creation time. Utilize it to manage and control access to your resources, ensuring that only trusted entities have the necessary permissions.

## Examples

### User details associated with each RAM group
Determine the areas in which users are associated with each RAM group in Alicloud. This can help in better understanding the group distribution and user management within your Alicloud environment.

```sql+postgres
select
  name as group_name,
  iam_user ->> 'UserName' as user_name,
  iam_user ->> 'DisplayName' as display_name,
  iam_user ->> 'JoinDate' as user_join_date
from
  alicloud_ram_group
  cross join jsonb_array_elements(users) as iam_user;
```

```sql+sqlite
select
  name as group_name,
  json_extract(iam_user.value, '$.UserName') as user_name,
  json_extract(iam_user.value, '$.DisplayName') as display_name,
  json_extract(iam_user.value, '$.JoinDate') as user_join_date
from
  alicloud_ram_group,
  json_each(users) as iam_user;
```

### List the policies attached to each RAM group
Explore the various policies attached to each RAM group, including the policy type, default version, and attachment date. This can help in understanding the security measures and access controls in place for each group.

```sql+postgres
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

```sql+sqlite
select
  name as group_name,
  json_extract(policies.value, '$.PolicyName') as policy_name,
  json_extract(policies.value, '$.PolicyType') as policy_type,
  json_extract(policies.value, '$.DefaultVersion') as policy_default_version,
  json_extract(policies.value, '$.AttachDate') as policy_attachment_date
from
  alicloud_ram_group,
  json_each(attached_policy) as policies;
```

### List of RAM groups with no users added to it
Determine the areas in which RAM groups have been created but no users have been added. This can help in identifying unused resources and optimizing resource allocation.

```sql+postgres
select
  name as group_name,
  create_date,
  users
from
  alicloud_ram_group
where
  users = '[]';
```

```sql+sqlite
select
  name as group_name,
  create_date,
  users
from
  alicloud_ram_group
where
  users = '[]';
```