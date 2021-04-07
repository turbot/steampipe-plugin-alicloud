select 
  route_table_id,
  title
from 
  alicloud_vpc_route_entry
where 
  route_table_id = '{{ output.route_table_id.value }}' 
  and destination_cidr_block = '172.11.1.1/32';