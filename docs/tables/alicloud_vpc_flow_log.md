---
title: "Steampipe Table: alicloud_vpc_flow_log - Query Alibaba Cloud VPC Flow Logs using SQL"
description: "Allows users to query VPC Flow Logs in Alibaba Cloud, providing insights into network traffic patterns and potential anomalies."
folder: "VPC"
---

# Table: alicloud_vpc_flow_log - Query Alibaba Cloud VPC Flow Logs using SQL

A VPC Flow Log in Alibaba Cloud is a feature that enables you to capture information about the IP traffic going to and from network interfaces in your Virtual Private Cloud (VPC). Flow logs can help you with a number of tasks, such as troubleshooting why specific traffic is not reaching an instance, which in turn helps you diagnose overly restrictive security group rules. You can also use flow logs as a security tool to monitor the traffic that is reaching your instances.

## Table Usage Guide

The `alicloud_vpc_flow_log` table provides insights into VPC Flow Logs within Alibaba Cloud. As a network engineer, explore flow log-specific details through this table, including the traffic patterns, network interface details, and associated metadata. Utilize it to uncover information about flow logs, such as those with specific traffic patterns, the network interfaces involved, and the verification of security group rules.

## Examples

### Basic info
Explore which resources within your Alibaba Cloud Virtual Private Cloud (VPC) have flow logs enabled, to gain insights into your network traffic for security monitoring and diagnostic purposes. This query is useful in identifying potential security risks and ensuring compliance with your organization's logging policies.

```sql+postgres
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

```sql+sqlite
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
Identify instances where flow logs are inactive in your Alicloud VPC. This can be useful for optimizing resource usage and ensuring that all active flow logs are necessary and functional.

```sql+postgres
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

```sql+sqlite
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
Explore which flow logs have been created for a specific type of resource. This can be particularly useful for managing and troubleshooting network traffic within your Virtual Private Cloud (VPC), allowing you to identify potential issues or inefficiencies.

```sql+postgres
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

```sql+sqlite
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
Analyze the flow logs to understand the recent activities in your Virtual Private Cloud (VPC). This is particularly useful for identifying any unusual network traffic patterns or potential security issues that have arisen in the past month.

```sql+postgres
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

```sql+sqlite
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
  creation_time >= datetime('now', '-30 day');
```