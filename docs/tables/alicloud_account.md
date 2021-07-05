# Table: alicloud_account

The Alibaba Cloud account is a container for your Alibaba Cloud resources. You create and manage your Alibaba Cloud resources in an Alibaba Cloud account.

## Examples

### Basic info

```sql
select
  alias,
  account_id,
  akas,
  title
from
  alicloud_account;
```