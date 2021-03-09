select id, local_name
from alicloud_compute_region
where id = '{{ output.current_region_id.value }}';