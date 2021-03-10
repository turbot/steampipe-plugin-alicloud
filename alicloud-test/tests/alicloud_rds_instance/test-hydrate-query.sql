select
  db_instance_ip_array_name,
  db_instance_ip_array_attribute,
  security_ip_type
from
  alicloud_rds_instance
where
  db_instance_id = '{{ output.db_instance_id.value }}';