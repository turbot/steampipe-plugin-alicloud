---
title: "Steampipe Table: alicloud_vpc - Query Alibaba Cloud Virtual Private Clouds using SQL"
description: "Allows users to query Virtual Private Clouds in Alibaba Cloud, specifically the VPC ID, creation time, status, and other details, providing insights into network configurations and potential anomalies."
folder: "VPC"
---

# Table: alicloud_vpc - Query Alibaba Cloud Virtual Private Clouds using SQL

Alibaba Cloud Virtual Private Cloud (VPC) is a private, isolated network environment based on Alibaba Cloud. It allows users to launch Alibaba Cloud resources in a virtual network that they define. With VPC, users can customize their network configuration, such as IP address range, subnet creation, route table, and network gateway.

## Table Usage Guide

The `alicloud_vpc` table provides insights into Virtual Private Clouds within Alibaba Cloud. As a network administrator, explore VPC-specific details through this table, including VPC ID, creation time, status, and other details. Utilize it to uncover information about your network configurations, such as IP address range, subnet creation, route table, and network gateway.

## Examples

### Find default VPCs
Determine the areas in which default VPCs are being used across different accounts and regions. This can be useful for managing network accessibility and understanding the distribution of resources in a cloud environment.

```sql+postgres
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

```sql+sqlite
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
Determine the characteristics of your network such as host address, broadcast address, netmask, and network address to better understand and manage your network infrastructure.

```sql+postgres
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

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### List VPCs with public CIDR blocks
Identify instances where your VPCs have CIDR blocks that are publicly accessible. This is useful for assessing potential security risks and ensuring that your network configurations align with best practices.

```sql+postgres
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

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### Get the VSwitches details for VPCs
Explore the status and available IP addresses of virtual switches within specific virtual private clouds. This is useful for managing and optimizing network resources in a cloud environment.

```sql+postgres
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

```sql+sqlite
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