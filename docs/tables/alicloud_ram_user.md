# Table: alicloud_ram_user

Alibaba Cloud RAM users can login to the console or use access keys programatically.

## Examples

### Basic user info

```sql
select
  id,
  name,
  display_name
from
  alicloud_ram_user;
```

### Agents and admins (paid seats) who have not logged in for 30 days

```sql
select
  name,
  last_login_at
from
  alicloud_ram_user
where
  last_login_at < current_date - interval '30 days';
```
