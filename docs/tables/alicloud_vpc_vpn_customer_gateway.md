# Table: alicloud_vpc_vpn_customer_gateway

A customer gateway device is a physical or software appliance that you own or manage in your on-premises network (on your side of a Site-to-Site VPN connection). You or your network administrator must configure the device to work with the VPN connection.

## Examples

### Basic info
Explore the details of your VPN customer gateway in Alibaba Cloud's VPC service. This query can be used to understand when and why each gateway was created, aiding in resource management and auditing processes.

```sql+postgres
select
  name,
  customer_gateway_id,
  description,
  create_time
from
  alicloud_vpc_vpn_customer_gateway;
```

```sql+sqlite
select
  name,
  customer_gateway_id,
  description,
  create_time
from
  alicloud_vpc_vpn_customer_gateway;
```

### Get the IP address of each customer gateway
Explore which customer gateways are associated with specific IP addresses to better manage network connections and troubleshoot potential issues. This query is beneficial for maintaining secure and efficient connectivity within your virtual private cloud (VPC).

```sql+postgres
select
  name,
  customer_gateway_id,
  ip_address
from
  alicloud_vpc_vpn_customer_gateway;
```

```sql+sqlite
select
  name,
  customer_gateway_id,
  ip_address
from
  alicloud_vpc_vpn_customer_gateway;
```