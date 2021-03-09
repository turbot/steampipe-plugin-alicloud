select zone_id
from alicloud_compute_zone
where zone_id = 'dummy-{{ output.zone_id.value }}';