
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
  version = "=1.117.0"
  region  = var.alicloud_region
}

data "alicloud_caller_identity" "current" {}

data "null_data_source" "resource" {
  inputs = {
    scope = "acs:::${data.alicloud_caller_identity.current.account_id}"
  }
}

resource "alicloud_key_pair" "named_test_resource" {
  key_name = "terraform-test-key-pair"
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "fingerprint" {
  value = alicloud_key_pair.named_test_resource.finger_print
}
output "key_name" {
  value = alicloud_key_pair.named_test_resource.key_name
}

output "resource_aka" {
  value = "acs:ecs:us-east-1:${data.alicloud_caller_identity.current.account_id}:keypair/${alicloud_key_pair.named_test_resource.id}"
}
