# Table: alicloud_ram_role

A RAM role is a virtual RAM identity that you can create in your Alibaba Cloud account. A RAM role does not have a specific logon password or AccessKey pair. A RAM user can be used only after the RAM user is assumed by a trusted entity.

## Examples

### List the policies attached to the roles

```sql
select
  name as group_name,
  policies ->> 'PolicyName' as policy_name,
  policies ->> 'PolicyType' as policy_type,
  policies ->> 'DefaultVersion' as policy_default_version,
  policies ->> 'AttachDate' as policy_attachment_date
from
  alicloud_ram_role
  cross join jsonb_array_elements(attached_policy) as policies
order by group_name;
```

### Find all roles having Administrator access

```sql
select
  name as user_name,
  policies ->> 'PolicyName' as policy_name
from
  alicloud_ram_role
  cross join jsonb_array_elements(attached_policy) as policies
where 
  policies ->> 'PolicyName' = 'AdministratorAccess';
```
