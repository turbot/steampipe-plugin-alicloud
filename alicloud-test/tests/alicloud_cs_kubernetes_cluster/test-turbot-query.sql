select name, akas, title
from alicloud_cs_kubernetes_cluster
where cluster_id = '{{ output.cluster_id.value }}';