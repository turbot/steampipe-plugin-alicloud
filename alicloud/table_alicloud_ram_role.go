package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type roleInfo = struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
	MaxSessionDuration       int64
}

//// TABLE DEFINITION

func tableAlicloudRAMRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_role",
		Description: "Resource Access Management roles who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listRAMRoles,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			ShouldIgnoreError: isNotFoundError([]string{"EntityNotExist.Role"}),
			Hydrate:           getRAMRole,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the RAM role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RoleName"),
			},
			{
				Name:        "arn",
				Description: "The Alibaba Cloud Resource Name (ARN) of the RAM role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "role_id",
				Description: "The ID of the RAM role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "description",
				Description: "The description of the RAM role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "max_session_duration",
				Description: "The maximum session duration of the RAM role.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "create_date",
				Description: "The time when the RAM role was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "update_date",
				Description: "The time when the RAM role was modified.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "assume_role_policy_document",
				Description: "The content of the policy that specifies one or more entities entrusted to assume the RAM role.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMRole,
				Transform:   transform.FromField("AssumeRolePolicyDocument").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "assume_role_policy_document_std",
				Description: "The standard content of the policy that specifies one or more entities entrusted to assume the RAM role.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMRole,
				Transform:   transform.FromField("AssumeRolePolicyDocument").Transform(policyToCanonical),
			},
			{
				Name:        "attached_policy",
				Description: "A list of policies attached to a RAM role.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMRolePolicies,
				Transform:   transform.FromField("Policies.Policy"),
			},

			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Arn").Transform(ensureStringArray),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RoleName"),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromConstant("global"),
			},
			{
				Name:        "account_id",
				Description: ColumnDescriptionAccount,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID")},
		},
	}
}

//// LIST FUNCTION

func listRAMRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_role.listRAMRoles", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListRolesRequest()
	request.Scheme = "https"

	for {
		response, err := client.ListRoles(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ram_role.listRAMRoles", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Roles.Role {
			plugin.Logger(ctx).Warn("listRAMRoles", "item", i)
			d.StreamListItem(ctx, roleInfo{i.RoleId, i.RoleName, i.Arn, i.Description, "", i.CreateDate, i.UpdateDate, i.MaxSessionDuration})
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getRAMRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMRole")

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_role.getRAMRole", "connection_error", err)
		return nil, err
	}

	var name string
	if h.Item != nil {
		i := h.Item.(roleInfo)
		name = i.RoleName
	} else {
		name = d.EqualsQuals["name"].GetStringValue()
	}

	request := ram.CreateGetRoleRequest()
	request.Scheme = "https"
	request.RoleName = name

	response, err := client.GetRole(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_role.getRAMRole", "query_error", err, "request", request)
		return nil, err
	}

	data := response.Role
	return roleInfo{data.RoleId, data.RoleName, data.Arn, data.Description, data.AssumeRolePolicyDocument, data.CreateDate, data.UpdateDate, data.MaxSessionDuration}, nil
}

func getRAMRolePolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMRolePolicies")
	data := h.Item.(roleInfo)

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMRolePolicies", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListPoliciesForRoleRequest()
	request.Scheme = "https"
	request.RoleName = data.RoleName

	response, err := client.ListPoliciesForRole(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMRolePolicies", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}
