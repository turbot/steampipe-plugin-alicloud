# Table: alicloud_vpc_flow_log

Virtual Private Cloud (VPC) provides the flow log feature to capture information about inbound and outbound traffic on an elastic network interface (ENI). You can use the flow log feature to verify access control lists (ACLs) rules, monitor network traffic, and troubleshoot network errors.

## Examples

### Basic info

```sql
select
  name,
  flow_log_id,
  status,
  resource_type,
  resource_id
from
  alicloud_vpc_flow_log;
```


### List flow logs where resource type is VPC.

```sql
select
  name,
  flow_log_id,
  status,
  resource_type,
  resource_id
from
  alicloud_vpc_flow_log
where
  resource_type = 'VPC';
```