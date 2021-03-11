# Table: alicloud_ecs_key_pair

An SSH key pair is a secure and convenient authentication method provided by Alibaba Cloud for instance logon. An SSH key pair consists of a public key and a private key. You can use SSH key pairs to log on to only Linux instances.

## Examples

### Basic info

```sql
select
  name,
  key_pair_finger_print,
  creation_time,
  resource_group_id
from
  alicloud_ecs_key_pair;
```

### List of available keypairs older than 30 days

```sql
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

### Access key count by Account ID

```sql
select
  account_id,
  count (name) as access_key_count
from
  alicloud_ecs_key_pair
group by
  account_id;
```

