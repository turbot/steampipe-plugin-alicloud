# Table: alicloud_slb_load_balancer

Server Load Balancer (SLB) distributes network traffic across groups of backend servers to improve the service capability and application availability. It Includes Layer 4 Network Load Balancer (NLB), Layer 7 Application Load Balancer (ALB), and Classic Load Balancer (CLB). It Is the Official Cloud-Native Gateway of Alibaba Cloud.

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

### Get VPC details associated with SLB load balancers

```sql
select
  s.load_balancer_name,
  s.load_balancer_id,
  s.vpc_id,
  v.is_default,
  v.cidr_block
from
  alicloud_vpc_dhcp_options_set as s,
  alicloud_vpc as v;
```

### List SLB load balancers that have deletion protection enabled

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

### List SLB load balancers that are created in the last 30 days

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