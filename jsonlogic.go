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
				return opEqualStrict(value, data)
			case "==":
				return opEqual(value, data)
			case "!==":
				return opNotEqualStrict(value, data)
			case "!=":
				return opNotEqual(value, data)
			case "!!":
				return opDoubleNot(value, data)
			case "!":
				return opNot(value, data)
			case "<":
				return opSmallerThan(value, data)
			case ">":
				return opGreaterThan(value, data)
			case ">=":
				return opGreaterEqThan(value, data)
			case "<=":
				return opSmallerEqThan(value, data)
			case "+":
				return opSum(value, data)
			case "-":
				return opSub(value, data)
			case "*":
				return opMult(value, data)
			case "/":
				return opDiv(value, data)
			case "%":
				return opMod(value, data)
			case "and":
				return opAnd(value, data)
			case "or":
				return opOr(value, data)
			case "merge":
				return opMerge(value, data)
			case "in":
				return opIn(value, data)
			case "substr":
				return opSubstr(value, data)
			case "cat":
				return opCat(value, data)
			case "log":
				return opLog(value)
			case "var":
				return opVar(value, data)
			case "if", "?:": // "?:" is an undocumented alias of 'if'
				return opIf(value, data)
			case "max":
				return opMax(value, data)
			case "all":
				return opAll(value, data)
			case "none":
				return opNone(value, data)
			case "some":
				return opSome(value, data)
			case "missing":
				return opMissing(value, data)
			case "missing_some":
				return opMissingSome(value, data)
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
