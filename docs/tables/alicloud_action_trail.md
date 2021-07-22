# Table: alicloud_action_trail

Alibaba Cloud ActionTrail is a service that monitors and records the actions of your Alibaba Cloud account, including the access to and use of cloud products and services through the Alibaba Cloud console, API operations, and SDKs. ActionTrail records these actions as events. You can download these events from the ActionTrail console or configure ActionTrail to deliver these events to Log Service Logstores or Object Storage Service (OSS) buckets. Then, you can perform behavior analysis, security analysis, resource change tracking, and compliance auditing based on the events.

## Examples

### Basic info

```sql
select
  name,
  home_region,
  event_rw,
  status,
  trail_region
from
  alicloud_action_trail;
```

### List enabled trails

```sql
select
  name,
  home_region,
  event_rw,
  status,
  trail_region
from
  alicloud_action_trail
where
  status = 'Enable';
```

### List multi-account trails

```sql
select
  name,
  home_region,
  is_organization_trail,
  status,
  trail_region
from
  alicloud_action_trail
where
  is_organization_trail;
```

### List shadow trails

```sql
select
  name,
  region,
  home_region
from
  alicloud_action_trail
where
  trail_region = 'All'
  and home_region <> region;
```
