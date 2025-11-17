-- CIS 2.10 Compliance Check Query
-- This query verifies that:
-- 1. ActionTrail is enabled and delivering to SLS
-- 2. An alert exists that monitors RAM policy changes
--
-- REGION CONSIDERATIONS:
-- - ActionTrail can deliver logs to SLS projects in specific regions
-- - SLS projects and alerts are region-specific
-- - The alert must exist in the same region as the SLS project receiving ActionTrail logs
-- - This query checks across all configured regions
--
-- IMPORTANT: When trail_region = "All":
--   - ActionTrail captures events from ALL regions
--   - But all logs are delivered to ONE SLS project in a specific region
--   - You only need ONE alert in that SLS project's region to monitor all logs
--   Example: trail_region="All", sls_region="cn-hangzhou"
--            → Need alert in cn-hangzhou (monitors logs from all regions)

with actiontrail_check as (
  select
    name as trail_name,
    status,
    sls_project_arn,
    sls_write_role_arn,
    home_region,
    trail_region,
    -- Extract region from SLS project ARN (format: acs:log:REGION:account-id:project/project-name)
    substring(sls_project_arn from 'acs:log:([^:]+):') as sls_region,
    -- Extract project name from SLS project ARN
    substring(sls_project_arn from 'project/([^/]+)') as sls_project_name
  from
    alicloud_action_trail
  where
    status = 'Enable'
    and sls_project_arn is not null
),
alert_check as (
  select
    project,
    region,
    name as alert_name,
    display_name,
    status as alert_status,
    query_obj ->> 'Query' as query_text
  from
    alicloud_sls_alert,
    jsonb_array_elements(query_list) as query_obj
  where
    -- Check if status is ENABLED (or null, which might mean enabled by default)
    (status = 'ENABLED' or status is null)
    and query_list is not null
    and (
      -- Check for ResourceManager or Ram service
      (query_obj ->> 'Query') ilike '%serviceName%ResourceManager%' or
      (query_obj ->> 'Query') ilike '%serviceName%Ram%' or
      (query_obj ->> 'Query') ilike '%event.serviceName%ResourceManager%' or
      (query_obj ->> 'Query') ilike '%event.serviceName%Ram%' or
      (query_obj ->> 'Query') ilike '%"event.serviceName": "ResourceManager"%' or
      (query_obj ->> 'Query') ilike '%"event.serviceName": "Ram"%'
    )
    and (
      -- Check for policy-related event names
      (query_obj ->> 'Query') ilike '%eventName%CreatePolicy%' or
      (query_obj ->> 'Query') ilike '%eventName%DeletePolicy%' or
      (query_obj ->> 'Query') ilike '%eventName%CreatePolicyVersion%' or
      (query_obj ->> 'Query') ilike '%eventName%UpdatePolicyVersion%' or
      (query_obj ->> 'Query') ilike '%eventName%SetDefaultPolicyVersion%' or
      (query_obj ->> 'Query') ilike '%eventName%DeletePolicyVersion%' or
      (query_obj ->> 'Query') ilike '%event.eventName%CreatePolicy%' or
      (query_obj ->> 'Query') ilike '%event.eventName%DeletePolicy%' or
      (query_obj ->> 'Query') ilike '%event.eventName%CreatePolicyVersion%' or
      (query_obj ->> 'Query') ilike '%event.eventName%UpdatePolicyVersion%' or
      (query_obj ->> 'Query') ilike '%event.eventName%SetDefaultPolicyVersion%' or
      (query_obj ->> 'Query') ilike '%event.eventName%DeletePolicyVersion%' or
      (query_obj ->> 'Query') ilike '%"event.eventName": "CreatePolicy"%' or
      (query_obj ->> 'Query') ilike '%"event.eventName": "DeletePolicy"%' or
      (query_obj ->> 'Query') ilike '%"event.eventName": "CreatePolicyVersion"%' or
      (query_obj ->> 'Query') ilike '%"event.eventName": "UpdatePolicyVersion"%' or
      (query_obj ->> 'Query') ilike '%"event.eventName": "SetDefaultPolicyVersion"%' or
      (query_obj ->> 'Query') ilike '%"event.eventName": "DeletePolicyVersion"%'
    )
),
-- Match ActionTrail with alerts in the same region
-- Note: Region matching is sufficient - alerts can query logs from any project in the same region
matched_pairs as (
  select distinct
    at.trail_name,
    at.sls_region,
    at.sls_project_name,
    ac.alert_name,
    ac.region as alert_region,
    ac.project as alert_project
  from
    actiontrail_check at
  inner join
    alert_check ac
  on
    -- Match by region (SLS project region = alert region)
    -- Trim and normalize to handle any whitespace issues
    trim(lower(coalesce(at.sls_region, ''))) = trim(lower(coalesce(ac.region, '')))
    -- Ensure both regions are not null/empty
    and at.sls_region is not null
    and ac.region is not null
    and at.sls_region != ''
    and ac.region != ''
)
select
  'CIS 2.10 Compliance Check' as control,
  case
    when (select count(*) from actiontrail_check) > 0
      and (select count(*) from alert_check) > 0
      and (select count(*) from matched_pairs) > 0
    then 'PASS'
    else 'FAIL'
  end as compliance_status,
  (select count(*) from actiontrail_check) as actiontrail_enabled_count,
  (select count(*) from alert_check) as ram_policy_alert_count,
  (select count(*) from matched_pairs) as matched_region_pairs,
  (select string_agg(distinct trail_name, ', ') from actiontrail_check) as enabled_trails,
  (select string_agg(distinct trail_region, ', ') from actiontrail_check) as trail_regions,
  (select string_agg(distinct alert_name, ', ') from alert_check) as matching_alerts,
  (select string_agg(distinct sls_region, ', ') from actiontrail_check where sls_region is not null) as actiontrail_sls_regions,
  (select string_agg(distinct region, ', ') from alert_check) as alert_regions;

