package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudAlidnsDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_alidns_domain",
		Description: "Alicloud Alidns Domain",
		List: &plugin.ListConfig{
			Hydrate: listAlidnsDomains,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "group_id", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    listAlidnsDomains,
			KeyColumns: plugin.SingleColumn("domain_name"),
		},
		GetMatrixItemFunc: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "domain_name",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain_id",
				Description: "The unique identifier of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "group_id",
				Description: "The ID of the group the domain belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "group_name",
				Description: "The name of the group the domain belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remark",
				Description: "Remarks for the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "The creation time of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_timestamp",
				Description: "The timestamp when the domain was created.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "record_count",
				Description: "The number of DNS records associated with the domain.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "instance_id",
				Description: "The instance ID of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ali_domain",
				Description: "Indicates whether the domain is registered with Alibaba Cloud.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "resource_group_id",
				Description: "The resource group ID of the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_end_time",
				Description: "The end time of the instance associated with the domain.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_expired",
				Description: "Indicates whether the instance associated with the domain has expired.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "version_name",
				Description: "The name of the domain version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "version_code",
				Description: "The code of the domain version.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "puny_code",
				Description: "The Punycode representation of the domain name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "registrant_email",
				Description: "The email address of the domain registrant.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "starmark",
				Description: "Indicates whether the domain is marked as a star domain.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "min_ttl",
				Description: "The minimum TTL (Time-To-Live) value for DNS records in the domain.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getAlidnsDomain,
			},
			{
				Name:        "line_type",
				Description: "The type of DNS record line, indicating the routing rules for the domain.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getAlidnsDomain,
			},
			{
				Name:        "in_clean",
				Description: "Indicates whether the domain is in a cleaning state.",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getAlidnsDomain,
			},
			{
				Name:        "dns_servers",
				Description: "The DNS servers associated with the domain.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "record_lines",
				Description: "The DNS record lines associated with the domain.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAlidnsDomain,
			},
			{
				Name:        "available_ttls",
				Description: "The list of available TTL (Time-To-Live) values for DNS records in the domain.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAlidnsDomain,
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Tags.Tag").Transform(dnsTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("DomainName").Transform(domainToAka),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DomainName"),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getAlidnsRegion,
				Transform:   transform.FromValue(),
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

func listAlidnsDomains(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// region := d.EqualsQualString(matrixKeyRegion)

	// Create service connection
	client, err := AliDNSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_alidns_domain.listAlidnsDomains", "connection_error", err)
		return nil, err
	}

	request := alidns.CreateDescribeDomainsRequest()
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)
	if d.EqualsQualString("group_id") != "" {
		request.GroupId = d.EqualsQualString("group_id")
	}

	count := 0
	for {
		response, err := client.DescribeDomains(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_alidns_domain.listAlidnsDomains", "query_error", err, "request", request)
			return nil, err
		}

		for _, domain := range response.Domains.Domain {
			d.StreamListItem(ctx, domain)
			count++
		}

		if count >= int(response.TotalCount) {
			break
		}

		pgNumber, err := request.PageNumber.GetValue()
		if err != nil {
			return nil, err
		}

		request.PageNumber = requests.NewInteger(pgNumber + 1)
	}

	return nil, nil
}

//// GET FUNCTION

func getAlidnsDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	domainName := d.EqualsQualString("domain_name")

	if domainName == "" {
		domain := h.Item.(alidns.DomainInDescribeDomains)
		domainName = domain.DomainName
	}

	// Create service connection
	client, err := AliDNSService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_alidns_domain.getAlidnsDomain", "connection_error", err)
		return nil, err
	}

	request := alidns.CreateDescribeDomainInfoRequest()
	request.DomainName = domainName

	response, err := client.DescribeDomainInfo(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_alidns_domain.getAlidnsDomain", "query_error", err, "request", request)
		return nil, err
	}

	return response, nil
}

//// HELPER FUNCTIONS

func domainToAka(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.Value.(string)
	region := d.HydrateItem.(map[string]interface{})["region"].(string)
	accountID := d.HydrateItem.(map[string]interface{})["account_id"].(string)

	aka := []string{"acs:alidns:" + region + ":" + accountID + ":domain/" + data}
	return aka, nil
}

func getAlidnsRegion(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	return region, nil
}
