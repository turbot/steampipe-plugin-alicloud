select id, local_name
from alicloud_compute_region
where akas::text = '["{{output.resource_aka.value}}"]';