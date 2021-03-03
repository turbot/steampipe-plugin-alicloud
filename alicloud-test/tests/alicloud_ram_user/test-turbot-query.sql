select name, akas, title
from alicloud_ram_user
where name = '{{ resourceName }}';