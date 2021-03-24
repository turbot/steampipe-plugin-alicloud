select name, cluster_id, state
from alicloud_elastic_container
where akas::text = '{{ output.resource_aka.value }}';