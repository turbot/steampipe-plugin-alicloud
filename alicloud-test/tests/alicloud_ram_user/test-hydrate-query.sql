select 
  name, 
  user_id, 
  mfa_enabled, 
  policy ->> 'PolicyName' as attached_policy_name,
  policy ->> 'PolicyType' as attached_policy_type,
  policy ->> 'DefaultVersion' as attached_policy_default_version,
  iam_group ->> 'GroupName' as attached_group_name
from 
  alicloud_ram_user,
  jsonb_array_elements(attached_policy) as policy,
  jsonb_array_elements(groups) as iam_group
where 
  name = '{{ resourceName }}';