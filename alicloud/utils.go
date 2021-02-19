package alicloud

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

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

func csvToStringArray(_ context.Context, d *transform.TransformData) (interface{}, error) {
	s := d.Value.(string)
	if s == "" {
		// Empty string should always be an empty array
		return []string{}, nil
	}
	sep := ","
	if d.Param != nil {
		sep = d.Param.(string)
	}
	return strings.Split(s, sep), nil
}

func getEnv(ctx context.Context) (region string, ak string, secret string, err error) {

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

	region, ok := os.LookupEnv("ALIBABACLOUD_REGION_ID")
	if !ok || region == "" {
		region, ok = os.LookupEnv("ALICLOUD_REGION_ID")
		if !ok || region == "" {
			region, ok = os.LookupEnv("ALICLOUD_REGION")
			if !ok || region == "" {
				err = errors.New("ALIBABACLOUD_REGION_ID, ALICLOUD_REGION_ID or ALICLOUD_REGION environment variable must be set")
				return
			}
		}
	}

	ak, ok = os.LookupEnv("ALIBABACLOUD_ACCESS_KEY_ID")
	if !ok || ak == "" {
		ak, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY_ID")
		if !ok || ak == "" {
			ak, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY")
			if !ok || ak == "" {
				err = errors.New("ALIBABACLOUD_ACCESS_KEY_ID, ALICLOUD_ACCESS_KEY_ID or ALICLOUD_ACCESS_KEY environment variable must be set")
				return
			}
		}
	}

	secret, ok = os.LookupEnv("ALIBABACLOUD_ACCESS_KEY_SECRET")
	if !ok || secret == "" {
		secret, ok = os.LookupEnv("ALICLOUD_ACCESS_KEY_SECRET")
		if !ok || secret == "" {
			secret, ok = os.LookupEnv("ALICLOUD_SECRET_KEY")
			if !ok || secret == "" {
				err = errors.New("ALIBABACLOUD_ACCESS_KEY_SECRET, ALICLOUD_ACCESS_KEY_SECRET or ALICLOUD_ACCESS_KEY environment variable must be set")
				return
			}
		}
	}

	return region, ak, secret, nil
}

func connectRam(ctx context.Context) (*ram.Client, error) {
	region, ak, secret, err := getEnv(ctx)
	if err != nil {
		return nil, err
	}
	return ram.NewClientWithAccessKey(region, ak, secret)
}

func connectVpc(ctx context.Context) (*vpc.Client, error) {
	region, ak, secret, err := getEnv(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewClientWithAccessKey(region, ak, secret)
}

func connectOss(ctx context.Context) (*oss.Client, error) {
	region, ak, secret, err := getEnv(ctx)
	if err != nil {
		return nil, err
	}
	return oss.New("oss-"+region+".aliyuncs.com", ak, secret)
}
