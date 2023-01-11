select
  bgp_peer_id,
  bgp_group_id,
  router_id,
  region,
  account_id
from
  alicloud_vpc_bgp_peer
where
  bgp_group_id = "{{ output.bgp_group_id.value }}";