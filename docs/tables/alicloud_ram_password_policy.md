# Table: alicloud_ram_password_policy

RAM password policies can be used to ensure password complexity. It is recommendedvthat the password policy require at least one uppercase letter.

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
