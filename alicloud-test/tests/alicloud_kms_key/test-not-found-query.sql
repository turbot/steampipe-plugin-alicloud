select key_id, key_state, description
from alicloud_kms_key
where key_arn = 'dummy-{{ output.resource_aka.value }}';