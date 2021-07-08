select akas, title
from alicloud_ram_policy
where policy_name = '{{ resourceName }}';