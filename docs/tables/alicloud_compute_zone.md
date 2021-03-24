# Table: alicloud_compute_zone

You can call this operation to query available Alibaba Cloud zones.

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

