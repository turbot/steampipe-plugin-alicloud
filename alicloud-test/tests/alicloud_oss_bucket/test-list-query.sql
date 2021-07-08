select
  name,
  arn,
  location,
  versioning,
  acl,
  server_side_encryption,
  lifecycle_rules,
  policy,
  tags_src,
  region,
  account_id
from
  alicloud_oss_bucket
where
  name = '{{ output.bucket_id.value }}';