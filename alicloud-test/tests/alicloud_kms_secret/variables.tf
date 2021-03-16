
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

# Create a new KMS secret.
resource "alicloud_kms_secret" "named_test_resource" {
  secret_name                   = var.resource_name
  description                   = "This is a test secret to validate the table outcome."
  secret_data                   = "Secret data."
  version_id                    = "000000000001"
  force_delete_without_recovery = true
  tags = {
    Name = var.resource_name
  }
}

output "resource_name" {
  value = alicloud_kms_secret.named_test_resource.id
}

output "resource_id" {
  value = alicloud_kms_secret.named_test_resource.id
}

output "resource_aka" {
  value = alicloud_kms_secret.named_test_resource.arn
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}
