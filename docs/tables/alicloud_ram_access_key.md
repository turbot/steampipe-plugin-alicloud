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

```sql
select
  access_key_id,
  user_name,
  create_date
from
  alicloud_ram_access_key;
```

### List of access keys which are inactive

```sql
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

```sql
select
  user_name,
  count (access_key_id) as access_key_count
from
  alicloud_ram_access_key
group by
  user_name;
```


### Access keys older than 90 days

```sql
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