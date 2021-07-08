select policy_name, description
from alicloud_ram_policy
where akas::text = '["{{output.resource_aka.value}}"]';