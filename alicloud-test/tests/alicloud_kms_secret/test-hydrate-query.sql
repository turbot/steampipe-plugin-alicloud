select 
  name,
  version ->> 'VersionId' as version_id
from 
  alicloud_kms_secret,
  jsonb_array_elements(version_ids) as version
where 
  name = '{{ resourceName }}';