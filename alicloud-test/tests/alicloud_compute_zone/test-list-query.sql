select zone_id, local_name, available_instance_types
from alicloud_compute_zone
where zone_id = '{{ output.zone_id.value }}';