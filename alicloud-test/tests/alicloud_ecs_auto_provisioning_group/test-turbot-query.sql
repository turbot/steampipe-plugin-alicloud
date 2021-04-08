select name, akas, title
from alicloud_ecs_auto_provisioning_group
where auto_provisioning_group_id = '{{ output.resource_id.value }}';