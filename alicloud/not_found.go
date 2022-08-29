package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

// function which returns an ErrorPredicate for ALicloud API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if serverErr, ok := err.(*errors.ServerError); ok {
			return helpers.StringSliceContains(notFoundErrors, serverErr.ErrorCode())
		}
		return false
	}
}
