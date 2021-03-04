select name, akas, title
from alicloud_ram_policy
where name = '{{ resourceName }}';