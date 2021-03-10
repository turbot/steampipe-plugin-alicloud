select
  name,
  version_detail -> 'LaunchTemplateData' ->> 'InstanceName' as instance_name,
  version_detail -> 'LaunchTemplateData' ->> 'InstanceType' as instance_type,
  version_detail -> 'LaunchTemplateData' ->> 'InternetChargeType' as instance_charge_type,
  version_detail -> 'LaunchTemplateData' ->> 'ImageId' as image_id,
  version_detail -> 'LaunchTemplateData' ->> 'VpcId' as vpc_id,
  version_detail -> 'LaunchTemplateData' ->> 'VSwitchId' as vswitch_id,
  version_detail -> 'LaunchTemplateData' ->> 'SecurityGroupId' as security_group_id
from 
  alicloud_ecs_launch_template,
  jsonb(latest_version_details) as version_detail
where 
  launch_template_id = '{{ output.resource_id.value }}';
