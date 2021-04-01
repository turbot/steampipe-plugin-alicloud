select distinct(region), akas, title
from alicloud_ecs_region
where region = '{{ output.current_region_id.value }}';