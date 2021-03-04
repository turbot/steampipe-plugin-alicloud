SELECT
  key_id,
  key_state,
  description,
  automatic_rotation,
  protection_level,
  akas,
  region,
  account_id
FROM
  alicloud_kms_key
WHERE
  key_id = '{{ output.resource_id.value }}';

