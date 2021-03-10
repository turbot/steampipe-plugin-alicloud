
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

data "alicloud_zones" "resource" {
  available_resource_creation = "VSwitch"
}

data "alicloud_instance_types" "resource" {
  availability_zone = data.alicloud_zones.resource.zones[0].id
  cpu_core_count    = 1
  memory_size       = 2
}

data "alicloud_images" "resource" {
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

resource "alicloud_vpc" "named_test_resource" {
  name       = var.resource_name
  cidr_block = "10.1.0.0/21"
}

resource "alicloud_vswitch" "named_test_resource" {
  vpc_id            = alicloud_vpc.named_test_resource.id
  cidr_block        = "10.1.1.0/24"
  availability_zone = data.alicloud_zones.resource.zones[0].id
  name              = var.resource_name
}

resource "alicloud_security_group" "named_test_resource" {
  name        = var.resource_name
  description = "foo"
  vpc_id      = alicloud_vpc.named_test_resource.id
}

resource "alicloud_security_group_rule" "named_test_resource" {
  type              = "ingress"
  ip_protocol       = "tcp"
  nic_type          = "intranet"
  policy            = "accept"
  port_range        = "22/22"
  priority          = 1
  security_group_id = alicloud_security_group.named_test_resource.id
  cidr_ip           = "0.0.0.0/0"
}

resource "alicloud_instance" "named_test_resource" {
  security_groups = [alicloud_security_group.named_test_resource.id]

  vswitch_id = alicloud_vswitch.named_test_resource.id

  instance_charge_type       = "PostPaid"
  instance_type              = data.alicloud_instance_types.resource.instance_types[0].id
  internet_charge_type       = "PayByTraffic"
  internet_max_bandwidth_out = 5

  system_disk_category = "cloud_efficiency"
  image_id             = data.alicloud_images.resource.images[0].id
  instance_name        = var.resource_name
}

resource "alicloud_route_entry" "named_test_resource" {
  name                  = "test"
  route_table_id        = alicloud_vpc.named_test_resource.route_table_id
  destination_cidrblock = "172.11.1.1/32"
  nexthop_type          = "Instance"
  nexthop_id            = alicloud_instance.named_test_resource.id
}

output "name" {
  value = alicloud_route_entry.named_test_resource.name
}

output "route_table_id" {
  value = alicloud_route_entry.named_test_resource.route_table_id
}

output "destination_cidrblock" {
  value = alicloud_route_entry.named_test_resource.destination_cidrblock
}
output "nexthop_type" {
  value = alicloud_route_entry.named_test_resource.nexthop_type
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:vpc:us-east-1:${data.alicloud_caller_identity.current.account_id}:route-entry/${alicloud_route_entry.named_test_resource.id}"
}
