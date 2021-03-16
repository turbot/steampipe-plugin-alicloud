select name, akas, title
from alicloud_ecs_keypair
where name = '{{output.key_name.value}}';