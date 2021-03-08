# Table: alicloud_kms_secret

Secret enables to manage secrets in a centralized manner throughout their lifecycle (creation, retrieval, updating, and deletion).

## Examples

### Basic secret info

```sql
select
  name,
  description,
  arn,
  secret_type,
  create_time
from
  alicloud_kms_secret;
```

### Get the rotation info of secret

```sql
select
  name,
  automatic_rotation,
  last_rotation_date,
  rotation_interval,
  next_rotation_date
from
  alicloud_kms_secret;
```

### Get the extended configuration info of secret

```sql
select
  name,
  extended_config -> 'CustomData' ->> 'DBName' as db_name,
  extended_config ->> 'DBInstanceId' as db_instance_id,
  extended_config ->> 'SecretSubType' as secret_sub_type
from
  alicloud_kms_secret;
```

### List of secrets without application tag key

```sql
select
  name,
  tags
from
  alicloud_kms_secret
where
  not tags :: JSONB ? 'application';
```
