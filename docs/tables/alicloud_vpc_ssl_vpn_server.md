---
title: "Steampipe Table: alicloud_vpc_ssl_vpn_server - Query Alicloud VPC SSL VPN Servers using SQL"
description: "Allows users to query Alicloud VPC SSL VPN Servers, providing detailed information about each SSL VPN server in the Alicloud Virtual Private Cloud (VPC)."
folder: "VPC"
---

# Table: alicloud_vpc_ssl_vpn_server - Query Alicloud VPC SSL VPN Servers using SQL

An Alicloud VPC SSL VPN Server is a resource within the Alicloud Virtual Private Cloud (VPC) that enables secure remote access to your private networks. It uses the SSL protocol to establish encrypted connections between remote users and your VPC. This service is critical for managing secure access to your VPC resources.

## Table Usage Guide

The `alicloud_vpc_ssl_vpn_server` table provides insights into SSL VPN servers within Alicloud VPC. As a network administrator, you can explore detailed information about each SSL VPN server, including its configuration, connection details, and associated network resources. Use this table to manage and monitor secure remote access to your VPC resources.

## Examples

### Basic info
Discover the segments that are utilizing the highest capacity in your VPN server. This can help in identifying potential bottlenecks and planning for necessary upgrades or changes to maintain optimal performance.

```sql+postgres
select
  name,
  ssl_vpn_server_id,
  cipher,
  max_connections,
  port,
  proto
from
  alicloud_vpc_ssl_vpn_server;
```

```sql+sqlite
select
  name,
  ssl_vpn_server_id,
  cipher,
  max_connections,
  port,
  proto
from
  alicloud_vpc_ssl_vpn_server;
```

### Get the SSL VPN servers that do not compress the transmitted data
Explore which SSL VPN servers are not compressing transmitted data, enabling you to identify potential areas for bandwidth optimization and improved performance.

```sql+postgres
select
  name,
  ssl_vpn_server_id,
  is_compressed
from
  alicloud_vpc_ssl_vpn_server
where
  not is_compressed;
```

```sql+sqlite
select
  name,
  ssl_vpn_server_id,
  is_compressed
from
  alicloud_vpc_ssl_vpn_server
where
  not is_compressed;
```

### List of all SSL VPN servers that do not use AES-256-CBC encryption
Discover the segments that are utilizing encryption standards other than AES-256-CBC for their SSL VPN servers. This can be particularly useful in identifying potential security risks and ensuring the highest level of data protection.

```sql+postgres
select
  name,
  ssl_vpn_server_id,
  cipher
from
  alicloud_vpc_ssl_vpn_server
where
  cipher <> 'AES-256-CBC';
```

```sql+sqlite
select
  name,
  ssl_vpn_server_id,
  cipher
from
  alicloud_vpc_ssl_vpn_server
where
  cipher <> 'AES-256-CBC';
```

### List of all SSL VPN servers for which Two-factor Authentication is not enabled
Determine the areas in which two-factor authentication is not enabled for SSL VPN servers. This can be valuable in identifying potential security vulnerabilities in your network infrastructure.

```sql+postgres
select
  name,
  ssl_vpn_server_id,
  enable_multi_factor_auth
from
  alicloud_vpc_ssl_vpn_server
where
  not enable_multi_factor_auth;
```

```sql+sqlite
select
  name,
  ssl_vpn_server_id,
  enable_multi_factor_auth
from
  alicloud_vpc_ssl_vpn_server
where
  not enable_multi_factor_auth;
```

### View SSL VPN server IP information
Explore the IP details of your SSL VPN servers, including their internet-facing IP and the client IP pool. This can be beneficial in network troubleshooting and security audits to ensure correct configuration and operation.

```sql+postgres
select
  name,
  ssl_vpn_server_id,
  internet_ip,
  client_ip_pool,
  local_subnet
from
  alicloud_vpc_ssl_vpn_server;
```

```sql+sqlite
select
  name,
  ssl_vpn_server_id,
  internet_ip,
  client_ip_pool,
  local_subnet
from
  alicloud_vpc_ssl_vpn_server;
```

### List of Client Certs for each SSL VPN server
Explore the relationship between your SSL VPN servers and their associated client certificates. This can help in managing and maintaining your network's security infrastructure by ensuring that each VPN server has the appropriate client certificates.

```sql+postgres
select
  s.name,
  s.ssl_vpn_server_id,
  c.name,
  c.ssl_vpn_client_cert_id
from
  alicloud_vpc_ssl_vpn_server as s
left join alicloud_vpc_ssl_vpn_client_cert as c
  on s.ssl_vpn_server_id = c.ssl_vpn_server_id

```

```sql+sqlite
select
  s.name,
  s.ssl_vpn_server_id,
  c.name,
  c.ssl_vpn_client_cert_id
from
  alicloud_vpc_ssl_vpn_server as s
left join alicloud_vpc_ssl_vpn_client_cert as c
  on s.ssl_vpn_server_id = c.ssl_vpn_server_id
```

### Count of Client Certs for each SSL VPN server
This example helps to analyze the distribution of client certificates across various SSL VPN servers. It's useful in understanding the load distribution and managing the capacity of each server efficiently.

```sql+postgres
select
  s.name,
  count(c.ssl_vpn_client_cert_id)
from
  alicloud_vpc_ssl_vpn_server as s
left join alicloud_vpc_ssl_vpn_client_cert as c
  on s.ssl_vpn_server_id = c.ssl_vpn_server_id
group by
  s.name;

```

```sql+sqlite
select
  s.name,
  count(c.ssl_vpn_client_cert_id)
from
  alicloud_vpc_ssl_vpn_server as s
left join alicloud_vpc_ssl_vpn_client_cert as c
  on s.ssl_vpn_server_id = c.ssl_vpn_server_id
group by
  s.name;
```