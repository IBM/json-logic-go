package jsonlogic

import (
	"encoding/json"
	"fmt"
	"os"
)

// StringToInterface converts a string json to an interface{}
func StringToInterface(input string) interface{} {
	b := []byte(input)
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return f
}

// Apply takes in an interface{} and applies its logic
func Apply(input interface{}) bool {
	inputmap := input.(map[string]interface{})

	for operator, value := range inputmap {
		switch operator {
		case "==":
			fmt.Println("detected ==")
			valuearray := value.([]interface{})
			value1 := valuearray[0]
			value2 := valuearray[1]
			return value1 == value2
		case "!=":
			fmt.Println("detected !=")
			valuearray := value.([]interface{})
			value1 := valuearray[0]
			value2 := valuearray[1]
			return value1 != value2
		default:
			fmt.Println("unrecognized operator", operator)
			return false
		}
	}

	return false
}
