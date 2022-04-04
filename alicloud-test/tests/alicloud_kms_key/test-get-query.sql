select
  key_id,
  key_state,
  creator,
  description,
  automatic_rotation,
  protection_level,
  key_usage,
  key_spec,
  origin,
  akas,
  region,
  account_id
from
  alicloud_kms_key
where
  key_id = '{{ output.resource_id.value }}';
