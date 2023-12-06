---
title: "Steampipe Table: alicloud_ecs_launch_template - Query Alibaba Cloud ECS Launch Templates using SQL"
description: "Allows users to query Alibaba Cloud ECS Launch Templates, providing detailed information about instances that can be launched from a template."
---

# Table: alicloud_ecs_launch_template - Query Alibaba Cloud ECS Launch Templates using SQL

Alibaba Cloud Elastic Compute Service (ECS) Launch Templates provide a way to save instance launch configurations, allowing for the rapid deployment of instances with pre-defined settings. These templates can include instance type, image, security group, and other instance-related parameters. Utilizing a launch template can streamline instance deployment and ensure consistency across instances.

## Table Usage Guide

The `alicloud_ecs_launch_template` table provides insights into ECS Launch Templates within Alibaba Cloud Elastic Compute Service. As a DevOps engineer, explore template-specific details through this table, including instance configurations, security settings, and associated metadata. Utilize it to uncover information about templates, such as those with specific instance types or security groups, and to streamline and standardize your instance deployment process.

## Examples

### Basic info
Explore the different launch templates available on Alibaba Cloud's Elastic Compute Service (ECS). This allows you to assess the default and latest version numbers of each template, aiding in version control and regional deployment strategies.

```sql+postgres
select
  name,
  launch_template_id,
  default_version_number,
  latest_version_number,
  region
from
  alicloud_ecs_launch_template;
```

```sql+sqlite
select
  name,
  launch_template_id,
  default_version_number,
  latest_version_number,
  region
from
  alicloud_ecs_launch_template;
```

### Get the current template version's configuration
Explore the latest configurations of your virtual server templates, including details like instance type, charge type, image ID, and associated security groups. This can help in understanding your current setup and planning for any necessary changes or updates.

```sql+postgres
select
  name,
  latest_version_details -> 'LaunchTemplateData' ->> 'InstanceName' as instance_name,
  latest_version_details -> 'LaunchTemplateData' ->> 'InstanceType' as instance_type,
  latest_version_details -> 'LaunchTemplateData' ->> 'InternetChargeType' as instance_charge_type,
  latest_version_details -> 'LaunchTemplateData' ->> 'ImageId' as image_id,
  latest_version_details -> 'LaunchTemplateData' ->> 'VpcId' as vpc_id,
  latest_version_details -> 'LaunchTemplateData' ->> 'VSwitchId' as v_switch_id,
  latest_version_details -> 'LaunchTemplateData' ->> 'SecurityGroupId' as security_group_id
from
  alicloud_ecs_launch_template;
```

```sql+sqlite
select
  name,
  json_extract(latest_version_details, '$.LaunchTemplateData.InstanceName') as instance_name,
  json_extract(latest_version_details, '$.LaunchTemplateData.InstanceType') as instance_type,
  json_extract(latest_version_details, '$.LaunchTemplateData.InternetChargeType') as instance_charge_type,
  json_extract(latest_version_details, '$.LaunchTemplateData.ImageId') as image_id,
  json_extract(latest_version_details, '$.LaunchTemplateData.VpcId') as vpc_id,
  json_extract(latest_version_details, '$.LaunchTemplateData.VSwitchId') as v_switch_id,
  json_extract(latest_version_details, '$.LaunchTemplateData.SecurityGroupId') as security_group_id
from
  alicloud_ecs_launch_template;
```

### List templates that use encrypted storage disk
Determine the areas in which encrypted storage disks are utilized within templates, focusing on those that are configured to be deleted along with the instance. This is useful for understanding your security measures and potential data loss risks.

```sql+postgres
select
  name,
  disk_config ->> 'Encrypted' as disk_encryption,
  disk_config ->> 'DeleteWithInstance' as delete_with_instance
from
  alicloud_ecs_launch_template,
  jsonb_array_elements(latest_version_details -> 'LaunchTemplateData' -> 'DataDisks' -> 'DataDisk') as disk_config
where
  (disk_config ->> 'Encrypted')::boolean
  and (disk_config ->> 'DeleteWithInstance')::boolean;
```

```sql+sqlite
select
  name,
  json_extract(disk_config.value, '$.Encrypted') as disk_encryption,
  json_extract(disk_config.value, '$.DeleteWithInstance') as delete_with_instance
from
  alicloud_ecs_launch_template,
  json_each(latest_version_details, '$.LaunchTemplateData.DataDisks.DataDisk') as disk_config
where
  json_extract(disk_config.value, '$.Encrypted') = 'true'
  and json_extract(disk_config.value, '$.DeleteWithInstance') = 'true';
```