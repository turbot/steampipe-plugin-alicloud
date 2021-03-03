select name, vpn_gateway_id, status
from alicloud_vpc_vpn_gateway
where vpn_gateway_id = 'dummy-{{ resourceName }}';