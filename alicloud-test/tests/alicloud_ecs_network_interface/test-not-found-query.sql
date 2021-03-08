select name, network_interface_id, type
from alicloud_ecs_network_interface
where network_interface_id = 'dummy-{{ resourceName }}';