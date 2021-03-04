select description, default_version,attachment_count
from alicloud_ram_policy
where name = '{{ resourceName }}';