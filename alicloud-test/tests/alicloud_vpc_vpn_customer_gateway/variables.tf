
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

# Create a VPN Customer Gateway.
resource "alicloud_vpn_customer_gateway" "named_test_resource" {
  name        = var.resource_name
  ip_address  = "43.104.22.228"
  description = "Test customer gateway to validate the table outcome."
}

output "resource_id" {
  value = alicloud_vpn_customer_gateway.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:customergateway/${alicloud_vpn_customer_gateway.named_test_resource.id}"
}
