select
  name,
  default_version
from
  alicloud_ram_policy
where
  name = '{{ resourceName }}';