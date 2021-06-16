## v0.0.11 [2021-06-10]

_What's new?_

- New tables added
  - [alicloud_account](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_account) ([#175](https://github.com/turbot/steampipe-plugin-alicloud/pull/175))

_Enhancements_

- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488)
- Updated: Add column `arn` to `alicloud_ecs_disk` table ([#152](https://github.com/turbot/steampipe-plugin-alicloud/pull/152))
- Updated: Add column `arn` to `alicloud_ecs_image` table ([#153](https://github.com/turbot/steampipe-plugin-alicloud/pull/153))
- Updated: Add column `arn` to `alicloud_ecs_instance` table ([#151](https://github.com/turbot/steampipe-plugin-alicloud/pull/151))
- Updated: Add column `arn` to `alicloud_ecs_security_group` table ([#154](https://github.com/turbot/steampipe-plugin-alicloud/pull/154))
- Updated: Add column `arn` to `alicloud_rds_instance` table ([#156](https://github.com/turbot/steampipe-plugin-alicloud/pull/156))
- Updated: Add columns `sql_collector_policy` and `sql_collector_retention` to `alicloud_rds_instance` table ([#176](https://github.com/turbot/steampipe-plugin-alicloud/pull/176))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.10](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v0210-2021-06-09)

## v0.0.10 [2021-04-29]

_What's new?_

- New tables added
  - [alicloud_cs_kubernetes_cluster](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_cs_kubernetes_cluster) ([#128](https://github.com/turbot/steampipe-plugin-alicloud/pull/128))

_Bug fixes_

- Fixed: Remove unsupported `mns_topic_arn` column from `alicloud_action_trail` table ([#139](https://github.com/turbot/steampipe-plugin-alicloud/pull/139))

## v0.0.9 [2021-04-15]

_What's new?_

- New tables added
  - [alicloud_action_trail](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_action_trail) ([#122](https://github.com/turbot/steampipe-plugin-alicloud/pull/122))
  - [alicloud_cas_certificate](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_cas_certificate) ([#101](https://github.com/turbot/steampipe-plugin-alicloud/pull/101))
  - [alicloud_vpc_route_entry](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_vpc_route_entry) ([#95](https://github.com/turbot/steampipe-plugin-alicloud/pull/95))

## v0.0.8 [2021-04-08]

_What's new?_

- New tables added
  - [alicloud_ecs_region](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_region) ([#91](https://github.com/turbot/steampipe-plugin-alicloud/pull/91))
  - [alicloud_ecs_zone](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_zone) ([#91](https://github.com/turbot/steampipe-plugin-alicloud/pull/91))

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
