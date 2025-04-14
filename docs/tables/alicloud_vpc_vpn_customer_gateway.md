---
title: "Steampipe Table: alicloud_vpc_vpn_customer_gateway - Query Alibaba Cloud VPN Customer Gateways using SQL"
description: "Allows users to query Alibaba Cloud VPN customer gateways, including gateway ID, name, IP address, and creation time."
folder: "VPC"
---

# Table: alicloud_vpc_vpn_customer_gateway - Query Alibaba Cloud VPN Customer Gateways using SQL

In Alibaba Cloud, a VPN customer gateway represents the on-premises gateway device used to establish secure VPN connections with a Virtual Private Cloud (VPC). It holds the public IP address and identification details of the customer's network endpoint.

## Table Usage Guide

The `alicloud_vpc_vpn_customer_gateway` table helps network engineers and cloud administrators query detailed information about customer gateways configured for VPN connections in Alibaba Cloud. Use this table to retrieve attributes such as customer gateway ID, name, IP address, description, and creation time. This information is essential for setting up and maintaining site-to-site VPN tunnels, ensuring secure hybrid cloud connectivity.

## Examples

### Basic info
Explore the details of your VPN customer gateway in Alibaba Cloud's VPC service. This query can be used to understand when and why each gateway was created, aiding in resource management and auditing processes.

```sql+postgres
select
  name,
  customer_gateway_id,
  description,
  create_time
from
  alicloud_vpc_vpn_customer_gateway;
```

```sql+sqlite
select
  name,
  customer_gateway_id,
  description,
  create_time
from
  alicloud_vpc_vpn_customer_gateway;
```

### Get the IP address of each customer gateway
Explore which customer gateways are associated with specific IP addresses to better manage network connections and troubleshoot potential issues. This query is beneficial for maintaining secure and efficient connectivity within your virtual private cloud (VPC).

```sql+postgres
select
  name,
  customer_gateway_id,
  ip_address
from
  alicloud_vpc_vpn_customer_gateway;
```

```sql+sqlite
select
  name,
  customer_gateway_id,
  ip_address
from
  alicloud_vpc_vpn_customer_gateway;
```