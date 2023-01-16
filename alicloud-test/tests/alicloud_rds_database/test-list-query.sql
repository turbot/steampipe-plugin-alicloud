select db_name, db_instance_id
from alicloud_rds_database
where db_name = '{{ resourceName }}';