select db_instance_id, backup_id, store_status
from alicloud_rds_backup
where backup_id = '{{ output.resource_id.value }}';