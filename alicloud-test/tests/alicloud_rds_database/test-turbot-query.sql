select title
from alicloud_rds_database
where db_instance_id = '{{ output.db_instance_id.value }}';