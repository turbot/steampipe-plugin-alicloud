select cluster_id, instance_id, instance_name
from alicloud_cs_kubernetes_cluster_node
where instance_id = '{{ output.instance_id.value }}';