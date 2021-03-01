# Table: alicloud_ram_user

Access keys are long-term credentials for a RAM user. You can use access keys to sign programmatic requests to the Alibaba Cloud CLI or API (directly or using the Alibaba Cloud SDK).

## Examples

### List of access keys with their corresponding user name and date of creation

```sql
select
  access_key_id,
  user_name,
  create_date
from
  alicloud_ram_access_key;
```

### List of access keys which are inactive

```sql
select
  access_key_id,
  user_name,
  status
from
  alicloud_ram_access_key
where
  status = 'Inactive';
```

### Access key count by user name

```sql
select
  user_name,
  count (access_key_id) as access_key_count
from
  alicloud_ram_access_key
group by
  user_name;
```


### Access keys older than 90 days

```sql
select
  access_key_id,
  user_name,
  status
  create_date,
  age(create_date)
from
  alicloud_ram_access_key
where
  create_date <= (current_date - interval '90' day)
order by
  create_date;
```