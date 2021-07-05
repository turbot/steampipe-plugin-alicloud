package alicloud

import (
	"context"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/sethvargo/go-retry"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

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
			KeyColumns:        plugin.AllColumns([]string{"name", "policy_type"}),
			ShouldIgnoreError: isNotFoundError([]string{"InvalidParameter.PolicyType", "EntityNotExist.Policy", "MissingParameter"}),
			Hydrate:           getRAMPolicy,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the policy.",
				Transform:   transform.FromField("PolicyName", "Policy.PolicyName"),
			},

			{
				Name:        "policy_type",
				Type:        proto.ColumnType_STRING,
				Description: "The type of the policy. Valid values: System and Custom.",
				Transform:   transform.FromField("PolicyType", "Policy.PolicyType"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The policy description",
				Transform:   transform.FromField("Description", "Policy.Description"),
			},
			{
				Name:        "default_version",
				Type:        proto.ColumnType_STRING,
				Description: "Deafult version of the policy",
				Transform:   transform.FromField("DefaultVersion", "Policy.DefaultVersion"),
			},
			{
				Name:        "create_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Policy creation date",
				Transform:   transform.FromField("CreateDate", "Policy.CreateDate"),
			},
			{
				Name:        "update_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Last time when policy got updated ",
				Transform:   transform.FromField("UpdateDate", "Policy.UpdateDate"),
			},
			{
				Name:        "attachment_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of references to the policy.",
				Transform:   transform.FromField("AttachmentCount", "Policy.AttachmentCount"),
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
				Description: "Contains the details about the policy.",
				Hydrate:     getRAMPolicy,
				Transform:   transform.FromField("DefaultPolicyVersion.PolicyDocument"),
			},
			{
				Name:        "policy_document_std",
				Type:        proto.ColumnType_JSON,
				Description: "Contains the policy document in a canonical form for easier searching.",
				Hydrate:     getRAMPolicy,
				Transform:   transform.FromField("DefaultPolicyVersion.PolicyDocument").Transform(policyToCanonical),
			},

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
				Transform:   transform.FromField("PolicyName", "Policy.PolicyName"),
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
				Transform:   transform.FromField("AccountID"),
			},
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
			d.StreamListItem(ctx, policy)
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
		i := h.Item.(ram.PolicyInListPolicies)
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
	var response *ram.GetPolicyResponse

	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(10, b), func(ctx context.Context) error {
		var err error
		response, err = client.GetPolicy(request)
		if err != nil {
			if serverErr, ok := err.(*errors.ServerError); ok {
				if serverErr.ErrorCode() == "Throttling.User" {
					return retry.RetryableError(err)
				}
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if response != nil && len(response.Policy.PolicyName) > 0 {
		return response, nil
	}

	return nil, nil
}

func getPolicyAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getPolicyAkas")
	data := policyName(h.Item)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ram::" + accountID + ":policy/" + data}, nil
}

func policyName(item interface{}) string {
	switch item.(type) {
	case ram.PolicyInListPolicies:
		return item.(ram.PolicyInListPolicies).PolicyName
	case *ram.GetPolicyResponse:
		return item.(*ram.GetPolicyResponse).Policy.PolicyName
	}
	return ""
}
