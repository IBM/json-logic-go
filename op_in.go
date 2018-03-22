package jsonlogic

import (
	"fmt"
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

	switch haystack.(type) {
	case string:
		return strings.Contains(haystack.(string), needle.(string)), nil
	case []interface{}:
		for _, word := range haystack.([]interface{}) {
			if needle == word {
				return true, nil
			}
		}
	default:
		return false, nil
	}
	return false, nil
}
