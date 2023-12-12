## v0.21.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#381](https://github.com/turbot/steampipe-plugin-alicloud/pull/381))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#381](https://github.com/turbot/steampipe-plugin-alicloud/pull/381))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-alicloud/blob/main/docs/LICENSE). ([#381](https://github.com/turbot/steampipe-plugin-alicloud/pull/381))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#378](https://github.com/turbot/steampipe-plugin-alicloud/pull/378))

## v0.20.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#360](https://github.com/turbot/steampipe-plugin-alicloud/pull/360))

## v0.20.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#356](https://github.com/turbot/steampipe-plugin-alicloud/pull/356))
- Recompiled plugin with Go version `1.21`. ([#356](https://github.com/turbot/steampipe-plugin-alicloud/pull/356))

## v0.19.0 [2023-06-20]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. ([#346](https://github.com/turbot/steampipe-plugin-alicloud/pull/346))

## v0.18.0 [2023-06-08]

_What's new?_

- Added `ignore_error_codes` config arg to provide users the ability to set a list of additional Alibaba Cloud error codes to ignore while running queries. For instance, to ignore some common access denied errors, which is helpful when running with limited permissions, set the argument `ignore_error_codes = ["AccessDenied", "Forbidden.Access", "Forbidden.NoPermission"]`. For more information, please see [Alibaba Cloud plugin configuration](https://hub.steampipe.io/plugins/turbot/alicloud#configuration) ([#344](https://github.com/turbot/steampipe-plugin-alicloud/pull/344))

## v0.17.0 [2023-05-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05) which fixes increased plugin initialization time due to multiple connections causing the schema to be loaded repeatedly. ([#342](https://github.com/turbot/steampipe-plugin-alicloud/pull/342))

## v0.16.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which adds go-getter support to dynamic tables. ([#340](https://github.com/turbot/steampipe-plugin-alicloud/pull/340))

## v0.15.0 [2023-03-15]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.2.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v520-2023-03-02) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#337](https://github.com/turbot/steampipe-plugin-alicloud/pull/337))

## v0.14.0 [2023-02-27]

_What's new?_

- New tables added
  - [alicloud_slb_load_balancer](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_slb_load_balancer) ([#321](https://github.com/turbot/steampipe-plugin-alicloud/pull/321))

_Enhancements_

- Added column `encryption_key` to `alicloud_rds_instance` table. ([#301](https://github.com/turbot/steampipe-plugin-alicloud/pull/301))
- Updated the `docs/index.md` file to include multi-account configuration information and examples. ([#334](https://github.com/turbot/steampipe-plugin-alicloud/pull/334) (Thanks [@vmdude](https://github.com/vmdude) for the contribution!!)

## v0.13.1 [2023-02-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.12](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v4112-2023-02-09) which fixes the query caching functionality. ([#331](https://github.com/turbot/steampipe-plugin-alicloud/pull/331))

## v0.13.0 [2023-01-20]

_Enhancements_

- Updated the `title` column of `alicloud_kms_key` table to first use the key alias if available, else fall back to the key ID. ([#328](https://github.com/turbot/steampipe-plugin-alicloud/pull/328))

_Bug fixes_

- Fixed the column `consistent_time` in `alicloud_rds_backup` table to correctly return data instead of an error. ([#327](https://github.com/turbot/steampipe-plugin-alicloud/pull/327))

## v0.12.0 [2023-01-19]

_Enhancements_

- Added column `security_group_configuration` to `alicloud_rds_instance` table. ([#319](https://github.com/turbot/steampipe-plugin-alicloud/pull/319))

## v0.11.0 [2023-01-17]

_What's new?_

- New tables added
  - [alicloud_rds_backup](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_backup) ([#300](https://github.com/turbot/steampipe-plugin-alicloud/pull/300))
  - [alicloud_rds_database](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_database) ([#299](https://github.com/turbot/steampipe-plugin-alicloud/pull/299))
  - [alicloud_vpc_flow_log](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_vpc_flow_log) ([#312](https://github.com/turbot/steampipe-plugin-alicloud/pull/312))

_Enhancements_

- Added column `ram_role` to `alicloud_ecs_instance` table. ([#313](https://github.com/turbot/steampipe-plugin-alicloud/pull/313))
- Updated the `alicloud_ecs_image` table to return the details of a particular image when `image_id` and `region` are passed in the `where` clause. ([#315](https://github.com/turbot/steampipe-plugin-alicloud/pull/315))

_Bug fixes_

- Fixed the `alicloud_ecs_image` table to return the details of an image instead of returning an error when multiple regions are configured in the `alicloud.spc` file. ([#315](https://github.com/turbot/steampipe-plugin-alicloud/pull/315))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.9](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v419-2022-11-30) which fixes hydrate function caching for aggregator connections. ([#316](https://github.com/turbot/steampipe-plugin-alicloud/pull/316))

## v0.10.1 [2022-11-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.8](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v418-2022-09-08) which increases the default open file limit. ([#294](https://github.com/turbot/steampipe-plugin-alicloud/pull/294))

## v0.10.0 [2022-09-06]

_Enhancements_

- Added retry mechanism for `alicloud_ram_user` table to handle API throttling errors.  ([#287](https://github.com/turbot/steampipe-plugin-alicloud/pull/287)) (Thanks to [@jurajsucik](https://github.com/jurajsucik) for the contribution!)

_Bug fixes_

- Fixed the `alicloud_cs_kubernetes_cluster` table to correctly return the value for column `cluster_id` instead of returning `null`. ([#290](https://github.com/turbot/steampipe-plugin-alicloud/pull/290))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#288](https://github.com/turbot/steampipe-plugin-alicloud/pull/288))
- Recompiled plugin with Go version `1.19`. ([#288](https://github.com/turbot/steampipe-plugin-alicloud/pull/288))

## v0.9.0 [2022-08-04]

_What's new?_

- Added support for `cn-beijing-finance-1`, `cn-shanghai-finance-1` and `cn-shenzhen-finance-1` regions. ([#284](https://github.com/turbot/steampipe-plugin-alicloud/pull/284)) (Thanks to [@jurajsucik](https://github.com/jurajsucik) for the contribution!)

_Bug fixes_

- Fixed `password_exist`, `password_active`, `mfa_active`, `user_last_logon`, `password_last_changed` and `password_next_rotation` columns in the `alicloud_ram_credential_report` table to return null instead of an error when the console login is disabled for a user. ([#284](https://github.com/turbot/steampipe-plugin-alicloud/pull/284)) (Thanks to [@jurajsucik](https://github.com/jurajsucik) for the contribution!)

## v0.8.0 [2022-07-14]

_Bug fixes_

- Fixed inconsistent table name in the `alicloud_vpc_vswitch` table which would cause intermittent caching issues. ([#281](https://github.com/turbot/steampipe-plugin-alicloud/pull/281))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#280](https://github.com/turbot/steampipe-plugin-alicloud/pull/280))

## v0.7.1 [2022-05-24]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#277](https://github.com/turbot/steampipe-plugin-alicloud/pull/277))

## v0.7.0 [2022-04-27]

_What's new?_

- New tables added
  - [alicloud_rds_instance_metric_cpu_utilization_hourly](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_rds_instance_metric_cpu_utilization_hourly) ([#272](https://github.com/turbot/steampipe-plugin-alicloud/pull/272))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#274](https://github.com/turbot/steampipe-plugin-alicloud/pull/274))
- Added support for native Linux ARM and Mac M1 builds. ([#275](https://github.com/turbot/steampipe-plugin-alicloud/pull/275))

## v0.6.0 [2022-04-07]

_Enhancements_

- Added column `arn` to `alicloud_ram_user`, `alicloud_ram_group` and `alicloud_vpc` tables ([#267](https://github.com/turbot/steampipe-plugin-alicloud/pull/267)) ([#268](https://github.com/turbot/steampipe-plugin-alicloud/pull/268)) ([#269](https://github.com/turbot/steampipe-plugin-alicloud/pull/269))
- Added column `virtual_mfa_devices` to `alicloud_ram_user` table ([#266](https://github.com/turbot/steampipe-plugin-alicloud/pull/266))

## v0.5.0 [2022-04-06]

_What's new?_

- New tables added
  - [alicloud_vpc_dhcp_options_set](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_vpc_dhcp_options_set) ([#259](https://github.com/turbot/steampipe-plugin-alicloud/pull/259))
  - [alicloud_ecs_disk_metric_read_iops_hourly](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_disk_metric_read_iops_hourly) ([#253](https://github.com/turbot/steampipe-plugin-alicloud/pull/253))
  - [alicloud_ecs_disk_metric_write_iops_hourly](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_disk_metric_write_iops_hourly) ([#254](https://github.com/turbot/steampipe-plugin-alicloud/pull/254))

_Enhancements_

- Added column `ipv6_cidr_blocks` to `alicloud_vpc` table ([#260](https://github.com/turbot/steampipe-plugin-alicloud/pull/260))
- Added column `attachments` to `alicloud_ecs_disk` table ([#256](https://github.com/turbot/steampipe-plugin-alicloud/pull/256))
- Added column `deletion_protection` to `alicloud_kms_key` table ([#248](https://github.com/turbot/steampipe-plugin-alicloud/pull/248))

## v0.4.0 [2022-03-25]

_What's new?_

- New tables added
  - [alicloud_ecs_instance_metric_cpu_utilization_hourly](https://hub.steampipe.io/plugins/turbot/alicloud/tables/alicloud_ecs_instance_metric_cpu_utilization_hourly) ([#244](https://github.com/turbot/steampipe-plugin-alicloud/pull/244))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v2.1.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v211--2022-03-10) ([#245](https://github.com/turbot/steampipe-plugin-alicloud/pull/245))

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
