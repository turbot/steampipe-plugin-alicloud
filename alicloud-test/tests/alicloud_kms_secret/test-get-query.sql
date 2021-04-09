select name, arn, secret_type, description, tags_src, account_id, region
from alicloud_kms_secret
where name = '{{ resourceName }}';