# Table: alicloud_vpc_vpn_customer_gateway

A customer gateway is a resource that is installed on the customer side and is often linked to the provider side.

## Examples

### Basic info

```sql
select
  name,
  id,
  description,
  create_time
from
  alicloud_vpc_vpn_customer_gateway;
```

### Get the IP address of each customer gateway

```sql
select
  name,
  id,
  ip_address
from
  alicloud_vpc_vpn_customer_gateway;
```
