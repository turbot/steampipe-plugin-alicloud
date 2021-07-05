# Table: alicloud_oss_bucket

An OSS bucket is the container used to store objects. All objects are contained in buckets. You can configure a variety of bucket properties such as the region, ACL, and storage class. You can create buckets of different storage classes to store data based on your requirements.

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
