select name, title, akas, account_id
from alicloud_vpc_network_acl
where network_acl_id = '{{ output.network_acl_id.value }}';