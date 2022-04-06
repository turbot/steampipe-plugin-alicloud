
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

resource "alicloud_vpc_dhcp_options_set" "named_test_resource" {
  dhcp_options_set_name        = var.resource_name
  dhcp_options_set_description = var.resource_name
  domain_name                  = "example.com"
  domain_name_servers          = "100.100.2.136"
}

output "dhcp_options_set_id" {
  value = alicloud_vpc_dhcp_options_set.named_test_resource.id
}

output "dhcp_options_set_stauts" {
  value = alicloud_vpc_dhcp_options_set.named_test_resource.status
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:dhcpoptionset/${alicloud_vpc_dhcp_options_set.named_test_resource.id}"
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}
