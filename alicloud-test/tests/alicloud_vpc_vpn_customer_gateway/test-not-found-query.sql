select name, customer_gateway_id, asn
from alicloud_vpc_vpn_customer_gateway
where customer_gateway_id = 'dummy-{{ resourceName }}';