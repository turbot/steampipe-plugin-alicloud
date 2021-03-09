
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

resource "alicloud_nat_gateway" "named_test_resource" {
  name                 = var.resource_name
  nat_type             = "Enhanced"
  specification        = "Small"
  description          = "This is a test NAT gateway to verify the table outcome."
  vpc_id               = alicloud_vpc.named_test_resource.id
  instance_charge_type = "PostPaid"
  vswitch_id           = alicloud_vswitch.named_test_resource.id

  depends_on           = [alicloud_vswitch.named_test_resource]
}

output "nat_gateway_id" {
  value = alicloud_nat_gateway.named_test_resource.id
}

output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
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
