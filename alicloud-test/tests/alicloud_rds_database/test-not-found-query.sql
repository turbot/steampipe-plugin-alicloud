select engine
from alicloud_rds_database
where db_instance_id = 'dummy{{ output.db_instance_id.value }}';