# Table: alicloud_vpc_ssl_vpn_client_cert

An SSL VPN client certificate is a digital certificate that is used by a SSL VPN client to authenticate and securely connect to an SSL VPN.

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