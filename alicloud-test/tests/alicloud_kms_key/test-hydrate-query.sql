select key_id, key_aliases
from alicloud_kms_key
where key_id = '{{ output.resource_id.value }}' and region = '{{ output.region.value }}';;