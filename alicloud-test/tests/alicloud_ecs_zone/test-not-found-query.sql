select zone_id
from alicloud_ecs_zone
where zone_id = 'dummy-{{ output.zone_id.value }}';