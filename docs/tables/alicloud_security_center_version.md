---
title: "Steampipe Table: alicloud_security_center_version - Query Alibaba Cloud Security Center Versions using SQL"
description: "Allows users to query Security Center Versions in Alibaba Cloud, specifically the version details to understand the security services provided by Alibaba Cloud."
---

# Table: alicloud_security_center_version - Query Alibaba Cloud Security Center Versions using SQL

Alibaba Cloud Security Center is a flagship security product that integrates both Server Guard and Threat Detection Service. It is a unified security management system that recognizes, analyzes, and alerts of security threats in real-time. Alibaba Cloud Security Center provides an array of security features to protect your environment and provides security insights for better visibility.

## Table Usage Guide

The `alicloud_security_center_version` table provides insights into Security Center Versions within Alibaba Cloud Security Center. As a security engineer, explore version-specific details through this table, including the version code, name, and associated metadata. Utilize it to understand the different versions available in Alibaba Cloud Security Center and the services provided by each version.

**Important Notes**
- Valid values are:
  - 1: Basic
  - 2: Enterprise
  - 3: Enterprise
  - 5: Advanced
  - 6: Basic Anti-Virus
- Both 2 and 3 indicate the Enterprise edition and have no differences.

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

### Ensure that Security Center is Advanced or Enterprise edition

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
