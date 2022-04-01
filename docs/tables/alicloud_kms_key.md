# Table: alicloud_kms_key

A kms key can help user to protect data security in the transmission process.

## Examples

### Basic info

```sql
select
  key_id,
  arn,
  key_state,
  description,
  creation_date,
  region
from
  alicloud_kms_key;
```

### List keys scheduled for deletion

```sql
select
  key_id,
  key_state,
  delete_date
from
  alicloud_kms_key
where
  key_state = 'PendingDeletion';
```

### List keys that have automatic key rotation suspended

```sql
select
  key_id,
  automatic_rotation
from
  alicloud_kms_key
where
  automatic_rotation = 'Suspended';
```

### Get the key alias info for each key

```sql
select
  alias ->> 'KeyId' as key_id,
  alias ->> 'AliasArn' as alias_arn,
  alias ->> 'AliasName' as alias_name
from
  alicloud_kms_key,
  jsonb_array_elements(key_aliases) as alias;
```

### Count of keys per region

```sql
select
  region,
  count(*)
from
  alicloud_kms_key
group by
  region;
```
## List keys that have deletion protection disabled

```sql
select
  key_id,
  key_state,
  description,
  creation_date
from
  alicloud_kms_key
where
  deletion_protection = 'Disabled';
```