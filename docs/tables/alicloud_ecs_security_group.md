# Table: alicloud_ecs_security_group

A security group is a logically isolated, mutually accessible group of instances within the same region that all share the same security requirements.

## Examples

### List of security groups where all instances within the security group are isolated from each other

```sql
select
  name,
  id,
  type,
  inner_access_policy
from
  alicloud_ecs_security_group
where
  inner_access_policy = 'drop';
```


### Get the security group rules of each security group

```sql
select
  name,
  id,
  p ->> 'IpProtocol' as ip_protocol_type,
  p ->> 'PortRange' as port_range,
  p ->> 'Direction' as direction,
  p ->> 'SourceCidrIp' as source_cidr_ip,
  p ->> 'SourcePortRange' as source_port_range
from
  alicloud_ecs_security_group,
  jsonb_array_elements(permissions) as p;
```


### List of all enterprise security groups

```sql
select
  name,
  id,
  region_id,
  type
from
  alicloud_ecs_security_group
where
  type = 'enterprise';
```


### Count of security groups by VPC ID

```sql
select
  vpc_id,
  count(*) as count
from
  alicloud_ecs_security_group
group by
  vpc_id;
```
