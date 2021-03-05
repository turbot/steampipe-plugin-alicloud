select name, akas, title
from alicloud_vpc_vpn_gateway
where vpn_gateway_id = '{{ output.vpn_gateway_id.value }}';