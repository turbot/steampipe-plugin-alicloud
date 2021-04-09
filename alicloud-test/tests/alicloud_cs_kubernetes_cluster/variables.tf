
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
  cidr_block = "172.16.0.0/12"
}

// According to the vswitch cidr blocks to launch several vswitches
resource "alicloud_vswitch" "named_test_resource" {
  vpc_id     = alicloud_vpc.named_test_resource.id
  cidr_block = "172.16.0.0/21"
  zone_id    = "us-east-1b"
}

resource "alicloud_cs_kubernetes" "named_test_resource" {
  master_vswitch_ids    = [alicloud_vswitch.named_test_resource.id, alicloud_vswitch.named_test_resource.id, alicloud_vswitch.named_test_resource.id]
  worker_vswitch_ids    = [alicloud_vswitch.named_test_resource.id, alicloud_vswitch.named_test_resource.id, alicloud_vswitch.named_test_resource.id]
  master_instance_types = ["ecs.c5.large", "ecs.c5.large", "ecs.c5.large"]
  worker_instance_types = ["ecs.c5.large", "ecs.c5.large", "ecs.c5.large"]
  worker_number         = "3"
  pod_cidr              = "10.10.0.0/24"
  service_cidr          = "192.168.0.0/24"
  password              = "Test1234!"
  name                  = var.resource_name
}

output "cluster_id" {
  value = alicloud_cs_kubernetes.named_test_resource.id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:cs::${data.alicloud_caller_identity.current.account_id}:container/${var.resource_name}"
}
