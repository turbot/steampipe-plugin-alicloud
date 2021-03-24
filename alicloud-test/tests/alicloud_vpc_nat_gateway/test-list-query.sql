select name, nat_gateway_id, nat_type, description
from alicloud_vpc_nat_gateway
where name = '{{ resourceName }}';