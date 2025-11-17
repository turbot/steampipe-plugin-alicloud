-- Simple test query to verify the alert exists
-- This works with the current data structure where status might be null

select
  project,
  name,
  display_name,
  status,
  query_obj ->> 'Query' as query_text
from
  alicloud_sls_alert,
  jsonb_array_elements(query_list) as query_obj
where
  query_list is not null
  and (
    (query_obj ->> 'Query') ilike '%event.serviceName%Ram%' or
    (query_obj ->> 'Query') ilike '%event.serviceName%ResourceManager%'
  )
  and (
    (query_obj ->> 'Query') ilike '%CreatePolicy%' or
    (query_obj ->> 'Query') ilike '%DeletePolicy%' or
    (query_obj ->> 'Query') ilike '%PolicyVersion%'
  );

