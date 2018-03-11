package jsonlogic

import (
	"encoding/json"
	"fmt"
)

// stringToInterface converts a string json to an interface{}
func stringToInterface(input string) interface{} {
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

// Apply takes in a string rule and an optional string data and applies its logic
func Apply(inputs ...string) interface{} {
	var rule, data interface{}
	if len(inputs) < 1 {
		//TODO: Expected behavior with no params?
		return nil
	}
	rule = stringToInterface(inputs[0])
	if len(inputs) > 1 {
		//We have data inputs
		data = stringToInterface(inputs[1])
	}

	return applyInterfaces(rule, data)
}

// applyInterfaces takes in an interface{} and an optional interface{} data set and applies its logic
func applyInterfaces(inputs ...interface{}) interface{} {
	var rule, data interface{}
	if len(inputs) < 1 {
		//TODO: Expected behavior with no params?
		return nil
	}
	rule = inputs[0]
	if len(inputs) > 1 {
		//We have data inputs
		data = inputs[1]
	}

	switch rule.(type) {
	case map[string]interface{}:
		//It's a rule
		inputmap := rule.(map[string]interface{})

		for operator, value := range inputmap {
			switch operator {
			case "===":
				fallthrough //golang does not support '===', so It's the same as '=='. To be discussed.
			case "==":
				valuearray := value.([]interface{})
				return applyInterfaces(valuearray[0], data) == applyInterfaces(valuearray[1], data)
			case "!==":
				fallthrough //golang does not support '!==', so It's the same as '!='. To be discussed.
			case "!=":
				valuearray := value.([]interface{})
				return applyInterfaces(valuearray[0], data) != applyInterfaces(valuearray[1], data)
			case "and":
				valuearray := value.([]interface{})
				for _, e := range valuearray {
					if applyInterfaces(e, data) == false {
						return false
					}
				}
				return true
			case "var":
				switch value.(type) {
				case []interface{}: // An array of values
					valuearray := value.([]interface{})
					value1 := applyInterfaces(valuearray[0], data)
					var value2 interface{}
					if len(valuearray) > 1 {
						value2 = applyInterfaces(valuearray[1], data)
					}
					return dataLookup(data, value1, value2)
				default: // A single value
					return dataLookup(data, applyInterfaces(value, data), nil)
				}

			default:
				fmt.Println("unrecognized operator", operator)
				return nil
			}
		}
		break
	default:
		//Non-rule
		return rule
	}

	return nil

}

func dataLookup(data interface{}, index interface{}, defaultValue interface{}) interface{} {
	switch data.(type) {
	case map[string]interface{}:
		switch index.(type) {
		case string:
			value, ok := data.(map[string]interface{})[index.(string)]
			if ok == true {
				return value
			}
			return defaultValue

		}
	case []interface{}:
		dataArray := data.([]interface{})
		switch index.(type) {
		case float64:
			indexInt := int(index.(float64))
			if len(dataArray) >= indexInt+1 {
				return dataArray[indexInt]
			}
			return defaultValue
		}
	}
	return defaultValue
}
