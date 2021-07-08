
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

resource "alicloud_oss_bucket" "named_test_resource" {
  bucket = var.resource_name
  acl    = "private"

  # Policy
  policy = <<POLICY
  {"Statement":
      [{"Action":
          ["oss:PutObject", "oss:GetObject", "oss:DeleteBucket"],
        "Effect":"Allow",
        "Resource":
            ["acs:oss:*:*:*"]}],
   "Version":"1"}
  POLICY

  # Lifecycle Rules
  lifecycle_rule {
    id      = "rule-days"
    prefix  = "path1/"
    enabled = true

    expiration {
      days = 365
    }
  }

  # Versioning
  versioning {
    status = "Suspended"
  }

  # SSE configuration
  server_side_encryption_rule {
    sse_algorithm = "AES256"
  }

  # Tags
  tags = {
    name = var.resource_name
  }
}

output "bucket_id" {
  value = alicloud_oss_bucket.named_test_resource.id
}

output "account_id" {
  value = data.alicloud_caller_identity.current.account_id
}

output "region" {
  value = var.alicloud_region
}

output "bucket_arn" {
  value = "arn:acs:oss:::${alicloud_oss_bucket.named_test_resource.id}"
}
