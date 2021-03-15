package alicloud

import (
	"context"
	"unsafe"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

type ramPolicyInfo = struct {
	ram.PolicyInListPolicies
	DefaultPolicyVersion ram.DefaultPolicyVersion
}

//// TABLE DEFINITION

func tableAlicloudRamPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "alicloud_ram_policy",
		Description:      "Alibaba Cloud RAM Policy",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: listRAMPolicy,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "policy_type"}),
			Hydrate:    getRAMPolicy,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PolicyName"),
				Description: "The name of the policy.",
			},

			{
				Name:        "policy_type",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PolicyType"),
				Description: "The type of the policy. Valid values: System and Custom.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The policy description",
			},
			{
				Name:        "default_version",
				Type:        proto.ColumnType_STRING,
				Description: "Deafult version of the policy",
			},
			{
				Name:        "create_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Policy creation date",
			},
			{
				Name:        "update_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Last time when policy got updated ",
			},
			{
				Name:        "attachment_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of references to the policy.",
			},
			{
				Name:        "version_id",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRAMPolicy,
				Transform:   transform.FromField("DefaultPolicyVersion.VersionId"),
				Description: "The ID of the default policy version.",
			},
			{
				Name:        "is_default_version",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getRAMPolicy,
				Transform:   transform.FromField("DefaultPolicyVersion.IsDefaultVersion"),
				Description: "An attribute in the DefaultPolicyVersion parameter. The value of the IsDefaultVersion parameter is true.",
			},
			{
				Name:        "policy_document",
				Type:        proto.ColumnType_JSON,
				Description: "The script of the default policy version.",
				Hydrate:     getRAMPolicy,
				Transform:   transform.FromValue(),
				//Transform:   transform.FromField("DefaultPolicyVersion.PolicyDocument"),
			},
			// {
			// 	Name:        "policy_document_std",
			// 	Type:        proto.ColumnType_JSON,
			// 	Description: "The policy document",
			// 	Transform:   transform.FromField("DefaultPolicyVersion.PolicyDocument").Transform(policyToCanonical),
			// 	Hydrate:     getRAMPolicy,
			// },

			// steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getPolicyAkas,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PolicyName"),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromConstant("global")},
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

func listRAMPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listRamPolicy", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListPoliciesRequest()
	request.Scheme = "https"
	for {
		response, err := client.ListPolicies(request)
		if err != nil {
			plugin.Logger(ctx).Error("listRamPolicy", "query_error", err, "request", request)
			return nil, err
		}
		for _, policy := range response.Policies.Policy {
			plugin.Logger(ctx).Warn("alicloud_ram.listRamPolicy", "item", policy)
			d.StreamListItem(ctx, ramPolicyInfo{policy, ram.DefaultPolicyVersion{}})
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

func getRAMPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMPolicy")

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_policy.getRAMPolicy", "connection_error", err)
		return nil, err
	}

	var name, policyType string
	if h.Item != nil {
		i := h.Item.(ramPolicyInfo)
		name = i.PolicyName
		policyType = i.PolicyType
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
		policyType = d.KeyColumnQuals["policy_type"].GetStringValue()
	}

	request := ram.CreateGetPolicyRequest()
	request.Scheme = "https"
	request.PolicyName = name
	request.PolicyType = policyType

	response, err := client.GetPolicy(request)
	if err != nil {
		plugin.Logger(ctx).Error("GetPolicy", "query_error", err, "request", request)
		return nil, err
	}

	if response != nil && len(response.Policy.PolicyName) > 0 {
		policyData := *(*ram.PolicyInListPolicies)(unsafe.Pointer(&response.Policy))
		return ramPolicyInfo{policyData, response.DefaultPolicyVersion}, nil
	}

	return nil, nil
}

func getPolicyAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getPolicyAkas")
	data := h.Item.(ramPolicyInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ram::" + accountID + ":policy/" + data.PolicyName}, nil
}
