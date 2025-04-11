---
title: "Steampipe Table: alicloud_vpc_nat_gateway - Query Alibaba Cloud NAT Gateways using SQL"
description: "Allows users to query Alibaba Cloud NAT Gateways, including gateway ID, name, status, type, VPC association, IP addresses, and billing details."
folder: "VPC"
---

# Table: alicloud_vpc_nat_gateway - Query Alibaba Cloud NAT Gateways using SQL

Alibaba Cloud NAT Gateways enable instances in a Virtual Private Cloud (VPC) to access the internet or other Alibaba Cloud services securely without exposing private IP addresses. NAT Gateways support source and destination NAT, offering scalable and managed outbound connectivity.

## Table Usage Guide

The `alicloud_vpc_nat_gateway` table helps network administrators and cloud architects query detailed information about NAT Gateways within Alibaba Cloud. Use this table to retrieve attributes such as NAT Gateway ID, name, status, specification type, associated VPC ID, public IP addresses, bandwidth settings, and charge type. This information is useful for managing external connectivity, securing network traffic, and optimizing resource allocation.

## Examples

### Basic info
Explore the status and billing method of your Alibaba Cloud Virtual Private Cloud (VPC) NAT gateways. This is useful for understanding the operational status and cost management of your NAT gateways across different regions.

```sql+postgres
select
  name,
  nat_gateway_id,
  vpc_id nat_type,
  status,
  description,
  billing_method,
  region,
  account_id
from
  alicloud_vpc_nat_gateway;
```

```sql+sqlite
select
  name,
  nat_gateway_id,
  vpc_id nat_type,
  status,
  description,
  billing_method,
  region,
  account_id
from
  alicloud_vpc_nat_gateway;
```

### List IP address details for NAT gateways
Determine the details of IP addresses associated with Network Address Translation (NAT) gateways to manage and monitor your network's internet connectivity and security.

```sql+postgres
select
  nat_gateway_id,
  address ->> 'IpAddress' as ip_address,
  address ->> 'AllocationId' as allocation_id
from
  alicloud_vpc_nat_gateway,
  jsonb_array_elements(ip_lists) as address;
```

```sql+sqlite
select
  nat_gateway_id,
  json_extract(address.value, '$.IpAddress') as ip_address,
  json_extract(address.value, '$.AllocationId') as allocation_id
from
  alicloud_vpc_nat_gateway,
  json_each(ip_lists) as address;
```

### List private network info for NAT gateways
Discover the segments that provide private network details for NAT gateways. This query can be used to assess the elements within your network infrastructure and optimize resource allocation based on bandwidth usage and zone distribution.

```sql+postgres
select
  name,
  nat_gateway_id,
  nat_gateway_private_info ->> 'EniInstanceId' as eni_instance_id,
  nat_gateway_private_info ->> 'IzNo' as nat_gateway_zone_id,
  nat_gateway_private_info ->> 'MaxBandwidth' as max_bandwidth,
  nat_gateway_private_info ->> 'PrivateIpAddress' as private_ip_address,
  nat_gateway_private_info ->> 'VswitchId' as vswitch_id
from
  alicloud_vpc_nat_gateway;
```

```sql+sqlite
select
  name,
  nat_gateway_id,
  json_extract(nat_gateway_private_info, '$.EniInstanceId') as eni_instance_id,
  json_extract(nat_gateway_private_info, '$.IzNo') as nat_gateway_zone_id,
  json_extract(nat_gateway_private_info, '$.MaxBandwidth') as max_bandwidth,
  json_extract(nat_gateway_private_info, '$.PrivateIpAddress') as private_ip_address,
  json_extract(nat_gateway_private_info, '$.VswitchId') as vswitch_id
from
  alicloud_vpc_nat_gateway;
```

### List NAT gateways that have traffic monitoring disabled
Identify instances where NAT gateways do not have traffic monitoring enabled. This can be useful in ensuring all gateways are properly configured for optimal security and performance.

```sql+postgres
select
  name,
  nat_gateway_id,
  ecs_metric_enabled
from
  alicloud_vpc_nat_gateway
where
  not ecs_metric_enabled;
```

```sql+sqlite
select
  name,
  nat_gateway_id,
  ecs_metric_enabled
from
  alicloud_vpc_nat_gateway
where
  not ecs_metric_enabled;
```

### List NAT gateways that have deletion protection disabled
Determine the areas in which NAT gateways lack deletion protection to enhance your network's security and prevent accidental data loss.

```sql+postgres
select
  name,
  nat_gateway_id,
  deletion_protection
from
  alicloud_vpc_nat_gateway
where
  not deletion_protection;
```

```sql+sqlite
select
  name,
  nat_gateway_id,
  deletion_protection
from
  alicloud_vpc_nat_gateway
where
  deletion_protection = 0;
```

### Count of NAT gateways per VPC ID
Assess the elements within your Alicloud Virtual Private Cloud (VPC) to understand the distribution of Network Address Translation (NAT) gateways. This allows for effective resource allocation and network planning.

```sql+postgres
select
  vpc_id,
  count(*) as nat_gateway_count
from
  alicloud_vpc_nat_gateway
group by
  vpc_id;
```

```sql+sqlite
select
  vpc_id,
  count(*) as nat_gateway_count
from
  alicloud_vpc_nat_gateway
group by
  vpc_id;
```