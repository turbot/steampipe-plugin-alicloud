---
title: "Steampipe Table: alicloud_kms_key - Query Alibaba Cloud Key Management Service Keys using SQL"
description: "Allows users to query Alibaba Cloud Key Management Service Keys, providing insights into key metadata, usage, and status."
---

# Table: alicloud_kms_key - Query Alibaba Cloud Key Management Service Keys using SQL

Alibaba Cloud Key Management Service (KMS) is a secure and easy-to-use service to create, control, and manage cryptographic keys used to secure your data. It provides centralized management of cryptographic keys, and offers a range of features including key rotation, key version management, and audit trails for key usage. KMS is integrated with other Alibaba Cloud services to help protect the data you store in these services and control the keys that decrypt it.

## Table Usage Guide

The `alicloud_kms_key` table provides insights into cryptographic keys within Alibaba Cloud Key Management Service (KMS). As a security engineer, you can explore key-specific details through this table, including key state, key spec, and associated metadata. Utilize it to uncover information about keys, such as their creation time, description, and the key material expiration status.

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