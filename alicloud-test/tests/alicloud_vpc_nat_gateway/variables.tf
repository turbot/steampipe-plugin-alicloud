
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

data "alicloud_enhanced_nat_available_zones" "named_test_resource" {
}

resource "alicloud_vswitch" "named_test_resource" {
  name              = var.resource_name
  availability_zone = data.alicloud_enhanced_nat_available_zones.named_test_resource.zones.0.zone_id
  cidr_block        = "10.10.0.0/20"
  vpc_id            = alicloud_vpc.named_test_resource.id
}

resource "alicloud_nat_gateway" "named_test_resource" {
  depends_on           = [alicloud_vswitch.named_test_resource]
  vpc_id               = alicloud_vpc.named_test_resource.id
  description          = "test nat gateway table"
  specification        = "Small"
  name                 = var.resource_name
  instance_charge_type = "PostPaid"
  vswitch_id           = alicloud_vswitch.named_test_resource.id
  nat_type             = "Enhanced"
}

output "nat_gateway_id" {
  value = alicloud_nat_gateway.named_test_resource.id
}

output "nat_gateway_name" {
  value = alicloud_nat_gateway.named_test_resource.name
}

output "vpc_id" {
  value = alicloud_nat_gateway.named_test_resource.vpc_id
}

output "nat_type" {
  value = alicloud_nat_gateway.named_test_resource.nat_type
}

output "specification" {
  value = alicloud_nat_gateway.named_test_resource.specification
}
output "description" {
  value = alicloud_nat_gateway.named_test_resource.description
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:natgateway/${alicloud_nat_gateway.named_test_resource.id}"
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}
