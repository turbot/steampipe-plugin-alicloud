select name, akas, title, account_id, tags
from alicloud_kms_secret
where name = '{{ resourceName }}';