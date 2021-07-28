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

func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	getCallerIdentityCached := plugin.HydrateFunc(getCallerIdentity).WithCache()
	getCallerIdentityData, err := getCallerIdentityCached(ctx, d, h)
	if err != nil {
		return nil, err
	}

	callerIdentity := getCallerIdentityData.(*sts.GetCallerIdentityResponse)
	commonColumnData := &alicloudCommonColumnData{
		AccountID: callerIdentity.AccountId,
	}

	return commonColumnData, nil
}

func getCallerIdentity(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cacheKey := "GetCallerIdentity"

	// if found in cache, return the result
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*sts.GetCallerIdentityResponse), nil
	}

	// Create service connection
	client, err := StsService(ctx, d)
	if err != nil {
		return nil, err
	}

	request := sts.CreateGetCallerIdentityRequest()
	request.Scheme = "https"

	callerIdentity, err := client.GetCallerIdentity(request)
	if err != nil {
		// let the cache know that we have failed to fetch this item
		return nil, err
	}

	// save to extension cache
	d.ConnectionManager.Cache.Set(cacheKey, callerIdentity)

	return callerIdentity, nil
}
