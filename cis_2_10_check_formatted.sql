with actiontrail_check as (
  select
    name as trail_name,
    status,
    sls_project_arn,
    sls_write_role_arn,
    home_region,
    trail_region,
    substring(sls_project_arn from 'acs:log:([^:]+):') as sls_region,
    substring(sls_project_arn from 'project/([^/]+)') as sls_project_name
  from
    alicloud_action_trail
  where
    status = 'Enable' and sls_project_arn is not null
), alert_check as (
  select
    project,
    region,
    name as alert_name,
    display_name,
    status as alert_status,
    coalesce(
      query_obj ->> 'Query',
      query_obj ->> 'query',
      query_obj::text
    ) as query_text
  from
    alicloud_sls_alert,
    jsonb_array_elements(query_list) as query_obj
  where
    (status = 'ENABLED' or status is null)
    and query_list is not null
    and (
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%serviceName%ResourceManager%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%serviceName%Ram%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.serviceName%ResourceManager%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.serviceName%Ram%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.serviceName": "ResourceManager"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.serviceName": "Ram"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.serviceName="ResourceManager"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.serviceName="Ram"%'
    ) and (
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%eventName%CreatePolicy%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%eventName%DeletePolicy%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%eventName%CreatePolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%eventName%UpdatePolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%eventName%SetDefaultPolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%eventName%DeletePolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName%CreatePolicy%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName%DeletePolicy%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName%CreatePolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName%UpdatePolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName%SetDefaultPolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName%DeletePolicyVersion%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.eventName": "CreatePolicy"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.eventName": "DeletePolicy"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.eventName": "CreatePolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.eventName": "UpdatePolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.eventName": "SetDefaultPolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%"event.eventName": "DeletePolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName="CreatePolicy"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName="DeletePolicy"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName="CreatePolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName="UpdatePolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName="SetDefaultPolicyVersion"%' or
      coalesce(query_obj ->> 'Query', query_obj ->> 'query', query_obj::text) ilike '%event.eventName="DeletePolicyVersion"%'
    )
), matched_pairs as (
  select distinct
    at.trail_name,
    at.sls_region,
    at.home_region,
    ac.alert_name,
    ac.region as alert_region
  from
    actiontrail_check at inner join alert_check ac on
    trim(lower(coalesce(at.sls_region, ''))) = trim(lower(coalesce(ac.region, '')))
    and at.sls_region is not null
    and ac.region is not null
    and at.sls_region != ''
    and ac.region != ''
)
select
  'acs:actiontrail:' || at.home_region || ':' || at.account_id || ':actiontrail/' || at.name as resource,
  case
    when at.status = 'Enable' and at.sls_project_arn is not null and exists (select 1 from matched_pairs mp where mp.trail_name = at.name)
    then 'ok'
    else 'alarm'
  end as status,
  case
    when at.status = 'Enable' and at.sls_project_arn is not null and exists (select 1 from matched_pairs mp where mp.trail_name = at.name)
    then at.name || ' is configured with ActionTrail enabled, delivering to SLS project in region ' || substring(at.sls_project_arn from 'acs:log:([^:]+):') || ', and has a RAM policy change monitoring alert configured'
    when at.status = 'Enable' and at.sls_project_arn is not null
    then at.name || ' is configured with ActionTrail enabled and delivering to SLS project in region ' || substring(at.sls_project_arn from 'acs:log:([^:]+):') || ', but no RAM policy change monitoring alert found in that region'
    when at.status = 'Enable' and at.sls_project_arn is null
    then at.name || ' is enabled but not configured to deliver logs to SLS'
    else at.name || ' is not enabled'
  end as reason
from
  alicloud_action_trail at
where
  at.status = 'Enable' and at.sls_project_arn is not null;

