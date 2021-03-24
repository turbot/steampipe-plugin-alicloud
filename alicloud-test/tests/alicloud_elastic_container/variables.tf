
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
  count      = var.vpc_id == "" ? 1 : 0
  cidr_block = var.vpc_cidr
}

// According to the vswitch cidr blocks to launch several vswitches
resource "alicloud_vswitch" "named_test_resource" {
  count             = length(var.vswitch_ids) > 0 ? 0 : length(var.vswitch_cidrs)
  vpc_id            = var.vpc_id == "" ? join("", alicloud_vpc.named_test_resource.*.id) : var.vpc_id
  cidr_block        = element(var.vswitch_cidrs, count.index)
  availability_zone = element(var.availability_zone, count.index)
}

resource "alicloud_cs_kubernetes" "named_test_resource" {
  count                 = 1
  master_vswitch_ids    = length(var.vswitch_ids) > 0 ? split(",", join(",", var.vswitch_ids)) : length(var.vswitch_cidrs) < 1 ? [] : split(",", join(",", alicloud_vswitch.named_test_resource.*.id))
  worker_vswitch_ids    = length(var.vswitch_ids) > 0 ? split(",", join(",", var.vswitch_ids)) : length(var.vswitch_cidrs) < 1 ? [] : split(",", join(",", alicloud_vswitch.named_test_resource.*.id))
  master_instance_types = var.master_instance_types
  worker_instance_types = var.worker_instance_types
  worker_number         = var.worker_number
  node_cidr_mask        = var.node_cidr_mask
  enable_ssh            = var.enable_ssh
  install_cloud_monitor = var.install_cloud_monitor
  cpu_policy            = var.cpu_policy
  proxy_mode            = var.proxy_mode
  password              = var.password
  pod_cidr              = var.pod_cidr
  service_cidr          = var.service_cidr
  name                  = var.resource_name

  dynamic "addons" {
    for_each = var.cluster_addons
    content {
      name     = lookup(addons.value, "name", var.cluster_addons)
      config   = lookup(addons.value, "config", var.cluster_addons)
      disabled = lookup(addons.value, "disabled", var.cluster_addons)
    }
  }
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
