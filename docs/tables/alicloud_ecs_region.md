---
title: "Steampipe Table: alicloud_ecs_region - Query Alicloud Elastic Compute Service Regions using SQL"
description: "Allows users to query Alicloud Elastic Compute Service (ECS) Regions, specifically providing information on all available regions within the Alicloud ECS."
---

# Table: alicloud_ecs_region - Query Alicloud Elastic Compute Service Regions using SQL

Alicloud Elastic Compute Service (ECS) is a high-performance, stable, reliable, and scalable IaaS-level service provided by Alibaba Cloud. ECS eliminates the need to invest in IT hardware up front and allows you to quickly scale computing resources on demand, making ECS more convenient and efficient than physical servers. ECS provides a variety of instance types that suit different business needs and help boost business growth.

## Table Usage Guide

The `alicloud_ecs_region` table provides insights into the regions available within Alibaba Cloud's Elastic Compute Service (ECS). As a system administrator or DevOps engineer, you can explore region-specific details through this table, including region ID, local name, and status. Utilize it to uncover information about regions, such as those that are active, to better manage resource allocation and understand geographical distribution of your ECS instances.

## Examples

### Basic info
Discover the segments that are active within different regional endpoints of the Alibaba Cloud Elastic Compute Service. This can help in assessing the performance and availability of various regions for better resource management.

```sql+postgres
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_ecs_region;
```

```sql+sqlite
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_ecs_region;
```

### Get details for a specific region
Explore which regions are active by checking their status, specifically focusing on the 'us-east-1' region. This provides insights into the operational status and local name of the region, which can be useful for managing resources and planning deployments.

```sql+postgres
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_ecs_region
where
  region = 'us-east-1';
```

```sql+sqlite
select
  region,
  local_name,
  region_endpoint,
  status
from
  alicloud_ecs_region
where
  region = 'us-east-1';
```