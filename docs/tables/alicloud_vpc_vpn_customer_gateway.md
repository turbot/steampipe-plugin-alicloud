# Table: alicloud_vpc_vpn_customer_gateway

A customer gateway device is a physical or software appliance that you own or manage in your on-premises network (on your side of a Site-to-Site VPN connection). You or your network administrator must configure the device to work with the VPN connection.

## Examples

### Basic info

```sql
select
  name,
  customer_gateway_id,
  description,
  create_time
from
  alicloud_vpc_vpn_customer_gateway;
```

### Get the IP address of each customer gateway

```sql
select
  name,
  customer_gateway_id,
  ip_address
from
  alicloud_vpc_vpn_customer_gateway;
```
