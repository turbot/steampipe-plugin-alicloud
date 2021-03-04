# Table: alicloud_vpc_vpn_gateway

VPN Gateway is an Internet-based service that establishes a connection between a VPC and on-premise data center.

## Examples

### Basic info

```sql
select
  name,
  vpn_gateway_id,
  status,
  description,
  internet_ip,
  billing_method,
  business_status,
  region
from
  alicloud_vpc_vpn_gateway;
```


### Get the VPC and VSwitch info of VPN gateway

```sql
select
  name,
  vpn_gateway_id,
  vpc_id vswitch_id
from
  alicloud_vpc_vpn_gateway;
```


### Get the vpn gateways where SSL VPN is enabled

```sql
select
  name,
  vpn_gateway_id,
  ssl_vpn,
  ssl_max_connections
from
  alicloud_vpc_vpn_gateway
where
  ssl_vpn = 'enable';
```


### VPN gateway count by VPC ID

```sql
select
  vpc_id,
  count(vpn_gateway_id) as vpn_gateway_count
from
  alicloud_vpc_vpn_gateway
group by
  vpc_id;
```


### List of VPN gateways without application tag key

```sql
select
  vpn_gateway_id,
  tags
from
  alicloud_vpc_vpn_gateway
where
  tags -> 'application' is null;
```


### List inactive VPN gateways

```sql
select
  vpn_gateway_id,
  status,
  create_time,
  jsonb_pretty(tags)
from
  alicloud_vpc_vpn_gateway
where
  status <> 'active';
```