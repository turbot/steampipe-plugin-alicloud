select name, title, akas
from alicloud_vpc_network_acl
where network_acl_id = 'dummy-{{ output.network_acl_id.value }}';