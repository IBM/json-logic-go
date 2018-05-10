package jsonlogic

import (
	"fmt"
	"reflect"
	"strings"
)

func opIn(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	var needle, haystack interface{}
	var err error
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for AND operator")
	}

	if len(valuearray) > 0 {
		needle, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			haystack, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}
	}

	haystackMeta := reflect.ValueOf(haystack)

	switch haystackMeta.Kind() {
	case reflect.String:
		return strings.Contains(haystack.(string), needle.(string)), nil
	case reflect.Slice, reflect.Array:
		len := haystackMeta.Len()
		for i := 0; i < len; i++ {
			if reflect.DeepEqual(needle, getValue(haystackMeta.Index(i))) {
				return true, nil
			}
		}
	default:
		return false, nil
	}

	return false, nil
}

func getValue(metaValue reflect.Value) interface{} {
	switch metaValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(metaValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(metaValue.Uint())
	case reflect.Float32, reflect.Float64:
		return metaValue.Float()
	case reflect.String:
		return metaValue.String()
	case reflect.Bool:
		return metaValue.Bool()
	case reflect.Map, reflect.Array, reflect.Struct, reflect.Slice, reflect.Interface:
		return metaValue.Interface()
	}

	return nil
}
