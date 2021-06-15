select host_name, instance_id, title, region
from alicloud_cms_monitor_host
where akas::text = '["{{ output.resource_aka.value }}"]' and instance_id = '{{  output.instance_id.value  }}'