# Table: alicloud_vpc_vpn_connection

An IPsec-VPN connection provides support to establish an encrypted communication tunnel between a VPN Gateway and a customer gateway.

## Examples

### Basic info

```sql
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
