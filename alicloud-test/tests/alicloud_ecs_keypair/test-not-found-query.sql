select name
from alicloud_ecs_keypair
where name = 'dummy-{{output.key_name.value}}';