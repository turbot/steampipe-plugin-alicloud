package alicloud

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ess"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vpc"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
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

	for _, i := range tags {
		sourceTags = append(sourceTags, resourceTags{i.TagKey, i.TagValue})
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

	if len(tags) == 0 {
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

	for _, i := range tags {
		sourceTags = append(sourceTags, resourceTags{i.TagKey, i.TagValue})
	}

	return sourceTags, nil
}

func essTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]ess.TagResource)

	if len(tags) == 0 {
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

	for _, i := range tags {
		sourceTags = append(sourceTags, resourceTags{i.TagKey, i.TagValue})
	}

	return sourceTags, nil
}

func kmsTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.Value.([]kms.Tag)

	if len(tags) == 0 {
		return nil, nil
	}

	turbotTags := map[string]string{}
	for _, i := range tags {
		turbotTags[i.TagKey] = i.TagValue
	}
	return turbotTags, nil
}

func GetBoolQualValue(quals plugin.KeyColumnQualMap, columnName string) (value *bool, exists bool) {
	exists = false
	if quals[columnName] == nil {
		return nil, exists
	}

	if quals[columnName].Quals == nil {
		return nil, exists
	}

	for _, qual := range quals[columnName].Quals {
		if qual.Value != nil {
			value := qual.Value
			boolValue := value.GetBoolValue()
			switch qual.Operator {
			case "<>":
				return types.Bool(!boolValue), true
			case "=":
				return types.Bool(boolValue), true
			}
			break
		}
	}
	return nil, exists
}

// GetStringQualValue :: Can be used to get equal value
func GetStringQualValue(quals plugin.KeyColumnQualMap, columnName string) (value *string, exists bool) {
	exists = false
	if quals[columnName] == nil {
		return nil, exists
	}

	if quals[columnName].Quals == nil {
		return nil, exists
	}

	for _, qual := range quals[columnName].Quals {
		if qual.Operator != "=" {
			return nil, exists
		}
		if qual.Value != nil {
			value := qual.Value
			// In case of IN caluse the qual value is usally of format vpcid = '[id1 id2]'
			// which can lead to generation of wrong filter
			if value.GetListValue() != nil {
				// Cannot assign array values
				return nil, exists
			} else {
				return types.String(value.GetStringValue()), true
			}
		}
	}
	return nil, exists
}

// GetStringQualValueList :: Can be used to get equal value as a list of strings
// supports only equal operator
func GetStringQualValueList(quals plugin.KeyColumnQualMap, columnName string) (values []string, exists bool) {
	exists = false
	if quals[columnName] == nil {
		return nil, exists
	}

	if quals[columnName].Quals == nil {
		return nil, exists
	}

	for _, qual := range quals[columnName].Quals {
		if qual.Operator != "=" {
			return nil, exists
		}
		if qual.Value != nil {
			value := qual.Value
			if value.GetListValue() != nil {
				for _, q := range value.GetListValue().Values {
					values = append(values, q.GetStringValue())
				}
				return values, true
			} else {
				values = append(values, value.GetStringValue())
				return values, true
			}
		}
	}
	return values, exists
}

type QueryFilterItem struct {
	Key    string
	Values []string
}

// QueryFilters is an array of filters items
type QueryFilters []QueryFilterItem

// To get the stringified value of QueryFilters
func (filters *QueryFilters) String() (string, error) {
	data, err := json.Marshal(filters)
	if err != nil {
		return "", fmt.Errorf("error marshalling filters: %v", err)
	}

	return string(data), nil
}
