-- Debug query to see the actual structure of query_list
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

-- Also check what fields are in each query object
select
  project,
  name,
  query_obj,
  jsonb_pretty(query_obj) as query_obj_pretty,
  query_obj ->> 'Query' as query_field,
  query_obj ->> 'query' as query_field_lower,
  query_obj::text as query_obj_text
from
  alicloud_sls_alert,
  jsonb_array_elements(query_list) as query_obj
where
  query_list is not null
limit 10;

