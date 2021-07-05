select name, trail_region,event_rw
from alicloud_action_trail
where akas::text = '["{{ output.resource_aka.value }}"]'