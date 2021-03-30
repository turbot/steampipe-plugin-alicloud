select name, akas, title
from alicloud_ecs_launch_template
where launch_template_id = '{{ output.resource_id.value }}';