variable "resource_name" {
  type        = string
  default     = "tf-testcontainer"
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

// If there is not specifying vpc_id, the module will launch a new vpc
resource "alicloud_vpc" "named_test_resource" {
  vpc_name   = var.resource_name
  cidr_block = "10.1.0.0/21"
}

// According to the vswitch cidr blocks to launch several vswitches
resource "alicloud_vswitch" "named_test_resource" {
  vpc_id     = alicloud_vpc.named_test_resource.id
  cidr_block = "10.1.1.0/24"
  zone_id    = "us-east-1b"
}

resource "alicloud_cs_managed_kubernetes" "named_test_resource" {
  name                         = var.resource_name
  count                        = 1
  cluster_spec                 = "ack.standard"
  is_enterprise_security_group = true
  worker_number                = 2
  password                     = "Hello1234"
  pod_cidr                     = "172.20.0.0/16"
  service_cidr                 = "172.21.0.0/20"
  worker_vswitch_ids           = [alicloud_vswitch.named_test_resource.id]
  instance_types               = ["ecs.c5.large"]
}

output "cluster_id" {
  value = alicloud_cs_managed_kubernetes.named_test_resource[0].id
}

output "instance_id" {
  value = alicloud_cs_managed_kubernetes.named_test_resource[0].worker_nodes[0].id
}

output "host_name" {
  value = alicloud_cs_managed_kubernetes.named_test_resource[0].worker_nodes[0].name
}

output "ip_address" {
  value = alicloud_cs_managed_kubernetes.named_test_resource[0].worker_nodes[0].private_ip
}

output "resource_name" {
  value = var.resource_name
}

output "region" {
  value = var.alicloud_region
}

output "resource_aka" {
  value = "arn:acs:cms:${var.alicloud_region}:${data.alicloud_caller_identity.current.account_id}:host/${alicloud_cs_managed_kubernetes.named_test_resource[0].worker_nodes[0].name}"
}
