package alicloud

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudCsKubernetesCluster(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_cs_kubernetes_cluster",
		Description: "Alicloud Container Service Kubernetes Cluster",
		List: &plugin.ListConfig{
			Hydrate: listCsKubernetesClusters,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"cluster_id", "region"}),
			Hydrate:    getCsKubernetesCluster,
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
				Name:        "created_at",
				Description: "The time when the cluster was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("created"),
			},
			{
				Name:      "capabilities",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("capabilities"),
			},
			{
				Name:        "cluster_healthy",
				Description: "The health status of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("cluster_healthy"),
			},
			{
				Name:      "cluster_spec",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("cluster_spec"),
			},
			{
				Name:        "cluster_type",
				Description: "The type of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("cluster_type"),
			},
			{
				Name:        "current_version",
				Description: "The version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("current_version"),
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
				Name:        "deletion_protection",
				Description: "Indicates whether deletion protection is enabled for the cluster.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("deletion_protection"),
			},
			{
				Name:        "docker_version",
				Description: "The version of Docker.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("docker_version"),
			},
			{
				Name:      "enabled_migration",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("enabled_migration"),
			},
			{
				Name:        "external_loadbalancer_id",
				Description: "The ID of the Server Load Balancer (SLB) instance deployed in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("external_loadbalancer_id"),
			},
			{
				Name:      "gw_bridge",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("gw_bridge"),
			},
			{
				Name:        "init_version",
				Description: "The initial version of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("init_version"),
			},
			{
				Name:        "instance_type",
				Description: "The Elastic Compute Service (ECS) instance type of cluster nodes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("instance_type"),
			},
			{
				Name:      "maintenance_info",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("maintenance_info"),
			},
			{
				Name:      "need_update_agent",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("need_update_agent"),
			},
			{
				Name:        "network_mode",
				Description: "The network type of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("network_mode"),
			},
			{
				Name:      "next_version",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("next_version"),
			},
			{
				Name:        "node_status",
				Description: "The status of cluster nodes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("node_status"),
			},
			{
				Name:      "outputs",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("outputs"),
			},
			{
				Name:      "parameters",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("parameters"),
			},
			{
				Name:        "port",
				Description: "Container port in Kubernetes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("port"),
			},
			{
				Name:        "private_zone",
				Description: "Indicates whether PrivateZone is enabled for the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("private_zone"),
			},
			{
				Name:        "profile",
				Description: "The identifier of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("profile"),
			},
			{
				Name:        "resource_group_id",
				Description: "The ID of the resource group to which the cluster belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("resource_group_id"),
			},
			{
				Name:      "service_discovery_types",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("service_discovery_types"),
			},
			{
				Name:        "subnet_cidr",
				Description: "The CIDR block of pods in the cluster.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("subnet_cidr"),
			},
			{
				Name:      "swarm_mode",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("swarm_mode"),
			},
			{
				Name:        "updated",
				Description: "The time when the cluster was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("updated"),
			},
			{
				Name:      "upgrade_components",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("upgrade_components"),
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
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("vswitch_cidr"),
			},
			{
				Name:        "worker_ram_role_name",
				Description: "The name of the RAM role for worker nodes in the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("worker_ram_role_name"),
			},
			{
				Name:        "zone_id",
				Description: "The ID of the zone where the cluster is deployed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("zone_id"),
			},
			{
				Name:        "cluster_log",
				Description: "The logs of a cluster.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCsKubernetesClusterLog,
				Transform:   transform.FromValue(),
			},
			{
				Name:      "maintenance_window",
				Type:      proto.ColumnType_JSON,
				Transform: transform.FromField("maintenance_window"),
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
				Name: "cluster_name_spaces",
				Type: proto.ColumnType_JSON,
				Hydrate: getCsKubernetesClusterNameSpaces,
				Transform: transform.FromValue(),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("name"),
			},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("tags").Transform(csKubernetesClusterAkaTagsToMap),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCsKubernetesClusterAka,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
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

func listCsKubernetesClusters(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := ContainerService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listCsKubernetesClusters", "connection_error", err)
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
			plugin.Logger(ctx).Error("listCsKubernetesClusters", "query_error", err, "request", request)
			return nil, err
		}
		var result map[string]interface{}
		json.Unmarshal([]byte(response.GetHttpContentString()), &result)
		clusters := result["clusters"].([]interface{})
		pageInfo := result["page_info"].(map[string]interface{})
		TotalCount := pageInfo["total_count"].(float64)
		PageNumber := pageInfo["page_number"].(float64)
		for _, cluster := range clusters {
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

func getCsKubernetesCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	matrixRegion := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getCsKubernetesCluster")

	// Create service connection
	client, err := ContainerService(ctx, d, matrixRegion)
	if err != nil {
		plugin.Logger(ctx).Error("getCsKubernetesCluster", "connection_error", err)
		return nil, err
	}
	region := d.KeyColumnQuals["region"].GetStringValue()

	if region != matrixRegion {
		return nil, nil
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
		plugin.Logger(ctx).Error("getCsKubernetesCluster", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.GetHttpContentString()) > 0 {
		var cluster map[string]interface{}
		json.Unmarshal([]byte(response.GetHttpContentString()), &cluster)
		return cluster, nil
	}

	return nil, nil
}

func getCsKubernetesClusterLog(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getCsKubernetesClusterLog")

	// Create service connection
	client, err := ContainerService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getCsKubernetesClusterLog", "connection_error", err)
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
		plugin.Logger(ctx).Error("getCsKubernetesClusterLog", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	if len(response.GetHttpContentString()) > 0 {
		return response.GetHttpContentString(), nil
	}

	return nil, nil
}

func getCsKubernetesClusterNameSpaces(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)
	plugin.Logger(ctx).Trace("getCsKubernetesClusterNameSpaces")

	var id string
	if h.Item != nil {
		clusterData := h.Item.(map[string]interface{})
		id = clusterData["cluster_id"].(string)
	} else {
		id = d.KeyColumnQuals["cluster_id"].GetStringValue()
	}

	accessKey := os.Getenv("ALICLOUD_ACCESS_KEY")
	secretAccess := os.Getenv("ALICLOUD_SECRET_KEY")
	client, err := sdk.NewClientWithAccessKey(region, accessKey, secretAccess)
	if err != nil {
		return nil, nil
	}

	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Scheme = "https"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/k8s/" + id + "/namespaces"
  request.Headers["Content-Type"] = "application/json"
  request.QueryParams["RegionId"] = region
	body := `{}`
	request.Content = []byte(body)
	
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return nil, nil
	}

	return response.GetHttpContentString(), nil
}

func getCsKubernetesClusterAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCsKubernetesClusterAka")

	data := h.Item.(map[string]interface{})

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cs:" + data["region_id"].(string) + ":" + accountID + ":container/" + data["name"].(string)}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func csKubernetesClusterAkaTagsToMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
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
