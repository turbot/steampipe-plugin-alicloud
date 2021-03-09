
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

# Declare the data source
data "alicloud_zones" "named_test_resource" {
  available_instance_type = "ecs.n4.large"
  available_disk_category = "cloud_ssd"
}

# # Create an ECS instance with the first matched zone
# resource "alicloud_instance" "named_test_resource" {
#   availability_zone = data.alicloud_zones.named_test_resource.zones.0.id
# }


output "zone_id" {
  value = data.alicloud_zones.named_test_resource.zones.0.id
}

output "local_name" {
  value = data.alicloud_zones.named_test_resource.zones.0.local_name
}

output "available_instance_types" {
  value = data.alicloud_zones.named_test_resource.zones.0.available_instance_types
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:ram::${data.alicloud_caller_identity.current.account_id}:zone/${var.resource_name}"
}
