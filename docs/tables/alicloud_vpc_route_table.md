# Table: alicloud_vpc_route_table

A route table contains a set of rules, called routes, that are used to determine where network traffic from your subnet or gateway is directed.

## Examples

### Basic info

```sql
select
  name,
  route_table_id,
  description,
  route_table_type,
  router_id,
  region
from
  alicloud_vpc_route_table;
```

### Get VPC and VSwitch info attached to route table

```sql
select
  name,
  route_table_id,
  jsonb_array_elements_text(vswitch_ids)
from
  alicloud_vpc_route_table;
```

### Routing details for each route table

```sql
select
  route_table_id,
  route_detail ->> 'Description' as description,
  route_detail ->> 'DestinationCidrBlock' as destination_CIDR_block,
  route_detail ->> 'InstanceId' as instance_id,
  route_detail ->> 'IpVersion' as ip_version,
  route_detail ->> 'NextHopOppsiteInstanceId' as next_hop_oppsite_instance_id,
  route_detail ->> 'NextHopOppsiteRegionId' as next_hop_oppsite_region_id,
  route_detail ->> 'NextHopOppsiteType' as next_hop_oppsite_type,
  route_detail ->> 'NextHopRegionId' as next_hop_region_id,
  route_detail ->> 'NextHopType' as next_hop_type,
  route_detail ->> 'RouteEntryId' as route_entry_id,
  route_detail ->> 'RouteEntryName' as route_entry_name,
  route_detail ->> 'RouteTableId' as route_table_id,
  route_detail ->> 'Status' as status
from
  alicloud_vpc_route_table,
  jsonb_array_elements(route_entries) as route_detail;
```

### List Route tables without application tags key

```sql
select
  name,
  route_table_id
from
  alicloud_vpc_route_table
where
  not tags :: JSONB ? 'application';
```
