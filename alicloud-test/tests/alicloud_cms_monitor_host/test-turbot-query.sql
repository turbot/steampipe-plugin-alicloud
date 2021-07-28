select host_name, akas, title
from alicloud_cms_monitor_host
where host_name = '{{ output.host_name.value }}';