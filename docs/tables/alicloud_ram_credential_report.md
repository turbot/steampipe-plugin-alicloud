# Table: alicloud_ram_credential_report

Retrieves a credential report for the Alibaba Cloud account. For more
information about the credential report, see [Generate and download user
credential
reports](https://partners-intl.aliyun.com/help/doc-detail/143477.htm) in the
RAM Guide.

_Please note_: This table requires a valid credential report to exist. To
generate it, please run the follow Aliyun CLI command:

`aliyun ims GenerateCredentialReport --endpoint ims.aliyuncs.com`

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
