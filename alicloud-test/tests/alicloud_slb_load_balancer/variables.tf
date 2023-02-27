
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
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "named_test_resource" {
  name              = var.resource_name
  availability_zone = "us-east-1b"
  cidr_block        = "10.10.0.0/20"
  vpc_id            = alicloud_vpc.named_test_resource.id
}

resource "alicloud_slb_load_balancer" "named_test_resource" {
  load_balancer_name = var.resource_name
  address_type       = "intranet"
  load_balancer_spec = "slb.s2.small"
  vswitch_id         = alicloud_vswitch.named_test_resource.id
  tags = {
    info = "create for internet"
  }
  instance_charge_type = "PayBySpec"
}


output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_id" {
  value = alicloud_slb_load_balancer.named_test_resource.id
}

output "resource_address" {
  value = alicloud_slb_load_balancer.named_test_resource.address
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}
