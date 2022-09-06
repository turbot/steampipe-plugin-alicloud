
variable "resource_name" {
  type        = string
  default     = "turbot-test-20200125-create-update"
  description = "Name of the resource used throughout the test."
}

# The Cas Certificate region only support cn-hangzhou, ap-south-1, me-east-1, eu-central-1, ap-northeast-1, ap-southeast-2.
variable "alicloud_region" {
  type        = string
  default     = "ap-south-1"
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

# Create Cetificate Key
resource "tls_private_key" "example" {
  algorithm = "RSA"
}

resource "tls_self_signed_cert" "example" {
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = "turbot.com"
    organization = "Turbot HQ Pvt. Ltd."
  }

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

# Add a new Certificate.
resource "alicloud_cas_certificate" "named_test_resource" {
  name = var.resource_name
  cert = tls_self_signed_cert.example.cert_pem
  key  = tls_private_key.example.private_key_pem
}

output "private_key" {
  value = replace(tls_private_key.example.private_key_pem, "\n", "\\n")
  sensitive = true
}

output "certificate_body" {
  value = replace(tls_self_signed_cert.example.cert_pem, "\n", "\\n")
}

output "certificate_id" {
  value = alicloud_cas_certificate.named_test_resource.id
}

output "resource_name" {
  value = var.resource_name
}
