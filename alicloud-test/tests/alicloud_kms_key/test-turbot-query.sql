select key_id, title, region, akas, account_id
from alicloud_kms_key
where key_arn = '{{ output.resource_aka.value }}';