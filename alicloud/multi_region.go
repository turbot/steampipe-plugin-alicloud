package alicloud

import (
	"context"
	"strings"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

const matrixKeyRegion = "region"

// BuildRegionList :: return a list of matrix items, one per region specified in the connection config
func BuildRegionList(_ context.Context, connection *plugin.Connection) []map[string]interface{} {
	// retrieve regions from connection config
	alicloudConfig := GetConfig(connection)

	if alicloudConfig.Regions != nil {
		regions := GetConfig(connection).Regions

		if len(getInvalidRegions(regions)) > 0 {
			panic("\n\nConnection config have invalid regions: " + strings.Join(getInvalidRegions(regions), ",") + ". Edit your connection configuration file and then restart Steampipe.")
		}

		// validate regions list
		matrix := make([]map[string]interface{}, len(regions))
		for i, region := range regions {
			matrix[i] = map[string]interface{}{matrixKeyRegion: region}
		}
		return matrix
	}

	return []map[string]interface{}{
		{matrixKeyRegion: GetDefaultRegion(connection)},
	}
}

func getInvalidRegions(regions []string) []string {
	alicloudRegions := []string{
		"cn-beijing", "cn-chengdu", "cn-guangzhou", "cn-hangzhou", "cn-heyuan", "cn-hongkong", "cn-huhehaote", "cn-qingdao", "cn-shanghai", "cn-shenzhen", "cn-wulanchabu", "cn-zhangjiakou", "ap-northeast-1", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-5", "eu-central-1", "eu-west-1", "me-east-1", "us-east-1", "us-west-1"}

	invalidRegions := []string{}
	for _, region := range regions {
		if !helpers.StringSliceContains(alicloudRegions, region) {
			invalidRegions = append(invalidRegions, region)
		}
	}
	return invalidRegions
}
