select name, akas, title
from alicloud_vpc_nat_gateway
where nat_gateway_id = '{{ output.nat_gateway_id.value }}';