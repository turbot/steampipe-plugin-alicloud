---
title: "Steampipe Table: alicloud_vpc_route_table - Query Alibaba Cloud VPC Route Tables using SQL"
description: "Allows users to query Alibaba Cloud VPC route tables, including route table ID, name, VPC association, route entries, and type."
folder: "VPC"
---

# Table: alicloud_vpc_route_table - Query Alibaba Cloud VPC Route Tables using SQL

Alibaba Cloud Virtual Private Cloud (VPC) route tables define the paths network traffic takes within and outside the VPC. Each VPC is automatically associated with a default route table, and additional custom route tables can be created to manage routing behavior for subnets.

## Table Usage Guide

The `alicloud_vpc_route_table` table enables network engineers and cloud administrators to query detailed information about route tables in Alibaba Cloud. Use this table to retrieve attributes such as route table ID, name, type (system or custom), associated VPC ID, and route entries. This information is crucial for managing network routing, diagnosing connectivity issues, and enforcing traffic flow rules in your cloud infrastructure.

## Examples

### Basic info
This query is used to gain insights into the configuration of your Alicloud Virtual Private Cloud (VPC) by examining the details of its routing tables. It can be particularly useful in managing network traffic and ensuring optimal routing strategies within your VPC environment.

```sql+postgres
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

```sql+sqlite
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

### Get VPC and VSwitch attachment info for each route table
Explore the relationships between your virtual private cloud (VPC) and virtual switch (VSwitch) by understanding their attachment details for each route table. This can be particularly useful in managing network routing and ensuring efficient data traffic within your cloud environment.

```sql+postgres
select
  name,
  route_table_id,
  vpc_id,
  jsonb_array_elements_text(vswitch_ids)
from
  alicloud_vpc_route_table;
```

```sql+sqlite
select
  name,
  route_table_id,
  vpc_id,
  json_each.value
from
  alicloud_vpc_route_table,
  json_each(vswitch_ids);
```

### Routing details for each route table
Explore the intricate details of your network routing configurations. This query can help you understand how data is being directed across your network, which can be critical for troubleshooting connectivity issues or optimizing network performance.

```sql+postgres
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

```sql+sqlite
select
  route_table_id,
  json_extract(route_detail.value, '$.Description') as description,
  json_extract(route_detail.value, '$.DestinationCidrBlock') as destination_CIDR_block,
  json_extract(route_detail.value, '$.InstanceId') as instance_id,
  json_extract(route_detail.value, '$.IpVersion') as ip_version,
  json_extract(route_detail.value, '$.NextHopOppsiteInstanceId') as next_hop_oppsite_instance_id,
  json_extract(route_detail.value, '$.NextHopOppsiteRegionId') as next_hop_oppsite_region_id,
  json_extract(route_detail.value, '$.NextHopOppsiteType') as next_hop_oppsite_type,
  json_extract(route_detail.value, '$.NextHopRegionId') as next_hop_region_id,
  json_extract(route_detail.value, '$.NextHopType') as next_hop_type,
  json_extract(route_detail.value, '$.RouteEntryId') as route_entry_id,
  json_extract(route_detail.value, '$.RouteEntryName') as route_entry_name,
  json_extract(route_detail.value, '$.RouteTableId') as route_table_id,
  json_extract(route_detail.value, '$.Status') as status
from
  alicloud_vpc_route_table,
  json_each(route_entries) as route_detail;
```

### List route tables without application tag key
Determine the areas in which route tables lack an application tag key. This can be useful for identifying potential gaps in your tagging strategy, ensuring that all resources are correctly tagged for better management and organization.

```sql+postgres
select
  name,
  route_table_id
from
  alicloud_vpc_route_table
where
  not tags :: JSONB ? 'application';
```

```sql+sqlite
select
  name,
  route_table_id
from
  alicloud_vpc_route_table
where
  json_extract(tags, '$.application') is null;
```