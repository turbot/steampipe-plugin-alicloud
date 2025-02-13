package alicloud

import (
	"context"
	"slices"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/sethvargo/go-retry"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudRamPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "alicloud_ram_policy",
		Description:      "Alibaba Cloud RAM Policy",
		DefaultTransform: transform.FromCamel(),
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"policy_name", "policy_type"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"InvalidParameter.PolicyType", "EntityNotExist.Policy", "MissingParameter"}),
			},
			Hydrate: getRAMPolicy,
		},
		List: &plugin.ListConfig{
			Hydrate: listRAMPolicies,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "policy_type", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "policy_name",
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
				Name:        "create_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Policy creation date",
				Transform:   transform.FromField("CreateDate", "Policy.CreateDate"),
			},
			{
				Name:        "attachment_count",
				Type:        proto.ColumnType_INT,
				Description: "The number of references to the policy.",
				Transform:   transform.FromField("AttachmentCount", "Policy.AttachmentCount"),
			},
			{
				Name:        "default_version",
				Type:        proto.ColumnType_STRING,
				Description: "Deafult version of the policy",
				Transform:   transform.FromField("DefaultVersion", "Policy.DefaultVersion"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The policy description",
				Transform:   transform.FromField("Description", "Policy.Description"),
			},
			{
				Name:        "update_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Last time when policy got updated ",
				Transform:   transform.FromField("UpdateDate", "Policy.UpdateDate"),
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

			// Steampipe standard columns
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

			// Alicloud standard columns
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

func listRAMPolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listRAMPolicies", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListPoliciesRequest()
	request.Scheme = "https"
	request.MaxItems = requests.NewInteger(1000)

	if value, ok := GetStringQualValue(d.Quals, "policy_type"); ok {
		request.PolicyType = *value

		// select policy_name, policy_type from alicloud.alicloud_ram_policy where policy_type = 'Custom1'
		// Message: PolicyType must be Custom/System but meet:Custom1
		if !slices.Contains([]string{"Custom", "System"}, *value) {
			return nil, nil
		}
	}

	// If the request no of items is less than the paging max limit
	// update limit to requested no of results.
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		pageSize, err := request.MaxItems.GetValue64()
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ecs_instance.listEcsInstance", "page_size_error", err)
			return nil, err
		}
		if *limit < pageSize {
			request.MaxItems = requests.NewInteger(int(*limit))
		}
	}

	for {
		// https://partners-intl.aliyun.com/help/doc-detail/28719.htm?spm=a2c63.p38356.b99.249.37d17aa2AscMLc
		response, err := client.ListPolicies(request)
		if err != nil {
			plugin.Logger(ctx).Error("listRAMPolicies", "query_error", err, "request", request)
			return nil, err
		}
		for _, policy := range response.Policies.Policy {
			d.StreamListItem(ctx, policy)
			// This will return zero if context has been cancelled (i.e due to manual cancellation) or
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
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
		i := h.Item.(ram.Policy)
		name = i.PolicyName
		policyType = i.PolicyType
	} else {
		name = d.EqualsQuals["policy_name"].GetStringValue()
		policyType = d.EqualsQuals["policy_type"].GetStringValue()
	}

	request := ram.CreateGetPolicyRequest()
	request.Scheme = "https"
	request.PolicyName = name
	request.PolicyType = policyType
	var response *ram.GetPolicyResponse

	b := retry.NewFibonacci(100 * time.Millisecond)

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
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ram::" + accountID + ":policy/" + data}, nil
}

func policyName(item interface{}) string {
	switch item := item.(type) {
	case ram.Policy:
		return item.PolicyName
	case *ram.GetPolicyResponse:
		return item.Policy.PolicyName
	}
	return ""
}
