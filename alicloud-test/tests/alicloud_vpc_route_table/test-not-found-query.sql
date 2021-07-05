select name, route_table_id
from alicloud_vpc_route_table
where route_table_id = 'dummy-{{ resourceName }}';