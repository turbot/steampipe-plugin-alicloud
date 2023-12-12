---
title: "Steampipe Table: alicloud_cms_monitor_host - Query Alibaba Cloud Monitor Hosts using SQL"
description: "Allows users to query Monitor Hosts in Alibaba Cloud, specifically providing insights into the performance of Elastic Compute Service (ECS) instances and custom hosts."
---

# Table: alicloud_cms_monitor_host - Query Alibaba Cloud Monitor Hosts using SQL

Alibaba Cloud Monitor Hosts is a feature within Alibaba Cloud Monitor that provides real-time monitoring of the performance of Elastic Compute Service (ECS) instances and custom hosts. It offers a centralized way to monitor and manage the performance of resources, ensuring smooth and efficient operation. Alibaba Cloud Monitor Hosts helps users stay informed about the health and performance of their resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `alicloud_cms_monitor_host` table provides insights into the performance of Elastic Compute Service (ECS) instances and custom hosts in Alibaba Cloud. As a system administrator or a DevOps engineer, you can explore host-specific details through this table, including the current status, network traffic, and associated metadata. Utilize it to uncover information about hosts, such as those with high CPU usage or network traffic, and to verify their performance.

## Examples

### Basic info
Explore which hosts are part of the Aliyun network and determine their geographical location and operating system. This information can help you understand the distribution and setup of your network, thus aiding in better resource allocation and security planning.

```sql+postgres
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

```sql+sqlite
select
  host_name,
  instance_id,
  is_aliyun_host as ali_uid,
  ip_group,
  operating_system,
  region
from
  alicloud_cms_monitor_host;
```

### Get the status of each host
This query allows you to assess the status of each host within your cloud management system. It provides valuable insights into which hosts are active, which are inactive, and which have automatic installation enabled, aiding in efficient system management and troubleshooting.

```sql+postgres
select
  host_name,
  m ->> 'InstanceId' as instance_id,
  m -> 'AutoInstall' as auto_install,
  m -> 'Status' as status
from
  alicloud_cms_monitor_host,
  jsonb_array_elements(monitoring_agent_status) as m;
```

```sql+sqlite
select
  host_name,
  json_extract(m.value, '$.InstanceId') as instance_id,
  json_extract(m.value, '$.AutoInstall') as auto_install,
  json_extract(m.value, '$.Status') as status
from
  alicloud_cms_monitor_host,
  json_each(monitoring_agent_status) as m;
```

### List hosts provided by Alibaba Cloud
Discover the segments that consist of hosts provided by Alibaba Cloud. This can be beneficial for understanding your cloud resource distribution and for managing your cloud infrastructure more effectively.

```sql+postgres
select
  host_name,
  is_aliyun_host
from
  alicloud_cms_monitor_host
where
  is_aliyun_host;
```

```sql+sqlite
select
  host_name,
  is_aliyun_host
from
  alicloud_cms_monitor_host
where
  is_aliyun_host = 1;
```