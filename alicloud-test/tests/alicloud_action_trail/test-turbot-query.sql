select name, akas, title
from alicloud_action_trail
where name = '{{ resourceName }}';