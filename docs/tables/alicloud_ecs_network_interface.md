---
title: "Steampipe Table: alicloud_ecs_network_interface - Query Alibaba Cloud Elastic Compute Service Network Interfaces using SQL"
description: "Allows users to query Network Interfaces in Alibaba Cloud Elastic Compute Service (ECS), retrieving details such as the network interface's ID, status, type, and associated security groups."
---

# Table: alicloud_ecs_network_interface - Query Alibaba Cloud Elastic Compute Service Network Interfaces using SQL

A Network Interface in Alibaba Cloud Elastic Compute Service (ECS) is a virtual network interface card (vNIC) that is attached to an instance. It provides the primary network connection for communication with network services and other instances. Each network interface is associated with a security group that controls the traffic to the instance.

## Table Usage Guide

The `alicloud_ecs_network_interface` table provides insights into Network Interfaces within Alibaba Cloud Elastic Compute Service (ECS). As a network administrator or cloud engineer, explore network interface-specific details through this table, including its status, type, and associated security groups. Utilize it to uncover information about network interfaces, such as those with specific security groups, the status of each interface, and the type of network interface.

## Examples

### Basic ENI info

```sql
select
  network_interface_id,
  type,
  description,
  status,
  instance_id,
  private_ip_address,
  associated_public_ip_address,
  mac_address
from
  alicloud_ecs_network_interface;
```

### Find all ENIs with private IPs that are in a given subnet (10.66.0.0/16)

```sql
select
  network_interface_id,
  type,
  description,
  private_ip_address,
  associated_public_ip_address,
  mac_address
from
  alicloud_ecs_network_interface
where
  private_ip_address <<= '10.66.0.0/16';
```

### Count of ENIs by interface type

```sql
select
  type,
  count(type) as count
from
  alicloud_ecs_network_interface
group by
  type
order by
  count desc;
```

### Security groups attached to each ENI

```sql
select
  network_interface_id as eni,
  sg
from
  alicloud_ecs_network_interface
  cross join jsonb_array_elements(security_group_ids) as sg
order by
  eni;
```

### Find ENIs for a specific instance
```sql
select
  network_interface_id as eni,
  instance_id, 
  status,
  type,
  description,
  private_ip_address,
  associated_public_ip_address,
  mac_address
from
  alicloud_ecs_network_interface
where 
  instance_id = 'i-0xi8u2s0ezl5auigem8t'
```