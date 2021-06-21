# Table: alicloud_cs_kubernetes_cluster_node

A node is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the Master. A node can have multiple pods, and the Kubernetes master automatically handles scheduling the pods across the nodes in the cluster.

## Examples

### Basic info

```sql
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

```sql
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

```sql
select
  instance_id,
  count(*) as node_count
from
  alicloud_cs_kubernetes_cluster_node
group by
  instance_id;
```
