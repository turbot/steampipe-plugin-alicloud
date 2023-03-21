select db_instance_id, backup_id, store_status
from alicloud_rds_backup
where db_instance_id = '{{ output.db_instance_id.value }}';