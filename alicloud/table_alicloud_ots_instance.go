package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ots"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableAlicloudOtsInstance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ots_instance",
		Description: "Alicloud Ots Instance",
		List: &plugin.ListConfig{
			Hydrate: listOtsInstances,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("instance_name"),
			Hydrate:    getOtsIntance,
		},
		GetMatrixItem: BuildRegionList,
		Columns: []*plugin.Column{
			{
				Name:        "instance_name",
				Description: "The name of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The instance current status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cluster_type",
				Description: "The specifications of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_time",
				Description: "Instance creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "The description of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "network",
				Description: "The network type of the instance.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "read_capacity",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "timestamp",
				Description: "",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "write_capacity",
				Description: "",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "quota",
				Description: "",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "tags_src",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("TagInfos.TagInfo"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("TagInfos.TagInfo").Transform(otsInstanceTurbotTags),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InstanceName"),
			},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getOtsInstanceAkas,
				Transform:   transform.FromValue(),
			},

			// Alicloud common columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegionId"),
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

func listOtsInstances(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := TableStoreService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("listOtsInstances", "connection_error", err)
		return nil, err
	}
	request := ots.CreateListInstanceRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(50)
	request.PageNum = requests.NewInteger(1)

	count := 0
	for {
	response, err := client.ListInstance(request)
	if err != nil {
		plugin.Logger(ctx).Error("listOtsInstances", "query_error", err, "request", request)
		return nil, err
	}
	for _, instanceInfo := range response.InstanceInfos.InstanceInfo {
		d.StreamListItem(ctx, instanceInfo)
			count++
	}
		if count >= response.TotalCount {
			break
		}
		request.PageNum = requests.NewInteger(response.PageNum + 1)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getOtsIntance(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := plugin.GetMatrixItem(ctx)[matrixKeyRegion].(string)

	// Create service connection
	client, err := TableStoreService(ctx, d, region)
	if err != nil {
		plugin.Logger(ctx).Error("getOtsIntance", "connection_error", err)
		return nil, err
	}

	name := d.KeyColumnQuals["instance_name"].GetStringValue()

	// Handle empty instance name
	if name == "" {
		return nil, nil
	}
	request := ots.CreateGetInstanceRequest()
	request.Scheme = "https"
	request.InstanceName = name

	response, err := client.GetInstance(request)
	if err != nil {
		plugin.Logger(ctx).Error("getVpc", "query_error", err, "request", request)
		return nil, err
	}

	return response.InstanceInfo, nil
}

func getOtsInstanceAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	instanceInfo := h.Item.(ots.InstanceInfo)

	// Get project details
	commonData, err := getCommonColumns(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ots:" + "" + ":" + accountID + ":instance/" + instanceInfo.InstanceName}, nil
}

//// TRANSFORM FUNCTIONS

func otsInstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]ots.TagInfo)

	if tags == nil || len(tags) == 0 {
		return nil, nil
	}

	turbotTags := map[string]string{}
	for _, i := range tags {
		turbotTags[i.TagKey] = i.TagValue
	}
	return turbotTags, nil
}
