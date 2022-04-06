select name, akas, title
from alicloud_vpc_dhcp_options_set
where dhcp_options_set_id = '{{ output.dhcp_options_set_id.value }}';