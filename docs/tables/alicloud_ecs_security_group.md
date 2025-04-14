---
title: "Steampipe Table: alicloud_ecs_security_group - Query Alibaba Cloud ECS Security Groups using SQL"
description: "Allows users to query Alibaba Cloud ECS Security Groups, providing insights into the security settings of Elastic Compute Service (ECS) instances."
folder: "ECS"
---

# Table: alicloud_ecs_security_group - Query Alibaba Cloud ECS Security Groups using SQL

An Alibaba Cloud ECS Security Group is a logical group that segregates ECS instances in different security domains. It acts as a virtual firewall to control inbound and outbound traffic for one or more ECS instances. It is a crucial component for managing the security of Alibaba Cloud ECS instances.

## Table Usage Guide

The `alicloud_ecs_security_group` table provides insights into the security configurations of Alibaba Cloud ECS instances. As a security analyst, you can use this table to explore the security group settings for each ECS instance, including inbound and outbound rules, and associated metadata. Use this table to identify instances with potentially risky security settings, such as open ports or unrestricted IP access.

## Examples

### List of security groups where all instances within the security group are isolated from each other
Determine the areas in which security groups are configured such that all instances within them are isolated from each other. This can be useful for identifying potential security risks and ensuring robust access control.

```sql+postgres
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

```sql+sqlite
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
Explore the security settings of your system by identifying the rules of each security group. This is useful for auditing security measures and ensuring appropriate access controls are in place.

```sql+postgres
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

```sql+sqlite
select
  name,
  security_group_id,
  json_extract(p.value, '$.IpProtocol') as ip_protocol_type,
  json_extract(p.value, '$.PortRange') as port_range,
  json_extract(p.value, '$.Direction') as direction,
  json_extract(p.value, '$.SourceCidrIp') as source_cidr_ip,
  json_extract(p.value, '$.SourcePortRange') as source_port_range
from
  alicloud_ecs_security_group,
  json_each(permissions) as p;
```


### List of all enterprise security groups
Explore which enterprise-level security groups are active in various regions. This can be useful for maintaining oversight of your security measures and ensuring they align with your company's standards.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand the distribution of security groups across different VPCs. This can aid in managing network access and security configurations more effectively.

```sql+postgres
select
  vpc_id,
  count(*) as count
from
  alicloud_ecs_security_group
group by
  vpc_id;
```

```sql+sqlite
select
  vpc_id,
  count(*) as count
from
  alicloud_ecs_security_group
group by
  vpc_id;
```


### Get the security group rules that allow inbound public access to all tcp or udp ports
This query helps to identify security group rules that could potentially expose your system to threats by allowing unrestricted inbound public access via any TCP or UDP port. This information is crucial for assessing potential vulnerabilities and taking steps to enhance your system's security.

```sql+postgres
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

```sql+sqlite
select
  name,
  security_group_id,
  json_extract(p.value, '$.IpProtocol') as ip_protocol_type,
  json_extract(p.value, '$.PortRange') as port_range,
  json_extract(p.value, '$.Direction') as direction,
  json_extract(p.value, '$.SourceCidrIp') as source_cidr_ip,
  json_extract(p.value, '$.SourcePortRange') as source_port_range
from
  alicloud_ecs_security_group,
  json_each(permissions) as p
where 
   json_extract(p.value, '$.IpProtocol') in ('TCP', 'UDP', 'ALL') 
   and json_extract(p.value, '$.Direction') = 'ingress'
   and json_extract(p.value, '$.SourceCidrIp') = '0.0.0.0/0'
   and (
     json_extract(p.value, '$.PortRange') = '-1/-1'
     or json_extract(p.value, '$.PortRange') = '1/65535'
   );
```

### Get the security group rules that allow inbound public access to all tcp or udp ports, along with instances attached to them
This query is useful for identifying potential security vulnerabilities in your system. It reveals the security group rules that permit unrestricted inbound public access to all TCP or UDP ports, and also lists the instances associated with these rules. This can help in strengthening your system's security by pinpointing areas of potential weakness.

```sql+postgres
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

```sql+sqlite
select
  i.name,
  i.instance_id,
  sg.name,
  sg.security_group_id,
  json_extract(p.value, '$.IpProtocol') as ip_protocol_type,
  json_extract(p.value, '$.PortRange') as port_range,
  json_extract(p.value, '$.Direction') as direction,
  json_extract(p.value, '$.SourceCidrIp') as source_cidr_ip,
  json_extract(p.value, '$.SourcePortRange') as source_port_range
from
  alicloud_ecs_security_group as sg,
  json_each(permissions) as p,
  alicloud_ecs_instance as i,
  json_each(i.security_group_ids) as instance_sg
where 
   json_extract(p.value, '$.IpProtocol') in ('TCP', 'UDP', 'ALL') 
   and json_extract(p.value, '$.Direction') = 'ingress'
   and json_extract(p.value, '$.SourceCidrIp') = '0.0.0.0/0'
   and (
     json_extract(p.value, '$.PortRange') = '-1/-1'
     or json_extract(p.value, '$.PortRange') = '1/65535'
   )
   and instance_sg.value = sg.security_group_id;
```