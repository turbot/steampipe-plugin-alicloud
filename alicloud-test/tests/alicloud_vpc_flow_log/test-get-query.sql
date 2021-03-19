select name, flow_log_id,status
from alicloud_vpc_flow_log
where name = '{{ resourceName }}';