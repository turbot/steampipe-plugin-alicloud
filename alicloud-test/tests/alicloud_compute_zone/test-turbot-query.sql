select zone_id, akas, title
from alicloud_compute_zone
where zone_id = '{{ output.zone_id.value }}';