---
title: "Steampipe Table: alicloud_vpc - Query Alibaba Cloud Virtual Private Clouds using SQL"
description: "Allows users to query Virtual Private Clouds in Alibaba Cloud, specifically the VPC ID, creation time, status, and other details, providing insights into network configurations and potential anomalies."
---

# Table: alicloud_vpc - Query Alibaba Cloud Virtual Private Clouds using SQL

Alibaba Cloud Virtual Private Cloud (VPC) is a private, isolated network environment based on Alibaba Cloud. It allows users to launch Alibaba Cloud resources in a virtual network that they define. With VPC, users can customize their network configuration, such as IP address range, subnet creation, route table, and network gateway.

## Table Usage Guide

The `alicloud_vpc` table provides insights into Virtual Private Clouds within Alibaba Cloud. As a network administrator, explore VPC-specific details through this table, including VPC ID, creation time, status, and other details. Utilize it to uncover information about your network configurations, such as IP address range, subnet creation, route table, and network gateway.

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
  region
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
  region
from
  alicloud_vpc
where
  not cidr_block <<= '10.0.0.0/8'
  and not cidr_block <<= '192.168.0.0/16'
  and not cidr_block <<= '172.16.0.0/12';
```

### Get the VSwitches details for VPCs

```sql
select
  vpc.vpc_id,
  vswitch.vswitch_id,
  vswitch.cidr_block,
  vswitch.status,
  vswitch.available_ip_address_count,
  vswitch.zone_id
from
  alicloud_vpc as vpc
  join alicloud_vpc_vswitch as vswitch on vpc.vpc_id = vswitch.vpc_id
order by 
  vpc.vpc_id;
```
