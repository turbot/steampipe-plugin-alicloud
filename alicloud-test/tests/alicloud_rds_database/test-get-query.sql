select db_name, db_instance_id, engine
from alicloud_rds_database
where db_instance_id = '{{ output.db_instance_id.value }}' and db_name = '{{ resourceName }}';