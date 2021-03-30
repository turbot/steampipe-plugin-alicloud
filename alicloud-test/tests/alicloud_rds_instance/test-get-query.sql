select engine, engine_version, instance_type, monitoring_period, region, db_instance_id
from alicloud_rds_instance
where db_instance_id = '{{ output.db_instance_id.value }}';