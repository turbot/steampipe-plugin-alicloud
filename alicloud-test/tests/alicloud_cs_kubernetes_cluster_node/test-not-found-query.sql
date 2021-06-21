select node_name, cluster_id, instance_id
from alicloud_cs_kubernetes_cluster_node
where node_name = 'dummy-{{ resourceName }}';