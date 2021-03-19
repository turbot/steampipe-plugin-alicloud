# Table: alicloud_vpc_network_acl

A network access control list (ACL) is an optional layer of security for traffic control in your VPC. You can associate a network ACL with a VSwitch to regulate access for one or more subnets. Similar to the rules of security groups, a user can configure custom rules for network ACLs.

Network ACLs are stateless. After you configure the inbound rules, you need to configure the corresponding outbound rules for certain requests to have a response.

## Examples

### Basic info

```sql
select
  name,
  network_acl_id,
  status,
  vpc_id,
  description,
  region_id
from
  alicloud_vpc_network_acl;
```

### List the VSwitch associated with network ACLs

```sql
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


### List the inbound rule info of network ACLs

```sql
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


### List the outbound rule info of network ACLs

```sql
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


### Count of network ACLs per region

```sql
select
  region,
  count(*) network_acl_count
from
  alicloud_vpc_network_acl
group by
  region;
```