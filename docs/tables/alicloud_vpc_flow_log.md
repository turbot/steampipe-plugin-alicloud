---
title: "Steampipe Table: alicloud_vpc_flow_log - Query Alibaba Cloud VPC Flow Logs using SQL"
description: "Allows users to query VPC Flow Logs in Alibaba Cloud, providing insights into network traffic patterns and potential anomalies."
---

# Table: alicloud_vpc_flow_log - Query Alibaba Cloud VPC Flow Logs using SQL

A VPC Flow Log in Alibaba Cloud is a feature that enables you to capture information about the IP traffic going to and from network interfaces in your Virtual Private Cloud (VPC). Flow logs can help you with a number of tasks, such as troubleshooting why specific traffic is not reaching an instance, which in turn helps you diagnose overly restrictive security group rules. You can also use flow logs as a security tool to monitor the traffic that is reaching your instances.

## Table Usage Guide

The `alicloud_vpc_flow_log` table provides insights into VPC Flow Logs within Alibaba Cloud. As a network engineer, explore flow log-specific details through this table, including the traffic patterns, network interface details, and associated metadata. Utilize it to uncover information about flow logs, such as those with specific traffic patterns, the network interfaces involved, and the verification of security group rules.

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