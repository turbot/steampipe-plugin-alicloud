select region
from alicloud_compute_region
where region = 'dummy-{{ output.current_region_id.value }}';