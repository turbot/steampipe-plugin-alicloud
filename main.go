package main

import (
	"github.com/turbot/steampipe-plugin-alicloud/alicloud"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: alicloud.Plugin})
}
