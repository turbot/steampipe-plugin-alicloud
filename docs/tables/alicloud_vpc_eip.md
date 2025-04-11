---
title: "Steampipe Table: alicloud_vpc_eip - Query Alibaba Cloud Elastic IP Addresses (EIPs) using SQL"
description: "Allows users to query Alibaba Cloud Elastic IP addresses (EIPs), including IP address, status, instance association, bandwidth, and billing details."
folder: "VPC"
---

# Table: alicloud_vpc_eip - Query Alibaba Cloud Elastic IP Addresses (EIPs) using SQL

Alibaba Cloud Elastic IP Addresses (EIPs) provide public IP addressing that can be dynamically associated with ECS instances, NAT gateways, or other cloud resources. EIPs support flexible networking and high availability for internet-facing applications.

## Table Usage Guide

The `alicloud_vpc_eip` table allows network engineers and cloud administrators to query detailed information about Elastic IP addresses in Alibaba Cloud. Use this table to retrieve data such as EIP address, allocation ID, status, associated instance or resource, bandwidth settings, internet charge type, and creation time. This information is essential for tracking public IP usage, optimizing bandwidth allocation, and managing cost and connectivity for cloud-based services.

## Examples

### Basic info
Explore the status and regional distribution of Elastic IP addresses in your Alibaba Cloud Virtual Private Cloud (VPC) to better manage and optimize your network resources.

```sql+postgres
select
  name,
  allocation_id,
  arn,
  description,
  ip_address,
  status,
  region
from
  alicloud_vpc_eip;
```

```sql+sqlite
select
  name,
  allocation_id,
  arn,
  description,
  ip_address,
  status,
  region
from
  alicloud_vpc_eip;
```


### Get the info of instance bound to eip
Discover the segments that are linked to a specific Elastic IP in order to understand its allocation, the type of instance it's attached to, and its regional location. This is useful for managing resources and bandwidth within your network infrastructure.

```sql+postgres
select
  name,
  allocation_id,
  instance_type,
  instance_id,
  instance_region_id,
  bandwidth
from
  alicloud_vpc_eip;
```

```sql+sqlite
select
  name,
  allocation_id,
  instance_type,
  instance_id,
  instance_region_id,
  bandwidth
from
  alicloud_vpc_eip;
```


### List all the available eips
Determine the areas in which Elastic IP addresses are available for use. This is useful for identifying potential resources in your Alicloud Virtual Private Cloud that are not currently being utilized.

```sql+postgres
select
  name,
  allocation_id,
  instance_type,
  status
from
  alicloud_vpc_eip
where
  status = 'Available';
```

```sql+sqlite
select
  name,
  allocation_id,
  instance_type,
  status
from
  alicloud_vpc_eip
where
  status = 'Available';
```


### Get the eips where hd monitoring is off
Identify instances where HD monitoring is turned off for Elastic IP addresses within your Alicloud Virtual Private Cloud. This can be crucial for optimizing your network performance and security measures.

```sql+postgres
select
  name,
  allocation_id,
  hd_monitor_status
from
  alicloud_vpc_eip
where
  hd_monitor_status = 'OFF';
```

```sql+sqlite
select
  name,
  allocation_id,
  hd_monitor_status
from
  alicloud_vpc_eip
where
  hd_monitor_status = 'OFF';
```