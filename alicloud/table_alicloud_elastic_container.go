package alicloud

import (
	"context"
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudElasticContainer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_elastic_container",
		Description: "Elastic Elastic Container",
		List: &plugin.ListConfig{
			Hydrate: listElasticContainers,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("cluster_id"),
			Hydrate:    getElasticContainer,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("name"),
			},
			{
				Name:        "cluster_id",
				Description: "The ID of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("cluster_id"),
			},
			{
				Name:        "state",
				Description: "The status of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("state"),
			},
			{
				Name:        "size",
				Description: "The number of nodes in the cluster.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("size"),
			},
			{
				Name:        "cluster_type",
				Description: "The type of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("cluster_type"),
			},
			{
				Name:        "created",
				Description: "The time when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("created"),
			},
			{
				Name:        "updated",
				Description: "The time when the cluster was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("updated"),
			},
			{
				Name:        "init_version",
				Description: "The initial version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("init_version"),
			},
			{
				Name:        "current_version",
				Description: "The version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("current_version"),
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the cluster belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("resource_group_id"),
			},
			{
				Name:        "instance_type",
				Description: "The Elastic Compute Service (ECS) instance type of cluster nodes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("instance_type"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC used by the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("vpc_id"),
			},
			{
				Name:        "vswitch_id",
				Description: "The IDs of VSwitches.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("vswitch_id"),
			},
			{
				Name:        "vswitch_cidr",
				Description: "The CIDR block of VSwitches.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("vswitch_cidr"),
			},
			{
				Name:        "data_disk_category",
				Description: "The type of data disks.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("data_disk_category"),
			},
			{
				Name:        "data_disk_size",
				Description: "The size of a data disk.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("data_disk_size"),
			},
			{
				Name:        "zone_id",
				Description: "The ID of the zone where the cluster is deployed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("zone_id"),
			},
			{
				Name:        "network_mode",
				Description: "The network type of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("network_mode"),
			},
			{
				Name:        "subnet_cidr",
				Description: "The CIDR block of pods in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("subnet_cidr"),
			},
			{
				Name:        "external_loadbalancer_id",
				Description: "The ID of the Server Load Balancer (SLB) instance deployed in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("external_loadbalancer_id"),
			},
			{
				Name:        "port",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("port"),
			},
			{
				Name:        "node_status",
				Description: "The status of cluster nodes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("node_status"),
			},
			{
				Name:        "cluster_healthy",
				Description: "The health status of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("cluster_healthy"),
			},
			{
				Name:        "docker_version",
				Description: "The version of Docker.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("docker_version"),
			},
			{
				Name:        "swarm_mode",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("swarm_mode"),
			},
			{
				Name:        "gw_bridge",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("gw_bridge"),
			},
			{
				Name:        "upgrade_components",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("upgrade_components"),
			},
			{
				Name:        "next_version",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("next_version"),
			},
			{
				Name:        "private_zone",
				Description: "Indicates whether PrivateZone is enabled for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("private_zone"),
			},
			{
				Name:        "service_discovery_types",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("service_discovery_types"),
			},
			{
				Name:        "profile",
				Description: "The identifier of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("profile"),
			},
			{
				Name:        "deletion_protection",
				Description: "Indicates whether deletion protection is enabled for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("deletion_protection"),
			},
			{
				Name:        "cluster_spec",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("cluster_spec"),
			},
			{
				Name:        "capabilities",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("capabilities"),
			},
			{
				Name:        "enabled_migration",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("enabled_migration"),
			},
			{
				Name:        "need_update_agent",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("need_update_agent"),
			},
			{
				Name:        "outputs",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("outputs"),
			},
			{
				Name:        "parameters",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("parameters"),
			},
			{
				Name:        "worker_ram_role_name",
				Description: "The name of the RAM role for worker nodes in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("worker_ram_role_name"),
			},
			{
				Name:        "maintenance_info",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("maintenance_info"),
			},
			{
				Name:        "maintenance_window",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("maintenance_window"),
			},
			{
				Name:        "master_url",
				Description: "The endpoints that are open for connections to the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("master_url"),
			},
			{
				Name:        "meta_data",
				Description: "The metadata of the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("meta_data"),
			},
			{
				Name:        "cluster_log",
				Description: "The logs of a cluster.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getClusterLog,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "tags_src",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("tags"),
			},

			// steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("tags").Transform(containerTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getElasticContainerAka,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("name"),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("region_id"),
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

func listElasticContainers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := CSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listElasticContainers", "connection_error", err)
		return nil, err
	}
	request := cs.CreateDescribeClustersV1Request()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNumber = requests.NewInteger(1)

	count := 0
	for {
		response, err := client.DescribeClustersV1(request)
		if err != nil {
			plugin.Logger(ctx).Error("listElasticContainers", "query_error", err, "request", request)
			return nil, err
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(response.GetHttpContentString()), &result)
		clusters := result["clusters"].([]interface{})
		pageInfo := result["page_info"].(map[string]interface{})
		TotalCount := pageInfo["total_count"].(float64)
		PageNumber := pageInfo["page_number"].(float64)
		for _, cluster := range clusters {
			plugin.Logger(ctx).Warn("listElasticContainers", "item", cluster)
			clusterAsMap := cluster.(map[string]interface{})
			d.StreamListItem(ctx, clusterAsMap)
			count++
		}
		if count >= int(TotalCount) {
			break
		}
		request.PageNumber = requests.NewInteger(int(PageNumber) + 1)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getElasticContainer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getElasticContainer")

	// Create service connection
	client, err := CSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getElasticContainer", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		clusterData := h.Item.(map[string]interface{})
		id = clusterData["cluster_id"].(string)
	} else {
		id = d.KeyColumnQuals["cluster_id"].GetStringValue()
	}

	request := cs.CreateDescribeClusterDetailRequest()
	request.Scheme = "https"
	request.ClusterId = id

	response, err := client.DescribeClusterDetail(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("getElasticContainer", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.GetHttpContentString()) > 0 {
		var cluster map[string]interface{}
		json.Unmarshal([]byte(response.GetHttpContentString()), &cluster)
		return cluster, nil
	}

	return nil, nil
}

func getClusterLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getClusterLog")

	// Create service connection
	client, err := CSService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getClusterLog", "connection_error", err)
		return nil, err
	}

	var id string
	if h.Item != nil {
		clusterData := h.Item.(map[string]interface{})
		id = clusterData["cluster_id"].(string)
	} else {
		id = d.KeyColumnQuals["cluster_id"].GetStringValue()
	}

	request := cs.CreateDescribeClusterLogsRequest()
	request.Scheme = "https"
	request.ClusterId = id

	response, err := client.DescribeClusterLogs(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("getClusterLog", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.GetHttpContentString()) > 0 {
		return response.GetHttpContentString(), nil
	}

	return nil, nil
}

func getElasticContainerAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getElasticContainerAka")

	data := h.Item.(map[string]interface{})

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cs:" + data["region_id"].(string) + ":" + accountID + ":container/" + data["cluster_id"].(string)}

	return akas, nil
}

func containerTagsToMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]interface{})
	if tags == nil {
		return nil, nil
	}

	if len(tags) == 0 {
		return nil, nil
	}
	turbotTagsMap := map[string]string{}
	for _, i := range tags {
		tagDetails := i.(map[string]interface{})
		turbotTagsMap[tagDetails["key"].(string)] = tagDetails["value"].(string)
	}

	return turbotTagsMap, nil
}
