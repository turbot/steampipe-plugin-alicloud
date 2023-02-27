select load_balancer_name, load_balancer_id, load_balancer_status, vpc_id
from alicloud_slb_load_balancer
where load_balancer_id = '{{ output.resource_id.value }}';