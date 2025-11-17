-- Debug query to see why matching is failing
-- Run this to see the extracted values

with actiontrail_check as (
  select
    name as trail_name,
    status,
    sls_project_arn,
    home_region,
    trail_region,
    -- Extract region from SLS project ARN
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
    status as alert_status
  from
    alicloud_sls_alert
  where
    (status = 'ENABLED' or status is null)
    and query_list is not null
)
-- Show all ActionTrail entries with extracted values
select
  'ActionTrail' as source,
  trail_name,
  sls_project_arn,
  sls_region,
  sls_project_name,
  home_region,
  trail_region
from
  actiontrail_check
union all
-- Show all alerts
select
  'Alert' as source,
  alert_name as trail_name,
  project as sls_project_arn,
  region as sls_region,
  project as sls_project_name,
  region as home_region,
  null as trail_region
from
  alert_check
order by source, trail_name;

