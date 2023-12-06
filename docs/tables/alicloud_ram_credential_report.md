---
title: "Steampipe Table: alicloud_ram_credential_report - Query Alicloud RAM Credential Reports using SQL"
description: "Allows users to query Alicloud RAM Credential Reports, providing insights into the credential security status of RAM users."
---

# Table: alicloud_ram_credential_report - Query Alicloud RAM Credential Reports using SQL

Alicloud RAM (Resource Access Management) is a service that helps you manage user identities and control their access to your resources. It allows you to create and manage multiple identities under your Alicloud account and grant permissions to these identities to access your Alicloud resources. The RAM Credential Report is a document that provides information about the credential security status of RAM users.

## Table Usage Guide

The `alicloud_ram_credential_report` table provides insights into the credential security status of RAM users within Alicloud RAM. As a security administrator, explore user-specific details through this table, including password status, MFA device bindings, and access key usage. Utilize it to uncover information about users, such as those with high-risk passwords or inactive MFA devices, and to monitor the usage of access keys.

**Important Notes**
- This table requires a valid credential report to exist. To generate it, please run the follow Aliyun CLI command:
  - `aliyun ims GenerateCredentialReport --endpoint ims.aliyuncs.com`

## Examples

### List users that have logged into the console in the past 90 days

```sql
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

### List users that have NOT logged into the console in the past 90 days

```sql
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

### List users with console access that have never logged in to the console

```sql
select
  user_name
from
  alicloud_ram_credential_report
where
  password_exist
  and user_last_logon is null;
```

### Find access keys older than 90 days

```sql
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

### Find users that have a console password but do not have MFA enabled

```sql
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

### Check if root login has MFA enabled

```sql
select
  user_name,
  mfa_active
from
  alicloud_ram_credential_report
where
  user_name = '<root>';
```
