select name, akas, title
from alicloud_vpc_vpn_customer_gateway
where customer_gateway_id = '{{ output.resource_id.value }}';