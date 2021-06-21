# Table: alicloud_cms_monitor_host

Cloud Monitor provides the host monitoring feature to monitor hosts by using the Cloud Monitor agents that are installed on the hosts. The host monitoring feature allows to monitor Elastic Compute Service (ECS) or physical Linux and Windows instances of Alibaba Cloud.

## Examples

### Basic info

```sql
select
  host_name,
  instance_id,
  is_aliyun_host ali_uid,
  ip_group,
  operating_system,
  region
from
  alicloud_cms_monitor_host;
```

### Get the status of each host

```sql
select
  host_name,
  m ->> 'InstanceId' as instance_id,
  m -> 'AutoInstall' as auto_install,
  m -> 'Status' as status
from
  alicloud_cms_monitor_host,
  jsonb_array_elements(monitoring_agent_status) as m;
```

### List hosts provided by Alibaba Cloud

```sql
select
  host_name,
  is_aliyun_host
from
  alicloud_cms_monitor_host
where
  is_aliyun_host;
```
