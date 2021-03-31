
variable "resource_name" {
  type        = string
  default     = "turbot-test-20200125-create-update"
  description = "Name of the resource used throughout the test."
}

variable "alicloud_region" {
  type        = string
  default     = "us-east-1"
  description = "Alicloud region used for the test."
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

resource "alicloud_vpc" "default" {
  name       = var.resource_name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id            = alicloud_vpc.default.id
  cidr_block        = "172.16.0.0/24"
  availability_zone = "us-east-1b"
  name              = var.resource_name
}

resource "alicloud_security_group" "default" {
  name   = var.resource_name
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_auto_provisioning_group" "named_test_resource" {
  auto_provisioning_group_name  = var.resource_name
  launch_template_id            = alicloud_launch_template.named_test_resource.id
  total_target_capacity         = "5"
  pay_as_you_go_target_capacity = "0"
  spot_target_capacity          = "5"
  terminate_instances           = true
  launch_template_config {
    instance_type     = "ecs.t5-lc2m1.nano"
    vswitch_id        = alicloud_vswitch.default.id
    weighted_capacity = "1"
    max_price         = "1"
  }
}

resource "alicloud_launch_template" "named_test_resource" {
  name                          = var.resource_name
  image_id                      = "aliyun_2_1903_x64_20G_alibase_20210120.vhd"
  host_name                     = "turbot"
  instance_charge_type          = "PostPaid"
  instance_name                 = var.resource_name
  instance_type                 = "ecs.t5-lc2m1.nano"
  internet_charge_type          = "PayByTraffic"
  internet_max_bandwidth_out    = 0
  io_optimized                  = "none"
  network_type                  = "vpc"
  security_enhancement_strategy = "Active"
  security_group_id             = alicloud_security_group.default.id
  system_disk_category          = "cloud_efficiency"
  system_disk_description       = "test disk"
  system_disk_name              = "hello"
  system_disk_size              = 20
  vswitch_id                    = alicloud_vswitch.default.id
  vpc_id                        = alicloud_vpc.default.id
  zone_id                       = "us-east-1b"
}

output "vswitch_id" {
  value = alicloud_vswitch.default.id
}

output "launch_template_id" {
  value = alicloud_launch_template.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_id" {
  value = alicloud_auto_provisioning_group.named_test_resource.id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:ecs:us-east-1:${data.alicloud_caller_identity.current.account_id}:auto-provisioning-group/${alicloud_auto_provisioning_group.named_test_resource.id}"
}
