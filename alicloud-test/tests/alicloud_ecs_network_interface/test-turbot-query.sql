select name, akas, title, tags
from alicloud_ecs_network_interface
where network_interface_id = '{{ output.network_interface_id.value }}';