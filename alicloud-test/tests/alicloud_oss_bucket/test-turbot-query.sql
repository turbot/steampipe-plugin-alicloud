select
  title,
  tags,
  akas
from
  alicloud_oss_bucket
where
  name = '{{ output.bucket_id.value }}';