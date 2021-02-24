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

## Installation

Download and install the latest Alibaba Cloud plugin:

```bash
steampipe plugin install alicloud
```

## Configure API Token

[Create a RAM user with an access key pair](https://partners-intl.aliyun.com/help/doc-detail/116401.htm).
Read permissions are required for tables to work.

Set the access keys in environment variables:

```bash
export ALICLOUD_ACCESS_KEY=LTAI8GBwJQkMs697FDp3GASK
export ALICLOUD_SECRET_KEY=0tIWnS1RvXvcBj4oR4QNb8mbtt85aBr
export ALICLOUD_REGION=us-east-1
```

Environment variables are loaded in this order of precedence, aligning with the aliyun CLI (first)
and Terraform (second):

| Priority | Access Key | Secret Key | Region |
| 1 | `ALIBABACLOUD_ACCESS_KEY_ID` | `ALIBABACLOUD_ACCESS_KEY_SECRET` | `ALIBABACLOUD_REGION_ID` |
| 2 | `ALICLOUD_ACCESS_KEY_ID` | `ALICLOUD_ACCESS_KEY_SECRET` | `ALICLOUD_REGION_ID` |
| 3 | `ALICLOUD_ACCESS_KEY` | `ALICLOUD_SECRET_KEY` | `ALICLOUD_REGION` |

Steampipe does not yet automatically load `aliyun` configuration files.

## Your first query

```bash
~ $ steampipe query
Welcome to Steampipe v0.1.0
Type ".inspect" for more information.
> select name, display_name, create_date from alicloud_ram_user;
+--------------+----------------+---------------------+
|     name     |  display_name  |     create_date     |
+--------------+----------------+---------------------+
| pam          | Pam            | 2021-01-27 02:52:22 |
| jim          | Jim            | 2021-01-27 02:52:22 |
| michael      | Michael        | 2021-01-27 02:52:22 |
| dwight       | Dwight         | 2021-01-27 02:52:22 |
+--------------+----------------+---------------------+
```
