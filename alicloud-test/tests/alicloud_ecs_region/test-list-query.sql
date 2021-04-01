select distinct(region)
from alicloud_ecs_region
where region = '{{ output.current_region_id.value }}';