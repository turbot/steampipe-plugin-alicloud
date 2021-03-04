# Table: alicloud_vpc_ssl_vpn_server

A Secure Socket Layer Virtual Private Network (SSL VPN) lets remote users access Web applications, client-server apps, and internal network utilities and directories without the need for specialized client software. SSL VPN provide safe communication for all types of device traffic across public networks and private networks.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
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

```sql
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

```sql
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

```sql
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
