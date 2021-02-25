package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

// Plugin creates this (alicloud) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-alicloud",
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"alicloud_ecs_disk":                tableAlicloudEcsDisk(ctx),
			"alicloud_ecs_image":               tableAlicloudEcsImage(ctx),
			"alicloud_ecs_instance":            tableAlicloudEcsInstance(ctx),
			"alicloud_ecs_security_group":      tableAlicloudEcsSecurityGroup(ctx),
			"alicloud_ecs_snapshot":            tableAlicloudEcsSnapshot(ctx),
			"alicloud_kms_secret":              tableAlicloudKmsSecret(ctx),
			"alicloud_oss_bucket":              tableAlicloudOssBucket(ctx),
			"alicloud_ram_access_key":          tableAlicloudRAMAccessKey(ctx),
			"alicloud_ram_group":               tableAlicloudRAMGroup(ctx),
			"alicloud_ram_password_policy":     tableAlicloudRamPasswordPolicy(ctx),
			"alicloud_ram_role":                tableAlicloudRAMRole(ctx),
			"alicloud_ram_security_preference": tableAlicloudRAMSecurityPreference(ctx),
			"alicloud_ram_user":                tableAlicloudRAMUser(ctx),
			"alicloud_vpc":                     tableAlicloudVpc(ctx),
			"alicloud_vpc_vswitch":             tableAlicloudVpcVSwitch(ctx),
		},
	}
	return p
}
