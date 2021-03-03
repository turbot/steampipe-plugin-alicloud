# Table: alicloud_ecs_network_interface

An elastic network interface (ENI) is a virtual network interface controller (NIC) that can be bound to a VPC-type ECS instance. You can use ENIs to deploy high availability clusters and perform low-cost failover and fine-grained network management.

## Examples

### Basic ENIs info

```sql
select
  network_interface_id,
  type,
  description,
  private_ip_address,
  associated_public_ip_address,
  mac_address
from
  alicloud_ecs_network_interface;
```

### Find all ENIs with private IPs that are in a given subnet (10.66.0.0/16)

```sql
select
  network_interface_id,
  type,
  description,
  private_ip_address,
  associated_public_ip_address,
  mac_address
from
  alicloud_ecs_network_interface
where
  private_ip_address :: cidr <= '10.66.0.0/16';
```

### Count of ENIs by interface type

```sql
select
  type,
  count(type) as count
from
  alicloud_ecs_network_interface
group by
  type
order by
  count desc;
```

### Security groups attached to each ENI

```sql
select
  network_interface_id as eni,
  sg
from
  alicloud_ecs_network_interface
  cross join jsonb_array_elements(security_group_ids) as sg
order by
  eni;
```
