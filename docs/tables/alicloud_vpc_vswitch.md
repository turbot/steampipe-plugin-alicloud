# Table: alicloud_vpc_vswitch

Alicloud VSwitch is a logical subdivision of an IP network. It enables dividing a network into two or more networks.

## Examples

### Basic info

```sql
select
  name,
  vswitch_id,
  status,
  cidr_block,
  zone_id,
  is_default
from
  alicloud_vpc_vswitch;
```


### Get the number of available IP addresses in each VSwitch

```sql
select
  name,
  vswitch_id,
  available_ip_address_count,
  power(2, 32 - masklen(cidr_block :: cidr)) -1 as raw_size
from
  alicloud_vpc_vswitch;
```

### Route Table info associated with VSwitch

```sql
select
  name,
  vswitch_id,
  route_table ->> 'RouteTableId' as route_table_id,
  route_table ->> 'RouteTableType' as route_table_type,
  route_table -> 'RouteEntrys' -> 'RouteEntry' as route_entry
from
  alicloud_vpc_vswitch;
```


### VSwitch count by VPC ID

```sql
select
  vpc_id,
  count(vswitch_id) as vswitch_count
from
  alicloud_vpc_vswitch
group by
  vpc_id;
```