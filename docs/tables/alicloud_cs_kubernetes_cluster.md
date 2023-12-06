---
title: "Steampipe Table: alicloud_cs_kubernetes_cluster - Query Alibaba Cloud Container Service Kubernetes Clusters using SQL"
description: "Allows users to query Kubernetes Clusters in Alibaba Cloud Container Service, providing detailed information about cluster configurations, versions, and statuses."
---

# Table: alicloud_cs_kubernetes_cluster - Query Alibaba Cloud Container Service Kubernetes Clusters using SQL

Alibaba Cloud Container Service for Kubernetes (ACK) is a fully-managed service compatible with Kubernetes to help users with cluster creation and operation. It integrates virtualization, storage, network, and security services, enabling micro-service applications to be deployed, managed, and scaled in a more efficient, secure, and stable manner. ACK supports multiple Kubernetes application deployment models, including monolithic applications, micro-services, and serverless applications.

## Table Usage Guide

The `alicloud_cs_kubernetes_cluster` table provides insights into Kubernetes Clusters within Alibaba Cloud Container Service (ACK). As a DevOps engineer, explore cluster-specific details through this table, including cluster configurations, versions, and statuses. Utilize it to uncover information about clusters, such as those with specific configurations, the versions of Kubernetes they are running, and their current operational status.

## Examples

### Basic info
Explore which Kubernetes clusters are currently active, their size, and type within the Alibaba Cloud service. This can be useful to manage resources and understand the distribution of different types of clusters.

```sql+postgres
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_cs_kubernetes_cluster;
```

```sql+sqlite
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_cs_kubernetes_cluster;
```

### List running clusters
Determine the areas in which active clusters are operating to manage resources and ensure optimal performance. This query is useful for maintaining system efficiency and preventing overutilization of resources.

```sql+postgres
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_cs_kubernetes_cluster
where
  state = 'running';
```

```sql+sqlite
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_cs_kubernetes_cluster
where
  state = 'running';
```

### List managed Kubernetes clusters
Discover the segments that are utilizing Managed Kubernetes clusters. This is useful for assessing the distribution of cluster types and identifying areas for potential optimization or consolidation.

```sql+postgres
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_cs_kubernetes_cluster
where
  cluster_type = 'ManagedKubernetes';
```

```sql+sqlite
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_cs_kubernetes_cluster
where
  cluster_type = 'ManagedKubernetes';
```