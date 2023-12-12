---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/alicloud.svg"
brand_color: "#FF6600"
display_name: Alibaba Cloud
name: alicloud
description: Steampipe plugin for querying Alibaba Cloud servers, databases, networks, and other resources.
og_description: Query Alibaba Cloud with SQL! Open source CLI. No DB required. 
og_image: "/images/plugins/turbot/alicloud-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Alibaba Cloud + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[Alibaba Cloud](https://alibabacloud.com) provides on-demand cloud computing platforms and APIs to authenticated customers on a metered pay-as-you-go basis. 

For example:

```sql
select
  name,
  display_name,
  mfa_enabled
from
  alicloud_ram_user;

+---------+----------------+-------------+
| name    | display_name   | mfa_enabled |
+---------+----------------+-------------+
| pam     | pam_beesly     | false       |
| creed   | creed_bratton  | true        |
| stanley | stanley_hudson | false       |
| michael | michael_scott  | false       |
| dwight  | dwight_schrute | true        |
+---------+----------------+-------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/alicloud/tables)**

## Get started

### Install

Download and install the latest Alibaba Cloud plugin:

```bash
steampipe plugin install alicloud
```

### Credentials

| Item | Description |
| - | - |
| Credentials | [Create API keys](https://www.alibabacloud.com/help/doc-detail/53045.htm) and add to `~/.steampipe/config/alicloud.spc` |
| Permissions | Minimally grant the user `AliyunOSSReadOnlyAccess`  |
| Radius | Each connection represents a single Alibaba Cloud account. |
| Resolution |  1. Credentials specified in connection argument file.<br />2. Credentials specified in environment variables. |
| Region Resolution | If `regions` is not specified, Steampipe will use the single default region. |

### Configuration

Installing the latest alicloud plugin will create a config file (`~/.steampipe/config/alicloud.spc`) with a single connection named `alicloud`:

```hcl
connection "alicloud" {
  plugin = "alicloud"

  # You may connect to one or more regions. If `regions` is not specified,
  # Steampipe will use a single default region using the below resolution
  # order:
  # The `ALIBABACLOUD_REGION_ID`, `ALICLOUD_REGION_ID` or `ALICLOUD_REGION` environment variable
  # regions = ["us-east-1", "ap-south-1"]

  # If no credentials are specified, the plugin will use the environment variables
  # resolver to get the current credentials.
  # Alternatively, you may set static credentials with the `access_key` and `secret_key` arguments.
  # access_key  	= "LTAI4GBVFakeKey09Kxezv66"
  # secret_key  	= "6iNPvThisIsNotARealSecretk1sZF"

  # List of additional Alicloud error codes to ignore for all queries.
  # By default, common not found error codes are ignored and will still be ignored even if this argument is not set.
  # ignore_error_codes = ["AccessDenied", "Forbidden.Access", "Forbidden.NoPermission"]
}
```

## Multi-Account Connections

You may create multiple alicloud connections:

```hcl
connection "alicloud_dev" {
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisISNotARealKeyVVhh"
  access_key  = "ASIA42DZSWFYSN2PFHPJ"
  regions     = ["eu-central-1" , "cn-hangzhou"]
}

connection "alicloud_qa" {
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisAintARealKeyVVhh"
  access_key  = "ASIA42DZSWFYS42PFJHP"
  regions     = ["cn-hangzhou"]
}

connection "alicloud_prod" {
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisAintARealKeyVVhh"
  access_key  = "ASIA42DZSWFYS42PFJHP"
  regions     = ["cn-hangzhou"]
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from alicloud_qa.alicloud_account;
```

You can multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "alicloud_all" {
  plugin      = "alicloud"
  type        = "aggregator"
  connections = ["alicloud_dev", "alicloud_qa", "alicloud_prod"]
}
```

Querying tables from this connection will return results from the `alicloud_dev`, `alicloud_qa`, and `alicloud_prod` connections:

```sql
select * from alicloud_all.alicloud_account;
```

Alternatively, you can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphabetically, so that it is the first connection in the search path (i.e. `alicloud_all` comes before `alicloud_dev`):

```sql
select * from alicloud_account;
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the Alicloud plugin connections whose names begin with `alicloud_`:

```hcl
connection "alicloud_all" {
  type        = "aggregator"
  plugin      = "alicloud"
  connections = ["alicloud_*"]
}
```

Aggregators are powerful, but they are not infinitely scalable. Like any other Steampipe connection, they query APIs and are subject to API limits and throttling. Consider as an example and aggregator that includes 3 Alicloud connections, where each connection queries 33 regions ([28 for `Alibaba Cloud public cloud`, 4 for `Alibaba Finance Cloud` and 1 for `Alibaba Gov Cloud`](https://www.alibabacloud.com/help/en/basics-for-beginners/latest/regions-and-zones)). This means you essentially run the same list API calls 99 times! When using aggregators, it is especially important to:

- Query only what you need! `select * from alicloud_oss_bucket` must make a list API call in each connection, and then 5 API calls *for each bucket*, where `select name, versioning from alicloud_oss_bucket` would only require a single API call per bucket.
- Consider extending the [cache TTL](https://steampipe.io/docs/reference/config-files#connection-options). The default is currently 300 seconds (5 minutes). Obviously, anytime Steampipe can pull from the cache, its is faster and less impactful to the APIs. If you don't need the most up-to-date results, increase the cache TTL!

## Specify static credentials using environment variables

Steampipe supports three different naming conventions for Alicloud authentication environment variables, checking for existence in the following order:

### Aliyun CLI format

```sh
export ALIBABACLOUD_ACCESS_KEY_ID=ASIA42DZSWFYS42PFJHP  
export ALIBABACLOUD_ACCESS_KEY_SECRET=gMCYsoGqjfThisAintARealKeyVVhh
export ALIBABACLOUD_REGION_ID=cn-east-1
```

### Terraform format

```sh
export ALICLOUD_ACCESS_KEY_ID=ASIA42DZSWFYS42PFJHP  
export ALICLOUD_ACCESS_KEY_SECRET=gMCYsoGqjfThisAintARealKeyVVhh
export ALICLOUD_REGION_ID=cn-east-1
```

### Steampipe format

```sh
export ALICLOUD_ACCESS_KEY=ASIA42DZSWFYS42PFJHP  
export ALICLOUD_SECRET_KEY=gMCYsoGqjfThisAintARealKeyVVhh
export ALICLOUD_REGION=cn-east-1
```

If regions is not specified, Steampipe will use the single default region.


