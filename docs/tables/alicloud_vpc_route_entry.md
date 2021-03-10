# Table: alicloud_vpc_route_entry

After you create a VPC, the system automatically adds the following system routes to the route table:
A route entry with a destination CIDR block of 100.64.0.0/10. This route is used for communication among cloud resources within the VPC.
Route entries whose destination CIDR blocks are the same as the CIDR blocks of the VSwitches in the VPC. These routes are used for communication among cloud resources within VSwitches.

## Examples

### Basic info

```sql
select
  name,
  route_table_id,
  description,
  instance_id,
  route_entry_id,
  destination_cidr_block,
  type,
  status
from
  alicloud_vpc_route_entry;
```

### Get route entry details for a particular route table

```sql
select
  name,
  route_table_id,
  description,
  instance_id,
  route_entry_id,
  destination_cidr_block,
  type,
  status
from
  alicloud_vpc_route_entry where route_table_id ='****';
```
