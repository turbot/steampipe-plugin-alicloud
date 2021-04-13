select name, secret_type
from alicloud_kms_secret
where arn = '{{ output.resource_aka.value }}';