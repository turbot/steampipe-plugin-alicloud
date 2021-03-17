select name, network_interface_id, type, owner_id, service_managed, description, mac_address, vswitch_id, vpc_id, zone_id, security_group_ids, tags_src, region, account_id
from alicloud_ecs_network_interface
where network_interface_id = '{{ output.network_interface_id.value }}';