---
title: "Steampipe Table: alicloud_vpc_vpn_connection - Query Alibaba Cloud VPN Connections using SQL"
description: "Allows users to query Alibaba Cloud VPN connections, including connection ID, name, status, associated VPN gateway, customer gateway, and tunnel configuration."
folder: "VPC"
---

# Table: alicloud_vpc_vpn_connection - Query Alibaba Cloud VPN Connections using SQL

Alibaba Cloud VPN Gateway supports secure, encrypted communication between on-premises networks and Alibaba Cloud VPCs. VPN connections link a customer gateway with a VPN gateway and use IPsec to establish secure tunnels over the internet.

## Table Usage Guide

The `alicloud_vpc_vpn_connection` table enables network administrators and cloud architects to query detailed information about site-to-site VPN connections in Alibaba Cloud. Use this table to retrieve values such as VPN connection ID, name, status, associated VPN gateway and customer gateway IDs, encryption settings, and tunnel options. This data is vital for managing hybrid cloud connectivity, enforcing encryption standards, and monitoring the health of secure network links.

## Examples

### Basic info
Explore the status of your VPN connections to determine their operational condition and identify the local and remote subnets they are connected to. This can be helpful in troubleshooting network connectivity issues or planning network expansions.

```sql+postgres
select
  name,
  vpn_connection_id,
  status,
  local_subnet,
  remote_subnet,
  vpn_gateway_id
from
  alicloud_vpc_vpn_connection;
```

```sql+sqlite
select
  name,
  vpn_connection_id,
  status,
  local_subnet,
  remote_subnet,
  vpn_gateway_id
from
  alicloud_vpc_vpn_connection;
```

### Get the vpn connections which are not healthy
Identify instances where VPN connections are not in a healthy state. This is useful for troubleshooting network issues and ensuring secure and reliable connectivity.

```sql+postgres
select
  name,
  vpn_connection_id,
  vco_health_check ->> 'Status' as health_check_status,
  status
from
  alicloud_vpc_vpn_connection
where vco_health_check ->> 'Status' = 'failed';
```

```sql+sqlite
select
  name,
  vpn_connection_id,
  json_extract(vco_health_check, '$.Status') as health_check_status,
  status
from
  alicloud_vpc_vpn_connection
where json_extract(vco_health_check, '$.Status') = 'failed';
```

### Get the BGP configuration information of vpn connections
Assess the elements within your VPN connections to understand the status and configuration of Border Gateway Protocol (BGP). This is useful for monitoring the health and performance of your VPN connections.

```sql+postgres
select
  name,
  vpn_connection_id,
  vpn_bgp_config ->> 'EnableBgp' as enable_bgp,
  vpn_bgp_config ->> 'LocalAsn' as local_asn,
  vpn_bgp_config ->> 'LocalBgpIp' as local_bgp_ip,
  vpn_bgp_config ->> 'PeerAsn' as peer_asn,
  vpn_bgp_config ->> 'PeerBgpIp' as peer_bgp_ip,
  vpn_bgp_config ->> 'Status' as status,
  vpn_bgp_config ->> 'TunnelCidr' as tunnel_cidr
from
  alicloud_vpc_vpn_connection;
```

```sql+sqlite
select
  name,
  vpn_connection_id,
  json_extract(vpn_bgp_config, '$.EnableBgp') as enable_bgp,
  json_extract(vpn_bgp_config, '$.LocalAsn') as local_asn,
  json_extract(vpn_bgp_config, '$.LocalBgpIp') as local_bgp_ip,
  json_extract(vpn_bgp_config, '$.PeerAsn') as peer_asn,
  json_extract(vpn_bgp_config, '$.PeerBgpIp') as peer_bgp_ip,
  json_extract(vpn_bgp_config, '$.Status') as status,
  json_extract(vpn_bgp_config, '$.TunnelCidr') as tunnel_cidr
from
  alicloud_vpc_vpn_connection;
```


### Get the vpn connections where NAT traversal feature is enabled
Identify instances where the NAT traversal feature is enabled in VPN connections. This can be useful to ensure secure and efficient data communication in scenarios where private networks are interconnected over the internet.

```sql+postgres
select
  name,
  vpn_connection_id,
  enable_nat_traversal
from
  alicloud_vpc_vpn_connection
where enable_nat_traversal;
```

```sql+sqlite
select
  name,
  vpn_connection_id,
  enable_nat_traversal
from
  alicloud_vpc_vpn_connection
where enable_nat_traversal = 1;
```