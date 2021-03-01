select name, user_id, display_name
from alicloud_ram_user
where name = 'dummy-{{ resourceName }}';