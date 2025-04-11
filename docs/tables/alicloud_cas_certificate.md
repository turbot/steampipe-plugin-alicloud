---
title: "Steampipe Table: alicloud_cas_certificate - Query Alibaba Cloud CAS Certificates using SQL"
description: "Allows users to query Alibaba Cloud CAS Certificates, specifically to retrieve information about the certificates such as their status, domain, issuer, validity period, and more."
folder: "CAS"
---

# Table: alicloud_cas_certificate - Query Alibaba Cloud CAS Certificates using SQL

Alibaba Cloud's Certificate Authority Service (CAS) is a platform that provides digital certificate services. The service is designed to help users secure online data transmission, establish SSL encrypted sessions and enhance the security of their websites, applications and services. It provides a range of certificate types, including DV, OV, and EV SSL certificates.

## Table Usage Guide

The `alicloud_cas_certificate` table provides insights into the digital certificates within Alibaba Cloud's Certificate Authority Service (CAS). As a security engineer, you can explore certificate-specific details through this table, including the certificate's status, domain, issuer, and validity period. Utilize it to uncover information about certificates, such as those that are expired or nearing expiration, the domains they are associated with, and the entities that issued them.

## Examples

### Basic info
Explore which certificates are issued by Alicloud CAS by determining their names, IDs, and associated organization names. This can help in managing and tracking the certificates used in your infrastructure.

```sql+postgres
select
  name,
  id,
  org_name,
  issuer
from
  alicloud_cas_certificate;
```

```sql+sqlite
select
  name,
  id,
  org_name,
  issuer
from
  alicloud_cas_certificate;
```

### List expired certificates
Explore which certificates have expired to ensure your systems remain secure and up-to-date. This is crucial as expired certificates can lead to security vulnerabilities and system downtime.

```sql+postgres
select
  name,
  id,
  issuer,
  expired
from
  alicloud_cas_certificate
where
  expired;
```

```sql+sqlite
select
  name,
  id,
  issuer,
  expired
from
  alicloud_cas_certificate
where
  expired = 1;
```

### List third-party certificates
Discover the segments that contain third-party certificates in the Alicloud CAS service. This can be useful to identify certificates not purchased through Alicloud, potentially highlighting areas of cost savings or security risks.

```sql+postgres
select
  name,
  id,
  issuer,
  buy_in_aliyun
from
  alicloud_cas_certificate
where
  not buy_in_aliyun;
```

```sql+sqlite
select
  name,
  id,
  issuer,
  buy_in_aliyun
from
  alicloud_cas_certificate
where
  buy_in_aliyun = 0;
```