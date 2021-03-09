
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

# Create a Network Interface
resource "alicloud_network_interface" "named_test_resource" {
  name              = var.resource_name
  description       = "Test network interface to validate the table outcome."
  vswitch_id        = alicloud_vswitch.named_test_resource.id
  security_groups   = [alicloud_security_group.named_test_resource.id]
  private_ip        = "192.168.0.2"
  private_ips_count = 3

  tags = {
    name = var.resource_name
  }
}

output "network_interface_id" {
  value = alicloud_network_interface.named_test_resource.id
}

output "mac_address" {
  value = alicloud_network_interface.named_test_resource.mac
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

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:ecs:us-east-1b:${data.alicloud_caller_identity.current.account_id}:eni/${alicloud_network_interface.named_test_resource.id}"
}
