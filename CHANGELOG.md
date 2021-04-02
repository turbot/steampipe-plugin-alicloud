## v0.0.7 [2021-04-02]

_Bug fixes_

- Fixed: `Table definitions & examples` link now points to the correct location ([#133](https://github.com/turbot/steampipe-plugin-alicloud/pull/133))

## v0.0.6 [2021-04-01]

_What's new?_

- New tables added
  - [alicloud_ecs_auto_provisioning_group](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_auto_provisioning_group) ([#107](https://github.com/turbot/steampipe-plugin-alicloud/pull/107))
  - [alicloud_ecs_launch_template](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_launch_template) ([#92](https://github.com/turbot/steampipe-plugin-alicloud/pull/92))

## v0.0.5 [2021-03-25]

_What's new?_

- New tables added
  - [alicloud_vpc_network_acl](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_vpc_network_acl) ([#48](https://github.com/turbot/steampipe-plugin-alicloud/pull/48))

_Enhancements_

- Updated: Add `parameters` column to `alicloud_rds_instance` table ([#121](https://github.com/turbot/steampipe-plugin-alicloud/pull/121))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v026-2021-03-18)

## v0.0.4 [2021-03-18]

_What's new?_

- New tables added
  - [alicloud_ecs_key_pair](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_key_pair) ([#90](https://github.com/turbot/steampipe-plugin-alicloud/pull/90))
  - [alicloud_kms_secret](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_kms_secret) ([#11](https://github.com/turbot/steampipe-plugin-alicloud/pull/11))
  - [alicloud_ram_credential_report](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ram_credential_report) ([#94](https://github.com/turbot/steampipe-plugin-alicloud/pull/94))
  - [alicloud_vpc_nat_gateway](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_vpc_nat_gateway) ([#28](https://github.com/turbot/steampipe-plugin-alicloud/pull/28))

_Enhancements_

- Updated: Add `assume_role_policy_document_std` column to `alicloud_ram_role` table ([#113](https://github.com/turbot/steampipe-plugin-alicloud/pull/113))
- Updated: Add `ssl_status` and `tde_status` columns to `alicloud_rds_instance` table ([#118](https://github.com/turbot/steampipe-plugin-alicloud/pull/118))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v024-2021-03-16)

_Bug fixes_

- Fixed: Simplified security IP related columns in `alicloud_rds_instance` table ([#117](https://github.com/turbot/steampipe-plugin-alicloud/pull/117))
  - Columns added:
    - security_ips
    - security_ips_src
  - Columns removed:
    - db_instance_ip_array_attribute
    - db_instance_ip_array_name
    - security_ip_list
    - security_ip_type
    - whitelist_network_type
- Fixed: `logging` column in `alicloud_oss_bucket` table now returns the correct data instead of `null` ([#110](https://github.com/turbot/steampipe-plugin-alicloud/pull/110))

## v0.0.3 [2021-03-11]

_What's new?_

- New tables added
  - [alicloud_kms_key](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_kms_key)
  - [alicloud_rds_instance](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_instance)
  - [alicloud_vpc_route_table](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_vpc_route_table)

## v0.0.2 [2021-03-04]

_What's new?_

- New tables added
  - alicloud_ecs_autoscaling_group
  - alicloud_ecs_network_interface
  - alicloud_vpc_eip
  - alicloud_vpc_ssl_vpn_client_cert
  - alicloud_vpc_ssl_vpn_server
  - alicloud_vpc_vpn_connection
  - alicloud_vpc_vpn_customer_gateway
  - alicloud_vpc_vpn_gateway

## v0.0.1 [2021-02-25]

_What's new?_

- Initial release with tables for RAM, OSS, ECS and VPC service.
