package alicloud

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type alicloudConfig struct {
	Regions   []string `cty:"regions"`
	AccessKey *string  `cty:"access_key"`
	SecretKey *string  `cty:"secret_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"regions": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"access_key": {
		Type: schema.TypeString,
	},
	"secret_key": {
		Type: schema.TypeString,
	},
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
