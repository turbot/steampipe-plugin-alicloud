
variable "resource_name" {
  type        = string
  default     = "turbot-test-20200125-create-update"
  description = "Name of the resource used throughout the test."
}

variable "alicloud_region" {
  type        = string
  default     = "cn-hangzhou"
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

resource "alicloud_oss_bucket" "named_test_resource" {
  bucket = var.resource_name
  acl    = "private"
}

# Create a new actiontrail trail.
resource "alicloud_actiontrail_trail" "named_test_resource" {
  trail_name      = var.resource_name
  oss_bucket_name = var.resource_name
  event_rw        = "All"
  trail_region    = var.alicloud_region
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "region_name"{
  value = var.alicloud_region
}
output "resource_aka" {
  value = "acs:actiontrail:cn-hangzhou:${data.alicloud_caller_identity.current.account_id}:actiontrail/${var.resource_name}"
}
