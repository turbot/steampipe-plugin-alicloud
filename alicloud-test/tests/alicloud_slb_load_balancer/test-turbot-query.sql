select load_balancer_name, title
from alicloud_slb_load_balancer
where load_balancer_id = '{{ output.resource_id.value }}';