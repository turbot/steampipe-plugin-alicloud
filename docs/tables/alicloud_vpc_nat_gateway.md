# Table: alicloud_vpc_nat_gateway

VPN Gateway is an Internet-based service that establishes a connection between a VPC and on-premise data center.

## Examples

### Basic info

```sql
select
  name,
  nat_gateway_id,
  vpc_id nat_type,
  status,
  description,
  billing_method,
  region,
  account_id
from
  alicloud_vpc_nat_gateway;
```

### Get the Private network info of NAT gateway

```sql
select
  name,
  nat_gateway_id,
  nat_gateway_private_info ->> 'EniInstanceId' as eni_instance_id,
  nat_gateway_private_info ->> 'IzNo' as nat_gateway_zone_id,
  nat_gateway_private_info ->> 'MaxBandwidth' as eni_instance_id,
  nat_gateway_private_info ->> 'PrivateIpAddress' as private_ip_address,
  nat_gateway_private_info ->> 'VswitchId' as vswitch_id
from
  alicloud_vpc_nat_gateway;
```


### Get the NAT gateway where traffic monitoring feature is not enabled

```sql
select
  name,
  nat_gateway_id,
  ecs_metric_enabled
from
  alicloud_vpc_nat_gateway
where
  not ecs_metric_enabled;
```


### Get the NAT gateway where deletion protection is not enabled

```sql
select
  name,
  nat_gateway_id,
  deletion_protection
from
  alicloud_vpc_nat_gateway
where
  not deletion_protection;
```


### Count of NAT gateway per VPC id

```sql
select
  nat_gateway_id,
  count(*)
from
  alicloud_vpc_nat_gateway
group by
  nat_gateway_id;
```