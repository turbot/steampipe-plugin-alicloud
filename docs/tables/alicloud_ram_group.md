# Table: alicloud_ram_group

A RAM group is a collection of RAM users. Groups let you specify permissions for multiple users, which makes it easier to manage the permissions for those users.

## Examples

### User details associated with each RAM group

```sql
select
  name as group_name,
  iam_user ->> 'UserName' as user_name,
  iam_user ->> 'DisplayName' as display_name,
  iam_user ->> 'JoinDate' as user_join_date
from
  alicloud_ram_group
  cross join jsonb_array_elements(users) as iam_user;
```

### List the policies attached to each RAM group

```sql
select
  name as group_name,
  policies ->> 'PolicyName' as policy_name,
  policies ->> 'PolicyType' as policy_type,
  policies ->> 'DefaultVersion' as policy_default_version,
  policies ->> 'AttachDate' as policy_attachment_date
from
  alicloud_ram_group,
  jsonb_array_elements(attached_policy) as policies;
```

### List of RAM groups with no users added to it

```sql
select
  name as group_name,
  create_date,
  users
from
  alicloud_ram_group
where
  users = '[]';
```
