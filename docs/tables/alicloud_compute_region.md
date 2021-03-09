# Table: alicloud_compute_region

You can call this operation to query available Alibaba Cloud regions.

## Examples

### Region basic info

```sql
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_compute_region;
```

### Details of a particular region

```sql
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_compute_region where region = 'us-east-1';
```

