# Table: alicloud_ecs_autoscaling_group

Auto Scaling is a service to automatically adjust computing resources based on your volume of user requests. When the demand for computing resources increase, Auto Scaling automatically adds ECS instances to serve additional user requests, or alternatively removes instances in the case of decreased user requests.

## Examples

### Basic auto scaling group info

```sql
select
  name,
  load_balancer_ids,
  default_cooldown,
  max_size,
  min_size
from
  alicloud_ecs_autoscaling_group;
```

### Instances' information attached to the autoscaling group

```sql
select
  name as autoscaling_group_name,
  ins_detail ->> 'InstanceId' as instance_id,
  ins_detail ->> 'CreationType' as instance_creation_type,
  ins_detail ->> 'HealthStatus' as health_status,
  ins_detail -> 'LaunchTemplateId' as launch_template_id,
  ins_detail -> 'LaunchTemplateVersion' as launch_template_version,
  ins_detail ->> 'ProtectedFromScaleIn' as protected_from_scale_in
from
  alicloud_ecs_autoscaling_group
  cross join jsonb_array_elements(scaling_instances) as ins_detail;
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
