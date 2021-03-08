# Table: alicloud_ram_policy

Permissions are specified by a statement within a policy that allows or denies access to a specific Alibaba Cloud resource.
A policy defines a set of permissions that are described based on the policy structure and syntax. A policy can accurately describe the authorized resource sets, authorized operation sets, and authorization conditions.

## Examples

### List of Policies

```sql
select
  name,policy_type,description,default_version,policy_document
from
  alicloud_ram_policy;
```

### List of Policies where Policy type is System Policy

```sql
select
  name,policy_type,description,default_version,policy_document
from
  alicloud_ram_policy where policy_type = 'System';
```

### List of Policies where Policy type is Custom Policy

```sql
select
  name,policy_type,description,default_version,policy_document
from
  alicloud_ram_policy where policy_type = 'Custom';
```