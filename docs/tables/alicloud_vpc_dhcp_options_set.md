# Table: alicloud_vpc_dhcp_options_set

Dynamic Host Configuration Protocol (DHCP) is a network management protocol. DHCP provides a standard for passing configuration information to servers in a TCP/IP network.

## Examples

### Basic info

```sql
select
  name,
  dhcp_options_set_id,
  associate_vpc_count
  status,
  description,
  domain_name,
  region,
  account_id
from
  alicloud_vpc_dhcp_options_set;
```

### List vpcs that are associate with dhcp options set

```sql
select
  name,
  dhcp_options_set_id,
  vpc ->> "VpcId"  as vpc_id
from
  alicloud_vpc_dhcp_options_set,
  jsonb_array_elements(associate_vpcs) as vpc;
```

### Count number of vpc associate with dhcp options set

```sql
select
  name,
  dhcp_options_set_id,
  associate_vpc_count
from
  alicloud_vpc_dhcp_options_set;
```

### List dhcp option sets which are in use

```sql
select
  name,
  nat_gateway_id,
  status
from
  alicloud_vpc_dhcp_options_set
where
  status = 'InUse';
```