select key_id, arn
from alicloud_kms_key
where arn = '{{ output.resource_aka.value }}';