---
title: "Steampipe Table: alicloud_ecs_auto_provisioning_group - Query Alicloud ECS Auto Provisioning Groups using SQL"
description: "Allows users to query Alicloud ECS Auto Provisioning Groups, providing detailed information about the configuration, status, and instance details of each group."
---

# Table: alicloud_ecs_auto_provisioning_group - Query Alicloud ECS Auto Provisioning Groups using SQL

Alicloud Elastic Compute Service (ECS) Auto Provisioning Groups are a feature that enables automatic creation and release of ECS instances based on specified rules. This feature helps to simplify capacity management and reduce costs by automatically adjusting the number of instances based on the real-time business needs.

## Table Usage Guide

The `alicloud_ecs_auto_provisioning_group` table provides insights into the auto provisioning groups within Alicloud Elastic Compute Service (ECS). As a system administrator or DevOps engineer, you can explore group-specific details through this table, including configuration, status, and instance information. Use it to manage your ECS resources efficiently, ensuring optimal capacity and cost-effectiveness.

## Examples

### Basic info
Explore which auto-provisioning groups in the Alibaba Cloud Elastic Compute Service are active or inactive. This can help in managing resources and ensuring optimal performance.

```sql+postgres
select
  name,
  auto_provisioning_group_id,
  state,
  status
from
  alicloud_ecs_auto_provisioning_group;
```

```sql+sqlite
select
  name,
  auto_provisioning_group_id,
  state,
  status
from
  alicloud_ecs_auto_provisioning_group;
```

### Get instance details for a specific group
Explore the specific details of instances within a certain group. This query is beneficial for understanding the settings and statuses of these instances, allowing for more informed management and troubleshooting decisions.

```sql+postgres
select
  apg.name as auto_provisioning_group_name,
  apg.launch_template_id as launch_template_id,
  apg.launch_template_version as launch_template_version,
  i.instance_type,
  i.os_name_en,
  i.private_ip_address,
  i.public_ip_address,
  ins_detail ->> 'InstanceId' as instance_id,
  ins_detail ->> 'InstanceType' as instance_type,
  ins_detail ->> 'Status' as instance_status,
  ins_detail ->> 'NetworkType' as instance_network_type
from
  alicloud_ecs_auto_provisioning_group as apg,
  jsonb_array_elements(apg.instances) as ins_detail,
  alicloud_ecs_instance as i
where
  ins_detail ->> 'InstanceId' = i.instance_id
  and apg.name = 'my_group';
```

```sql+sqlite
select
  apg.name as auto_provisioning_group_name,
  apg.launch_template_id as launch_template_id,
  apg.launch_template_version as launch_template_version,
  i.instance_type,
  i.os_name_en,
  i.private_ip_address,
  i.public_ip_address,
  json_extract(ins_detail.value, '$.InstanceId') as instance_id,
  json_extract(ins_detail.value, '$.InstanceType') as instance_type,
  json_extract(ins_detail.value, '$.Status') as instance_status,
  json_extract(ins_detail.value, '$.NetworkType') as instance_network_type
from
  alicloud_ecs_auto_provisioning_group as apg,
  json_each(apg.instances) as ins_detail,
  alicloud_ecs_instance as i
where
  json_extract(ins_detail.value, '$.InstanceId') = i.instance_id
  and apg.name = 'my_group';
```

### List inactive groups
Discover the segments that are not currently active within your auto provisioning groups. This could be useful for auditing resources, managing costs, or identifying potential issues in your infrastructure.

```sql+postgres
select
  name,
  auto_provisioning_group_id,
  status
from
  alicloud_ecs_auto_provisioning_group
where
  status <> 'active';
```

```sql+sqlite
select
  name,
  auto_provisioning_group_id,
  status
from
  alicloud_ecs_auto_provisioning_group
where
  status != 'active';
```