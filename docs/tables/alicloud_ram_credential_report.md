---
title: "Steampipe Table: alicloud_ram_credential_report - Query Alicloud RAM Credential Reports using SQL"
description: "Allows users to query Alicloud RAM Credential Reports, providing insights into the credential security status of RAM users."
folder: "RAM"
---

# Table: alicloud_ram_credential_report - Query Alicloud RAM Credential Reports using SQL

Alicloud RAM (Resource Access Management) is a service that helps you manage user identities and control their access to your resources. It allows you to create and manage multiple identities under your Alicloud account and grant permissions to these identities to access your Alicloud resources. The RAM Credential Report is a document that provides information about the credential security status of RAM users.

## Table Usage Guide

The `alicloud_ram_credential_report` table provides insights into the credential security status of RAM users within Alicloud RAM. As a security administrator, explore user-specific details through this table, including password status, MFA device bindings, and access key usage. Utilize it to uncover information about users, such as those with high-risk passwords or inactive MFA devices, and to monitor the usage of access keys.

## Examples

### List users that have logged into the console in the past 90 days
Determine the areas in which users have been active in the past 90 days, focusing on those who have logged into the console. This can help in understanding user engagement and identifying patterns in user activity.

```sql+postgres
select
  user_name,
  user_last_logon
from
  alicloud_ram_credential_report
where
  password_exist
  and password_active
  and user_last_logon > (current_date - interval '90' day);
```

```sql+sqlite
select
  user_name,
  user_last_logon
from
  alicloud_ram_credential_report
where
  password_exist = 1
  and password_active = 1
  and user_last_logon > date('now', '-90 day');
```

### List users that have NOT logged into the console in the past 90 days
Determine the areas in which users have not been active for over 90 days, focusing on those with existing and active passwords. This is useful for identifying potentially dormant or unused accounts, helping to maintain security and efficiency within your system.

```sql+postgres
select
  user_name,
  user_last_logon,
  age(user_last_logon)
from
  alicloud_ram_credential_report
where
  password_exist
  and password_active
  and user_last_logon <= (current_date - interval '90' day)
order by
  user_last_logon;
```

```sql+sqlite
select
  user_name,
  user_last_logon,
  julianday('now') - julianday(user_last_logon) as age
from
  alicloud_ram_credential_report
where
  password_exist
  and password_active
  and date(user_last_logon) <= date(julianday('now'), '-90 day')
order by
  user_last_logon;
```

### List users with console access that have never logged in to the console
Determine the users who have console access but have never actually logged in. This can help identify unused accounts, enabling better management of user access and improving security by eliminating potential vulnerabilities.

```sql+postgres
select
  user_name
from
  alicloud_ram_credential_report
where
  password_exist
  and user_last_logon is null;
```

```sql+sqlite
select
  user_name
from
  alicloud_ram_credential_report
where
  password_exist = 1
  and user_last_logon is null;
```

### Find access keys older than 90 days
Identify instances where user access keys are older than 90 days to ensure secure and up-to-date access management. This is useful for maintaining security standards and preventing potential unauthorized access.

```sql+postgres
select
  user_name,
  access_key_1_last_rotated,
  age(access_key_1_last_rotated) as access_key_1_age,
  access_key_2_last_rotated,
  age(access_key_2_last_rotated) as access_key_2_age
from
  alicloud_ram_credential_report
where
  access_key_1_last_rotated <= (current_date - interval '90' day)
  or access_key_2_last_rotated <= (current_date - interval '90' day)
order by
  user_name;
```

```sql+sqlite
select
  user_name,
  access_key_1_last_rotated,
  julianday('now') - julianday(access_key_1_last_rotated) as access_key_1_age,
  access_key_2_last_rotated,
  julianday('now') - julianday(access_key_2_last_rotated) as access_key_2_age
from
  alicloud_ram_credential_report
where
  julianday('now') - julianday(access_key_1_last_rotated) >= 90
  or julianday('now') - julianday(access_key_2_last_rotated) >= 90
order by
  user_name;
```

### Find users that have a console password but do not have MFA enabled
Determine the areas in which users have an active console password but lack multi-factor authentication (MFA). This query is useful for identifying potential security risks within your Alicloud resource access management.

```sql+postgres
select
  user_name,
  mfa_active,
  password_exist,
  password_active
from
  alicloud_ram_credential_report
where
  password_exist
  and password_active
  and not mfa_active;
```

```sql+sqlite
select
  user_name,
  mfa_active,
  password_exist,
  password_active
from
  alicloud_ram_credential_report
where
  password_exist = 1
  and password_active = 1
  and mfa_active = 0;
```

### Check if root login has MFA enabled
Determine if multi-factor authentication (MFA) is activated for the root login, enhancing security by requiring an additional verification step during authentication.

```sql+postgres
select
  user_name,
  mfa_active
from
  alicloud_ram_credential_report
where
  user_name = '<root>';
```

```sql+sqlite
select
  user_name,
  mfa_active
from
  alicloud_ram_credential_report
where
  user_name = '<root>';
```