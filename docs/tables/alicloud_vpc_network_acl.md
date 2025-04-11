---
title: "Steampipe Table: alicloud_vpc_network_acl - Query Alibaba Cloud Network ACLs using SQL"
description: "Allows users to query Alibaba Cloud Network Access Control Lists (ACLs), including ACL ID, name, VPC association, ingress and egress rules, and creation time."
folder: "VPC"
---

# Table: alicloud_vpc_network_acl - Query Alibaba Cloud Network ACLs using SQL

Alibaba Cloud Network Access Control Lists (ACLs) provide stateless, subnet-level traffic filtering to enhance security within Virtual Private Clouds (VPCs). Network ACLs support both inbound and outbound rules and can be associated with one or more subnets.

## Table Usage Guide

The `alicloud_vpc_network_acl` table allows cloud security engineers and network administrators to query detailed information about Network ACLs in Alibaba Cloud. Use this table to retrieve data such as the ACL ID, name, associated VPC ID, creation time, and configured ingress and egress rules. This information is essential for auditing traffic control configurations, enforcing security boundaries, and ensuring compliance with organizational network policies.

## Examples

### Basic info
Explore the status and regional location of your network access controls within the Alibaba Cloud Virtual Private Cloud (VPC) service. This can help you manage and secure your network resources effectively.

```sql+postgres
select
  name,
  network_acl_id,
  status,
  vpc_id,
  description,
  region
from
  alicloud_vpc_network_acl;
```

```sql+sqlite
select
  name,
  network_acl_id,
  status,
  vpc_id,
  description,
  region
from
  alicloud_vpc_network_acl;
```

### List the VSwitches associated with each network ACL
Determine the associations between network access control lists (ACLs) and virtual switches in your network. This query is useful for assessing network security and understanding connection statuses within your virtual private cloud (VPC).

```sql+postgres
select
  network_acl_id,
  vpc_id,
  association ->> 'ResourceId' as vswitch_id,
  association ->> 'Status' as association_status
from
  alicloud_vpc_network_acl,
  jsonb_array_elements(resources) as association
where
  association ->> 'ResourceType' = 'VSwitch';
```

```sql+sqlite
select
  network_acl_id,
  vpc_id,
  json_extract(association.value, '$.ResourceId') as vswitch_id,
  json_extract(association.value, '$.Status') as association_status
from
  alicloud_vpc_network_acl,
  json_each(resources) as association
where
  json_extract(association.value, '$.ResourceType') = 'VSwitch';
```

### Get inbound rule info for each network ACL
Explore the details of inbound rules for each network access control list (ACL), allowing you to understand the security settings and policies applied to your virtual private cloud (VPC). This can be useful in identifying potential security vulnerabilities or for auditing purposes.

```sql+postgres
select
  name,
  network_acl_id,
  vpc_id,
  i ->> 'NetworkAclEntryId' as network_acl_entry_id,
  i ->> 'NetworkAclEntryName' as network_acl_entry_name,
  i ->> 'Description' as description,
  i ->> 'EntryType' as entry_type,
  i ->> 'Policy' as policy,
  i ->> 'Port' as port,
  i ->> 'Protocol' as protocol,
  i ->> 'SourceCidrIp' as source_cidr_ip
from
  alicloud_vpc_network_acl,
  jsonb_array_elements(ingress_acl_entries) as i;
```

```sql+sqlite
select
  name,
  network_acl_id,
  vpc_id,
  json_extract(i.value, '$.NetworkAclEntryId') as network_acl_entry_id,
  json_extract(i.value, '$.NetworkAclEntryName') as network_acl_entry_name,
  json_extract(i.value, '$.Description') as description,
  json_extract(i.value, '$.EntryType') as entry_type,
  json_extract(i.value, '$.Policy') as policy,
  json_extract(i.value, '$.Port') as port,
  json_extract(i.value, '$.Protocol') as protocol,
  json_extract(i.value, '$.SourceCidrIp') as source_cidr_ip
from
  alicloud_vpc_network_acl,
  json_each(ingress_acl_entries) as i;
```

### Get outbound rule info for each network ACL
Explore the outbound rules for each network access control list (ACL) to understand the restrictions placed on outgoing network traffic. This is crucial for maintaining network security by controlling which traffic is allowed to leave your network.

```sql+postgres
select
  name,
  network_acl_id,
  vpc_id,
  i ->> 'NetworkAclEntryId' as network_acl_entry_id,
  i ->> 'NetworkAclEntryName' as network_acl_entry_name,
  i ->> 'Description' as description,
  i ->> 'EntryType' as entry_type,
  i ->> 'Policy' as policy,
  i ->> 'Port' as port,
  i ->> 'Protocol' as protocol,
  i ->> 'DestinationCidrIp' as destination_cidr_ip
from
  alicloud_vpc_network_acl,
  jsonb_array_elements(egress_acl_entries) as i;
```

```sql+sqlite
select
  name,
  network_acl_id,
  vpc_id,
  json_extract(i.value, '$.NetworkAclEntryId') as network_acl_entry_id,
  json_extract(i.value, '$.NetworkAclEntryName') as network_acl_entry_name,
  json_extract(i.value, '$.Description') as description,
  json_extract(i.value, '$.EntryType') as entry_type,
  json_extract(i.value, '$.Policy') as policy,
  json_extract(i.value, '$.Port') as port,
  json_extract(i.value, '$.Protocol') as protocol,
  json_extract(i.value, '$.DestinationCidrIp') as destination_cidr_ip
from
  alicloud_vpc_network_acl,
  json_each(egress_acl_entries) as i;
```