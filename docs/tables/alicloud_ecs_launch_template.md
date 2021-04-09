# Table: alicloud_ecs_launch_template

A launch template helps you quickly create an ECS instance. A template contains configurations that you can use to create instances for various scenarios with specific requirements.

A template can include any configurations except passwords. It can include key pairs, RAM roles, instance type, and network configurations.

You can create multiple versions of each template. Each version can contain different configurations. You can then create an instance using any version of the template.

## Examples

### Basic info

```sql
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

```sql
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

### List templates that use encrypted storage disk

```sql
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
