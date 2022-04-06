select name, dhcp_options_set_id, status, description, account_id, region
from alicloud_vpc_dhcp_options_set
where dhcp_options_set_id = '{{ output.dhcp_options_set_id.value }}';