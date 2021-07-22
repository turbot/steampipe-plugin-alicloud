select name, trail_region,event_rw, home_region
from alicloud_action_trail
where akas::text = '["{{ output.resource_aka.value }}"]' and region = '{{ output.region_name.value }}'