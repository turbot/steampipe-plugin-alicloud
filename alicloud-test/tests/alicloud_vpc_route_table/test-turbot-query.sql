select name, akas, title, tags
from alicloud_vpc_route_table
where route_table_id = '{{ output.resource_id.value }}';