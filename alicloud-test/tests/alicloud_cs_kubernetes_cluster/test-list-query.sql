select name, cluster_id, state
from alicloud_cs_kubernetes_cluster
where akas::text = '["{{ output.resource_aka.value }}"]';