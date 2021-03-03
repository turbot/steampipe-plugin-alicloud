# Table: alicloud_vpc_eip

A virtual private cloud service that provides an isolated cloud network to operate resources in a secure environment.

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


### Get the instance info that bound to eip

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

### Get all the available eips

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


### Get all the available eips

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


### Get the Eips where Hd Monitoring is off

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