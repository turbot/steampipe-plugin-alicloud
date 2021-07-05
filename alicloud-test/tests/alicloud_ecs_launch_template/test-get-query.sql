select name, launch_template_id, created_by, default_version_number, latest_version_number, region, account_id
from alicloud_ecs_launch_template
where launch_template_id = '{{ output.resource_id.value }}';