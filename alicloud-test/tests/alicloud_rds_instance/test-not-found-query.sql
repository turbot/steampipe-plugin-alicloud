select engine
from alicloud_rds_instance
where name = 'dummy-{{ resourceName }}';