package jsonlogic

import (
	"fmt"
	"reflect"
)

//Per http://jsonlogic.com/truthy.html

func truthy(input interface{}) (bool, error) {
	switch input.(type) {
	case bool:
		return input.(bool), nil
	case float64:
		return (input.(float64) != 0), nil
	case int:
		return (input.(int) != 0), nil
	case []interface{}: //Real world are interfaces, but tests use real types
		return (len(input.([]interface{})) != 0), nil
	case string:
		return (len(input.(string)) != 0), nil
	case nil:
		return false, nil
	case map[string]interface{}:
		return true, nil
	default:
		return false, fmt.Errorf("truthy unexpected type %v", reflect.TypeOf(input))
	}
}
