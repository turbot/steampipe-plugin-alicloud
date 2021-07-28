package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudCmsMonitorHost(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_cms_monitor_host",
		Description: "Alicloud Cloud Monitor Host",
		List: &plugin.ListConfig{
			Hydrate: listCmsMonitorHosts,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"host_name", "instance_id"}),
			Hydrate:    getCmsMonitorHost,
		},
		Columns: []*plugin.Column{
			{
				Name:        "host_name",
				Description: "The name of the host.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_type_family",
				Description: "The type of the ECS instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "agent_version",
				Description: "The version of the Cloud Monitor agent.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_aliyun_host",
				Description: "Indicates whether the host is provided by Alibaba Cloud.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "eip_id",
				Description: "The ID of the EIP.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "eip_address",
				Description: "The elastic IP address (EIP) of the host.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ali_uid",
				Description: "The ID of the Alibaba Cloud account.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "ip_group",
				Description: "The IP address of the host.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "nat_ip",
				Description: "The IP address of the Network Address Translation (NAT) gateway.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "network_type",
				Description: "The type of the network.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "operating_system",
				Description: "The operating system.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "serial_number",
				Type:        proto.ColumnType_STRING,
				Description: "The serial number of the host. A host that is not provided by Alibaba Cloud has a serial number instead of an instance ID.",
			},
			{
				Name:        "monitoring_agent_status",
				Description: "The status of the Cloud Monitor agent.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCmsMonitoringAgentStatus,
				Transform:   transform.FromValue(),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("HostName"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCmsMonitoringHostAka,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
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

func listCmsMonitorHosts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := CmsService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listCmsMonitorHosts", "connection_error", err)
		return nil, err
	}
	request := cms.CreateDescribeMonitoringAgentHostsRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeMonitoringAgentHosts(request)
		if err != nil {
			plugin.Logger(ctx).Error("listCmsMonitorHosts", "query_error", err, "request", request)
			return nil, err
		}
		for _, host := range response.Hosts.Host {
			plugin.Logger(ctx).Warn("listCmsMonitorHosts", "item", host)
			d.StreamListItem(ctx, host)
			count++
		}
		if count >= response.Total {
			break
		}
		request.PageNumber = requests.NewInteger(response.PageNumber + 1)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCmsMonitorHost(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCmsMonitorHost")
	// Create service connection
	client, err := CmsService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getCmsMonitorHost", "connection_error", err)
		return nil, err
	}

	hostName := d.KeyColumnQuals["host_name"].GetStringValue()
	instanceId := d.KeyColumnQuals["instance_id"].GetStringValue()

	// handle empty hostName or instanceId in get call
	if hostName == "" || instanceId == "" {
		return nil, nil
	}

	request := cms.CreateDescribeMonitoringAgentHostsRequest()
	request.Scheme = "https"
	request.HostName = hostName
	request.InstanceIds = instanceId

	response, err := client.DescribeMonitoringAgentHosts(request)
	if err != nil {
		plugin.Logger(ctx).Error("getCmsMonitorHost", "query_error", err, "request", request)
		return nil, err
	}

	return response.Hosts.Host[0], nil
}

func getCmsMonitoringAgentStatus(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCmsMonitoringAgentStatus")

	// Create service connection
	client, err := CmsService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getCmsMonitoringAgentStatus", "connection_error", err)
		return nil, err
	}

	id := h.Item.(cms.Host).InstanceId

	request := cms.CreateDescribeMonitoringAgentStatusesRequest()
	request.Scheme = "https"
	request.InstanceIds = id

	response, err := client.DescribeMonitoringAgentStatuses(request)
	if err != nil {
		plugin.Logger(ctx).Error("getCmsMonitoringAgentStatus", "query_error", err, "request", request)
		return nil, err
	}

	return response.NodeStatusList.NodeStatus, nil
}

func getCmsMonitoringHostAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCmsMonitoringHostAka")

	data := h.Item.(cms.Host)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"arn:acs:cms:" + data.Region + ":" + accountID + ":host/" + data.HostName}

	return akas, nil
}
