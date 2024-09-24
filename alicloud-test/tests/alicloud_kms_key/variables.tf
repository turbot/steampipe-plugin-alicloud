
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

# Create a KMS Key
resource "alicloud_kms_key" "named_test_resource" {
  description            = "This is a test key used to validate the table outcome."
  pending_window_in_days = "7"
  key_state              = "Enabled"
  key_usage              = "ENCRYPT/DECRYPT"
  protection_level       = "SOFTWARE"
  automatic_rotation     = "Disabled"
}

# Create a KMS Alias
resource "alicloud_kms_alias" "named_test_resource" {
  alias_name = "alias/${var.resource_name}"
  key_id     = alicloud_kms_key.named_test_resource.id
}

output "resource_aka" {
  value = alicloud_kms_key.named_test_resource.arn
}

output "pending_window_in_days" {
  value = alicloud_kms_key.named_test_resource.pending_window_in_days
}

output "resource_id" {
  value = alicloud_kms_key.named_test_resource.id
}

output "key_state" {
  value = alicloud_kms_key.named_test_resource.key_state
}

output "key_usage" {
  value = alicloud_kms_key.named_test_resource.key_usage
}

output "protection_level" {
  value = alicloud_kms_key.named_test_resource.protection_level
}

output "automatic_rotation" {
  value = alicloud_kms_key.named_test_resource.automatic_rotation
}

output "resource_name" {
  value = var.resource_name
}

output "region" {
  value = var.alicloud_region
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}
