select
  policy_name,
  default_version
from
  alicloud_ram_policy
where
  policy_name = '{{ resourceName }}' and policy_type = '{{ output.policy_type.value }}';