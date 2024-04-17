package alicloud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	ims "github.com/alibabacloud-go/ims-20190815/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/provider"
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// so it was not in cache - create service
	svc, err := ram.NewClientWithOptions(region, cfg.Config, cfg.Creds)
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

	credCfg, err := getCredentialSessionCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	cfg := credCfg.(*CredentialConfig)

	// so it was not in cache - create service
	svc, err := slb.NewClientWithOptions(region, cfg.Config, cfg.Creds)
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

	// so it was not in cache - create service
	svc, err := sts.NewClientWithOptions(region, cfg.Config, cfg.Creds)
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

	// so it was not in cache - create service
	svc, err := vpc.NewClientWithOptions(region, cfg.Config, cfg.Creds)
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

// Credential configuration
type CredentialConfig struct {
	Creds         auth.Credential
	DefaultRegion string
	Config        *sdk.Config
}

type Config struct {
	Profiles []Profile `json:"profiles"`
}

// Profile structure
type Profile struct {
	Name            string `json:"name"`
	Mode            string `json:"mode"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	StsToken        string `json:"sts_token"`
	StsRegion       string `json:"sts_region"`
	RamRoleName     string `json:"ram_role_name"`
	RamRoleArn      string `json:"ram_role_arn"`
	RamSessionName  string `json:"ram_session_name"`
	SourceProfile   string `json:"source_profile"`
	PublicKeyId     string `json:"public_key_id"`
	PrivateKey      string `json:"private_key"`
	KeyPairName     string `json:"key_pair_name"`
	ExpiredSeconds  int    `json:"expired_seconds"`
	Verified        string `json:"verified"`
	RegionId        string `json:"region_id"`
	OutputFormat    string `json:"output_format"`
	Language        string `json:"language"`
	Site            string `json:"site"`
	RetryTimeout    int    `json:"retry_timeout"`
	ConnectTimeout  int    `json:"connect_timeout"`
	RetryCount      int    `json:"retry_count"`
	ProcessCommand  string `json:"process_command"`
	CredentialsUri  string `json:"credentials_uri"`
}

// Get credential from the profile configuration for Alicloud CLI
func getProfileConfigurations(_ context.Context, d *plugin.QueryData) (*CredentialConfig, error) {
	alicloudConfig := GetConfig(d.Connection)
	if alicloudConfig.Profile != nil {
		defaultRegion := GetDefaultRegion(d.Connection)
		defaultConfig := sdk.NewConfig()
		profile := alicloudConfig.Profile

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get home directory: %v", err)
		}
		// Credential path is ~/.aliyun/config.json
		configPath := filepath.Join(homeDir, ".aliyun", "config.json")

		config, err := loadConfig(configPath)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		configuredProfiles := make(map[string]*Profile, 0)

		for _, p := range config.Profiles {
			configuredProfiles[p.Name] = &p
		}

		profileConfig := configuredProfiles[*profile]

		if profileConfig == nil {
			return nil, fmt.Errorf("profile with name '%s' is not configured", *profile)
		}

		creds := getCredentialBasedOnProfile(profileConfig)
		if creds == nil {
			return nil, fmt.Errorf("unsupported authentication mode '%s'", profileConfig.Mode)
		}
		return &CredentialConfig{creds, defaultRegion, defaultConfig}, nil
	}
	return nil, nil
}

// Load the alicloud credential
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return &config, nil
}

// We can configure the profile with following supported authentication methods:
// https://github.com/aliyun/aliyun-cli/blob/master/README.md#configure-authentication-methods
func getCredentialBasedOnProfile(profileConfig *Profile) interface{} {
	switch profileConfig.Mode {
	case "AK":
		return &credentials.AccessKeyCredential{
			AccessKeyId:     profileConfig.AccessKeyId,
			AccessKeySecret: profileConfig.AccessKeySecret,
		}
	case "StsToken":
		return &credentials.StsTokenCredential{
			AccessKeyId:       profileConfig.AccessKeyId,
			AccessKeySecret:   profileConfig.AccessKeySecret,
			AccessKeyStsToken: profileConfig.StsToken,
		}
	case "EcsRamRole":
		return &credentials.EcsRamRoleCredential{
			RoleName: profileConfig.RamRoleName,
		}
	case "RamRoleArn":
		return &credentials.RamRoleArnCredential{
			AccessKeyId:           profileConfig.AccessKeyId,
			AccessKeySecret:       profileConfig.AccessKeySecret,
			RoleArn:               profileConfig.RamRoleArn,
			RoleSessionName:       profileConfig.RamSessionName,
			RoleSessionExpiration: profileConfig.ExpiredSeconds,
		}
		//// Commenting out for the time being, will uncomment it as per user's request.
		//// This type of authentication is not supported by Alicloud CLI
		//// Supported authentication modes: AK, StsToken, RamRoleArn, and EcsRamRole
		//// https://www.alibabacloud.com/help/en/cli/interactive-configuration-or-fast-configuration#concept-508451/section-pq4-04b-7an
		// case "RsaKeyPair":
		// 	return &credentials.RsaKeyPairCredential{
		// 		PublicKeyId:       profileConfig.PublicKeyId,
		// 		PrivateKey:        profileConfig.PrivateKey,
		// 		SessionExpiration: profileConfig.ExpiredSeconds,
		// }
	}
	return nil
}

var getCredentialSessionCached = plugin.HydrateFunc(getCredentialSessionUncached).Memoize()

func getCredentialSessionUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	var connectionCfg *CredentialConfig

	config := GetConfig(d.Connection)
	defaultRegion := GetDefaultRegion(d.Connection)
	defaultConfig := sdk.NewConfig() // initialize with default config

	// Create client based on environment credential support
	// Supported environment variables: https://github.com/aliyun/aliyun-cli/blob/master/README.md#supported-environment-variables
	envProvider := provider.NewInstanceCredentialsProvider()
	envCreds, err := envProvider.Resolve()
	if err != nil {
		return nil, err
	}
	if envCreds != nil {
		return &CredentialConfig{envCreds, defaultRegion, defaultConfig}, nil
	}

	// Profile based client
	if config.Profile != nil {
		return getProfileConfigurations(ctx, d)
	}

	// Access key and Secret Key from environment variable
	accessKey, secretKey, err := getEnv(ctx, d)
	if err != nil {
		return nil, err
	}
	if accessKey != "" && secretKey != "" {
		creds := credentials.NewAccessKeyCredential(accessKey, secretKey)
		connectionCfg = &CredentialConfig{creds, defaultRegion, defaultConfig}
		return connectionCfg, nil
	}

	return nil, nil
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
