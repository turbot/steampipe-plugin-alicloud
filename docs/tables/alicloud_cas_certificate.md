# Table: alicloud_cas_certificate

Alibaba Cloud SSL Certificates Service allows customers to directly apply, purchase and manage SSL certificates on Alibaba Cloud.

This service is offered in cooperation with qualified certificate authorities. From this platform, customers can select the expected certificate authority and its certificate products to enjoy full-site HTTPS security solutions.

## Examples

### Basic certificate info

```sql
select
  name,
  id,
  org_name,
  issuer
from
  alicloud_cas_certificate;
```

### List expired certificates

```sql
select
  name,
  id,
  issuer,
  expired
from
  alicloud_cas_certificate
where
  expired;
```

### List of third-party certificates

```sql
select
  name,
  id,
  issuer,
  buy_in_aliyun
from
  alicloud_cas_certificate
where
  not buy_in_aliyun;
```
