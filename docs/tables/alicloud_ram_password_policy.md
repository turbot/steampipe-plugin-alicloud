---
title: "Steampipe Table: alicloud_ram_password_policy - Query Alibaba Cloud RAM Password Policies using SQL"
description: "Allows users to query Alibaba Cloud RAM Password Policies, providing comprehensive details about password policies applied to Alibaba Cloud RAM users."
---

# Table: alicloud_ram_password_policy - Query Alibaba Cloud RAM Password Policies using SQL

Alibaba Cloud's Resource Access Management (RAM) is a service that helps manage user identities and access control. You can create and manage multiple identities under your Alibaba Cloud account, and control the access of these identities to your Alibaba Cloud resources. RAM allows you to grant fine-grained permissions and authorized access methods to users under your Alibaba Cloud account in a secure and controllable manner.

## Table Usage Guide

The `alicloud_ram_password_policy` table provides insights into password policies within Alibaba Cloud Resource Access Management (RAM). As a security analyst, you can explore password policy details through this table, including minimum password length, password complexity requirements, and password change frequency. Use it to ensure that password policies comply with your organization's security standards and to identify any potential security risks.

## Examples

### Ensure RAM password policy requires at least one uppercase letter (CIS v1.1.7)

```sql
select
  require_uppercase_characters,
  case require_uppercase_characters
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy requires at least one lowercase letter (CIS v1.1.8)

```sql
select
  require_lowercase_characters,
  case require_lowercase_characters
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy requires at least one symbol (CIS v1.1.9)

```sql
select
  require_symbols,
  case require_symbols
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy require at least one number (CIS v1.1.10)

```sql
select
  require_numbers,
  case require_numbers
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy requires minimum length of 14 or greater (CIS v1.1.11)

```sql
select
  minimum_password_length,
  case minimum_password_length >= 14
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```
