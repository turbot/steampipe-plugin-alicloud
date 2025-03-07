package alicloud

import (
	"context"
	"slices"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const matrixKeyRegion = "region"

// BuildRegionList :: return a list of matrix items, one per region specified in the connection config
func BuildRegionList(_ context.Context, d *plugin.QueryData) []map[string]interface{} {
	// retrieve regions from connection config
	alicloudConfig := GetConfig(d.Connection)

	if alicloudConfig.Regions != nil {
		regions := GetConfig(d.Connection).Regions

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
		{matrixKeyRegion: GetDefaultRegion(d.Connection)},
	}
}

func getInvalidRegions(regions []string) []string {
	alicloudRegions := []string{
		"cn-beijing", "cn-beijing-finance-1", "cn-chengdu", "cn-guangzhou", "cn-hangzhou", "cn-heyuan", "cn-hongkong", "cn-huhehaote", "cn-qingdao", "cn-shanghai", "cn-shanghai-finance-1", "cn-shenzhen", "cn-shenzhen-finance-1", "cn-wulanchabu", "cn-zhangjiakou", "ap-northeast-1", "ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-5", "ap-southeast-6", "ap-southeast-7", "eu-central-1", "eu-west-1", "me-east-1", "me-central-1", "us-east-1", "us-west-1", "cn-wuhan-lr", "cn-nanjing", "cn-fuzhou",
	}

	invalidRegions := []string{}
	for _, region := range regions {
		if !slices.Contains(alicloudRegions, region) {
			invalidRegions = append(invalidRegions, region)
		}
	}
	return invalidRegions
}
