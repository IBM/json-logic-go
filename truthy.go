package jsonlogic

import (
	"log"
	"reflect"
)

//Per http://jsonlogic.com/truthy.html

func truthy(input interface{}) bool {
	switch input.(type) {
	case bool:
		return input.(bool)
	case float64:
		return (input.(float64) != 0)
	case int:
		return (input.(int) != 0)
	case []interface{}: //Real world are interfaces, but tests use real types
		return len(input.([]interface{})) != 0
	case string:
		return len(input.(string)) != 0
	case nil:
		return false
	case map[string]interface{}:
		return true
	default:
		log.Println("truthy unexpected type", reflect.TypeOf(input))
		return true
	}
}
