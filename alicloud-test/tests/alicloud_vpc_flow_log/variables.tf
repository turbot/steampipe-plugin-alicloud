
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

resource "alicloud_vpc" "named_test_resource" {
  cidr_block = "192.168.0.0/24"
  name       = var.resource_name
}

resource "alicloud_vpc_flow_log" "named_test_resource" {
  depends_on     = ["alicloud_vpc.named_test_resource"]
  resource_id    = alicloud_vpc.named_test_resource.id
  resource_type  = "VPC"
  traffic_type   = "All"
  log_store_name = var.resource_name
  project_name   = var.resource_name
  flow_log_name  = var.resource_name
  status         = "Active"
}

output "resource_id" {
  value = alicloud_vpc_flow_log.named_test_resource.id
}

output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
}

output "router_id" {
  value = alicloud_vpc.named_test_resource.router_id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:flowlog/${alicloud_vpc_flow_log.named_test_resource.id}"
}
