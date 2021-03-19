select name, akas, title
from alicloud_vpc_flow_log
where name = '{{ resourceName }}';