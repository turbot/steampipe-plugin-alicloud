select name, launch_template_id
from alicloud_ecs_launch_template
where launch_template_id = 'dummy-{{ resourceName }}';