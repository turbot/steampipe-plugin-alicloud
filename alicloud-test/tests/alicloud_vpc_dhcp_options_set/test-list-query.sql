select name, dhcp_options_set_id
from alicloud_vpc_dhcp_options_set
where name = '{{ resourceName }}';