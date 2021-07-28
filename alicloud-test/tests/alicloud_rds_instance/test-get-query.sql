select engine, engine_version, db_instance_type, db_instance_id
from alicloud_rds_instance
where db_instance_id = '{{ output.db_instance_id.value }}';