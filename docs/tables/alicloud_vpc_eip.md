# Table: alicloud_vpc_eip

An independent public IP resource that decouples ECS and public IP resources, allowing you to flexibly manage public IP resources.

## Examples

### Basic info

```sql
select
  name,
  allocation_id,
  descritpion,
  ip_address,
  status,
  region
from
  alicloud_vpc_eip;
```

### Get the info of instance bound to eip

```sql
select
  name,
  allocation_id,
  instance_type instance_id,
  instance_region_id,
  bandwidth
from
  alicloud_vpc_eip;
```

### List all the available eips

```sql
select
  name,
  allocation_id,
  instance_type,
  status
from
  alicloud_vpc_eip
where
  status = 'Available';
```

### Get the eips where hd monitoring is off

```sql
select
  name,
  allocation_id,
  hd_monitor_status
from
  alicloud_vpc_eip
where
  hd_monitor_status = 'OFF';
```
