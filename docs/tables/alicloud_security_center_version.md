# Table: alicloud_security_center_version

Security Center is a centralized security management system that dynamically identifies and analyzes security threats, and generates alerts when threats are detected. Security Center provides multiple features to ensure the security of cloud resources and servers in data centers.

Alicloud Security Center Version provides the details of the purchased edition of Security Center.

## Examples

### Basic info

```sql
select
  version,
  is_trial_version,
  is_over_balance,
  region
from
  alicloud_security_center_version;
```

### Ensure that Security Center is Advanced or Enterprise Edition

```sql
select
  version,
  is_trial_version,
  is_over_balance,
  region
from
  alicloud_security_center_version
where
version in ('2','3','5');
```
