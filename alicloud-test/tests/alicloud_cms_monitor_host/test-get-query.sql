select host_name, instance_id, title, region
from alicloud_cms_monitor_host
where host_name = '{{ output.host_name.value }}' and instance_id = '{{ output.instance_id.value }}';