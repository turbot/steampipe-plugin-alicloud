select region_id
from alicloud_compute_region
where region_id = 'dummy-{{ output.current_region_id.value }}';