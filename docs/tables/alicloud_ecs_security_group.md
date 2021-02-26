# Table: alicloud_ecs_security_group

A security group is a logically isolated, mutually accessible group of instances within the same region that all share the same security requirements.

## Examples

### List of security groups where all instances within the security group are isolated from each other

```sql
select
  name,
  security_group_id,
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
  security_group_id,
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
  security_group_id,
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


### Get the security group rules that allow inbound public access to all tcp or udp ports

```sql
select
  name,
  security_group_id,
  p ->> 'IpProtocol' as ip_protocol_type,
  p ->> 'PortRange' as port_range,
  p ->> 'Direction' as direction,
  p ->> 'SourceCidrIp' as source_cidr_ip,
  p ->> 'SourcePortRange' as source_port_range
from
  alicloud_ecs_security_group,
  jsonb_array_elements(permissions) as p
where 
   p ->> 'IpProtocol' in ('TCP', 'UDP', 'ALL') 
   and p ->> 'Direction' = 'ingress'
   and p ->> 'SourceCidrIp' = '0.0.0.0/0'
   and (
     p ->> 'PortRange' = '-1/-1'
     or p ->> 'PortRange' = '1/65535'
   );
```

### Get the security group rules that allow inbound public access to all tcp or udp ports, along with instances attached to them

```sql
select
  i.name,
  i.instance_id,
  sg.name,
  sg.security_group_id,
  p ->> 'IpProtocol' as ip_protocol_type,
  p ->> 'PortRange' as port_range,
  p ->> 'Direction' as direction,
  p ->> 'SourceCidrIp' as source_cidr_ip,
  p ->> 'SourcePortRange' as source_port_range
from
  alicloud_ecs_security_group as sg,
  jsonb_array_elements(permissions) as p,
  alicloud_ecs_instance as i,
  jsonb_array_elements_text(i.security_group_ids) as instance_sg
where 
   p ->> 'IpProtocol' in ('TCP', 'UDP', 'ALL') 
   and p ->> 'Direction' = 'ingress'
   and p ->> 'SourceCidrIp' = '0.0.0.0/0'
   and (
     p ->> 'PortRange' = '-1/-1'
     or p ->> 'PortRange' = '1/65535'
   )
   and instance_sg = sg.security_group_id;
```