---
title: "Steampipe Table: alicloud_vpc_vswitch - Query Alibaba Cloud VPC VSwitches using SQL"
description: "Allows users to query Alibaba Cloud VPC VSwitches, providing detailed information about each VSwitch within a Virtual Private Cloud."
folder: "VPC"
---

# Table: alicloud_vpc_vswitch - Query Alibaba Cloud VPC VSwitches using SQL

A VSwitch is a basic network device of a VPC in Alibaba Cloud, which connects different cloud product instances. It is similar to a traditional switch in a data center, which provides a communication channel for cloud product instances in the same network segment. VSwitches are used to create an isolated network environment, which can be associated with different cloud resources.

## Table Usage Guide

The `alicloud_vpc_vswitch` table provides insights into VSwitches within Alibaba Cloud Virtual Private Cloud (VPC). As a network administrator, explore VSwitch-specific details through this table, including its ID, status, creation time, and associated metadata. Utilize it to uncover information about VSwitches, such as their availability zones, CIDR blocks, and the VPCs they belong to.

## Examples

### Basic info
Explore the status and default settings of various zones within your virtual private cloud (VPC) to assess their configuration and better manage your resources.

```sql+postgres
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

```sql+sqlite
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
Explore the capacity of each VSwitch by determining the number of available IP addresses. This can aid in resource allocation and network planning.

```sql+postgres
select
  name,
  vswitch_id,
  available_ip_address_count,
  power(2, 32 - masklen(cidr_block :: cidr)) -1 as raw_size
from
  alicloud_vpc_vswitch;
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### Route Table info associated with VSwitch
Explore the relationship between your VSwitch and its associated route tables in your Alicloud VPC setup. This allows you to gain insights into your network's routing configurations and better manage your network traffic flow.

```sql+postgres
select
  name,
  vswitch_id,
  route_table ->> 'RouteTableId' as route_table_id,
  route_table ->> 'RouteTableType' as route_table_type,
  route_table -> 'RouteEntrys' -> 'RouteEntry' as route_entry
from
  alicloud_vpc_vswitch;
```

```sql+sqlite
select
  name,
  vswitch_id,
  json_extract(route_table, '$.RouteTableId') as route_table_id,
  json_extract(route_table, '$.RouteTableType') as route_table_type,
  json_extract(route_table, '$.RouteEntrys.RouteEntry') as route_entry
from
  alicloud_vpc_vswitch;
```


### VSwitch count by VPC ID
Determine the number of virtual switches within each Virtual Private Cloud to effectively manage network resources and ensure optimal distribution.

```sql+postgres
select
  vpc_id,
  count(vswitch_id) as vswitch_count
from
  alicloud_vpc_vswitch
group by
  vpc_id;
```

```sql+sqlite
select
  vpc_id,
  count(vswitch_id) as vswitch_count
from
  alicloud_vpc_vswitch
group by
  vpc_id;
```