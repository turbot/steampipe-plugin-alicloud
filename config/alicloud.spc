connection "alicloud" {
  plugin = "alicloud"

  # You may connect to one or more regions. If `regions` is not specified,
  # Steampipe will use a single default region using the below resolution
  # order:
  #  1. The `ALIBABACLOUD_REGION_ID`, `ALICLOUD_REGION_ID` or `ALICLOUD_REGION` environment variable
  #regions     = ["us-east-1", "ap-south-1"]

  # If no credentials are specified, the plugin will use the environment variables
  # resolver to get the current credentials.
  #  Alternatively, you may set static credentials with the `access_key` and `secret_key` arguments.

}
