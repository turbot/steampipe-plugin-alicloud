-- First, check if we have any alerts with query_list
select
  project,
  name,
  status,
  query_list is not null as has_query_list,
  jsonb_typeof(query_list) as query_list_type,
  jsonb_array_length(query_list) as query_list_length
from
  alicloud_sls_alert
limit 10;

-- Then, see the actual structure
select
  project,
  name,
  status,
  query_list,
  jsonb_pretty(query_list) as query_list_pretty
from
  alicloud_sls_alert
where
  query_list is not null
limit 5;

-- Expand the array and see what each element looks like
select
  project,
  name,
  status,
  query_obj,
  jsonb_typeof(query_obj) as element_type,
  jsonb_pretty(query_obj) as element_pretty,
  query_obj ->> 'Query' as query_field_upper,
  query_obj ->> 'query' as query_field_lower,
  query_obj::text as element_as_text
from
  alicloud_sls_alert,
  jsonb_array_elements(query_list) as query_obj
where
  query_list is not null
limit 10;

