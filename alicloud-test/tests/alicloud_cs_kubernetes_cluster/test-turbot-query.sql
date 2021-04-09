select name, akas, title
from alicloud_elastic_container
where cluster_id = '{{ output.cluster_id.value }}';