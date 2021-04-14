select name, trail_region, status, oss_bucket_name, region
from alicloud_action_trail
where name = '{{ resourceName }}';