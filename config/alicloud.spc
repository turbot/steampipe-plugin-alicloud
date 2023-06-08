connection "alicloud" {
  plugin = "alicloud"

  # You may connect to one or more regions. If `regions` is not specified,
  # Steampipe will use a single default region using the below resolution
  # order:
  #  1. The `ALIBABACLOUD_REGION_ID`, `ALICLOUD_REGION_ID` or `ALICLOUD_REGION` environment variable
  #regions     = ["us-east-1", "ap-south-1"]

  # If no credentials are specified, the plugin will use the environment variables
  # resolver to get the current credentials.
  # Alternatively, you may set static credentials with the `access_key` and `secret_key` arguments.
  #access_key  	= "LTAI4GBVFakeKey09Kxezv66"

  # List of additional Alicloud error codes to ignore for all queries.
  # By default, common not found error codes are ignored and will still be ignored even if this argument is not set.
  # ignore_error_codes = ["Forbidden.Access", "Forbidden.NoPermission"]

}
