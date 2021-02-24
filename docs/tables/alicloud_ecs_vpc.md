# Table: alicloud_vpc

A VPC is a virtual network in Alicloud.

## Examples

### Find default VPCs

```sql
select
  name,
  vpc_id,
  is_default,
  cidr_block,
  status,
  account_id,
  region_id
from
  alicloud_vpc
where
  is_default;
```

### Show CIDR details

```sql
select
  vpc_id,
  cidr_block,
  host(cidr_block),
  broadcast(cidr_block),
  netmask(cidr_block),
  network(cidr_block)
from
  alicloud_vpc;
```

### List VPCs with public CIDR blocks

```sql
select
  vpc_id,
  cidr_block,
  status,
  region_id
from
  alicloud_vpc
where
  not cidr_block <<= '10.0.0.0/8'
  and not cidr_block <<= '192.168.0.0/16'
  and not cidr_block <<= '172.16.0.0/12';
```

### Get the VSwitches details

```sql
select
  vswitch.vswitch_id,
  vpc.vpc_id,
  vswitch.cidr_block,
  vswitch.status,
  vswitch.available_ip_address_count,
  vswitch.zone_id
from
  alicloud_vpc as vpc
  join alicloud_vpc_vswitch as vswitch on vpc.vpc_id = vswitch.vpc_id;
```
