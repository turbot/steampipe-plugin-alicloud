# Table: alicloud_ecs_auto_provisioning_group

Auto Provisioning is a service to quickly deploy an instance cluster that consists of preemptible and pay-as-you-go instances. It supports one-click deployment of instance clusters with specified billing methods, zones, and instance families.

You can use auto provisioning groups to provide stable computing power, alleviate the instability caused by the reclaiming of preemptible instances, and eliminate the need to manually create instances.

## Examples

### Basic auto provisioning group info

```sql
select
  name,
  auto_provisioning_group_id,
  state,
  status
from
  alicloud_ecs_auto_provisioning_group;
```

### Auto Provisioning group instance details

```sql
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
  and apg.name = 'js_as_1';
```

### List of Auto Provisioning group for which are not active

```sql
select
  name,
  auto_provisioning_group_id,
  status
from
  alicloud_ecs_auto_provisioning_group
where
  status <> 'active';
```
