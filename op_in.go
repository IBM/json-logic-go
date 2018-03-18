package jsonlogic

import (
	"strings"
)

func opIn(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	var needle, haystack interface{}
	var err error
	needle, err = applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}

	haystack, err = applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
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
