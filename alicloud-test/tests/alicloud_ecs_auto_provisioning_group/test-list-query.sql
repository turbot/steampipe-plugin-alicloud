select name, auto_provisioning_group_id, auto_provisioning_group_type
from alicloud_ecs_auto_provisioning_group
where name = '{{ resourceName }}';