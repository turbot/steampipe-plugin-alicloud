---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/alicloud.svg"
brand_color: "#FF6600"
display_name: Alibaba Cloud
name: alicloud
description: Steampipe plugin for querying Alibaba Cloud servers, databases, networks, and other resources.
---

# Alibaba Cloud

Query your Alibaba Cloud infrastructure including servers, database, networks and other resources.

### Installation

To download and install the latest Alibaba Cloud plugin:

```bash
steampipe plugin install alicloud
```

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

### Scope

Each `alicloud` connection is scoped to a single Alibaba Cloud account, with a single set of credentials. You may configure multiple `alicloud` connections if desired, with each connecting to a different account. Each `alicloud` connection may be configured for multiple regions.

### Configuration Arguments

The Alicloud plugin allows you set static credentials with the `access_key` and `secret_key` arguments. You may select one or more regions with the `regions` argument. A connection may connect to multiple regions, however be aware that performance may be negatively affected by both the number of regions and the latency to them.

```hcl
# credentials via key pair
connection "alicloud_account_x" {
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisISNotARealKeyVVhh"
  access_key  = "ASIA3ODZSWFYSN2PFHPJ"
  regions     = ["us-east-1" , "ap-south-1"]
}
```

If no credentials are specified, the plugin will get the current credentials from environment variables:

```hcl
# default
connection "alicloud" {
  plugin      = "alicloud"
}
```

The Alicloud credential resolution order is:

1. Credentials specified in connection argument file.
2. Credentials specified in environment variables.
   Environment variables are loaded in this order of precedence, aligning with the aliyun CLI (first) and Terraform (second):

| Priority |         Access Key         |           Secret Key           |         Region         |
| :------: | :------------------------: | :----------------------------: | :--------------------: |
|    1     | ALIBABACLOUD_ACCESS_KEY_ID | ALIBABACLOUD_ACCESS_KEY_SECRET | ALIBABACLOUD_REGION_ID |
|    2     |   ALICLOUD_ACCESS_KEY_ID   |   ALICLOUD_ACCESS_KEY_SECRET   |   ALICLOUD_REGION_ID   |
|    3     |    ALICLOUD_ACCESS_KEY     |      ALICLOUD_SECRET_KEY       |    ALICLOUD_REGION     |

If `regions` is not specified, Steampipe will use a single default region using the resolution order as mentioned for `Region` in above table:

Steampipe will require read access in order to query your Alicloud resources.[Create a RAM user with an access key pair](https://partners-intl.aliyun.com/help/doc-detail/116401.htm).

**Note:** Read permissions are required for tables to work.

_Steampipe does not yet automatically load `aliyun` configuration files._
