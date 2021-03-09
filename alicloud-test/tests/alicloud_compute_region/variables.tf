
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

data "alicloud_regions" "named_test_resource" {
  current = true
}

output "current_region_id" {
  value = data.alicloud_regions.named_test_resource.regions.0.id
}

output "local_name" {
  value = data.alicloud_regions.named_test_resource.regions.0.local_name
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:ecs::${data.alicloud_caller_identity.current.account_id}:zone/${data.alicloud_regions.named_test_resource.regions.0.id}"
}
