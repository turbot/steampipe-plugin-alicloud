select name, vpn_gateway_id, status, auto_propagate, billing_method, business_status, description, internet_ip, spec, ssl_max_connections, ssl_vpn, vswitch_id, vpc_id, region, account_id
from alicloud_vpc_vpn_gateway
where vpn_gateway_id = '{{ output.vpn_gateway_id.value }}';