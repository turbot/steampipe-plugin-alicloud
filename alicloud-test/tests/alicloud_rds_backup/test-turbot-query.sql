select title
from alicloud_rds_backup
where db_instance_id = '{{ output.db_instance_id.value }}' and backup_id = '{{ output.resource_id.value }}';