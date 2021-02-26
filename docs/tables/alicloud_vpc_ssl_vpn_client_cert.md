# Table: alicloud_vpc_ssl_vpn_client_cert

An SSL VPN certificate is a server digital certificate that signed both the server certificate and the user certificate to ensure the data security.

## Examples

### Basic info

```sql
select
  name,
  ssl_vpn_client_cert_id,
  status
from
  alicloud_vpc_ssl_vpn_client_cert;
```

### List of expired certificates

```sql
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

```sql
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

```sql
select
  ssl_vpn_server_id,
  count (ssl_vpn_client_cert_id) as certificate_count
from
  alicloud_vpc_ssl_vpn_client_cert
group by
  ssl_vpn_server_id;
```
