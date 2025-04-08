---
title: "Steampipe Table: alicloud_account - Query Alibaba Cloud Accounts using SQL"
description: "Allows users to query Alibaba Cloud Accounts, specifically the account details such as account ID, account name, and account type."
folder: "Account"
---

# Table: alicloud_account - Query Alibaba Cloud Accounts using SQL

An Alibaba Cloud Account is a basic organizational unit of Alibaba Cloud resources. It is used to sign up for and manage Alibaba Cloud products and services, and to manage resource access permissions. It is also used to manage billing by setting up payment methods and managing invoices.

## Table Usage Guide

The `alicloud_account` table provides insights into Alibaba Cloud Accounts. As a Cloud Administrator, explore account-specific details through this table, including account ID, account name, and account type. Utilize it to uncover information about accounts, such as those with specific account types, the account names, and the verification of account IDs.

## Examples

### Basic info
Explore the different aliases, account IDs, and titles within your Alicloud account. This can be particularly useful to understand the structure and organization of your account, aiding in efficient management and oversight.

```sql+postgres
select
  alias,
  account_id,
  akas,
  title
from
  alicloud_account;
```

```sql+sqlite
select
  alias,
  account_id,
  akas,
  title
from
  alicloud_account;
```