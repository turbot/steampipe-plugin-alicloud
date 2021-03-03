select name, network_interface_id, owner_id
from alicloud_ecs_network_interface
where name = '{{ resourceName }}';