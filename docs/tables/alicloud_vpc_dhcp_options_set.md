---
title: "Steampipe Table: alicloud_vpc_dhcp_options_set - Query Alibaba Cloud VPC DHCP Options Sets using SQL"
description: "Allows users to query DHCP options sets in Alibaba Cloud VPC, including domain name servers, domain names, lease times, and associated configurations."
folder: "VPC"
---

# Table: alicloud_vpc_dhcp_options_set - Query Alibaba Cloud VPC DHCP Options Sets using SQL

Alibaba Cloud Virtual Private Cloud (VPC) enables the creation of isolated network environments. DHCP options sets allow you to define custom configurations such as domain name servers and domain names that are automatically assigned to ECS instances within a VPC.

## Table Usage Guide

The `alicloud_vpc_dhcp_options_set` table enables network engineers and cloud administrators to query detailed information about DHCP options sets configured in Alibaba Cloud VPCs. Use this table to retrieve values such as DHCP options set ID, domain name servers, domain name, lease time, and description. This information is helpful for managing network behavior, standardizing instance configuration, and enforcing internal DNS policies across your cloud environment.

## Examples

### Basic info
Explore the status and associated count of your Virtual Private Cloud (VPC) configurations with this query. It helps in understanding the deployment of your VPC resources across different regions and accounts, providing insights for better resource management and planning.

```sql+postgres
select
  name,
  dhcp_options_set_id,
  associate_vpc_count,
  status,
  description,
  domain_name,
  region,
  account_id
from
  alicloud_vpc_dhcp_options_set;
```

```sql+sqlite
select
  name,
  dhcp_options_set_id,
  associate_vpc_count,
  status,
  description,
  domain_name,
  region,
  account_id
from
  alicloud_vpc_dhcp_options_set;
```

### List VPCs that are associated with DHCP options set
Identify the VPCs that are linked with specific DHCP options sets. This is useful in managing network configurations and ensuring proper communication between devices within your virtual private cloud.

```sql+postgres
select
  name,
  dhcp_options_set_id,
  vpc ->> "VpcId"  as vpc_id
from
  alicloud_vpc_dhcp_options_set,
  jsonb_array_elements(associate_vpcs) as vpc;
```

```sql+sqlite
select
  name,
  dhcp_options_set_id,
  json_extract(vpc.value, '$.VpcId') as vpc_id
from
  alicloud_vpc_dhcp_options_set,
  json_each(associate_vpcs) as vpc;
```

### Count the number of VPCs associated with DHCP options set
Determine the quantity of virtual private clouds (VPCs) linked with a Dynamic Host Configuration Protocol (DHCP) options set. This can be useful when assessing the extent of network configurations within your system.

```sql+postgres
select
  name,
  dhcp_options_set_id,
  associate_vpc_count
from
  alicloud_vpc_dhcp_options_set;
```

```sql+sqlite
select
  name,
  dhcp_options_set_id,
  associate_vpc_count
from
  alicloud_vpc_dhcp_options_set;
```

### List DHCP option sets that are in use
Identify the DHCP option sets that are currently in use. This can help assess network settings and ensure they are configured correctly.

```sql+postgres
select
  name,
  nat_gateway_id,
  status
from
  alicloud_vpc_dhcp_options_set
where
  status = 'InUse';
```

```sql+sqlite
select
  name,
  nat_gateway_id,
  status
from
  alicloud_vpc_dhcp_options_set
where
  status = 'InUse';
```