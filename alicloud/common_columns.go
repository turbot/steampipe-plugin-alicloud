package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// struct to store the common column data
type alicloudCommonColumnData struct {
	AccountID string
}

func getCommonColumns(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	cacheKey := "commonColumnData"
	var commonColumnData *alicloudCommonColumnData
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		commonColumnData = cachedData.(*alicloudCommonColumnData)
	} else {

		client, err := connectSts(ctx)
		if err != nil {
			return nil, err
		}

		request := sts.CreateGetCallerIdentityRequest()
		request.Scheme = "https"

		callerIdentity, err := client.GetCallerIdentity(request)
		if err != nil {
			return nil, err
		}

		commonColumnData = &alicloudCommonColumnData{
			// extract partition from arn
			AccountID: callerIdentity.AccountId,
		}

		// save to extension cache
		d.ConnectionManager.Cache.Set(cacheKey, commonColumnData)
	}

	plugin.Logger(ctx).Trace("getCommonColumns: ", "commonColumnData", commonColumnData)

	return commonColumnData, nil
}
