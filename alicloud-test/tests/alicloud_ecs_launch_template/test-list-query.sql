select name, launch_template_id
from alicloud_ecs_launch_template
where name = '{{ resourceName }}';