select name, arn, rotation_interval, version_ids, title, account_id, tags
from alicloud_kms_secret
where name = '{{ resourceName }}';