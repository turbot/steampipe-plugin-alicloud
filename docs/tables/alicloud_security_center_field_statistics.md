---
title: "Steampipe Table: alicloud_security_center_field_statistics - Query Alibaba Cloud Security Center Field Statistics using SQL"
description: "Allows users to query Security Center Field Statistics in Alibaba Cloud, providing insights into the statistical data of different fields."
---

# Table: alicloud_security_center_field_statistics - Query Alibaba Cloud Security Center Field Statistics using SQL

Alibaba Cloud Security Center is a flagship security product that integrates both Server Guard and Threat Detection Service. It is a unified security management system that recognizes, analyzes, and alerts of security threats in real-time. This service provides a wide range of capabilities including visualized security monitoring, security alerting, and security protection capabilities to secure users' servers hosted on Alibaba Cloud.

## Table Usage Guide

The `alicloud_security_center_field_statistics` table provides insights into the statistical data of different fields within Alibaba Cloud Security Center. As a security analyst, explore field-specific details through this table, including the number of alerts and vulnerabilities. Utilize it to uncover information about the fields, such as those with high alert rates, helping in identifying potential security threats.

## Examples

### Basic info
Analyze the settings to understand the distribution and status of assets in different regions on Alicloud Security Center. This helps in identifying regions with a high number of unprotected instances, thereby aiding in enhancing the overall security posture.

```sql+postgres
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

```sql+sqlite
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