# Table: alicloud_elastic_container

Alibaba Cloud Container Service for Kubernetes (ACK) integrates virtualization, storage, networking, and security capabilities. ACK allows you to deploy applications in high-performance and scalable containers and provides full lifecycle management of enterprise-class containerized applications.

## Examples

### Basic info

```sql
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_elastic_container;
```

### List clusters which are running

```sql
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_elastic_container
where
  state = 'running';
```

### List clusters which are managed by kubernetes

```sql
select
  name,
  cluster_id,
  state,
  size,
  cluster_type
from
  alicloud_elastic_container
where
  cluster_type = 'ManagedKubernetes';
```