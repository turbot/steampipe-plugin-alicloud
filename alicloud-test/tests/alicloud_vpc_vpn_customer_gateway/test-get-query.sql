select name, customer_gateway_id, description, ip_address, region, account_id
from alicloud_vpc_vpn_customer_gateway
where customer_gateway_id = '{{ output.resource_id.value }}';