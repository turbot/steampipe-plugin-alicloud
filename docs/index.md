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
---

# Alibaba Cloud + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Alibaba Cloud](https://alibabacloud.com) provides on-demand cloud computing platforms and APIs to authenticated customers on a metered pay-as-you-go basis. 

For example:

```sql
select
  name,
  display_name,
  mfa_enabled
from
  alicloud_ram_user;
```

```
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
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisISNotARealKeyVVhh"
  access_key  = "ASIA42DZSWFYSN2PFHPJ"
  regions     = ["us-east-1" , "ap-south-1"]
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-alicloud
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)


## Advanced configuration options

For users with multiple accounts and more complex authentication use cases, here are some examples of advanced configuration options:

### Specify multiple accounts 
A common configuration is to have multiple connections to different accounts:

```hcl
connection "ali_account_aaa" {
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisISNotARealKeyVVhh"
  access_key  = "ASIA42DZSWFYSN2PFHPJ"
  regions     = ["us-east-1" , "ap-south-1"]
}

connection "ali_account_bbb" {
  plugin      = "alicloud"
  secret_key  = "gMCYsoGqjfThisAintARealKeyVVhh"
  access_key  = "ASIA42DZSWFYS42PFJHP"
  regions     = ["cn-east-1"]
}

```

### Specify static credentials using environment variables 
Steampipe supports three different naming conventions for Alicloud authentication environment variables, checking for existence in the following order:

1. aliyun CLI format
```sh
export ALIBABACLOUD_ACCESS_KEY_ID=ASIA42DZSWFYS42PFJHP  
export ALIBABACLOUD_ACCESS_KEY_SECRET=gMCYsoGqjfThisAintARealKeyVVhh
export ALIBABACLOUD_REGION_ID=cn-east-1
``` 

2. Terraform format
```sh
export ALICLOUD_ACCESS_KEY_ID=ASIA42DZSWFYS42PFJHP  
export ALICLOUD_ACCESS_KEY_SECRET=gMCYsoGqjfThisAintARealKeyVVhh
export ALICLOUD_REGION_ID=cn-east-1
``` 

3. Steampipe format
```sh
export ALICLOUD_ACCESS_KEY=ASIA42DZSWFYS42PFJHP  
export ALICLOUD_SECRET_KEY=gMCYsoGqjfThisAintARealKeyVVhh
export ALICLOUD_REGION=cn-east-1
``` 

If regions is not specified, Steampipe will use the single default region.
