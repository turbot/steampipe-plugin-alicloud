# Table: alicloud_ram_security_preference

Alibaba Cloud RAM users security preference provides better security to user.

## Examples

### Basic security preference info

```sql
select
  allow_user_to_change_password,
  allow_user_to_manage_access_keys,
  allow_user_to_manage_mfa_devices,
  allow_user_to_manage_public_keys,
  enable_save_mfa_ticket,
  login_session_duration
from
  alicloud_ram_security_preference;
```

### Check if user have access to change password

```sql
select
  allow_user_to_change_password
from
  alicloud_ram_security_preference;
```

### Check if user have access to manage public access key

```sql
select
  allow_user_to_manage_public_keys
from
  alicloud_ram_security_preference;
```

### Get the log on session duration of User

```sql
select
  login_session_duration
from
  alicloud_ram_security_preference;
```
