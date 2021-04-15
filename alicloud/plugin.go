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
			"alicloud_action_trail":                tableAlicloudActionTrail(ctx),
			"alicloud_cas_certificate":             tableAlicloudUserCertificate(ctx),
			"alicloud_ecs_auto_provisioning_group": tableAlicloudEcsAutoProvisioningGroup(ctx),
			"alicloud_ecs_autoscaling_group":       tableAlicloudEcsAutoscalingGroup(ctx),
			"alicloud_ecs_disk":                    tableAlicloudEcsDisk(ctx),
			"alicloud_ecs_image":                   tableAlicloudEcsImage(ctx),
			"alicloud_ecs_instance":                tableAlicloudEcsInstance(ctx),
			"alicloud_ecs_key_pair":                tableAlicloudEcskeyPair(ctx),
			"alicloud_ecs_launch_template":         tableAlicloudEcsLaunchTemplate(ctx),
			"alicloud_ecs_network_interface":       tableAlicloudEcsEni(ctx),
			"alicloud_ecs_region":                  tableAlicloudEcsRegion(ctx),
			"alicloud_ecs_security_group":          tableAlicloudEcsSecurityGroup(ctx),
			"alicloud_ecs_snapshot":                tableAlicloudEcsSnapshot(ctx),
			"alicloud_ecs_zone":                    tableAlicloudEcsZone(ctx),
			"alicloud_kms_key":                     tableAlicloudKmsKey(ctx),
			"alicloud_kms_secret":                  tableAlicloudKmsSecret(ctx),
			"alicloud_oss_bucket":                  tableAlicloudOssBucket(ctx),
			"alicloud_ram_access_key":              tableAlicloudRAMAccessKey(ctx),
			"alicloud_ram_credential_report":       tableAlicloudRAMCredentialReport(ctx),
			"alicloud_ram_group":                   tableAlicloudRAMGroup(ctx),
			"alicloud_ram_password_policy":         tableAlicloudRamPasswordPolicy(ctx),
			"alicloud_ram_role":                    tableAlicloudRAMRole(ctx),
			"alicloud_ram_security_preference":     tableAlicloudRAMSecurityPreference(ctx),
			"alicloud_ram_user":                    tableAlicloudRAMUser(ctx),
			"alicloud_rds_instance":                tableAlicloudRdsInstance(ctx),
			"alicloud_vpc":                         tableAlicloudVpc(ctx),
			"alicloud_vpc_eip":                     tableAlicloudVpcEip(ctx),
			"alicloud_vpc_nat_gateway":             tableAlicloudVpcNatGateway(ctx),
			"alicloud_vpc_network_acl":             tableAlicloudVpcNetworkACL(ctx),
			"alicloud_vpc_route_entry":             tableAlicloudVpcRouteEntry(ctx),
			"alicloud_vpc_route_table":             tableAlicloudVpcRouteTable(ctx),
			"alicloud_vpc_ssl_vpn_client_cert":     tableAlicloudVpcSslVpnClientCert(ctx),
			"alicloud_vpc_ssl_vpn_server":          tableAlicloudVpcSslVpnServer(ctx),
			"alicloud_vpc_vpn_connection":          tableAlicloudVpcVpnConnection(ctx),
			"alicloud_vpc_vpn_customer_gateway":    tableAlicloudVpcVpnCustomerGateway(ctx),
			"alicloud_vpc_vpn_gateway":             tableAlicloudVpcVpnGateway(ctx),
			"alicloud_vpc_vswitch":                 tableAlicloudVpcVSwitch(ctx),
		},
	}
	return p
}
