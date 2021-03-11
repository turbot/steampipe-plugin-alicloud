# Table: alicloud_kms_secret

Secret enables to manage secrets in a centralized manner throughout their lifecycle (creation, retrieval, updating, and deletion).

## Examples

### Basic info

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

### List secrets that do not have automatic rotation enabled

```sql
select
  name,
  secret_type automatic_rotation
from
  alicloud_kms_secret
where
  automatic_rotation <> 'Enabled';
```

### List secrets that have not been rotated within the last 30 days

```sql
select
  name,
  secret_type,
  automatic_rotation
from
  alicloud_kms_secret
where
  last_rotation_date < (current_date - interval '30' day);
```

### Get the extended configuration info for each secret

```sql
select
  name,
  extended_config -> 'CustomData' ->> 'DBName' as db_name,
  extended_config ->> 'DBInstanceId' as db_instance_id,
  extended_config ->> 'SecretSubType' as secret_sub_type
from
  alicloud_kms_secret;
```

### List secrets without application tag key

```sql
select
  name,
  tags
from
  alicloud_kms_secret
where
  not tags :: JSONB ? 'application';
```
