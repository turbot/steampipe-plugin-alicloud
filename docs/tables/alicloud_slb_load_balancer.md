---
title: "Steampipe Table: alicloud_slb_load_balancer - Query Alibaba Cloud SLB Load Balancers using SQL"
description: "Allows users to query Alibaba Cloud Server Load Balancers (SLB), including load balancer ID, name, status, type, region, VPC information, and IP configuration."
folder: "SLB"
---

# Table: alicloud_slb_load_balancer - Query Alibaba Cloud SLB Load Balancers using SQL

Alibaba Cloud Server Load Balancer (SLB) automatically distributes incoming traffic across multiple backend servers to improve application availability and fault tolerance. SLB supports different load balancing protocols and deployment types, enabling flexible and reliable traffic distribution in cloud environments.

## Table Usage Guide

The `alicloud_slb_load_balancer` table enables cloud engineers, DevOps teams, and network administrators to query detailed information about Server Load Balancers in Alibaba Cloud. Use this table to retrieve attributes such as load balancer ID, name, status, type (public or internal), region, associated VPC, IP address settings, and listener configurations. This data supports efficient traffic management, health monitoring, and secure configuration of network-facing services.


## Examples

### Basic info
Explore the status and details of your load balancers to understand their configuration and network type. This can help in managing network traffic and ensuring efficient distribution of workloads across resources.

```sql+postgres
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
  address,
  address_type,
  vpc_id,
  network_type
from
  alicloud_slb_load_balancer;
```

```sql+sqlite
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
  address,
  address_type,
  vpc_id,
  network_type
from
  alicloud_slb_load_balancer;
```

### Get VPC details associated with SLB load balancers
Determine the areas in which your SLB load balancers are associated with specific VPC details. This can help you gain insights into your load balancing configurations and how they interact with your virtual private cloud settings.

```sql+postgres
select
  s.load_balancer_name,
  s.load_balancer_id,
  s.vpc_id,
  v.is_default,
  v.cidr_block
from
  alicloud_slb_load_balancer as s,
  alicloud_vpc as v;
```

```sql+sqlite
select
  s.load_balancer_name,
  s.load_balancer_id,
  s.vpc_id,
  v.is_default,
  v.cidr_block
from
  alicloud_slb_load_balancer as s,
  alicloud_vpc as v;
```

### List SLB load balancers that have deletion protection enabled
Identify the instances where load balancers have deletion protection enabled. This is useful in ensuring the prevention of accidental deletion of critical load balancers.

```sql+postgres
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
  delete_protection
from
  alicloud_slb_load_balancer
where
  delete_protection = 'on';
```

```sql+sqlite
The query provided is already compatible with SQLite. It does not use any PostgreSQL-specific functions or data types that need to be converted. Therefore, the SQLite query is the same as the PostgreSQL query:

```sql
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
  delete_protection
from
  alicloud_slb_load_balancer
where
  delete_protection = 'on';
```
```

### List SLB load balancers created in the last 30 days
Explore which load balancers have been created in the past month. This can be useful in monitoring recent activity and ensuring proper load distribution across your network.

```sql+postgres
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status
from
  alicloud_slb_load_balancer
where
  create_time >= now() - interval '30' day;
```

```sql+sqlite
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status
from
  alicloud_slb_load_balancer
where
  create_time >= datetime('now', '-30 day');
```