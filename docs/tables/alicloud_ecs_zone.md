# Table: alicloud_ecs_zone

A zone is a physical area with independent power grids and networks in a region.
Zones within the same region have access to each other, but faults within a single zone will not affect the others.

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
  alicloud_ecs_zone;
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
  alicloud_ecs_zone where zone_id = 'ap-south-1b';
```