package alicloud

import (
	"context"
	"strings"

	"github.com/turbot/go-kit/helpers"
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
		"cn-beijing", "cn-beijing-finance-1", "cn-chengdu", "cn-guangzhou", "cn-hangzhou", "cn-heyuan", "cn-hongkong", "cn-huhehaote", "cn-qingdao", "cn-shanghai", "cn-shanghai-finance-1", "cn-shenzhen", "cn-shenzhen-finance-1", "cn-wulanchabu", "cn-zhangjiakou", "ap-northeast-1","ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ap-southeast-5", "eu-central-1", "eu-west-1", "me-east-1", "us-east-1", "us-west-1"}

	invalidRegions := []string{}
	for _, region := range regions {
		if !helpers.StringSliceContains(alicloudRegions, region) {
			invalidRegions = append(invalidRegions, region)
		}
	}
	return invalidRegions
}
