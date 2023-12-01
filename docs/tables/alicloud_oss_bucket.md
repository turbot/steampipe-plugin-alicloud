---
title: "Steampipe Table: alicloud_oss_bucket - Query Alibaba Cloud Object Storage Service Buckets using SQL"
description: "Allows users to query Alibaba Cloud Object Storage Service (OSS) Buckets, providing detailed information about each OSS bucket such as its name, location, storage class, and creation time."
---

# Table: alicloud_oss_bucket - Query Alibaba Cloud Object Storage Service Buckets using SQL

Alibaba Cloud Object Storage Service (OSS) is a cost-effective, highly secure, and easy-to-use object storage service that enables you to store, back up, and archive large amounts of data in the cloud. OSS is designed to store and retrieve any type of data, at any time, from anywhere on the web. It provides massive, secure, durable, and highly available storage capacity.

## Table Usage Guide

The `alicloud_oss_bucket` table provides insights into OSS buckets within Alibaba Cloud Object Storage Service. As a cloud architect or developer, explore bucket-specific details through this table, including the bucket's name, location, storage class, and creation time. Utilize it to manage and analyze your OSS buckets, such as identifying buckets that are using outdated storage classes or located in regions with higher costs.

## Examples

### List of buckets where versioning is not enabled

```sql
select
  name,
  arn,
  region,
  account_id,
  versioning
from
  alicloud_oss_bucket
where
  versioning <> 'Enabled';
```

### List of buckets which do not have default encryption enabled

```sql
select
  name,
  server_side_encryption
from
  alicloud_oss_bucket
where
  server_side_encryption ->> 'SSEAlgorithm' = '';
```

### List of buckets where public access to bucket is not blocked

```sql
select
  name,
  acl
from
  alicloud_oss_bucket
where
  acl <> 'private';
```

### List of buckets where server access logging destination is same as the source bucket

```sql
select
  name,
  logging ->> 'TargetBucket' as target_bucket
from
  alicloud_oss_bucket
where
  logging ->> 'TargetBucket' = name;
```

### List of buckets without owner tag key

```sql
select
  name,
  tags
from
  alicloud_oss_bucket
where
  tags ->> 'owner' is null;
```

### List of Bucket policy statements that grant external access

```sql
select
  title,
  p as principal,
  a as action,
  s ->> 'Effect' as effect,
  s -> 'Condition' as conditions
from
  alicloud_oss_bucket,
  jsonb_array_elements(policy -> 'Statement') as s,
  jsonb_array_elements_text(s -> 'Principal') as p,
  jsonb_array_elements_text(s -> 'Action') as a
where
  s ->> 'Effect' = 'Allow'
  and (
    p != account_id
    or p = '*'
  );
```

### List of buckets with no lifecycle policy

```sql
select
  name,
  arn,
  region,
  account_id,
  lifecycle_rules
from
  alicloud_oss_bucket
where
  lifecycle_rules is null;
```
