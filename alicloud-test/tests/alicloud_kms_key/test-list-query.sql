select key_id, key_arn
from alicloud_kms_key
where key_arn = '{{ output.resource_aka.value }}';