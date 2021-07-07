select policy_name
from alicloud_ram_policy
where policy_name = 'dummy-{{ resourceName }}';