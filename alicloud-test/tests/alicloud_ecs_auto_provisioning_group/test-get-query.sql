select name, auto_provisioning_group_id, auto_provisioning_group_type, launch_template_id, launch_template_version, terminate_instances, terminate_instances_with_expiration, launch_template_configs, region, account_id
from alicloud_ecs_auto_provisioning_group
where auto_provisioning_group_id = '{{ output.resource_id.value }}';