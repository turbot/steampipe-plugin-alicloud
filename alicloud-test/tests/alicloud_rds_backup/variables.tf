variable "resource_name" {
  type        = string
  default     = "tf-testaccdbinstance"
  description = "Name of the resource used throughout the test."
}

variable "alicloud_region" {
  type        = string
  default     = "us-east-1"
  description = "Alicloud region used for the test."
}

variable "creation" {
  default = "Rds"
}

provider "alicloud" {
  region = var.alicloud_region
}

data "alicloud_caller_identity" "current" {}

data "null_data_source" "resource" {
  inputs = {
    scope = "acs:::${data.alicloud_caller_identity.current.account_id}"
  }
}

data "alicloud_zones" "resource" {
  available_resource_creation = var.creation
}


resource "alicloud_vpc" "named_test_resource" {
  vpc_name   = var.resource_name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "named_test_resource" {
  vpc_id       = alicloud_vpc.named_test_resource.id
  zone_id      = data.alicloud_zones.resource.zones[0].id
  cidr_block   = "172.16.0.0/24"
  vswitch_name = var.resource_name
}

resource "alicloud_db_instance" "named_test_resource" {
  engine               = "MySQL"
  engine_version       = "5.7"
  instance_type        = "rds.mysql.s2.large"
  instance_storage     = "10"
  vswitch_id           = alicloud_vswitch.named_test_resource.id
  instance_charge_type = "Postpaid"
  instance_name        = var.resource_name
  monitoring_period    = "60"
}

resource "alicloud_rds_backup" "named_test_resource" {
  db_instance_id = alicloud_db_instance.named_test_resource.id
}

output "db_instance_id" {
  value = alicloud_db_instance.named_test_resource.id
}

output "port" {
  value = alicloud_db_instance.named_test_resource.port
}

output "zone_id" {
  value = alicloud_db_instance.named_test_resource.zone_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_id" {
  value = alicloud_rds_backup.named_test_resource.backup_id
}

output "store_status" {
  value = alicloud_rds_backup.named_test_resource.store_status
}
