select name, nat_gateway_id, vpc_id, account_id, region
from alicloud_vpc_nat_gateway
where nat_gateway_id = '{{ output.nat_gateway_id.value }}';