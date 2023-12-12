# Table: alicloud_slb_load_balancer

Server Load Balancer (SLB) distributes network traffic across groups of backend servers to improve the service capability and application availability. It Includes Layer 4 Network Load Balancer (NLB), Layer 7 Application Load Balancer (ALB), and Classic Load Balancer (CLB). It is the Official Cloud-Native Gateway of Alibaba Cloud.

## Examples

### Basic info
Explore the status and details of your load balancers to understand their configuration and network type. This can help in managing network traffic and ensuring efficient distribution of workloads across resources.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which your SLB load balancers are associated with specific VPC details. This can help you gain insights into your load balancing configurations and how they interact with your virtual private cloud settings.

```sql+postgres
select
  s.load_balancer_name,
  s.load_balancer_id,
  s.vpc_id,
  v.is_default,
  v.cidr_block
from
  alicloud_slb_load_balancer as s,
  alicloud_vpc as v;
```

```sql+sqlite
select
  s.load_balancer_name,
  s.load_balancer_id,
  s.vpc_id,
  v.is_default,
  v.cidr_block
from
  alicloud_slb_load_balancer as s,
  alicloud_vpc as v;
```

### List SLB load balancers that have deletion protection enabled
Identify the instances where load balancers have deletion protection enabled. This is useful in ensuring the prevention of accidental deletion of critical load balancers.

```sql+postgres
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

```sql+sqlite
The query provided is already compatible with SQLite. It does not use any PostgreSQL-specific functions or data types that need to be converted. Therefore, the SQLite query is the same as the PostgreSQL query:

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
```

### List SLB load balancers created in the last 30 days
Explore which load balancers have been created in the past month. This can be useful in monitoring recent activity and ensuring proper load distribution across your network.

```sql+postgres
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status
from
  alicloud_slb_load_balancer
where
  create_time >= now() - interval '30' day;
```

```sql+sqlite
select
  load_balancer_name,
  load_balancer_id,
  load_balancer_status
from
  alicloud_slb_load_balancer
where
  create_time >= datetime('now', '-30 day');
```