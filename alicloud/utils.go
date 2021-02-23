package alicloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
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

func ecsTagsToMap(tags []ecs.Tag) (map[string]string, error) {
	var turbotTagsMap map[string]string

	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[i.TagKey] = i.TagValue
	}

	return turbotTagsMap, nil
}
