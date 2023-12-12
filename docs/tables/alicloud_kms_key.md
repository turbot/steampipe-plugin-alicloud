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
Explore which encryption keys in your Alicloud account are currently in use and where. This query can help you manage your keys effectively by providing information about their state, creation date, and the region they are located in.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that are marked for deletion in the near future. This is useful for preemptively managing resources and ensuring system integrity by preventing unexpected loss of access to important keys.

```sql+postgres
select
  key_id,
  key_state,
  delete_date
from
  alicloud_kms_key
where
  key_state = 'PendingDeletion';
```

```sql+sqlite
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
Explore which encryption keys have had their automatic rotation feature suspended. This is useful for maintaining security standards, as keys that are not regularly rotated may pose a risk.

```sql+postgres
select
  key_id,
  automatic_rotation
from
  alicloud_kms_key
where
  automatic_rotation = 'Suspended';
```

```sql+sqlite
select
  key_id,
  automatic_rotation
from
  alicloud_kms_key
where
  automatic_rotation = 'Suspended';
```

### Get the key alias info for each key
Determine the alias details for each encryption key to manage and track your keys effectively. This helps in identifying and organizing your keys while maintaining security standards.

```sql+postgres
select
  alias ->> 'KeyId' as key_id,
  alias ->> 'AliasArn' as alias_arn,
  alias ->> 'AliasName' as alias_name
from
  alicloud_kms_key,
  jsonb_array_elements(key_aliases) as alias;
```

```sql+sqlite
select
  json_extract(alias.value, '$.KeyId') as key_id,
  json_extract(alias.value, '$.AliasArn') as alias_arn,
  json_extract(alias.value, '$.AliasName') as alias_name
from
  alicloud_kms_key,
  json_each(key_aliases) as alias;
```

### Count of keys per region
Example 1: "Count of keys per region"
Explore which regions have the most keys in your AliCloud Key Management Service. This can help you understand the distribution of your keys and identify regions with a high concentration of keys.

Example 2: "List keys that have deletion protection disabled"
Identify instances where keys in your AliCloud Key Management Service have deletion protection disabled. This can be useful in maintaining security standards and avoiding accidental data loss.

```sql+postgres
select
  region,
  count(*)
from
  alicloud_kms_key
group by
  region;
```

```sql+sqlite
select
  region,
  count(*)
from
  alicloud_kms_key
group by
  region;
```

## List keys that have deletion protection disabled

```sql+postgres
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

```sql+sqlite
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