# Table: alicloud_vpc_network_acl

A network access control list (ACL) is an optional layer of security for traffic control in your VPC. You can associate a network ACL with a VSwitch to regulate access for one or more subnets. Similar to the rules of security groups, a user can configure custom rules for network ACLs.

Network ACLs are stateless. After you configure the inbound rules, you need to configure the corresponding outbound rules for certain requests to have a response.

## Examples

### List the attached VPC IDs for each network ACL

```sql
select
  network_acl_id,
  vpc_id
from
  alicloud_vpc_network_acl;
```

### VSwitch associated with each network ACL

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
