---
title: "Steampipe Table: alicloud_ecs_autoscaling_group - Query Alibaba Cloud Elastic Compute Service Autoscaling Groups using SQL"
description: "Allows users to query Autoscaling Groups in Alibaba Cloud Elastic Compute Service (ECS), specifically the configuration, capacity, and detailed information about each autoscaling group."
---

# Table: alicloud_ecs_autoscaling_group - Query Alibaba Cloud Elastic Compute Service Autoscaling Groups using SQL

Alibaba Cloud Elastic Compute Service (ECS) Autoscaling Groups are a collection of ECS instances that are created, managed, and released automatically according to the specified scaling rules and strategies. They help in maintaining application availability and improve cost-effectiveness by automatically adjusting the number of ECS instances according to the network traffic.

## Table Usage Guide

The `alicloud_ecs_autoscaling_group` table provides insights into Autoscaling Groups within Alibaba Cloud Elastic Compute Service (ECS). As a system administrator or DevOps engineer, you can explore group-specific details through this table, including configuration, capacity, and detailed information about each autoscaling group. Use it to manage your ECS instances effectively, ensuring optimal resource allocation and cost-effectiveness.

## Examples

### Basic auto scaling group info
This example helps to understand the configuration and capacity details of an auto scaling group in Alibaba Cloud. It can be used to monitor and manage the scaling of resources, ensuring optimal performance and cost-effectiveness.

```sql+postgres
select
  name,
  load_balancer_ids,
  default_cooldown,
  active_capacity,
  desired_capacity,
  min_size,
  max_size
from
  alicloud_ecs_autoscaling_group;
```

```sql+sqlite
select
  name,
  load_balancer_ids,
  default_cooldown,
  active_capacity,
  desired_capacity,
  min_size,
  max_size
from
  alicloud_ecs_autoscaling_group;
```

### Autoscaling group instance details
Explore the specific details of instances within a particular autoscaling group. This can help in managing and optimizing resources by analyzing the health status, instance type, and other relevant details of each instance.

```sql+postgres
select
  asg.name as autoscaling_group_name,
  i.instance_type,
  i.os_name_en,
  i.private_ip_address,
  i.public_ip_address,
  ins_detail ->> 'InstanceId' as instance_id,
  ins_detail ->> 'CreationType' as instance_creation_type,
  ins_detail ->> 'HealthStatus' as health_status,
  ins_detail ->> 'ScalingConfigurationId' as scaling_configuration_id,
  ins_detail ->> 'ScalingGroupId' as scaling_group_id,
  ins_detail -> 'LaunchTemplateId' as launch_template_id,
  ins_detail -> 'LaunchTemplateVersion' as launch_template_version
from
  alicloud_ecs_autoscaling_group as asg,
  jsonb_array_elements(asg.scaling_instances) as ins_detail,
  alicloud_ecs_instance as i
where
  ins_detail ->> 'InstanceId' = i.instance_id
  and asg.name = 'js_as_1';
```

```sql+sqlite
select
  asg.name as autoscaling_group_name,
  i.instance_type,
  i.os_name_en,
  i.private_ip_address,
  i.public_ip_address,
  json_extract(ins_detail.value, '$.InstanceId') as instance_id,
  json_extract(ins_detail.value, '$.CreationType') as instance_creation_type,
  json_extract(ins_detail.value, '$.HealthStatus') as health_status,
  json_extract(ins_detail.value, '$.ScalingConfigurationId') as scaling_configuration_id,
  json_extract(ins_detail.value, '$.ScalingGroupId') as scaling_group_id,
  json_extract(ins_detail.value, '$.LaunchTemplateId') as launch_template_id,
  json_extract(ins_detail.value, '$.LaunchTemplateVersion') as launch_template_version
from
  alicloud_ecs_autoscaling_group as asg,
  json_each(asg.scaling_instances) as ins_detail,
  alicloud_ecs_instance as i
where
  json_extract(ins_detail.value, '$.InstanceId') = i.instance_id
  and asg.name = 'js_as_1';
```

### List of Autoscaling Group for which deletion protection is not enabled
Explore which autoscaling groups in your Alicloud ECS setup lack deletion protection. This is beneficial to identify potential risk areas and take necessary measures to prevent accidental deletions.

```sql+postgres
select
  name,
  scaling_group_id,
  group_deletion_protection
from
  alicloud_ecs_autoscaling_group
where
  not group_deletion_protection;
```

```sql+sqlite
select
  name,
  scaling_group_id,
  group_deletion_protection
from
  alicloud_ecs_autoscaling_group
where
  group_deletion_protection = 0;
```