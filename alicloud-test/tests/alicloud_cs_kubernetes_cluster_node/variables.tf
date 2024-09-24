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
  name         = var.resource_name
  cluster_spec = "ack.pro.small"
  # version can not be defined in variables.tf.
  # version            = "1.26.3-aliyun.1"
  worker_vswitch_ids = [alicloud_vswitch.named_test_resource.id]
  new_nat_gateway    = true
  proxy_mode         = "ipvs"
  service_cidr       = "192.168.0.0/16"

  addons {
    name = "terway-eniip"
  }
  addons {
    name = "csi-plugin"
  }
  addons {
    name = "csi-provisioner"
  }
  addons {
    name = "logtail-ds"
    config = jsonencode({
      IngressDashboardEnabled = "true"
    })
  }
  addons {
    name = "nginx-ingress-controller"
    config = jsonencode({
      IngressSlbNetworkType = "internet"
    })
    # to disable install nginx-ingress-controller automatically
    # disabled = true
  }
  addons {
    name = "arms-prometheus"
  }
  addons {
    name = "ack-node-problem-detector"
    config = jsonencode({
      # sls_project_name = "your-sls-project"
    })
  }
}

output "cluster_id" {
  value = alicloud_cs_managed_kubernetes.named_test_resource.id
}

output "instance_id" {
  value = alicloud_cs_managed_kubernetes.named_test_resource.worker_nodes[0].id
}

output "ip_address" {
  value = alicloud_cs_managed_kubernetes.named_test_resource.worker_nodes[0].private_ip
}

output "resource_name" {
  value = var.resource_name
}

output "region" {
  value = var.alicloud_region
}

output "resource_aka" {
  value = "acs:cs:${var.alicloud_region}:${data.alicloud_caller_identity.current.account_id}:node/${var.resource_name}"
}
