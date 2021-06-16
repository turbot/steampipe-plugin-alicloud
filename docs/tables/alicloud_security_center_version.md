# Table: alicloud_security_center_version

Security center is a centralized security management system that dynamically identifies and analyzes security threats, and generates alerts when threats are detected. It provides multiple features to ensure the security of cloud resources and servers in data centers.

Alicloud security center version provides the details of the purchased edition of security center.Valid values:

1: Basic
2: Enterprise
3: Enterprise
5: Advanced
6: Basic Anti-Virus

Note: Both 2 and 3 indicate the Enterprise edition and have no differences.

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
