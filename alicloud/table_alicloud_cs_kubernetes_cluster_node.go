package alicloud

import (
	"context"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudCsKubernetesClusterNode(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_cs_kubernetes_cluster_node",
		Description: "Alicloud Container Service Kubernetes Cluster Node",
		List: &plugin.ListConfig{
			Hydrate:       listCsKubernetesClusterNodes,
			ParentHydrate: listCsKubernetesClusters,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"cluster_id", "instance_id"}),
			Hydrate:    getCsKubernetesClusterNode,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "node_name",
				Description: "The name of the node in the ACK cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cluster_id",
				Description: "The ID of the cluster that the node pool belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(clusterIdFromInstanceName),
			},
			{
				Name:        "state",
				Description: "The states of the nodes in the node pool.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The time when the node was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "expired_time",
				Description: "The expiration time of the node.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "node_status",
				Description: "Indicates whether the node is ready in the ACK cluster. Valid values: true, false.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_aliyun_node",
				Description: "Indicates whether the instance is provided by Alibaba Cloud.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "source",
				Description: "Indicates how the nodes in the node pool were initialized. The nodes can be manually created or created by using Resource Orchestration Service (ROS).",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_id",
				Description: "The ID of the ECS instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_name",
				Description: "The name of the node. This name contains the ID of the cluster to which the node is deployed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_role",
				Description: "The role of the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_status",
				Description: "The state of the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_type",
				Description: "The instance type of the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_charge_type",
				Description: "The billing method of the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "instance_type_family",
				Description: "The ECS instance family of the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "ip_address",
				Description: "The IP address of the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "image_id",
				Description: "The ID of the system image that is used by the node.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "host_name",
				Description: "The name of the host.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "nodepool_id",
				Description: "The ID of the node pool.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "error_message",
				Description: "The error message generated when the node was created.",
				Type:        proto.ColumnType_STRING,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NodeName"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCsKubernetesClusterNodeAka,
				Transform:   transform.FromValue(),
			},

			// Alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(clusterNodeRegion),
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

func listCsKubernetesClusterNodes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := GetDefaultRegion(d.Connection)

	// Create service connection
	client, err := ContainerService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listCsKubernetesClusterNodes", "connection_error", err)
		return nil, err
	}

	clusterId := h.Item.(map[string]interface{})["cluster_id"].(string)
	request := cs.CreateDescribeClusterNodesRequest()
	request.Scheme = "https"
	request.ClusterId = clusterId

	response, err := client.DescribeClusterNodes(request)
	if err != nil {
		plugin.Logger(ctx).Error("listCsKubernetesClusterNodes", "query_error", err, "request", request)
		return nil, err
	}
	for _, node := range response.Nodes {
		d.StreamListItem(ctx, node)
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCsKubernetesClusterNode(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	matrixRegion := GetDefaultRegion(d.Connection)
	plugin.Logger(ctx).Trace("getCsKubernetesClusterNode")

	// Create service connection
	client, err := ContainerService(ctx, d, matrixRegion)
	if err != nil {
		plugin.Logger(ctx).Error("getCsKubernetesClusterNode", "connection_error", err)
		return nil, err
	}

	clusterId := d.KeyColumnQuals["cluster_id"].GetStringValue()
	instanceId := d.KeyColumnQuals["instance_id"].GetStringValue()

	// handle empty clusterId or instanceId in get call
	if clusterId == "" || instanceId == "" {
		return nil, nil
	}

	request := cs.CreateDescribeClusterNodesRequest()
	request.Scheme = "https"
	request.ClusterId = clusterId

	response, err := client.DescribeClusterNodes(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("getCsKubernetesClusterNode", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	for _, item := range response.Nodes {
		if item.InstanceId == instanceId {
			return item, nil
		}
	}

	return nil, nil
}

func getCsKubernetesClusterNodeAka(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCsKubernetesClusterNodeAka")

	nodeName := h.Item.(cs.Node).NodeName

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	akas := []string{"acs:cs:" + strings.Split(nodeName, ".")[0] + ":" + accountID + ":node/" + nodeName}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func clusterIdFromInstanceName(_ context.Context, d *transform.TransformData) (interface{}, error) {
	instanceName := d.HydrateItem.(cs.Node).InstanceName

	return strings.Split(instanceName, "-")[4], nil
}

func clusterNodeRegion(_ context.Context, d *transform.TransformData) (interface{}, error) {
	nodeName := d.HydrateItem.(cs.Node).NodeName

	return strings.Split(nodeName, ".")[0], nil
}
