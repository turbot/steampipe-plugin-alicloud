# Table: alicloud_ecs_region

Elastic Compute resources are hosted in multiple locations worldwide. These locations are composed of regions and zones.
A region is a geographic area where a data center resides.

## Examples

### Region basic info

```sql
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_ecs_region;
```

### Details of a particular region

```sql
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_ecs_region where region = 'us-east-1';
```
