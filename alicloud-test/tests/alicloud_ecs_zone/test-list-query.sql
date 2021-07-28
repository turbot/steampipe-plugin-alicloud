select zone_id, region
from alicloud_ecs_zone
where zone_id = '{{ output.zone_id.value }}';