
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

# Create a new RAM user.
resource "alicloud_ram_user" "named_test_resource" {
  name         = var.resource_name
  display_name = var.resource_name
  mobile       = "86-18600008888"
  email        = "${var.resource_name}.test@yoyo.com"
  comments     = "This is a test user."
  force        = true
}

# Create a new RAM group.
resource "alicloud_ram_group" "named_test_resource" {
  name     = var.resource_name
  comments = "This is a group comments."
  force    = true
}

# Create a RAM Group membership.
resource "alicloud_ram_group_membership" "named_test_resource" {
  group_name = alicloud_ram_group.named_test_resource.name

  user_names = [
    alicloud_ram_user.named_test_resource.name,
  ]
}

# Create a RAM Policy.
resource "alicloud_ram_policy" "named_test_resource" {
  policy_name     = var.resource_name
  policy_document = <<EOF
  {
    "Statement": [
      {
        "Action": [
          "oss:ListObjects",
          "oss:GetObject"
        ],
        "Effect": "Allow",
        "Resource": [
          "acs:oss:*:*:mybucket",
          "acs:oss:*:*:mybucket/*"
        ]
      }
    ],
      "Version": "1"
  }
  EOF
  description     = "This is a policy test"
  force           = true
}

# Create a RAM User Policy attachment.
resource "alicloud_ram_user_policy_attachment" "attach" {
  policy_name = alicloud_ram_policy.named_test_resource.name
  user_name   = alicloud_ram_user.named_test_resource.name
  policy_type = alicloud_ram_policy.named_test_resource.type
}

output "user_id" {
  value = alicloud_ram_user.named_test_resource.id
}

output "membership_id" {
  value = alicloud_ram_group_membership.named_test_resource.id
}

output "policy_attachemnt_id" {
  value = alicloud_ram_policy.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "resource_name" {
  value = var.resource_name
}

output "resource_aka" {
  value = "acs:ram::${data.alicloud_caller_identity.current.account_id}:user/${var.resource_name}"
}
