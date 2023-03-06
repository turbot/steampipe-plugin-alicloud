package alicloud

import (
	"context"
	"fmt"
	"os"

	ims "github.com/alibabacloud-go/ims-20190815/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/actiontrail"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sas"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// AutoscalingService returns the service connection for Alicloud Autoscaling service
func AutoscalingService(ctx context.Context, d *plugin.QueryData) (*ess.Client, error) {
	region := d.EqualsQualString(matrixKeyRegion)

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

// CasService returns the service connection for Alicloud SSL service
func CasService(ctx context.Context, d *plugin.QueryData, region string) (*cas.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed CasService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("cas-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*cas.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := cas.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// CmsService returns the service connection for Alicloud CMS service
func CmsService(ctx context.Context, d *plugin.QueryData) (*cms.Client, error) {
	region := GetDefaultRegion(d.Connection)

	if region == "" {
		return nil, fmt.Errorf("region must be passed CmsService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("cms-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*cms.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := cms.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// ECSService returns the service connection for Alicloud ECS service
func ECSService(ctx context.Context, d *plugin.QueryData) (*ecs.Client, error) {
	region := d.EqualsQualString(matrixKeyRegion)

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

// ECSRegionService returns the service connection for Alicloud ECS Region service
func ECSRegionService(ctx context.Context, d *plugin.QueryData, region string) (*ecs.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed ECSRegionService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("ecsregion-%s", region)
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

// IMSService returns the service connection for Alicloud IMS service
func IMSService(ctx context.Context, d *plugin.QueryData) (*ims.Client, error) {
	region := GetDefaultRegion(d.Connection)
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("ims-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*ims.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	config := &rpc.Config{}
	config.AccessKeyId = &ak
	config.AccessKeySecret = &secret
	config.RegionId = &region

	// so it was not in cache - create service
	svc, err := ims.NewClient(config)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// KMSService returns the service connection for Alicloud KMS service
func KMSService(ctx context.Context, d *plugin.QueryData) (*kms.Client, error) {
	region := d.EqualsQualString(matrixKeyRegion)

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

// SLBService returns the service connection for Alicloud Server Load Balancer service
func SLBService(ctx context.Context, d *plugin.QueryData) (*slb.Client, error) {
	region := GetDefaultRegion(d.Connection)

	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("ram-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*slb.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := slb.NewClientWithAccessKey(region, ak, secret)
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
func VpcService(ctx context.Context, d *plugin.QueryData) (*vpc.Client, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	if region == "" {
		return nil, fmt.Errorf("region must be passed VpcService")
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

// OssService returns the service connection for Alicloud OSS service
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

	if alicloudConfig.Regions != nil {
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

	// https://github.com/aliyun/aliyun-cli/blob/master/CHANGELOG.md#3040
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

	return accessKey, secretKey, nil
}

// RDSService returns the service connection for Alicloud RDS service
func RDSService(ctx context.Context, d *plugin.QueryData, region string) (*rds.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed RDSService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("rds-%s", region)
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

// ActionTrailService returns the service connection for Alicloud ActionTrail service
func ActionTrailService(ctx context.Context, d *plugin.QueryData) (*actiontrail.Client, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	if region == "" {
		return nil, fmt.Errorf("region must be passed ActionTrailService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("actiontrail-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*actiontrail.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := actiontrail.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// ContainerService returns the service connection for Alicloud Container service
func ContainerService(ctx context.Context, d *plugin.QueryData) (*cs.Client, error) {
	region := GetDefaultRegion(d.Connection)

	if region == "" {
		return nil, fmt.Errorf("region must be passed ContainerService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("cs-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*cs.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := cs.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// SecurityCenterService returns the service connection for Alicloud Security Center service
func SecurityCenterService(ctx context.Context, d *plugin.QueryData, region string) (*sas.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed SecurityCenterService")
	}

	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("sas-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*sas.Client), nil
	}

	ak, secret, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}

	// so it was not in cache - create service
	svc, err := sas.NewClientWithAccessKey(region, ak, secret)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}
