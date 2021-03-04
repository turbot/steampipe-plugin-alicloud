select name, description
from alicloud_ram_policy
where name = '{{ resourceName }}';