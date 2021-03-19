select name, network_acl_id, vpc_id, region, description
from alicloud_vpc_network_acl
where name = '{{ resourceName }}';