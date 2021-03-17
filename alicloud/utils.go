package alicloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

// Constants for Standard Column Descriptions
const (
	ColumnDescriptionAkas    = "Array of globally unique identifier strings (also known as) for the resource."
	ColumnDescriptionTags    = "A map of tags for the resource."
	ColumnDescriptionTitle   = "Title of the resource."
	ColumnDescriptionAccount = "The Alicloud Account ID in which the resource is located."
	ColumnDescriptionRegion  = "The Alicloud region in which the resource is located."
)

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

func modifyEcsSourceTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]ecs.Tag)

	type resourceTags = struct {
		TagKey   string
		TagValue string
	}
	var sourceTags []resourceTags

	if tags != nil {
		for _, i := range tags {
			sourceTags = append(sourceTags, resourceTags{i.TagKey, i.TagValue})
		}
	}

	return sourceTags, nil
}

func ecsTagsToMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]ecs.Tag)

	if tags == nil {
		return nil, nil
	}

	if len(tags) == 0 {
		return nil, nil
	}

	turbotTagsMap := map[string]string{}
	for _, i := range tags {
		turbotTagsMap[i.TagKey] = i.TagValue
	}

	return turbotTagsMap, nil
}

func vpcTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]vpc.Tag)

	if tags == nil || len(tags) == 0 {
		return nil, nil
	}

	turbotTags := map[string]string{}
	for _, i := range tags {
		turbotTags[i.Key] = i.Value
	}
	return turbotTags, nil
}

func modifyEssSourceTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]ess.TagResource)

	type resourceTags = struct {
		TagKey   string
		TagValue string
	}
	var sourceTags []resourceTags

	if tags != nil {
		for _, i := range tags {
			sourceTags = append(sourceTags, resourceTags{i.TagKey, i.TagValue})
		}
	}

	return sourceTags, nil
}

func essTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]ess.TagResource)

	if tags == nil || len(tags) == 0 {
		return nil, nil
	}

	turbotTags := map[string]string{}
	for _, i := range tags {
		turbotTags[i.TagKey] = i.TagValue
	}
	return turbotTags, nil
}

func zoneToRegion(_ context.Context, d *transform.TransformData) (interface{}, error) {
	region := d.Value.(string)
	return region[:len(region)-1], nil
}

func modifyKmsSourceTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]kms.Tag)

	type resourceTags = struct {
		TagKey   string
		TagValue string
	}
	var sourceTags []resourceTags

	if tags != nil {
		for _, i := range tags {
			sourceTags = append(sourceTags, resourceTags{i.TagKey, i.TagValue})
		}
	}

	return sourceTags, nil
}

func kmsTagsToMap(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]kms.Tag)

	if tags == nil {
		return nil, nil
	}

	if len(tags) == 0 {
		return nil, nil
	}

	turbotTagsMap := map[string]string{}
	for _, i := range tags {
		turbotTagsMap[i.TagKey] = i.TagValue
	}

	return turbotTagsMap, nil
}
