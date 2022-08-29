select cluster_id, instance_id
from alicloud_cs_kubernetes_cluster_node
where cluster_id = '{{ output.cluster_id.value }}' and instance_id = '{{ output.instance_id.value }}';