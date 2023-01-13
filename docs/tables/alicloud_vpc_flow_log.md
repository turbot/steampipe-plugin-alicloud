# Table: alicloud_vpc_flow_log

Virtual Private Cloud (VPC) provides the flow log feature to capture information about inbound and outbound traffic of an elastic network interface (ENI).

## Examples

### Basic info

```sql
select
  name,
  flow_log_id,
  creation_time,
  resource_type,
  description,
  status,
  region,
  account_id
from
  alicloud_vpc_flow_log;
```

### List flow logs that are inactive

```sql
select
  name,
  flow_log_id,
  creation_time,
  resource_type,
  status
from
  alicloud_vpc_flow_log
where
  status = 'Inactive';
```

### List flow logs by resource type

```sql
select
  name,
  flow_log_id,
  creation_time,
  resource_type,
  project_name,
  log_store_name
from
  alicloud_vpc_flow_log
where
  resource_type = 'VPC';
```

### List flow logs created in the last 30 days

```sql
select
  name,
  flow_log_id,
  creation_time,
  resource_type,
  project_name,
  log_store_name
from
  alicloud_vpc_flow_log
where
  creation_time >= now() - interval '30' day;
```