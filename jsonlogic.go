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
			case "!!":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
					return truthy(applyInterfaces(valuearray[0], data))
				default:
					return truthy(value)
				}

			case "!":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
					return !(truthy(applyInterfaces(valuearray[0], data)))
				default:
					return !(truthy(value))
				}

			case "<":
				valuearray := value.([]interface{})
				if len(valuearray) == 3 {
					return smallerThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)) &&
						(smallerThan(applyInterfaces(valuearray[1], data), applyInterfaces(valuearray[2], data)))
				}
				return smallerThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))

			case ">":
				valuearray := value.([]interface{})
				if len(valuearray) == 3 {
					return greaterThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)) &&
						(greaterThan(applyInterfaces(valuearray[1], data), applyInterfaces(valuearray[2], data)))
				}
				return greaterThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))

			case ">=":
				valuearray := value.([]interface{})
				if len(valuearray) == 3 {
					return greaterEqualThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)) &&
						(greaterEqualThan(applyInterfaces(valuearray[1], data), applyInterfaces(valuearray[2], data)))
				}
				return greaterEqualThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))

			case "<=":
				valuearray := value.([]interface{})
				if len(valuearray) == 3 {
					return smallerEqualThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)) &&
						(smallerEqualThan(applyInterfaces(valuearray[1], data), applyInterfaces(valuearray[2], data)))
				}
				return smallerEqualThan(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))
			case "+":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
				default:
					return interFaceToFloat(applyInterfaces(value, data))
				}

				if len(valuearray) == 3 {
					return sumOp(sumOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)), applyInterfaces(valuearray[2], data))
				} else if len(valuearray) == 2 {
					return sumOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))
				} else {
					return interFaceToFloat(applyInterfaces(valuearray[0], data))
				}
			case "-":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
				default:
					return interFaceToFloat(applyInterfaces(value, data))
				}

				if len(valuearray) == 3 {
					return subOp(subOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)), applyInterfaces(valuearray[2], data))
				} else if len(valuearray) == 2 {
					return subOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))
				} else {
					return interFaceToFloat(applyInterfaces(valuearray[0], data))
				}
			case "*":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
				default:
					return interFaceToFloat(applyInterfaces(value, data))
				}

				if len(valuearray) == 3 {
					return multOp(multOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)), applyInterfaces(valuearray[2], data))
				} else if len(valuearray) == 2 {
					return multOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))
				} else {
					return interFaceToFloat(applyInterfaces(valuearray[0], data))
				}
			case "/":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
				default:
					return interFaceToFloat(applyInterfaces(value, data))
				}

				if len(valuearray) == 3 {
					return divOp(sumOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)), applyInterfaces(valuearray[2], data))
				} else if len(valuearray) == 2 {
					return divOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))
				} else {
					return interFaceToFloat(applyInterfaces(valuearray[0], data))
				}
			case "%":
				var valuearray []interface{}
				switch value.(type) {
				case []interface{}:
					valuearray = value.([]interface{})
				default:
					return interFaceToFloat(applyInterfaces(value, data))
				}

				if len(valuearray) == 3 {
					return modOp(modOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data)), applyInterfaces(valuearray[2], data))
				} else if len(valuearray) == 2 {
					return modOp(applyInterfaces(valuearray[0], data), applyInterfaces(valuearray[1], data))
				} else {
					return interFaceToFloat(applyInterfaces(valuearray[0], data))
				}
			case "and":
				return opAnd(value, data)
			case "or":
				return opOr(value, data)
			case "var":
				switch value.(type) {
				case []interface{}: // An array of values
					valuearray := value.([]interface{})
					if len(valuearray) > 0 {
						value1 := applyInterfaces(valuearray[0], data)
						var value2 interface{}
						if len(valuearray) > 1 {
							value2 = applyInterfaces(valuearray[1], data)
						}
						return dataLookup(data, value1, value2)
					}
					//TODO: Expected behavior for empty array?
					return false

				default: // A single value
					return dataLookup(data, applyInterfaces(value, data), nil)
				}
			case "if":
				return opIf(value, data)
			case "all":
				return opAll(value, data)
			case "none":
				return opNone(value, data)
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
