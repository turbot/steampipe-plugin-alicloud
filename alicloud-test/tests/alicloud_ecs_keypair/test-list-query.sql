select name, key_pair_finger_print
from alicloud_ecs_keypair
where akas::text = '["{{output.resource_aka.value}}"]';