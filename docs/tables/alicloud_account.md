# Table: alicloud_account

The Alibaba Cloud account is a container for your Alibaba Cloud resources. You create and manage your Alibaba Cloud resources in an Alibaba Cloud account.It has two accounts: International-Site Account and China-Site Account. In most cases the account type makes no difference when creating Alibaba Cloud resources.

## Examples

### Basic AliCloud account info

```sql
select
  alias,
  account_id,
  akas,
  title
from
  alicloud_account;
```