# Table: alicloud_ecs_autoscaling_group

Auto Scaling is a service to automatically adjust computing resources based on your volume of user requests. When the demand for computing resources increase, Auto Scaling automatically adds ECS instances to serve additional user requests, or alternatively removes instances in the case of decreased user requests.

## Examples

### Basic auto scaling group info

```sql
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

```sql
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

### List of Autoscaling Group for which deletion protection is not enabled

```sql
select
  name,
  scaling_group_id,
  group_deletion_protection
from
  alicloud_ecs_autoscaling_group
where
  not group_deletion_protection;
```
