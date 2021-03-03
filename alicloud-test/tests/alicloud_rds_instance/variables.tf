
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
  name       = var.resource_name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "named_test_resource" {
  vpc_id     = alicloud_vpc.named_test_resource.id
  cidr_block = "172.16.0.0/24"
  name       = var.resource_name
}

resource "alicloud_db_instance" "named_test_resource" {
  engine                         = "MySQL"
  engine_version                 = "5.6"
  instance_type                  = "rds.mysql.s2.large"
  instance_storage               = "30"
  instance_charge_type           = "Postpaid"
  instance_name                  = var.resource_name
  monitoring_period              = "60"
  port                           = "3306"
  connection_string              = "test"
  ssl_status                     = "N"
  db_instance_ip_array_name      = "testArray"
  db_instance_ip_array_attribute = "testAttr"
  security_ip_type               = "testIP"
}


output "db_instance_id" {
  value = alicloud_db_instance.named_test_resource.instance_name
}

output "port" {
  value = alicloud_db_instance.named_test_resource.port
}

output "connection_string" {
  value = alicloud_db_instance.named_test_resource.connection_string
}

output "ssl_status" {
  value = alicloud_db_instance.named_test_resource.ssl_status
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:rds::${data.alicloud_caller_identity.current.account_id}:instance/${var.resource_name}"
}
