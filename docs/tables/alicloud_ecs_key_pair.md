---
title: "Steampipe Table: alicloud_ecs_key_pair - Query Alibaba Cloud ECS Key Pairs using SQL"
description: "Allows users to query Alibaba Cloud ECS Key Pairs, specifically the key pair name, creation time, key pair fingerprint, and resource group ID."
folder: "ECS"
---

# Table: alicloud_ecs_key_pair - Query Alibaba Cloud ECS Key Pairs using SQL

Alibaba Cloud Elastic Compute Service (ECS) provides scalable, on-demand cloud servers for secure, flexible, and efficient application environments. ECS supports the key pairs method for logging on to an instance. A key pair consists of a public key and a private key. You can use key pairs to log on to your instances securely without entering a password.

## Table Usage Guide

The `alicloud_ecs_key_pair` table provides insights into key pairs within Alibaba Cloud Elastic Compute Service (ECS). As a system administrator or DevOps engineer, explore key pair-specific details through this table, including the key pair name, creation time, key pair fingerprint, and resource group ID. Utilize it to manage and monitor your key pairs, ensuring secure and efficient access to your instances.

## Examples

### Basic info
Determine the areas in which specific resources are being utilized by assessing the creation time and associated resource group. This can help in managing and optimizing resource allocation in a cloud environment.

```sql+postgres
select
  name,
  key_pair_finger_print,
  creation_time,
  resource_group_id
from
  alicloud_ecs_key_pair;
```

```sql+sqlite
select
  name,
  key_pair_finger_print,
  creation_time,
  resource_group_id
from
  alicloud_ecs_key_pair;
```

### List key pairs older than 30 days
Discover key pairs that have been in existence for over 30 days. This is useful for maintaining security and access control by regularly updating or changing key pairs.

```sql+postgres
select
  name,
  key_pair_finger_print,
  creation_time,
  age(creation_time)
from
  alicloud_ecs_key_pair
where
  creation_time <= (current_date - interval '30' day)
order by
  creation_time;
```

```sql+sqlite
select
  name,
  key_pair_finger_print,
  creation_time,
  julianday('now') - julianday(creation_time) as age
from
  alicloud_ecs_key_pair
where
  julianday(creation_time) <= julianday(date('now','-30 day'))
order by
  creation_time;
```