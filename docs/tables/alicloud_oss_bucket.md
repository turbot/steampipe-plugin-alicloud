---
title: "Steampipe Table: alicloud_oss_bucket - Query Alibaba Cloud Object Storage Service Buckets using SQL"
description: "Allows users to query Alibaba Cloud Object Storage Service (OSS) Buckets, providing detailed information about each OSS bucket such as its name, location, storage class, and creation time."
folder: "OSS"
---

# Table: alicloud_oss_bucket - Query Alibaba Cloud Object Storage Service Buckets using SQL

Alibaba Cloud Object Storage Service (OSS) is a cost-effective, highly secure, and easy-to-use object storage service that enables you to store, back up, and archive large amounts of data in the cloud. OSS is designed to store and retrieve any type of data, at any time, from anywhere on the web. It provides massive, secure, durable, and highly available storage capacity.

## Table Usage Guide

The `alicloud_oss_bucket` table provides insights into OSS buckets within Alibaba Cloud Object Storage Service. As a cloud architect or developer, explore bucket-specific details through this table, including the bucket's name, location, storage class, and creation time. Utilize it to manage and analyze your OSS buckets, such as identifying buckets that are using outdated storage classes or located in regions with higher costs.

## Examples

### List of buckets where versioning is not enabled
Discover the segments that have not enabled versioning in their storage buckets. This is useful to identify potential areas of risk, as versioning provides a means of recovery in case of accidental deletion or alteration of data.

```sql+postgres
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

```sql+sqlite
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
Explore which storage buckets lack default encryption, providing a useful way to identify potential security weaknesses in your data storage. This can help prioritize security enhancements and ensure data protection compliance.

```sql+postgres
select
  name,
  server_side_encryption
from
  alicloud_oss_bucket
where
  server_side_encryption ->> 'SSEAlgorithm' = '';
```

```sql+sqlite
select
  name,
  server_side_encryption
from
  alicloud_oss_bucket
where
  json_extract(server_side_encryption, '$.SSEAlgorithm') = '';
```

### List of buckets where public access to bucket is not blocked
Explore which storage buckets have public access enabled, which could potentially expose sensitive data. This is useful for identifying and mitigating security risks associated with unauthorized data access.

```sql+postgres
select
  name,
  acl
from
  alicloud_oss_bucket
where
  acl <> 'private';
```

```sql+sqlite
select
  name,
  acl
from
  alicloud_oss_bucket
where
  acl <> 'private';
```

### List of buckets where server access logging destination is same as the source bucket
Determine the areas in which server access logging destinations are identical to their source buckets. This is useful for identifying potential security risks, as it could indicate a lack of segregation between log data and source data.

```sql+postgres
select
  name,
  logging ->> 'TargetBucket' as target_bucket
from
  alicloud_oss_bucket
where
  logging ->> 'TargetBucket' = name;
```

```sql+sqlite
select
  name,
  json_extract(logging, '$.TargetBucket') as target_bucket
from
  alicloud_oss_bucket
where
  json_extract(logging, '$.TargetBucket') = name;
```

### List of buckets without owner tag key
Explore which AliCloud OSS buckets lack an assigned owner. This can be crucial in managing resources and ensuring accountability within your cloud storage environment.

```sql+postgres
select
  name,
  tags
from
  alicloud_oss_bucket
where
  tags ->> 'owner' is null;
```

```sql+sqlite
select
  name,
  tags
from
  alicloud_oss_bucket
where
  json_extract(tags, '$.owner') is null;
```

### List of Bucket policy statements that grant external access
Identify instances where your OSS bucket policies may be granting external access. This is beneficial for assessing potential security vulnerabilities and ensuring that your data is protected.

```sql+postgres
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

```sql+sqlite
select
  title,
  p.value as principal,
  a.value as action,
  json_extract(s.value, '$.Effect') as effect,
  json_extract(s.value, '$.Condition') as conditions
from
  alicloud_oss_bucket,
  json_each(policy, '$.Statement') as s,
  json_each(s.value, '$.Principal') as p,
  json_each(s.value, '$.Action') as a
where
  json_extract(s.value, '$.Effect') = 'Allow'
  and (
    p.value != account_id
    or p.value = '*'
  );
```

### List of buckets with no lifecycle policy
Explore which storage buckets are missing a lifecycle policy, allowing you to identify potential areas of risk and implement necessary changes to enhance data management. This is particularly useful in maintaining compliance and optimizing storage costs.

```sql+postgres
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

```sql+sqlite
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