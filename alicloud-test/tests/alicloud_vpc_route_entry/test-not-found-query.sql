select 
  name, 
  route_table_id
from 
  alicloud_vpc_route_entry
where 
  route_table_id = 'dummy-{{ output.route_table_id.value }}';