select name, akas, title
from alicloud_vpc_route_entry
where route_table_id = '{{ output.route_table_id.value }}';