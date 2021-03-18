select name
from alicloud_ecs_key_pair
where name = 'dummy-{{output.key_name.value}}';