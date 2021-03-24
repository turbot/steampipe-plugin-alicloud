select name, cluster_id, size
from alicloud_elastic_container
where cluster_id = '{{ output.cluster_id.value }}';