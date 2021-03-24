select name, akas, title
from alicloud_ecs_key_pair
where name = '{{output.key_name.value}}';