---
title: "Steampipe Table: alicloud_vpc_vpn_gateway - Query Alicloud VPC VPN Gateways using SQL"
description: "Allows users to query Alicloud VPC VPN Gateways, providing details such as the gateway's ID, name, description, status, and more."
folder: "VPC"
---

# Table: alicloud_vpc_vpn_gateway - Query Alicloud VPC VPN Gateways using SQL

An Alicloud VPC VPN Gateway is a component of Alibaba Cloud's Virtual Private Cloud (VPC) service. It is used to establish a secure, encrypted communication tunnel between a VPC and an on-premises data center or between VPCs. It supports both IPsec-VPN connections and GRE-VPN connections, and is designed to facilitate secure and convenient cloud network deployment.

## Table Usage Guide

The `alicloud_vpc_vpn_gateway` table provides insights into VPN Gateways within Alibaba Cloud's VPC service. As a network administrator or cloud architect, you can explore gateway-specific details through this table, including its ID, description, status, bandwidth, and associated VPC information. Use it to monitor the status of your VPN gateways, analyze bandwidth usage, and manage your secure network connections.

## Examples

### Basic info
Explore the status and billing methods of your VPN gateways across different regions. This is useful for managing resource allocation and understanding the operational health of your network infrastructure.

```sql+postgres
select
  name,
  vpn_gateway_id,
  status,
  description,
  internet_ip,
  billing_method,
  business_status,
  region
from
  alicloud_vpc_vpn_gateway;
```

```sql+sqlite
select
  name,
  vpn_gateway_id,
  status,
  description,
  internet_ip,
  billing_method,
  business_status,
  region
from
  alicloud_vpc_vpn_gateway;
```


### Get the VPC and VSwitch info of VPN gateway
Determine the areas in which the VPN gateway is connected by identifying the associated VPC and VSwitch. This aids in network management and understanding the connectivity of your virtual private network.

```sql+postgres
select
  name,
  vpn_gateway_id,
  vpc_id vswitch_id
from
  alicloud_vpc_vpn_gateway;
```

```sql+sqlite
select
  name,
  vpn_gateway_id,
  vpc_id as vswitch_id
from
  alicloud_vpc_vpn_gateway;
```


### Get the vpn gateways where SSL VPN is enabled
Determine the areas in your network where SSL VPN is enabled, allowing you to assess security measures and manage potential vulnerabilities. This query is useful for identifying and mitigating potential security risks in your network infrastructure.

```sql+postgres
select
  name,
  vpn_gateway_id,
  ssl_vpn,
  ssl_max_connections
from
  alicloud_vpc_vpn_gateway
where
  ssl_vpn = 'enable';
```

```sql+sqlite
select
  name,
  vpn_gateway_id,
  ssl_vpn,
  ssl_max_connections
from
  alicloud_vpc_vpn_gateway
where
  ssl_vpn = 'enable';
```


### VPN gateway count by VPC ID
Identify the number of VPN gateways associated with each VPC to better manage network resources and optimize security configurations.

```sql+postgres
select
  vpc_id,
  count(vpn_gateway_id) as vpn_gateway_count
from
  alicloud_vpc_vpn_gateway
group by
  vpc_id;
```

```sql+sqlite
select
  vpc_id,
  count(vpn_gateway_id) as vpn_gateway_count
from
  alicloud_vpc_vpn_gateway
group by
  vpc_id;
```


### List of VPN gateways without application tag key
Discover the segments that are missing application tags in VPN gateways. This is useful for identifying untagged resources that may need to be categorized for better resource management.

```sql+postgres
select
  vpn_gateway_id,
  tags
from
  alicloud_vpc_vpn_gateway
where
  tags -> 'application' is null;
```

```sql+sqlite
select
  vpn_gateway_id,
  tags
from
  alicloud_vpc_vpn_gateway
where
  json_extract(tags, '$.application') is null;
```


### List inactive VPN gateways
Identify instances where VPN gateways are not active in your Alicloud VPC. This can be useful to audit and manage your network resources effectively by pinpointing potential network vulnerabilities or unnecessary costs.

```sql+postgres
select
  vpn_gateway_id,
  status,
  create_time,
  jsonb_pretty(tags)
from
  alicloud_vpc_vpn_gateway
where
  status <> 'active';
```

```sql+sqlite
select
  vpn_gateway_id,
  status,
  create_time,
  tags
from
  alicloud_vpc_vpn_gateway
where
  status <> 'active';
```