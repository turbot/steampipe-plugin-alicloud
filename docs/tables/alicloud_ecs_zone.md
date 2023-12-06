---
title: "Steampipe Table: alicloud_ecs_zone - Query Alibaba Cloud Elastic Compute Service Zones using SQL"
description: "Allows users to query Alibaba Cloud Elastic Compute Service Zones, providing detailed information about each zone's availability and capabilities."
---

# Table: alicloud_ecs_zone - Query Alibaba Cloud Elastic Compute Service Zones using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides fast memory and flexible compute power. An ECS Zone is a physical area with independent power grids and networks inside an Alibaba Cloud region. Zones are designed to ensure that failures are isolated within Zones and are physically separated within a typical metropolitan region.

## Table Usage Guide

The `alicloud_ecs_zone` table provides insights into the zones within Alibaba Cloud Elastic Compute Service (ECS). As a system administrator, explore zone-specific details through this table, including zone availability, network type, and resource specifications. Utilize it to uncover information about zones, such as their capacities and capabilities, the network types they support, and the resources available within them.

## Examples

### Basic info
Explore which resources, volume categories, and instance types are available in specific zones of your Alicloud Elastic Compute Service. This can help you make informed decisions about where to deploy your resources based on the capabilities of each zone.

```sql+postgres
select
  zone_id,
  local_name,
  available_resource_creation,
  available_volume_categories,
  available_instance_types
from
  alicloud_ecs_zone;
```

```sql+sqlite
select
  zone_id,
  local_name,
  available_resource_creation,
  available_volume_categories,
  available_instance_types
from
  alicloud_ecs_zone;
```

### Get details for a specific region
Determine the resources available in a specific geographical region. This is useful for planning and optimizing resource allocation for your operations in that region.

```sql+postgres
select
  zone_id,
  local_name,
  available_resource_creation,
  available_volume_categories,
  available_instance_types
from
  alicloud_ecs_zone
where
  zone_id = 'ap-south-1b';
```

```sql+sqlite
select
  zone_id,
  local_name,
  available_resource_creation,
  available_volume_categories,
  available_instance_types
from
  alicloud_ecs_zone
where
  zone_id = 'ap-south-1b';
```