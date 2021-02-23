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
			"alicloud_ecs_disk":                tableAlicloudEcsDisk(ctx),
			"alicloud_ecs_security_group":      tableAlicloudEcsSecurityGroup(ctx),
			"alicloud_oss_bucket":              tableAlicloudOssBucket(ctx),
			"alicloud_ram_access_key":          tableAlicloudRamAccessKey(ctx),
			"alicloud_ram_group":               tableAlicloudRamGroup(ctx),
			"alicloud_ram_password_policy":     tableAlicloudRamPasswordPolicy(ctx),
			"alicloud_ram_role":                tableAlicloudRamRole(ctx),
			"alicloud_ram_security_preference": tableAlicloudRamSecurityPreference(ctx),
			"alicloud_ram_user":                tableAlicloudRamUser(ctx),
			"alicloud_vpc":                     tableAlicloudVpc(ctx),
			"alicloud_vpc_vswitch":             tableAlicloudVpcVSwitch(ctx),
		},
	}
	return p
}
