select instance_storage, instance_charge_type, monitoring_period
from alicloud_rds_instance
where akas::text = '{{ output.resource_aka.value }}';