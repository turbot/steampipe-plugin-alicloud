select name, user_id, display_name, email, mobile_phone, comments, region, account_id
from alicloud_ram_user
where name = '{{ resourceName }}';