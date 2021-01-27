package alicloud

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	//"github.com/aliyun/credentials-go/credentials"

	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func resourceInterfaceDescription(key string) string {
	switch key {
	case "akas":
		return "Array of globally unique identifier strings (also known as) for the resource."
	case "tags":
		return "A map of tags for the resource."
	case "title":
		return "Title of the resource."
	}
	return ""
}

func ensureStringArray(_ context.Context, d *transform.TransformData) (interface{}, error) {
	switch v := d.Value.(type) {
	case []string:
		return v, nil
	case string:
		return []string{v}, nil
	default:
		str := fmt.Sprintf("%v", d.Value)
		return []string{string(str)}, nil
	}
}

func connectRam(_ context.Context) (*ram.Client, error) {

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

	ak, ok := os.LookupEnv("ALIBABACLOUD_ACCESS_KEY_ID")
	if !ok || ak == "" {
		ak, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY_ID")
		if !ok || ak == "" {
			ak, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY")
			if !ok || ak == "" {
				return nil, errors.New("ALIBABACLOUD_ACCESS_KEY_ID, ALICLOUD_ACCESS_KEY_ID or ALICLOUD_ACCESS_KEY environment variable must be set")
			}
		}
	}

	secret, ok := os.LookupEnv("ALIBABACLOUD_ACCESS_KEY_SECRET")
	if !ok || secret == "" {
		secret, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY_SECRET")
		if !ok || secret == "" {
			secret, ok = os.LookupEnv("ALICLOUD_SECRET_KEY")
			if !ok || secret == "" {
				return nil, errors.New("ALIBABACLOUD_ACCESS_KEY_SECRET, ALICLOUD_ACCESS_KEY_SECRET or ALICLOUD_ACCESS_KEY environment variable must be set")
			}
		}
	}

	region, ok := os.LookupEnv("ALIBABACLOUD_REGION_ID")
	if !ok || region == "" {
		region, ok = os.LookupEnv("ALICLOUD_REGION_ID")
		if !ok || region == "" {
			region, ok = os.LookupEnv("ALICLOUD_REGION")
			if !ok || region == "" {
				return nil, errors.New("ALIBABACLOUD_REGION_ID, ALICLOUD_REGION_ID or ALICLOUD_REGION environment variable must be set")
			}
		}
	}

	client, err := ram.NewClientWithAccessKey(region, ak, secret)

	return client, err
}

/*
func getCreds(_ context.Context) (credentials.Credential, error) {
	return credentials.NewCredential(nil)
}

func connectRam(ctx context.Context) (*ram.Client, error) {
	creds, err := getCreds(ctx)
	if err != nil {
		return nil, err
	}
	accessKeyId, err := creds.GetAccessKeyId()
	accessSecret, err := creds.GetAccessKeySecret()
	client, err := ram.NewClientWithAccessKey("us-east-1", *accessKeyId, *accessSecret)
	//client, err := ram.NewClientWithAccessKey("us-east-1", os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABACLOUD_ACCESS_KEY_SECRET"))
	//client, err := ram.NewClient()
	return client, err
}

*/
