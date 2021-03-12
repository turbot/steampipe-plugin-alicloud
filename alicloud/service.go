package alicloud

import (
	"context"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

// AutoscalingService returns the service connection for Alicloud Autoscaling service
func AutoscalingService(ctx context.Context, d *plugin.QueryData, region string) (*ess.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed AutoscalingService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("ess-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*ess.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := ess.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// ECSService returns the service connection for Alicloud ECS service
func ECSService(ctx context.Context, d *plugin.QueryData, region string) (*ecs.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed ECSService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("ecs-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*ecs.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := ecs.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// KMSService returns the service connection for Alicloud KMS service
func KMSService(ctx context.Context, d *plugin.QueryData, region string) (*kms.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed KMSService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("kms-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*kms.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := kms.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// RAMService returns the service connection for Alicloud RAM service
func RAMService(ctx context.Context, d *plugin.QueryData) (*ram.Client, error) {
	region := GetDefaultRegion(d.Connection)
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("ram-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*ram.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := ram.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// StsService returns the service connection for Alicloud STS service
func StsService(ctx context.Context, d *plugin.QueryData) (*sts.Client, error) {
	region := GetDefaultRegion(d.Connection)
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("sts-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*sts.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := sts.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// VpcService returns the service connection for Alicloud VPC service
func VpcService(ctx context.Context, d *plugin.QueryData, region string) (*vpc.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed ECSService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("vpc-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*vpc.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := vpc.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// OssService returns the service connection for Alicloud VPC service
func OssService(ctx context.Context, d *plugin.QueryData, region string) (*oss.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed OssService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("oss-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*oss.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := oss.New("oss-"+region+".aliyuncs.com", ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// GetDefaultRegion returns the default region used
func GetDefaultRegion(connection *plugin.Connection) string {
	// get alicloud config info
	alicloudConfig := GetConfig(connection)

	var regions []string
	var region string

	if &alicloudConfig != nil && alicloudConfig.Regions != nil {
		regions = alicloudConfig.Regions
	}

	if len(regions) > 0 {
		// Set the first region in regions list to be default region
		region = regions[0]
		// check if it is a valid region
		if len(getInvalidRegions([]string{region})) > 0 {
			panic("\n\nConnection config have invalid region: " + region + ". Edit your connection configuration file and then restart Steampipe")
		}
		return region
	}

	if region == "" {
		region = os.Getenv("ALIBABACLOUD_REGION_ID")
		if region == "" {
			region = os.Getenv("ALICLOUD_REGION_ID")
			if region == "" {
				region = os.Getenv("ALICLOUD_REGION")
			}
		}
	}

	if region == "" {
		region = "cn-hangzhou"
	}

	return region
}

func getEnv(_ context.Context, d *plugin.QueryData) (secretKey string, accessKey string, err error) {

	// https://gitea.com/aliyun/aliyun-cli/src/branch/master/CHANGELOG.md#3-0-40
	// The CLI order of preference is:
	// 1. ALIBABACLOUD_ACCESS_KEY_ID / ALIBABACLOUD_ACCESS_KEY_SECRET / ALIBABACLOUD_REGION_ID
	// 2. ALICLOUD_ACCESS_KEY_ID / ALICLOUD_ACCESS_KEY_SECRET / ALICLOUD_REGION_ID
	// 3. ACCESS_KEY_ID / ACCESS_KEY_SECRET / REGION
	//
	// The Go SDK and Terraform do:
	// 1. ALICLOUD_ACCESS_KEY / ALICLOUD_SECRET_KEY / ALICLOUD_REGION
	//
	// So, Steampipe will do:
	// 1. ALIBABACLOUD_ACCESS_KEY_ID / ALIBABACLOUD_ACCESS_KEY_SECRET / ALIBABACLOUD_REGION_ID
	// 2. ALICLOUD_ACCESS_KEY_ID / ALICLOUD_ACCESS_KEY_SECRET / ALICLOUD_REGION_ID
	// 3. ALICLOUD_ACCESS_KEY / ALICLOUD_SECRET_KEY / ALICLOUD_REGION

	// get alicloud config info
	alicloudConfig := GetConfig(d.Connection)

	if &alicloudConfig != nil {
		if alicloudConfig.AccessKey != nil {
			accessKey = *alicloudConfig.AccessKey
		} else {
			var ok bool
			if accessKey, ok = os.LookupEnv("ALIBABACLOUD_ACCESS_KEY_ID"); !ok {
				if accessKey, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY_ID"); !ok {
					if accessKey, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY"); !ok {
						panic("\n'access_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
					}
				}
			}
		}

		if alicloudConfig.SecretKey != nil {
			secretKey = *alicloudConfig.SecretKey
		} else {
			var ok bool
			if secretKey, ok = os.LookupEnv("ALIBABACLOUD_ACCESS_KEY_SECRET"); !ok {
				if secretKey, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY_SECRET"); !ok {
					if secretKey, ok = os.LookupEnv("ALICLOUD_SECRET_KEY"); !ok {
						panic("\n'secret_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
					}
				}
			}
		}
	}

	return accessKey, secretKey, nil
}

// RDSService returns the service connection for Alicloud RDS service
func RDSService(ctx context.Context, d *plugin.QueryData, region string) (*rds.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed RDSService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("vpc-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*rds.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := rds.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// VpcService returns the service connection for Alicloud VPC service
func CommonService(ctx context.Context, d *plugin.QueryData, region string) (*sdk.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed CommonService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("sdk-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*sdk.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := sdk.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}
