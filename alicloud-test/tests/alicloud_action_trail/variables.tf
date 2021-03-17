
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

# Create a new actiontrail trail.
resource "alicloud_actiontrail_trail" "named_test_resource" {
  trail_name      = var.resource_name
  oss_bucket_name = "cis-test2"
  event_rw        = "All"
  trail_region    = "cn-hangzhou"
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:actiontrail:us-east-1:${data.alicloud_caller_identity.current.account_id}:actiontrail/${var.resource_name}"
}
