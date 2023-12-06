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
Assess the elements within your Alicloud RAM password policy to verify if it mandates the inclusion of at least one uppercase letter. This aids in enhancing password security, aligning with the CIS v1.1.7 benchmark.

```sql+postgres
select
  require_uppercase_characters,
  case require_uppercase_characters
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

```sql+sqlite
select
  require_uppercase_characters,
  case require_uppercase_characters
    when 1 then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy requires at least one lowercase letter (CIS v1.1.8)
Assess the security of your password policy by determining if it necessitates the inclusion of at least one lowercase letter. This can help enhance your system's protection by ensuring passwords are more complex and harder to guess.

```sql+postgres
select
  require_lowercase_characters,
  case require_lowercase_characters
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

```sql+sqlite
select
  require_lowercase_characters,
  case require_lowercase_characters
    when 1 then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy requires at least one symbol (CIS v1.1.9)
This example helps in assessing the security of your password policy by determining whether it mandates the inclusion of at least one symbol. This is crucial for enhancing password strength and reducing the risk of unauthorized access.

```sql+postgres
select
  require_symbols,
  case require_symbols
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

```sql+sqlite
select
  require_symbols,
  case require_symbols
    when 1 then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy require at least one number (CIS v1.1.10)
Assess the elements within your Alicloud RAM password policy to ensure it mandates the inclusion of at least one numerical value, providing a simple pass or fail status. This aids in maintaining robust security standards as per the CIS v1.1.10 guidelines.

```sql+postgres
select
  require_numbers,
  case require_numbers
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

```sql+sqlite
select
  require_numbers,
  case require_numbers
    when 1 then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

### Ensure RAM password policy requires minimum length of 14 or greater (CIS v1.1.11)
Determine the strength of your password policy by checking if it requires a minimum length of 14 characters or more. This can help ensure your system's security by enforcing robust password requirements.

```sql+postgres
select
  minimum_password_length,
  case minimum_password_length >= 14
    when true then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```

```sql+sqlite
select
  minimum_password_length,
  case when minimum_password_length >= 14
    then 'pass'
    else 'fail'
  end as status
from
  alicloud_ram_password_policy;
```