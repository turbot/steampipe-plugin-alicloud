---
title: "Steampipe Table: alicloud_vpc_route_entry - Query Alicloud VPC Route Entries using SQL"
description: "Allows users to query Alicloud VPC Route Entries, providing detailed information on each route entry within the specified VPC."
---

# Table: alicloud_vpc_route_entry - Query Alicloud VPC Route Entries using SQL

Alicloud VPC Route Entry is a routing rule that determines the next hop for a network packet. It is part of Alicloud's Virtual Private Cloud (VPC) service and plays a crucial role in directing traffic within the VPC. The route entry determines the path that network traffic takes based on the destination IP address of the traffic.

## Table Usage Guide

The `alicloud_vpc_route_entry` table provides insights into the routing rules within Alicloud's Virtual Private Cloud (VPC). As a network administrator, explore route entry-specific details through this table, including the destination CIDR block, next hop type, and associated metadata. Utilize it to uncover information about route entries, such as those with specific next hop types, the traffic direction for each entry, and the status of each route entry.

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
