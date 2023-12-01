---
title: "Steampipe Table: alicloud_cas_certificate - Query Alibaba Cloud CAS Certificates using SQL"
description: "Allows users to query Alibaba Cloud CAS Certificates, specifically to retrieve information about the certificates such as their status, domain, issuer, validity period, and more."
---

# Table: alicloud_cas_certificate - Query Alibaba Cloud CAS Certificates using SQL

Alibaba Cloud's Certificate Authority Service (CAS) is a platform that provides digital certificate services. The service is designed to help users secure online data transmission, establish SSL encrypted sessions and enhance the security of their websites, applications and services. It provides a range of certificate types, including DV, OV, and EV SSL certificates.

## Table Usage Guide

The `alicloud_cas_certificate` table provides insights into the digital certificates within Alibaba Cloud's Certificate Authority Service (CAS). As a security engineer, you can explore certificate-specific details through this table, including the certificate's status, domain, issuer, and validity period. Utilize it to uncover information about certificates, such as those that are expired or nearing expiration, the domains they are associated with, and the entities that issued them.

## Examples

### Basic info

```sql
select
  name,
  id,
  org_name,
  issuer
from
  alicloud_cas_certificate;
```

### List expired certificates

```sql
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

### List third-party certificates

```sql
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
