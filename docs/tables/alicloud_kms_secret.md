---
title: "Steampipe Table: alicloud_kms_secret - Query Alicloud Key Management Service Secrets using SQL"
description: "Allows users to query Alicloud Key Management Service Secrets, specifically the detailed information of the secrets including their version stages, rotation configuration, and recovery window."
---

# Table: alicloud_kms_secret - Query Alicloud Key Management Service Secrets using SQL

Alicloud Key Management Service (KMS) Secrets is a feature of the Alicloud KMS that helps manage the lifecycle of secrets. It provides a secure and convenient method to create, use, and manage secrets, including database passwords, API keys, and other sensitive information. It also supports secret versioning and rotation to enhance the security of applications.

## Table Usage Guide

The `alicloud_kms_secret` table provides insights into secrets within Alicloud Key Management Service (KMS). As a security engineer, explore secret-specific details through this table, including their lifecycle stages, rotation configurations, and recovery windows. Utilize it to uncover information about secrets, such as their current status, the last time they were accessed, and whether they are scheduled for deletion.

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
