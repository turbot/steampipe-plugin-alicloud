select region
from alicloud_ecs_region
where region = 'dummy-{{ output.current_region_id.value }}';