---
title: "Steampipe Table: alicloud_security_center_version - Query Alibaba Cloud Security Center Versions using SQL"
description: "Allows users to query Security Center Versions in Alibaba Cloud, specifically the version details to understand the security services provided by Alibaba Cloud."
folder: "Security Center"
---

# Table: alicloud_security_center_version - Query Alibaba Cloud Security Center Versions using SQL

Alibaba Cloud Security Center is a flagship security product that integrates both Server Guard and Threat Detection Service. It is a unified security management system that recognizes, analyzes, and alerts of security threats in real-time. Alibaba Cloud Security Center provides an array of security features to protect your environment and provides security insights for better visibility.

## Table Usage Guide

The `alicloud_security_center_version` table provides insights into Security Center Versions within Alibaba Cloud Security Center. As a security engineer, explore version-specific details through this table, including the version code, name, and associated metadata. Utilize it to understand the different versions available in Alibaba Cloud Security Center and the services provided by each version.

## Examples

### Basic info
Explore which versions of the Alicloud Security Center are running as trial versions, are over balance, and identify the regions they are operating in. This can assist in managing resources and optimizing security measures.

```sql+postgres
select
  version,
  is_trial_version,
  is_over_balance,
  region
from
  alicloud_security_center_version;
```

```sql+sqlite
select
  version,
  is_trial_version,
  is_over_balance,
  region
from
  alicloud_security_center_version;
```

### Ensure that Security Center is Advanced or Enterprise edition
Discover the segments that are utilizing either the Advanced or Enterprise editions of the Security Center. This helps in understanding the deployment and usage of these premium editions across different regions, and whether they are trial versions or exceeding their balance.

```sql+postgres
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

```sql+sqlite
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