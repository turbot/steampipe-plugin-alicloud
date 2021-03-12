select name, route_table_id, description, route_table_type, router_id, router_type, vswitch_ids, vpc_id, owner_id, tags_src, region, account_id
from alicloud_vpc_route_table
where route_table_id = '{{ output.resource_id.value }}';