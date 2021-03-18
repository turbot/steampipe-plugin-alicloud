select name, key_pair_finger_print
from alicloud_ecs_key_pair
where name = '{{output.key_name.value}}';