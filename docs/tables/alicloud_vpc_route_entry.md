# Table: alicloud_vpc_route_entry

Routes are set of rules that are used to determine where network traffic from the vswitch or gateway is directed.

## Examples

### Basic info

```sql
select
  name,
  route_table_id,
  description,
  instance_id,
  route_entry_id,
  destination_cidr_block,
  type,
  status
from
  alicloud_vpc_route_entry;
```

### List custom route entries

```sql
select
  name,
  route_table_id,
  description,
  instance_id,
  route_entry_id,
  destination_cidr_block,
  type,
  status
from
  alicloud_vpc_route_entry
where
  type = 'Custom';
```

### List route entries that have a next hop type of VPN gateway

```sql
select
  name,
  route_table_id,
  description,
  route_entry_id,
  destination_cidr_block,
  type,
  status
from
  alicloud_vpc_route_entry,
  jsonb_array_elements(next_hops) as next_hop
where
  next_hop ->> 'NextHopType' = 'VpnGateway';
```
