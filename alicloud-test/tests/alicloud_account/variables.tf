variable "alicloud_region" {
  type        = string
  default     = "us-east-1"
  description = "Alicloud region used for the test."
}

provider "alicloud" {
  region = var.alicloud_region
}

data "alicloud_account" "current" {
}

output "current_account_id" {
  value = "${data.alicloud_account.current.id}"
}
