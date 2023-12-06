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
Explore the status and details of network interfaces within your Alicloud Elastic Compute Service. This can be useful to manage network configurations and troubleshoot connectivity issues.

```sql+postgres
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

```sql+sqlite
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
Explore which Elastic Network Interfaces (ENIs) with private IPs fall within a specific subnet. This is particularly useful in understanding network configurations and managing resources within a particular network range.

```sql+postgres
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

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### Count of ENIs by interface type
Analyze the variety of network interfaces in use to understand their distribution within your Alicloud Elastic Compute Service (ECS). This is useful for gauging network capacity and planning for potential upgrades or changes.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which security groups are associated with each Elastic Network Interface (ENI). This allows you to understand your network's security layout and identify potential vulnerabilities or areas for improvement.

```sql+postgres
select
  network_interface_id as eni,
  sg
from
  alicloud_ecs_network_interface
  cross join jsonb_array_elements(security_group_ids) as sg
order by
  eni;
```

```sql+sqlite
select
  network_interface_id as eni,
  sg.value as sg
from
  alicloud_ecs_network_interface,
  json_each(security_group_ids) as sg
order by
  eni;
```

### Find ENIs for a specific instance
Gain insights into the network interfaces associated with a specific instance, including their status, type, and IP addresses. This can be useful for troubleshooting connectivity issues or understanding the network configuration of a particular instance.
```sql+postgres
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

```sql+sqlite
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