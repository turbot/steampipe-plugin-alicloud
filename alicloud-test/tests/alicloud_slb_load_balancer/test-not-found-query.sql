select load_balancer_name, load_balancer_id
from alicloud_slb_load_balancer
where load_balancer_name = 'dummy-{{ resourceName }}';