# Table: alicloud_ram_user

Alibaba Cloud RAM users can login to the console or use access keys programmatically.

## Examples

### Basic user info

```sql
select
  user_id,
  name,
  display_name
from
  alicloud_ram_user;
```

### Users who have not logged in for 30 days

```sql
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date < current_date - interval '30 days';
```

### Users who have never logged in

```sql
select
  name,
  last_login_date
from
  alicloud_ram_user
where
  last_login_date is null;
```

### Groups details to which the RAM user belongs

```sql
select
  name as user_name,
  iam_group ->> 'GroupName' as group_name,
  iam_group ->> 'JoinDate' as join_date
from
  alicloud_ram_user,
  jsonb_array_elements(groups) as iam_group;
```

### List all the users having Administrator access

```sql
select
  name as user_name,
  policies ->> 'PolicyName' as policy_name,
  policies ->> 'PolicyType' as policy_type,
  policies ->> 'DefaultVersion' as policy_default_version,
  policies ->> 'AttachDate' as policy_attachment_date
from
  alicloud_ram_user,
  jsonb_array_elements(attached_policy) as policies
where
  policies ->> 'PolicyName' = 'AdministratorAccess';
```

### List all the users for whom MFA is not enabled

```sql
select
  name as user_name,
  user_id as user_id,
  mfa_enabled
from
  alicloud_ram_user
where
  not mfa_enabled;
```

### List all the users for whom role-based access control (RBAC) authorization is Enabled on Kubernetes Engine Clusters

```sql
select
  name as user_name,
  user_id as user_id
from
  alicloud_ram_user
where
  cs_user_permission <> '[]';
```
