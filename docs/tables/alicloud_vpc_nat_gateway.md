# Table: alicloud_vpc_nat_gateway

NAT gateways are enterprise-class gateways that provide the Source Network Address Translation (SNAT) and Destination Network Address Translation (DNAT) features. Each NAT gateway provides a throughput capacity of up to 10 Gbit/s. NAT gateways also support cross-zone disaster recovery.

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

### List IP address details for NAT gateways

```sql
select
  nat_gateway_id,
  address ->> 'IpAddress' as ip_address,
  address ->> 'AllocationId' as allocation_id
from
  alicloud_vpc_nat_gateway,
  jsonb_array_elements(ip_lists) as address;
```

### List private network info for NAT gateways

```sql
select
  name,
  nat_gateway_id,
  nat_gateway_private_info ->> 'EniInstanceId' as eni_instance_id,
  nat_gateway_private_info ->> 'IzNo' as nat_gateway_zone_id,
  nat_gateway_private_info ->> 'MaxBandwidth' as max_bandwidth,
  nat_gateway_private_info ->> 'PrivateIpAddress' as private_ip_address,
  nat_gateway_private_info ->> 'VswitchId' as vswitch_id
from
  alicloud_vpc_nat_gateway;
```

### List NAT gateways that have traffic monitoring disabled

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

### List NAT gateways that have deletion protection disabled

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
  vpc_id,
  count(*) as nat_gateway_count
from
  alicloud_vpc_nat_gateway
group by
  vpc_id;
```
