package jsonlogic

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// StringToInterface converts a string json to an interface{}
func StringToInterface(input string) interface{} {
	b := []byte(input)
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		fmt.Println("Unmarshal warning:", err)
		// The API supports sending just plain string. How to differentiate between an invalid json and a plain string?
		return input
	}

	return f
}

// Apply takes in an interface{} and applies its logic
func Apply(input interface{}) interface{} {
	fmt.Println("input is of type", reflect.TypeOf(input))
	switch input.(type) {
	case map[string]interface{}:
		//It's a rule
		inputmap := input.(map[string]interface{})

		for operator, value := range inputmap {
			switch operator {
			case "===":
				fallthrough //golang does not support '===', so It's the same as '=='. To be discussed.
			case "==":
				valuearray := value.([]interface{})
				value1 := valuearray[0]
				value2 := valuearray[1]
				return value1 == value2
			case "!==":
				fallthrough //golang does not support '!==', so It's the same as '!='. To be discussed.
			case "!=":
				valuearray := value.([]interface{})
				value1 := valuearray[0]
				value2 := valuearray[1]
				return value1 != value2
			default:
				fmt.Println("unrecognized operator", operator)
				return nil
			}
		}
		break
	default:
		//Non-rule
		return input
	}

	return nil

}
