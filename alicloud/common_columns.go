package alicloud

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	sts "github.com/alibabacloud-go/sts-20150401/client"
)

// struct to store the common column data
type alicloudCommonColumnData struct {
	AccountID string
}

// getCommonColumns:: helps to avoid multiple sts.GetCallerIdentity API calls in parallel where using it directly in column definitions
func getCommonColumns(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	getCallerIdentityData, err := getAccountDetails(ctx, d, h)
	if err != nil {
		return nil, err
	}

	callerIdentity := getCallerIdentityData.(*sts.GetCallerIdentityResponse)
	commonColumnData := &alicloudCommonColumnData{
		AccountID: *callerIdentity.Body.AccountId,
	}

	return commonColumnData, nil
}

func getAccountId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	getCallerIdentityData, err := getAccountDetails(ctx, d, h)
	if err != nil {
		return nil, err
	}

	callerIdentity := getCallerIdentityData.(*sts.GetCallerIdentityResponse)

	return callerIdentity.Body.AccountId, nil
}

var getAccountDetailsMemoize = plugin.HydrateFunc(getCallerIdentityUncached).Memoize(memoize.WithCacheKeyFunction(getAccountDetailsCacheKey))

func getAccountDetailsCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "GetCallerIdentity"
	return cacheKey, nil
}

func getAccountDetails(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	config, err := getAccountDetailsMemoize(ctx, d, h)
	if err != nil {
		return nil, err
	}

	c := config.(*sts.GetCallerIdentityResponse)

	return c, nil
}

func getCallerIdentityUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Create service connection
	client, err := StsService(ctx, d)
	if err != nil {
		return nil, err
	}

	callerIdentity, err := client.GetCallerIdentity()
	if err != nil {
		// let the cache know that we have failed to fetch this item
		return nil, err
	}

	return callerIdentity, nil
}
