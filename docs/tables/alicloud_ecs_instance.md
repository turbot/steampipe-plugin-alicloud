# Table: alicloud_ecs_instance

An ECS instance is a virtual machine that contains basic computing components such as the vCPU, memory, operating system, network, and disk.

## Examples

### Basic Instance Info
```sql
select
  instance_id,
  name,
  arn,
  status,
  instance_type,
  os_name_en,
  public_ip_address,
  private_ip_address,
  zone
from
  alicloud_ecs_instance;
```


### List stopped instances that you are still being charged for
```sql
select
  instance_id,
  name,
  status,
  stopped_mode,
  instance_type,
  os_name_en,
  public_ip_address,
  private_ip_address,
  zone
from
  alicloud_ecs_instance
where
  stopped_mode = 'KeepCharging';
```


### List linux instances
```sql
select
  instance_id,
  name,
  instance_type,
  os_name_en,
  zone
from
  alicloud_ecs_instance
where
  os_type = 'linux';
```


### Instance count in each zone

```sql
select
  zone as az,
  count(*)
from
  alicloud_ecs_instance
group by
  zone;
```

### Count the number of instances by instance type

```sql
select
  instance_type,
  count(instance_type) as count
from
  alicloud_ecs_instance
group by
  instance_type;
```

### List of instances without application tag key

```sql
select
  instance_id,
  tags
from
  alicloud_ecs_instance
where
  tags ->> 'application' is null;
```

### List of ECS instances provisioned with undesired(for example ecs.t5-lc2m1.nano and ecs.t6-c2m1.large is desired) instance type(s)

```sql
select
  instance_type,
  count(*) as count
from
  alicloud_ecs_instance
where
  instance_type not in ('ecs.t5-lc2m1.nano', 'ecs.t6-c2m1.large')
group by
  instance_type;
```

### List ECS instances having deletion protection safety feature disabled

```sql
select
  instance_id,
  deletion_protection
from
  alicloud_ecs_instance
where
  not deletion_protection;
```
