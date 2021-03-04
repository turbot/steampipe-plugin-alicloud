select name, customer_gateway_id
from alicloud_vpc_vpn_customer_gateway
where name = '{{ resourceName }}';