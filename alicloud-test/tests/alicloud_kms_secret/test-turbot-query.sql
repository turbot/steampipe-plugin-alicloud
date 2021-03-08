select name, akas, title, tags
from alicloud_kms_secret
where name = '{{ resourceName }}';