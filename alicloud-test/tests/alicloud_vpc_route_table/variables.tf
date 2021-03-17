
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

# Create a VPC
resource "alicloud_vpc" "named_test_resource" {
  cidr_block = "10.1.0.0/21"
  name       = var.resource_name
}

# Create a VSwitch
resource "alicloud_vswitch" "named_test_resource" {
  vpc_id            = alicloud_vpc.named_test_resource.id
  cidr_block        = "10.1.1.0/24"
  availability_zone = "us-east-1b"
  name              = var.resource_name
}

# Create a VPC Route Table
resource "alicloud_route_table" "named_test_resource" {
  vpc_id      = alicloud_vpc.named_test_resource.id
  name        = var.resource_name
  description = "This is a test route table to validate the table outcome."
  tags = {
    Name = var.resource_name
  }
}

resource "null_resource" "delay" {
  provisioner "local-exec" {
    command = "sleep 120"
  }
  triggers = {
    "before" = "${alicloud_route_table.named_test_resource.id}"
  }
}

# Create Attachement of VPC Route Table with VSwitch
resource "alicloud_route_table_attachment" "named_test_resource" {
  vswitch_id     = alicloud_vswitch.named_test_resource.id
  route_table_id = alicloud_route_table.named_test_resource.id
}

output "resource_id" {
  value = alicloud_route_table.named_test_resource.id
}

output "vpc_id" {
  value = alicloud_vpc.named_test_resource.id
}

output "router_id" {
  value = alicloud_vpc.named_test_resource.router_id
}

output "vswitch_id" {
  value = alicloud_vswitch.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:route-table/${alicloud_route_table.named_test_resource.id}"
}
