select id, akas, title
from alicloud_compute_region
where id = '{{ output.current_region_id.value }}';