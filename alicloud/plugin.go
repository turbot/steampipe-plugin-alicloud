package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-alicloud",
		DefaultTransform: transform.FromCamel().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"alicloud_ram_group": tableAlicloudRamGroup(ctx),
			"alicloud_ram_role":  tableAlicloudRamRole(ctx),
			"alicloud_ram_user":  tableAlicloudRamUser(ctx),
		},
	}
	return p
}
