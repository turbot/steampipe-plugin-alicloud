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

### List of KMS keys scheduled for deletion

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

### Get the keys where automatic key rotation is suspended

```sql
select
  key_id,
  automatic_rotation
from
  alicloud_kms_key
where
  automatic_rotation = 'Suspended';
```

### Get the key alias info

```sql
select
  alias ->> 'KeyId' as key_id,
  alias ->> 'AliasArn' as alias_arn,
  alias ->> 'AliasName' as alias_name
from
  alicloud_kms_key,
  jsonb_array_elements(key_aliases) as alias;
```

### Count of key per region

```sql
select
  region,
  count(*)
from
  alicloud_kms_key
group by
  region;
```
