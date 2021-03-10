select name, route_table_id, nexthop_type
from alicloud_vpc_route_entry
where name = '{{ resourceName }}';