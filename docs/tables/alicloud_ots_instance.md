# Table: alicloud_ots_instance

An OSS bucket is the container used to store objects. All objects are contained in buckets. You can configure a variety of bucket properties such as the region, ACL, and storage class. You can create buckets of different storage classes to store data based on your requirements.

## Examples

### Basic info

```sql
select
  instance_name,
  user_id,
  status,
  cluster_type,
  create_time
from
  alicloud_ots_instance;
```
