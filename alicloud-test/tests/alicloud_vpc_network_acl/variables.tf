
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
  cidr_block = "172.16.0.0/12"
  vpc_name   = var.resource_name
}

resource "alicloud_network_acl" "named_test_resource" {
  vpc_id      = alicloud_vpc.named_test_resource.id
  name        = var.resource_name
  description = "Test network acl table"
}

output "network_acl_id" {
  value = alicloud_network_acl.named_test_resource.id
}

output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:${var.alicloud_region}:${data.alicloud_caller_identity.current.account_id}:network-acl/${alicloud_network_acl.named_test_resource.id}"
}

output "region_id" {
  value = var.alicloud_region
}
