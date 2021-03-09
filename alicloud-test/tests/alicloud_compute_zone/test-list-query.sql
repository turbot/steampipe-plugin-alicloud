select zone_id, local_name, available_instance_types
from alicloud_compute_zone
where akas::text = '["{{output.resource_aka.value}}"]';