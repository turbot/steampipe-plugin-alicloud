select name, key_pair_finger_print
from alicloud_ecs_keypair
where name = '{{output.key_name.value}}';