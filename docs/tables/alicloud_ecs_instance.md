# Table: alicloud_ecs_instance

An ECS instance is a virtual machine that contains basic computing components such as the vCPU, memory, operating system, network, and disk.

## Examples

### Instance count in each zone

```sql
select
  zone as az,
  type as instance_type,
  count(*)
from
  alicloud_ecs_instance
group by
  zone,
  type;
```

### Count the number of instances by instance type

```sql
select
  type as instance_type,
  count(type) as count
from
  alicloud_ecs_instance
group by 
  type;
```

### List of instances without application tag key

```sql
select
  id as instance_id,
  tags
from
  alicloud_ecs_instance
where
  not tags :: JSONB ? 'application';
```

### List of ECS instances provisioned with undesired(for example ecs.t5-lc2m1.nano and ecs.t6-c2m1.large is desired) instance type(s)

```sql
select
  type as instance_type,
  count(*) as count
from
  alicloud_ecs_instance
where
  type not in ('ecs.t5-lc2m1.nano', 'ecs.t6-c2m1.large')
group by
  instance_type;
```

### List ECS instances having deletion protection safety feature enabled

```sql
select
  id as instance_id,
  deletion_protection
from
  alicloud_ecs_instance
where
  not deletion_protection;
```
