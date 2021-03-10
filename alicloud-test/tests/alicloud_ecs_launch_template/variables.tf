
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

# Create a VPC
resource "alicloud_vpc" "named_test_resource" {
  name       = var.resource_name
  cidr_block = "192.168.0.0/24"
}

# Create a VSwitch
resource "alicloud_vswitch" "named_test_resource" {
  name              = var.resource_name
  cidr_block        = "192.168.0.0/24"
  availability_zone = "us-east-1b"
  vpc_id            = alicloud_vpc.named_test_resource.id
}

# Create a VPC Security Group
resource "alicloud_security_group" "named_test_resource" {
  name   = var.resource_name
  vpc_id = alicloud_vpc.named_test_resource.id
}

# Create a Launch Template
resource "alicloud_launch_template" "named_test_resource" {
  name                          = var.resource_name
  description                   = "This is test launch template to validate the table outcome."
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
  security_group_id             = alicloud_security_group.named_test_resource.id
  system_disk_category          = "cloud_efficiency"
  system_disk_description       = "test disk"
  system_disk_name              = "hello"
  system_disk_size              = 20
  vswitch_id                    = alicloud_vswitch.named_test_resource.id
  vpc_id                        = alicloud_vpc.named_test_resource.id
  zone_id                       = "us-east-1b"

  tags = {
    name = var.resource_name
  }
}

output "resource_id" {
  value = alicloud_launch_template.named_test_resource.id
}

output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
}

output "vswitch_id" {
  value = alicloud_vswitch.named_test_resource.id
}

output "security_group_id" {
  value = alicloud_security_group.named_test_resource.id
}

output "caller_id" {
  value = data.alicloud_caller_identity.current.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:ecs:us-east-1:${data.alicloud_caller_identity.current.account_id}:launch-template/${alicloud_launch_template.named_test_resource.id}"
}
