---
title: "Steampipe Table: alicloud_alidns_domain - Query Alibaba Cloud DNS Domains using SQL"
description: "Allows users to query Alibaba Cloud DNS Domains, specifically to retrieve information about the domains such as their DNS records, TTL, registrant details, and more."
folder: "DNS"
---

# Table: alicloud_alidns_domain - Query Alibaba Cloud DNS Domains using SQL

Alibaba Cloud DNS (Alidns) is a scalable and high-performance Domain Name System (DNS) service provided by Alibaba Cloud. Alidns allows users to manage and configure DNS records for their domains, ensuring faster and more reliable domain resolution.

## Table Usage Guide

The `alicloud_alidns_domain` table provides insights into DNS domains managed within Alibaba Cloud DNS (Alidns). As a network administrator or DevOps engineer, you can explore domain-specific details through this table, including DNS records, TTL values, registrant details, and the DNS servers associated with the domains. Utilize it to manage and monitor domain configurations, detect misconfigurations, and ensure optimal domain performance.

## Examples

### Basic info
Retrieve basic information about all Alidns domains, including their names, IDs, and creation times. This helps in inventory management and tracking.

```sql+postgres
select
  domain_name,
  domain_id,
  create_time,
  registrant_email
from
  alicloud_alidns_domain;
```

```sql+sqlite
select
  domain_name,
  domain_id,
  create_time,
  registrant_email
from
  alicloud_alidns_domain;
```

### List expired domain instances
Identify expired domain instances to ensure timely renewals and avoid service disruptions.

```sql+postgres
select
  domain_name,
  domain_id,
  instance_expired,
  instance_end_time
from
  alicloud_alidns_domain
where
  instance_expired;
```

```sql+sqlite
select
  domain_name,
  domain_id,
  instance_expired,
  instance_end_time
from
  alicloud_alidns_domain
where
  instance_expired = 1;
```

### List star-marked domains
Explore domains marked as "star domains" for prioritization or special handling.

```sql+postgres
select
  domain_name,
  domain_id,
  starmark
from
  alicloud_alidns_domain
where
  starmark;
```

```sql+sqlite
select
  domain_name,
  domain_id,
  starmark
from
  alicloud_alidns_domain
where
  starmark = 1;
```

### List DNS servers for a domain
Retrieve the DNS servers associated with a specific domain to verify the configuration.

```sql+postgres
select
  domain_name,
  dns_servers
from
  alicloud_alidns_domain
where
  domain_name = 'example.com';
```

```sql+sqlite
select
  domain_name,
  dns_servers
from
  alicloud_alidns_domain
where
  domain_name = 'example.com';
```

### List domains with high record counts
Identify domains with a high number of DNS records for optimization and management.

```sql+postgres
select
  domain_name,
  record_count
from
  alicloud_alidns_domain
where
  record_count > 100;
```

```sql+sqlite
select
  domain_name,
  record_count
from
  alicloud_alidns_domain
where
  record_count > 100;
```

### Domains by resource group
Group domains by their resource group IDs to analyze resource allocation and usage.

```sql+postgres
select
  resource_group_id,
  count(*) as domain_count
from
  alicloud_alidns_domain
group by
  resource_group_id;
```

```sql+sqlite
select
  resource_group_id,
  count(*) as domain_count
from
  alicloud_alidns_domain
group by
  resource_group_id;
```

### Domains with remarks
Retrieve all domains that have remarks for additional metadata or notes.

```sql+postgres
select
  domain_name,
  remark
from
  alicloud_alidns_domain
where
  remark is not null;
```

```sql+sqlite
select
  domain_name,
  remark
from
  alicloud_alidns_domain
where
  remark is not null;
```