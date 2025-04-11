---
title: "Steampipe Table: alicloud_vpc_ssl_vpn_client_cert - Query Alibaba Cloud SSL VPN Client Certificates using SQL"
description: "Allows users to query Alibaba Cloud SSL VPN client certificates, including certificate ID, name, status, creation time, and associated SSL VPN server."
folder: "VPC"
---

# Table: alicloud_vpc_ssl_vpn_client_cert - Query Alibaba Cloud SSL VPN Client Certificates using SQL

Alibaba Cloud SSL VPN enables secure communication between remote clients and a Virtual Private Cloud (VPC) by using encrypted tunnels. SSL VPN client certificates are used to authenticate clients and establish secure VPN connections to the cloud environment.

## Table Usage Guide

The `alicloud_vpc_ssl_vpn_client_cert` table helps network administrators and security teams query detailed information about SSL VPN client certificates in Alibaba Cloud. Use this table to retrieve attributes such as client certificate ID, name, status, creation time, and associated SSL VPN server ID. This information is crucial for managing secure client access, monitoring certificate status, and ensuring compliance with access control policies.

## Examples

### Basic info
Determine the status of your VPN client certificates in your network. This is useful for ensuring security compliance and identifying any inactive or expired certificates.

```sql+postgres
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert;
```

```sql+sqlite
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert;
```

### List of expired certificates
Identify instances where SSL VPN client certificates have expired in your AliCloud VPC environment. This query is useful for maintaining security standards and for timely renewal of certificates.

```sql+postgres
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert
where
  status = 'expired';
```

```sql+sqlite
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert
where
  status = 'expired';
```

### List of certificates that will expire in one week
Identify instances where SSL VPN client certificates are nearing their expiration date. This is useful for ensuring timely renewal and maintaining uninterrupted VPN service.

```sql+postgres
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert
where
  status = 'expiring-soon';
```

```sql+sqlite
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert
where
  status = 'expiring-soon';
```

### Certificate count by SSL server
Determine the number of certificates associated with each SSL server to monitor your network's security. This can help in managing certificate distribution and identifying servers with unusually high or low certificate counts.

```sql+postgres
select
  ssl_vpn_server_id,
  count (ssl_vpn_client_cert_id) as certificate_count
from
  alicloud_vpc_ssl_vpn_client_cert
group by
  ssl_vpn_server_id;
```

```sql+sqlite
select
  ssl_vpn_server_id,
  count (ssl_vpn_client_cert_id) as certificate_count
from
  alicloud_vpc_ssl_vpn_client_cert
group by
  ssl_vpn_server_id;
```