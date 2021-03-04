# Table: alicloud_ecs_network_interface

An elastic network interface (ENI) is a virtual network interface controller (NIC) that can be bound to a VPC-type ECS instance. You can use ENIs to deploy high availability clusters and perform low-cost failover and fine-grained network management.

## Examples

### Basic ENI info

```sql
select
  network_interface_id,
  type,
  description,
  status,
  instance_id,
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
  private_ip_address <<= '10.66.0.0/16';
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

### Find ENIs for a specific instance
```sql
select
  network_interface_id as eni,
  instance_id, 
  status,
  type,
  description,
  private_ip_address,
  associated_public_ip_address,
  mac_address
from
  alicloud_ecs_network_interface
where 
  instance_id = 'i-0xi8u2s0ezl5auigem8t'
```