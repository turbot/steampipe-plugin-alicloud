---
title: "Steampipe Table: alicloud_cs_kubernetes_cluster_node - Query Alibaba Cloud Container Service Kubernetes Cluster Nodes using SQL"
description: "Allows users to query Kubernetes Cluster Nodes in Alibaba Cloud Container Service, providing detailed information about each node's configuration, status, and associated metadata."
folder: "CS"
---

# Table: alicloud_cs_kubernetes_cluster_node - Query Alibaba Cloud Container Service Kubernetes Cluster Nodes using SQL

Alibaba Cloud Container Service for Kubernetes (ACK) is a fully-managed service compatible with Kubernetes to help users focus on their applications rather than managing container infrastructure. It provides out-of-the-box Kubernetes native capabilities, simplifies the deployment of Kubernetes clusters, and offers high-performance and flexible management of containerized applications throughout their lifecycle.

## Table Usage Guide

The `alicloud_cs_kubernetes_cluster_node` table provides insights into Kubernetes Cluster Nodes within Alibaba Cloud Container Service (ACK). As a DevOps engineer or system administrator, you can explore node-specific details through this table, including configuration, status, and associated metadata. Utilize it to monitor the health and performance of your nodes, track resource usage, and manage your containerized applications more effectively.

## Examples

### Basic info
Explore the status and details of your Kubernetes nodes within the Alibaba Cloud Container Service. This allows you to monitor your infrastructure and identify any potential issues or changes.

```sql+postgres
select
  node_name,
  cluster_id,
  state,
  creation_time,
  instance_id
from
  alicloud_cs_kubernetes_cluster_node;
```

```sql+sqlite
select
  node_name,
  cluster_id,
  state,
  creation_time,
  instance_id
from
  alicloud_cs_kubernetes_cluster_node;
```

### List worker nodes
Identify instances where the role of a node in a Kubernetes cluster is 'Worker'. This allows you to quickly determine which nodes are performing worker tasks within your cluster.

```sql+postgres
select
  node_name,
  instance_id,
  instance_name,
  instance_role
from
  alicloud_cs_kubernetes_cluster_node
where
  instance_role = 'Worker';
```

```sql+sqlite
select
  node_name,
  instance_id,
  instance_name,
  instance_role
from
  alicloud_cs_kubernetes_cluster_node
where
  instance_role = 'Worker';
```

### Count the number of nodes per instance
Determine the distribution of nodes across different instances in your Kubernetes cluster. This can help identify any uneven distribution and manage workload balancing effectively.

```sql+postgres
select
  instance_id,
  count(*) as node_count
from
  alicloud_cs_kubernetes_cluster_node
group by
  instance_id;
```

```sql+sqlite
select
  instance_id,
  count(*) as node_count
from
  alicloud_cs_kubernetes_cluster_node
group by
  instance_id;
```