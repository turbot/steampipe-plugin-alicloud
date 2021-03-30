select zone_id, region
from alicloud_compute_zone
where zone_id = '{{ output.zone_id.value }}' and region = '{{ output.region_id.value }}';