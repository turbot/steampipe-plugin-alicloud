select name, flow_log_id
from alicloud_vpc_flow_log
where akas::text = '["{{ output.resource_aka.value }}"]'