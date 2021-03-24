select name, nat_gateway_id, nat_type, deletion_protection, auto_pay, billing_method, description, ecs_metric_enabled, vpc_id, account_id, region
from alicloud_vpc_nat_gateway
where nat_gateway_id = '{{ output.nat_gateway_id.value }}';