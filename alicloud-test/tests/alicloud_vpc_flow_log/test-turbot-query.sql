select akas, title
from alicloud_vpc_flow_log
where flow_log_id = '{{ output.resource_id.value }}';