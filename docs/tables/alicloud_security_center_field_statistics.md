# Table: alicloud_security_center_field_statistics

Security Center is a centralized security management system that dynamically identifies and analyzes security threats, and generates alerts when threats are detected. Security Center provides multiple features to ensure the security of cloud resources and servers in data centers.

Alicloud Security Center Field Statistics provides the statistics of assets.

## Examples

### Basic info

```sql
select
  general_asset_count,
  group_count,
  important_asset_count,
  instance_count,
  unprotected_instance_count,
  region
from
  alicloud_security_center_field_statistics;
```
