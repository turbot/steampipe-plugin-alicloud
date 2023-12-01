---
title: "Steampipe Table: alicloud_vpc_vpn_gateway - Query Alicloud VPC VPN Gateways using SQL"
description: "Allows users to query Alicloud VPC VPN Gateways, providing details such as the gateway's ID, name, description, status, and more."
---

# Table: alicloud_vpc_vpn_gateway - Query Alicloud VPC VPN Gateways using SQL

An Alicloud VPC VPN Gateway is a component of Alibaba Cloud's Virtual Private Cloud (VPC) service. It is used to establish a secure, encrypted communication tunnel between a VPC and an on-premises data center or between VPCs. It supports both IPsec-VPN connections and GRE-VPN connections, and is designed to facilitate secure and convenient cloud network deployment.

## Table Usage Guide

The `alicloud_vpc_vpn_gateway` table provides insights into VPN Gateways within Alibaba Cloud's VPC service. As a network administrator or cloud architect, you can explore gateway-specific details through this table, including its ID, description, status, bandwidth, and associated VPC information. Use it to monitor the status of your VPN gateways, analyze bandwidth usage, and manage your secure network connections.

## Examples

### Basic info

```sql
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

```sql
select
  name,
  vpn_gateway_id,
  vpc_id vswitch_id
from
  alicloud_vpc_vpn_gateway;
```


### Get the vpn gateways where SSL VPN is enabled

```sql
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

```sql
select
  vpc_id,
  count(vpn_gateway_id) as vpn_gateway_count
from
  alicloud_vpc_vpn_gateway
group by
  vpc_id;
```


### List of VPN gateways without application tag key

```sql
select
  vpn_gateway_id,
  tags
from
  alicloud_vpc_vpn_gateway
where
  tags -> 'application' is null;
```


### List inactive VPN gateways

```sql
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