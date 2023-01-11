select
  router_id,
  title
from
  alicloud_vpc_route_entry
where
  router_id = '{{ output.router_id.value }}';