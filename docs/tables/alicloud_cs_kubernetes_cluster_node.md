# Table: alicloud_cs_kubernetes_cluster

A Node is a worker machine in Kubernetes and may be either a virtual or a physical machine, depending on the cluster. Each Node is managed by the Master. A Node can have multiple pods, and the Kubernetes master automatically handles scheduling the pods across the Nodes in the cluster.

## Examples

### Basic info

```sql
select
  node_name,
  cluster_id,
  state,
  creation_time,
  instance_id,
from
  alicloud_cs_kubernetes_cluster_node;
```

### List of worker nodes

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

### List ecs instance info of nodes

```sql
select
  node_name,
  instance_id,
  instance_name,
  instance_role,
  instance_type,
  instance_charge_type,
  instance_type_family
from
  alicloud_cs_kubernetes_cluster_node;
```

### Count of node per instance

```sql
select
  instance_id,
  count(*) as node_count
from
  alicloud_cs_kubernetes_cluster_node
group by
  instance_id;
```
