select name, flow_log_id
from alicloud_vpc_flow_log
where name = 'dummy-{{ resourceName }}';