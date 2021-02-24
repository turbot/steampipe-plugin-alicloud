# Table: alicloud_ram_password_policy

RAM password policies can be used to ensure password complexity. It is recommended
that the password policy require at least one uppercase letter.

## Examples

### Ensure IAM password policy requires at least one uppercase letter (CIS v1.1.7)
```sql
select
  require_uppercase_characters
from
  alicloud_ram_password_policy;
```


### Ensure IAM password policy requires at least one lowercase letter (CIS v1.1.8)
```sql
select
  require_lowercase_characters
from
  alicloud_ram_password_policy;
```


### Ensure IAM password policy requires at least one symbol (CIS v1.1.9)
```sql
select
  require_symbols
from
  alicloud_ram_password_policy;
```


### Ensure IAM password policy require at least one number (CIS v1.1.10)
```sql
select
  require_numbers
from
  alicloud_ram_password_policy;
```


### Ensure IAM password policy requires minimum length of 14 or greater (CIS v1.1.11)
```sql
select
  minimum_password_length >= 14
from
  alicloud_ram_password_policy;
```


