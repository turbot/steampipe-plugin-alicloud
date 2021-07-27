select akas, title
from alicloud_rds_instance
where db_instance_id = '{{ output.db_instance_id.value }}';