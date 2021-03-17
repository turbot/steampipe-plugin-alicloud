select name, user_id, display_name
from alicloud_ram_user
where user_id = '{{ output.user_id.value }}';