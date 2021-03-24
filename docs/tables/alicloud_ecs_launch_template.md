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

### Get the configuration of current template version

```sql
select
  name,
  version_detail -> 'LaunchTemplateData' ->> 'InstanceName' as instance_name,
  version_detail -> 'LaunchTemplateData' ->> 'InstanceType' as instance_type,
  version_detail -> 'LaunchTemplateData' ->> 'InternetChargeType' as instance_charge_type,
  version_detail -> 'LaunchTemplateData' ->> 'ImageId' as image_id,
  version_detail -> 'LaunchTemplateData' ->> 'VpcId' as vpc_id,
  version_detail -> 'LaunchTemplateData' ->> 'VSwitchId' as v_switch_id,
  version_detail -> 'LaunchTemplateData' ->> 'SecurityGroupId' as security_group_id
from
  alicloud_ecs_launch_template,
  jsonb(latest_version_details) as version_detail;
```
