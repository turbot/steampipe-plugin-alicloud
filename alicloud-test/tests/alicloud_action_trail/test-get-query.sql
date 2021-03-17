select name, trail_region,status
from alicloud_action_trail
where name = '{{ resourceName }}';