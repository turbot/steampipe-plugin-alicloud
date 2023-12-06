# Table: alicloud_ram_security_preference

Alibaba Cloud RAM users security preference provides better security to user.

## Examples

### Basic security preference info
This query is useful to gain insights into the different security preferences and their settings in Alicloud RAM. It helps in assessing whether users have the permissions to change passwords, manage access keys, MFA devices, public keys, and the duration of login sessions, thereby aiding in understanding the security posture of your Alicloud environment.

```sql+postgres
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

```sql+sqlite
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
Explore which Alicloud users have the ability to change their passwords. This can be crucial for maintaining account security and ensuring users can manage their own credentials.

```sql+postgres
select
  allow_user_to_change_password
from
  alicloud_ram_security_preference;
```

```sql+sqlite
select
  allow_user_to_change_password
from
  alicloud_ram_security_preference;
```

### Check if user have access to manage public access key
Determine if users have the necessary permissions to manage public access keys. This can help in maintaining security by ensuring only authorized individuals can handle sensitive keys.

```sql+postgres
select
  allow_user_to_manage_public_keys
from
  alicloud_ram_security_preference;
```

```sql+sqlite
select
  allow_user_to_manage_public_keys
from
  alicloud_ram_security_preference;
```

### Get the log on session duration of User
Analyze the duration of user login sessions to understand their activity patterns and potential security risks.

```sql+postgres
select
  login_session_duration
from
  alicloud_ram_security_preference;
```

```sql+sqlite
select
  login_session_duration
from
  alicloud_ram_security_preference;
```