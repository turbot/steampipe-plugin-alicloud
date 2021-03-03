
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
  cidr_block = "172.16.0.0/12"
}

# Create a VSwitch
resource "alicloud_vswitch" "named_test_resource" {
  vpc_id            = alicloud_vpc.named_test_resource.id
  cidr_block        = "172.16.0.0/21"
  availability_zone = "us-east-1b"
}

# Create a VPN Gateway
resource "alicloud_vpn_gateway" "named_test_resource" {
  name                 = var.resource_name
  vpc_id               = alicloud_vpc.named_test_resource.id
  bandwidth            = "10"
  enable_ssl           = true
  instance_charge_type = "PostPaid"
  description          = "Test VPN gateway to validate the table outcome."
  vswitch_id           = alicloud_vswitch.named_test_resource.id
}

output "vpn_gateway_id" {
  value = alicloud_vpn_gateway.named_test_resource.id
}

output "internet_ip" {
  value = alicloud_vpn_gateway.named_test_resource.internet_ip
}

output "status" {
  value = alicloud_vpn_gateway.named_test_resource.status
}

output "business_status" {
  value = alicloud_vpn_gateway.named_test_resource.business_status
}

output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
}

output "vswitch_id" {
  value = alicloud_vswitch.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:vpngateway/${alicloud_vpn_gateway.named_test_resource.id}"
}
