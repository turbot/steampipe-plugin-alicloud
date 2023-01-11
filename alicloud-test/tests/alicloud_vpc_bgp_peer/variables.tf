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

data "alicloud_express_connect_physical_connections" "named_test_resource" {}

resource "alicloud_express_connect_virtual_border_router" "named_test_resource" {
  local_gateway_ip           = "10.0.0.1"
  peer_gateway_ip            = "10.0.0.2"
  peering_subnet_mask        = "255.255.255.252"
  physical_connection_id     = data.alicloud_express_connect_physical_connections.named_test_resource.connections.0.id
  virtual_border_router_name = "example_value"
  vlan_id                    = 120
  min_rx_interval            = 1000
  min_tx_interval            = 1000
  detect_multiplier          = 10
}

resource "alicloud_vpc_bgp_group" "named_test_resource" {
  auth_key       = "YourPassword+12345678"
  bgp_group_name = var.resource_name
  description    = "example_value"
  local_asn      = 64512
  peer_asn       = 1111
  router_id      = alicloud_express_connect_virtual_border_router.named_test_resource.id
}

resource "alicloud_vpc_bgp_peer" "named_test_resource" {
  bfd_multi_hop   = "10"
  bgp_group_id    = alicloud_vpc_bgp_group.named_test_resource.id
  enable_bfd      = true
  ip_version      = "IPV4"
  peer_ip_address = "1.1.1.1"
}

output "resource_id" {
  value = alicloud_vpc_bgp_peer.named_test_resource.id
}

output "resource_status" {
  value = alicloud_vpc_bgp_peer.named_test_resource.status
}

output "router_id" {
  value = alicloud_express_connect_virtual_border_router.named_test_resource.id
}

output "bgp_group_id" {
  value = alicloud_vpc_bgp_group.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "region_id" {
  value = var.alicloud_region
}

output "resource_name" {
  value = var.resource_name
}
