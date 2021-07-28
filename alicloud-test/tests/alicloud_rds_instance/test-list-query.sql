select db_instance_storage, zone_id
from alicloud_rds_instance
where akas::text = '["{{ output.resource_aka.value }}"]';