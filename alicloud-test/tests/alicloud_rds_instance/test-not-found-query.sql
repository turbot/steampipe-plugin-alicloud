select name, port
from alicloud_rds_instance
where name = 'dummy-{{ resourceName }}';