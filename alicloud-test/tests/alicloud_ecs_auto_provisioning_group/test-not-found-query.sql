select name, auto_provisioning_group_id, auto_provisioning_group_type
from alicloud_ecs_auto_provisioning_group
where auto_provisioning_group_id = 'dummy-{{ resourceName }}';