package alicloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type alicloudConfig struct {
	Regions          []string `hcl:"regions"`
	AccessKey        *string  `hcl:"access_key"`
	SecretKey        *string  `hcl:"secret_key"`
	IgnoreErrorCodes []string `hcl:"ignore_error_codes,optional"`
}

func ConfigInstance() interface{} {
	return &alicloudConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) alicloudConfig {
	if connection == nil || connection.Config == nil {
		return alicloudConfig{}
	}
	config, _ := connection.Config.(alicloudConfig)
	return config
}
