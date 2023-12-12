---
title: "Steampipe Table: alicloud_ram_access_key - Query Alibaba Cloud RAM Access Keys using SQL"
description: "Allows users to query Alibaba Cloud RAM Access Keys, specifically the AccessKey ID, status, and creation time, providing insights into the access keys of RAM users."
---

# Table: alicloud_ram_access_key - Query Alibaba Cloud RAM Access Keys using SQL

Alibaba Cloud Resource Access Management (RAM) is a service that helps you manage user identities and access permissions. With RAM, you can create and manage multiple identities under one Alibaba Cloud account, and control the access of these identities to your resources in an efficient and secure manner. RAM Access Keys are used by RAM users to access Alibaba Cloud services.

## Table Usage Guide

The `alicloud_ram_access_key` table provides insights into the access keys of RAM users within Alibaba Cloud Resource Access Management (RAM). As a security analyst, explore key-specific details through this table, including the AccessKey ID, status, and creation time. Utilize it to uncover information about access keys, such as those that are active or inactive, and the verification of their creation times.

## Examples

### List of access keys with their corresponding user name and date of creation
Discover the segments that have access keys, their corresponding user names, and creation dates. This can be useful in managing and tracking user access within your system.

```sql+postgres
select
  access_key_id,
  user_name,
  create_date
from
  alicloud_ram_access_key;
```

```sql+sqlite
select
  access_key_id,
  user_name,
  create_date
from
  alicloud_ram_access_key;
```

### List of access keys which are inactive
Determine the areas in which there are inactive access keys. This can be useful in maintaining security by identifying and managing unused access keys.

```sql+postgres
select
  access_key_id,
  user_name,
  status
from
  alicloud_ram_access_key
where
  status = 'Inactive';
```

```sql+sqlite
select
  access_key_id,
  user_name,
  status
from
  alicloud_ram_access_key
where
  status = 'Inactive';
```

### Access key count by user name
Identify instances where multiple access keys are associated with the same user in Alicloud. This can help in managing access keys effectively and improving security by limiting the number of access keys per user.

```sql+postgres
select
  user_name,
  count (access_key_id) as access_key_count
from
  alicloud_ram_access_key
group by
  user_name;
```

```sql+sqlite
select
  user_name,
  count(access_key_id) as access_key_count
from
  alicloud_ram_access_key
group by
  user_name;
```


### Access keys older than 90 days
Determine the instances where access keys have been in use for more than 90 days. This can be beneficial for managing security and access control, as older keys may pose a higher risk if not regularly updated or reviewed.

```sql+postgres
select
  access_key_id,
  user_name,
  status
  create_date,
  age(create_date)
from
  alicloud_ram_access_key
where
  create_date <= (current_date - interval '90' day)
order by
  create_date;
```

```sql+sqlite
select
  access_key_id,
  user_name,
  status,
  create_date,
  julianday('now') - julianday(create_date)
from
  alicloud_ram_access_key
where
  julianday('now') - julianday(create_date) >= 90
order by
  create_date;
```