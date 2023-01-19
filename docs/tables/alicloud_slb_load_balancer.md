# Table: alicloud_vpc_dhcp_options_set

Dynamic Host Configuration Protocol (DHCP) is a network management protocol. DHCP provides a standard for passing configuration information to servers in a TCP/IP network.

## Examples

### Basic info

```sql
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
  address,
  address_type,
  vpc_id,
  network_type
from
  alicloud_slb_load_balancer;
```

### Get VPC details that associated with SLB load balancers

```sql
select
  s.load_balancer_name,
  s.load_balancer_id,
  s.vpc_id,
  v.is_default,
  v.cidr_block
from
  alicloud_vpc_dhcp_options_set as s,
  alicloud_vpc as v
```

### List SLB load balancers that have delete protection enable

```sql
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
  delete_protection
from
  alicloud_slb_load_balancer
where
  delete_protection = 'on';
```

### List SLB load balancers that are created in last 30 days

```sql
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status,
from
  alicloud_vpc_dhcp_options_set
where
  create_time >= now() - interval '30' day;
```