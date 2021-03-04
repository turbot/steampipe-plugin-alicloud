select instance_storage, instance_charge_type, monitoring_period
from alicloud_rds_instance
where db_instance_id = '{{ output.db_instance_id.value }}';