# Table: alicloud_vpc_network_acl

A Network Access Control List (ACL) is an optional layer of security for traffic control in your VPC. You can associate a network ACL with a VSwitch to regulate access for one or more subnets. Similar to the rules of security groups, a user can configure custom rules for network ACLs.

Network ACLs are stateless. After you configure the inbound rules, you need to configure the corresponding outbound rules for certain requests to have a response.

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