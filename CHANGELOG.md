## v0.4.0 [2022-03-25]

_What's new?_

- New tables added
  - [alicloud_ecs_instance_metric_cpu_utilization_hourly](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_instance_metric_cpu_utilization_hourly) ([#244](https://github.com/turbot/steampipe-plugin-alicloud/pull/244))

_Enhancements_

- Recompiled `alicloud_ecs_instance_metric_cpu_utilization_hourly` table with [steampipe-plugin-sdk v2.1.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v211--2022-03-10) ([#245](https://github.com/turbot/steampipe-plugin-alicloud/pull/245))

## v0.3.1 [2021-12-15]

_Bug fixes_

- Fixed the `ContentMD5NotMatched` error response in the `alicloud_ram_user` table ([#240](https://github.com/turbot/steampipe-plugin-alicloud/pull/240))

## v0.3.0 [2021-11-23]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#236](https://github.com/turbot/steampipe-plugin-alicloud/pull/236))

## v0.2.0 [2021-10-26]

_Enhancements_

- Updated: Add new optional key quals, filter support, context cancellation handling, and limit support for `alicloud_ecs_instance`, `alicloud_kms_key`, `alicloud_ram_policy`, and `alicloud_vpc` tables ([#228](https://github.com/turbot/steampipe-plugin-alicloud/pull/228))
- Recompiled plugin with [steampipe-plugin-sdk v1.7.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v170--2021-10-18)

## v0.1.1 [2021-09-13]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.5.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v151--2021-09-13) ([#230](https://github.com/turbot/steampipe-plugin-alicloud/pull/230))

## v0.1.0 [2021-08-04]

_What's new?_

- New tables added
  - [alicloud_rds_instance_metric_cpu_utilization](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_instance_metric_cpu_utilization) ([#214](https://github.com/turbot/steampipe-plugin-alicloud/pull/214))
  - [alicloud_rds_instance_metric_cpu_utilization_daily](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_instance_metric_cpu_utilization_daily) ([#214](https://github.com/turbot/steampipe-plugin-alicloud/pull/214))

_Enhancements_

- Updated: Add retry mechanism for throttling in `alicloud_rds_instance` table ([#220](https://github.com/turbot/steampipe-plugin-alicloud/pull/220))
- Updated: Improve caching when getting common columns in all tables ([#215](https://github.com/turbot/steampipe-plugin-alicloud/pull/215))

_Bug fixes_

- Fixed: Add retry mechanism when getting intermittent API runtime error responses in metric statistic tables ([#225](https://github.com/turbot/steampipe-plugin-alicloud/pull/225))

## v0.0.14 [2021-07-22]

_What's new?_

- New tables added
  - [alicloud_ecs_disk_metric_read_iops](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_disk_metric_read_iops) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))
  - [alicloud_ecs_disk_metric_read_ops_daily](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_disk_metric_read_ops_daily) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))
  - [alicloud_ecs_disk_metric_write_iops](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_disk_metric_write_iops) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))
  - [alicloud_ecs_disk_metric_write_iops_daily](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_disk_metric_write_iops_daily) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))
  - [alicloud_ecs_instance_metric_cpu_utilization_daily](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_instance_metric_cpu_utilization_daily) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))
  - [alicloud_ram_policy](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ram_policy) ([#86](https://github.com/turbot/steampipe-plugin-alicloud/pull/86))
  - [alicloud_rds_instance_metric_connections](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_instance_metric_connections) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))
  - [alicloud_rds_instance_metric_connections_daily](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_instance_metric_connections_daily) ([#204](https://github.com/turbot/steampipe-plugin-alicloud/pull/204))

_Enhancements_

- Updated: Add column `arn` to `alicloud_vpc_eip` table ([#198](https://github.com/turbot/steampipe-plugin-alicloud/pull/198))
- Updated: Add column `arn` to `alicloud_ecs_snapshot` table ([#197](https://github.com/turbot/steampipe-plugin-alicloud/pull/197))
- Updated: Add column `lifecycle_rules` to `alicloud_oss_bucket` table ([#189](https://github.com/turbot/steampipe-plugin-alicloud/pull/189))
- Updated: `alicloud_action_trail` table now also returns shadow trails ([#208](https://github.com/turbot/steampipe-plugin-alicloud/pull/208))
- Recompiled plugin with [steampipe-plugin-sdk v1.3.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v131--2021-07-15)

## v0.0.13 [2021-06-24]

_What's new?_

- New tables added
  - [alicloud_cms_monitor_host](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_cms_monitor_host) ([#180](https://github.com/turbot/steampipe-plugin-alicloud/pull/180))
  - [alicloud_cs_kubernetes_cluster_node](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_cs_kubernetes_cluster_node) ([#182](https://github.com/turbot/steampipe-plugin-alicloud/pull/182))

## v0.0.12 [2021-06-17]

_What's new?_

- New tables added
  - [alicloud_security_center_field_statistics](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_security_center_field_statistics) ([#177](https://github.com/turbot/steampipe-plugin-alicloud/pull/177))
  - [alicloud_security_center_version](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_security_center_version) ([#173](https://github.com/turbot/steampipe-plugin-alicloud/pull/173))

_Enhancements_

- Updated: Add `name` column to get key columns in `alicloud_ecs_autoscaling_group` table ([#179](https://github.com/turbot/steampipe-plugin-alicloud/pull/179))
- Updated: Add column `arn` to `alicloud_oss_bucket` table ([#155](https://github.com/turbot/steampipe-plugin-alicloud/pull/155))
- Updated: Add column `cs_user_permissions` to `alicloud_ram_user` table ([#183](https://github.com/turbot/steampipe-plugin-alicloud/pull/183))
- Updated: Add columns `arn` and `cluster_namespace` to `alicloud_cs_kubernetes_cluster` table ([#166](https://github.com/turbot/steampipe-plugin-alicloud/pull/166))

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
