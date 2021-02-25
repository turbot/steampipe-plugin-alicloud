# Table: alicloud_vpc_vpn_ssl_server

A Secure Socket Layer Virtual Private Network (SSL VPN) lets remote users access Web applications, client-server apps, and internal network utilities and directories without the need for specialized client software. SSL VPN provide safe communication for all types of device traffic across public networks and private networks.

## Examples

### Basic info

```sql
select
  name,
  ssl_vpn_server_id,
  cipher,
  port,
  proto
from
  alicloud_vpc_vpn_ssl_server;
```

### List of all SSL VPN servers for which no encryption algorithm is used

```sql
select
  name,
  ssl_vpn_server_id,
  cipher,
from
  alicloud_vpc_vpn_ssl_server
where
  cipher = 'none';
```

### List of all SSL VPN servers for which Two-factor Authentication is not enabled

```sql
select
  name,
  ssl_vpn_server_id,
  enable_multi_factor_auth
from
  alicloud_vpc_vpn_ssl_server
where
  not enable_multi_factor_auth;
```
