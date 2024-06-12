connection "alicloud" {
  plugin = "alicloud"

  # You may connect to one or more regions. If `regions` is not specified,
  # Steampipe will use a single default region using the below resolution
  # order:
  # The `ALIBABACLOUD_REGION_ID`, `ALICLOUD_REGION_ID` or `ALICLOUD_REGION` environment variable
  # regions = ["us-east-1", "ap-south-1"]

  # If no credentials are specified, the plugin will use the Aliyun credentials
  # resolver to get the current credentials in the same manner as the CLI.
  # Alternatively, you may set static credentials with the `access_key`,
  # `secret_key`, and `session_token` arguments, or select a named profile
  # from an Aliyun credential file(`~/.aliyun/config.json`) with the `profile` argument.
  # Additionally, it can be configured via environment variables: ALIBABACLOUD_PROFILE, ALIBABA_CLOUD_PROFILE, or ALICLOUD_PROFILE.
  # profile = "myprofile"

  # If no credentials are specified, the plugin will use the environment variables
  # resolver to get the current credentials.
  # Alternatively, you may set static credentials with the `access_key` and `secret_key` arguments.
  # access_key  	= "LTAI4GBVFakeKey09Kxezv66"
  # secret_key  	= "6iNPvThisIsNotARealSecretk1sZF"

  # List of additional Alicloud error codes to ignore for all queries.
  # By default, common not found error codes are ignored and will still be ignored even if this argument is not set.
  # ignore_error_codes = ["AccessDenied", "Forbidden.Access", "Forbidden.NoPermission"]
}
