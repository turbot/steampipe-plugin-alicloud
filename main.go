package main

import (
	"github.com/turbot/steampipe-plugin-alicloud/alicloud"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: alicloud.Plugin})
}
