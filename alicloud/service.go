package alicloud

import (
	"context"
	"fmt"
	"os"

	// V2.0 SDK imports
	actiontrail "github.com/alibabacloud-go/actiontrail-20200706/client"
	alidns "github.com/alibabacloud-go/alidns-20150109/client"

	// cas "github.com/alibabacloud-go/cas-20200407/client" // Temporarily disabled due to tea-utils compatibility issues
	// cms "github.com/alibabacloud-go/cms-20190101/client" // Temporarily disabled due to code generation issues
	cs "github.com/alibabacloud-go/cs-20151215/client"
	openapiClient "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/client"
	ess "github.com/alibabacloud-go/ess-20220222/client"
	kms "github.com/alibabacloud-go/kms-20160120/client"
	ram "github.com/alibabacloud-go/ram-20150501/client"
	rds "github.com/alibabacloud-go/rds-20140815/client"
	sas "github.com/alibabacloud-go/sas-20181203/client"
	slb "github.com/alibabacloud-go/slb-20140515/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	teaRoaClient "github.com/alibabacloud-go/tea-roa/client"
	vpc "github.com/alibabacloud-go/vpc-20160428/client"

	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	credential "github.com/aliyun/credentials-go/credentials"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	ossCred "github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// AliDNSService returns the service connection for Alicloud DNS service - now using V2.0 SDK
func AliDNSService(ctx context.Context, d *plugin.QueryData) (*alidns.Client, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	if region == "" {
		return nil, fmt.Errorf("region must be passed AliDNSService")
	}

	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("alidns-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*alidns.Client), nil
	}

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := alidns.NewClient(clientCfg)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// AutoscalingService returns the service connection for Alicloud Autoscaling service - now using V2.0 SDK
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := ess.NewClient(clientCfg)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// CasService returns the service connection for Alicloud SSL service
// Temporarily disabled due to tea-utils compatibility issues
/*
func CasService(ctx context.Context, d *plugin.QueryData, region string) (*cas.Client, error) {
	if region == "" {
		return nil, fmt.Errorf("region must be passed CasService")
	}
	// have we already created and cached the service?
	serviceCacheKey := fmt.Sprintf("cas-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*cas.Client), nil
	}

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := cas.NewClient(clientCfg)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}
*/

// CmsService returns the service connection for Alicloud CMS service
// Temporarily disabled due to code generation issues
/*
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := cms.NewClient(cfg.Config)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}
*/

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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := ecs.NewClient(cfg.Config)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := ecs.NewClient(cfg.Config)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := kms.NewClient(cfg.Config)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := ram.NewClient(clientCfg)
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
	serviceCacheKey := fmt.Sprintf("slb-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*slb.Client), nil
	}

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := slb.NewClient(clientCfg)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := sts.NewClient(clientCfg)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := vpc.NewClient(cfg.Config)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

// OssService returns the service connection for Alicloud OSS service
func OssService(ctx context.Context, d *plugin.QueryData, region string) (*oss.Client, error) {
	// Validate the region parameter before proceeding
	if region == "" {
		return nil, fmt.Errorf("region must be provided to initialize the OSS service")
	}

	// Check if the OSS client is already cached to avoid redundant initialization
	serviceCacheKey := fmt.Sprintf("oss-%s", region)
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*oss.Client), nil
	}

	// Construct the OSS endpoint for the given region
	endpoint := "oss-" + region + ".aliyuncs.com"

	// Initialize OSS client configuration
	ossCfg := oss.NewConfig()
	ossCfg.WithEndpoint(endpoint)
	ossCfg.WithRegion(region)
	ossCfg.WithProxyFromEnvironment(true)

	// Retrieve cached credentials for authentication
	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cached credentials: %v", err)
	}

	cfg := credCfg.(*CredentialConfig)

	credentialP, err := cfg.Config.Credential.GetCredential()
	if err != nil {
		return nil, fmt.Errorf("failed to get access key secret: %v", err)
	}

	// Create OSS credentials provider using V2.0 SDK credential values
	ossCfg.CredentialsProvider = ossCred.NewStaticCredentialsProvider(*credentialP.AccessKeyId, *credentialP.AccessKeySecret, *credentialP.SecurityToken)

	// Initialize and return the OSS client
	svc := oss.NewClient(ossCfg)

	// Cache the service connection to optimize future requests
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := actiontrail.NewClient(cfg.Config)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	teaRoaC := &teaRoaClient.Config{}
	teaRoaC.SetCredential(cfg.Config.Credential)
	svc, err := cs.NewClient(teaRoaC)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	clientCfg := &openapiClient.Config{}
	clientCfg.SetCredential(cfg.Config.Credential)
	svc, err := sas.NewClient(clientCfg)
	if err != nil {
		return nil, err
	}

	// cache the service connection
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// Create V2.0 client using the individual service client
	svc, err := rds.NewClient(cfg.Config)
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

// https://github.com/aliyun/aliyun-cli/blob/master/README.md#supported-environment-variables
func getEnvForProfile(_ context.Context, d *plugin.QueryData) (profile string) {
	alicloudConfig := GetConfig(d.Connection)
	if alicloudConfig.Profile != nil {
		profile = *alicloudConfig.Profile
	} else {
		var ok bool
		if profile, ok = os.LookupEnv("ALIBABACLOUD_PROFILE"); !ok {
			if profile, ok = os.LookupEnv("ALIBABA_CLOUD_PROFILE"); !ok {
				if profile, ok = os.LookupEnv("ALICLOUD_PROFILE"); !ok {
					return ""
				}
			}
		}
	}
	return profile
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
					panic("\n'access_key' or 'profile' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
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
					panic("\n'secret_key' or 'profile' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
				}
			}
		}
	}

	return accessKey, secretKey, nil
}

// Credential configuration - now using V2.0 SDK
type CredentialConfig struct {
	Config        *rpc.Config
	DefaultRegion string
	Runtime       *util.RuntimeOptions
}

// Get credential from the profile configuration for Alicloud CLI - now using V2.0 SDK
func getProfileConfigurations(ctx context.Context, d *plugin.QueryData) (*CredentialConfig, error) {
	alicloudConfig := GetConfig(d.Connection)
	profile := alicloudConfig.Profile

	cfg, err := getCredentialConfigByProfile(*profile, d)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func getCredentialConfigByProfile(profile string, d *plugin.QueryData) (*CredentialConfig, error) {
	defaultRegion := GetDefaultRegion(d.Connection)
	config := GetConfig(d.Connection)

	// Initialize V2.0 client config with default settings
	defaultConfig := new(rpc.Config)
	defaultConfig.SetRegionId(defaultRegion)
	defaultConfig.SetProtocol("HTTPS")

	// Initialize runtime options
	runtime := new(util.RuntimeOptions)
	runtime.SetAutoretry(true)

	// Apply configuration options (maintaining same behavior as V1.0)
	if config.AutoRetry != nil {
		runtime.SetAutoretry(*config.AutoRetry)
	}
	if config.MaxRetryTime != nil {
		runtime.SetMaxIdleConns(*config.MaxRetryTime)
	}
	if config.Timeout != nil {
		runtime.SetReadTimeout(*config.Timeout * 1000) // Convert to milliseconds
	}

	// For V2.0 SDK, we use the default credential provider chain
	// This will automatically resolve profile-based credentials from ~/.alibabacloud/credentials
	// and fall back to environment variables if needed
	cred, err := credential.NewCredential(nil)
	if err != nil {
		// Fallback to environment variables if profile resolution fails
		accessKey, secretKey, err := getEnv(context.Background(), d)
		if err != nil {
			return nil, err
		}

		if accessKey != "" && secretKey != "" {
			defaultConfig.SetAccessKeyId(accessKey)
			defaultConfig.SetAccessKeySecret(secretKey)
			return &CredentialConfig{defaultConfig, defaultRegion, runtime}, nil
		}

		return nil, fmt.Errorf("unable to resolve credentials for profile: %s", profile)
	}

	// Set the credential in the config
	defaultConfig.SetCredential(cred)

	return &CredentialConfig{defaultConfig, defaultRegion, runtime}, nil
}

var getCredentialSessionCached = plugin.HydrateFunc(getCredentialSessionUncached).Memoize()

func getCredentialSessionUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var connectionCfg *CredentialConfig

	config := GetConfig(d.Connection)
	defaultRegion := GetDefaultRegion(d.Connection)

	// Initialize V2.0 client config with default settings
	defaultConfig := new(rpc.Config)
	defaultConfig.SetRegionId(defaultRegion)
	defaultConfig.SetProtocol("HTTPS")

	// Initialize runtime options
	runtime := new(util.RuntimeOptions)
	runtime.SetAutoretry(true)

	// Apply configuration options (maintaining same behavior as V1.0)
	if config.AutoRetry != nil {
		runtime.SetAutoretry(*config.AutoRetry)
	}
	if config.MaxRetryTime != nil {
		runtime.SetMaxIdleConns(*config.MaxRetryTime)
	}
	if config.Timeout != nil {
		runtime.SetReadTimeout(*config.Timeout * 1000) // Convert to milliseconds
	}

	// Profile based client
	if config.Profile != nil {
		return getProfileConfigurations(ctx, d)
	}

	profileEnv := getEnvForProfile(ctx, d)
	if profileEnv != "" {
		return getCredentialConfigByProfile(profileEnv, d)
	}

	// Access key and Secret Key from environment variable
	accessKey, secretKey, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}
	if accessKey != "" && secretKey != "" {
		defaultConfig.SetAccessKeyId(accessKey)
		defaultConfig.SetAccessKeySecret(secretKey)
		connectionCfg = &CredentialConfig{defaultConfig, defaultRegion, runtime}
		return connectionCfg, nil
	}

	return nil, nil
}
