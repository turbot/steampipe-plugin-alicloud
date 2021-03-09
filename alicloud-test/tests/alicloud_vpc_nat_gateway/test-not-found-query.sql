select name, nat_gateway_id
from alicloud_vpc_nat_gateway
where nat_gateway_id = 'dummy-{{ resourceName }}';