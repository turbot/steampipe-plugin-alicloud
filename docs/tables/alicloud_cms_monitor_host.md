# Table: alicloud_cms_monitor_host

Cloud Monitor provides the host monitoring feature to monitor hosts by using the Cloud Monitor agents that are installed on the hosts. The host monitoring feature allows to monitor the Elastic Compute Service (ECS) instances of Alibaba Cloud. Also use the host monitoring feature to monitor virtual machines (VMs) or physical machines from another vendor. The host monitoring feature supports hosts that run the Linux and Windows operating systems.

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

### List cloud monitor agent status

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

### List hosts provided by Alibaba Cloud.

```sql
select
  host_name,
  is_aliyun_host
from
  alicloud_cms_monitor_host
where
  is_aliyun_host;
```
