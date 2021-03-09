# Table: alicloud_compute_zone

When you call this operation, only a list of zones and some resource information related to each zone is returned. If you want to query instance types and disk categories that are available for purchase in a specified zone, we recommend that you call the DescribeAvailableResource operation.

## Examples

### Zone basic info

```sql
select
  zone_id,
  local_name,
  available_resource_creation,
  available_volume_categories,
  available_instance_types
from
  alicloud_compute_zone;
```

### Details of a particular zone

```sql
select
  zone_id,
  local_name,
  available_resource_creation,
  available_volume_categories,
  available_instance_types
from
  alicloud_compute_zone where zone_id = 'ap-south-1b';
```

