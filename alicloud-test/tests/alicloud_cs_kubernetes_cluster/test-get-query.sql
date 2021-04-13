select name, cluster_id, size
from alicloud_cs_kubernetes_cluster
where cluster_id = '{{ output.cluster_id.value }}' and region = '{{ output.region.value }}';