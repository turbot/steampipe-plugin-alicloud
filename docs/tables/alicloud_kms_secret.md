---
title: "Steampipe Table: alicloud_kms_secret - Query Alicloud Key Management Service Secrets using SQL"
description: "Allows users to query Alicloud Key Management Service Secrets, specifically the detailed information of the secrets including their version stages, rotation configuration, and recovery window."
folder: "KMS"
---

# Table: alicloud_kms_secret - Query Alicloud Key Management Service Secrets using SQL

Alicloud Key Management Service (KMS) Secrets is a feature of the Alicloud KMS that helps manage the lifecycle of secrets. It provides a secure and convenient method to create, use, and manage secrets, including database passwords, API keys, and other sensitive information. It also supports secret versioning and rotation to enhance the security of applications.

## Table Usage Guide

The `alicloud_kms_secret` table provides insights into secrets within Alicloud Key Management Service (KMS). As a security engineer, explore secret-specific details through this table, including their lifecycle stages, rotation configurations, and recovery windows. Utilize it to uncover information about secrets, such as their current status, the last time they were accessed, and whether they are scheduled for deletion.

## Examples

### Basic info
Explore the basic information of your encrypted data keys in Alibaba Cloud's Key Management Service. This allows you to understand the type of secrets you have, when they were created, and their overall descriptions, aiding in efficient key management.

```sql+postgres
select
  name,
  description,
  arn,
  secret_type,
  create_time
from
  alicloud_kms_secret;
```

```sql+sqlite
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
Uncover the details of encryption secrets that are not set to auto-renew, potentially exposing your system to security risks. This is useful for identifying and rectifying weak points in your security infrastructure.

```sql+postgres
select
  name,
  secret_type automatic_rotation
from
  alicloud_kms_secret
where
  automatic_rotation <> 'Enabled';
```

```sql+sqlite
select
  name,
  secret_type as automatic_rotation
from
  alicloud_kms_secret
where
  automatic_rotation != 'Enabled';
```

### List secrets that have not been rotated within the last 30 days
Explore which secrets have not been updated in the last month. This is useful for maintaining security standards and ensuring that sensitive information is regularly updated.

```sql+postgres
select
  name,
  secret_type,
  automatic_rotation
from
  alicloud_kms_secret
where
  last_rotation_date < (current_date - interval '30' day);
```

```sql+sqlite
select
  name,
  secret_type,
  automatic_rotation
from
  alicloud_kms_secret
where
  last_rotation_date < date('now','-30 day');
```

### Get the extended configuration info for each secret
This query is useful for gaining insights into the extended configuration details of each secret, such as the associated database name and instance ID, as well as the secret subtype. It can help in understanding and managing the security aspects of your cloud resources.

```sql+postgres
select
  name,
  extended_config -> 'CustomData' ->> 'DBName' as db_name,
  extended_config ->> 'DBInstanceId' as db_instance_id,
  extended_config ->> 'SecretSubType' as secret_sub_type
from
  alicloud_kms_secret;
```

```sql+sqlite
select
  name,
  json_extract(json_extract(extended_config, '$.CustomData'), '$.DBName') as db_name,
  json_extract(extended_config, '$.DBInstanceId') as db_instance_id,
  json_extract(extended_config, '$.SecretSubType') as secret_sub_type
from
  alicloud_kms_secret;
```

### List secrets without application tag key
Discover the segments that have secrets without an application tag key. This is useful to identify and manage secrets that may not be associated with a specific application.

```sql+postgres
select
  name,
  tags
from
  alicloud_kms_secret
where
  not tags :: JSONB ? 'application';
```

```sql+sqlite
select
  name,
  tags
from
  alicloud_kms_secret
where
  json_extract(tags, '$.application') is null;
```